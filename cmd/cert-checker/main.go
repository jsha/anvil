// Copyright 2015 ISRG.  All rights reserved
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"crypto/x509"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cactus/go-statsd-client/statsd"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/codegangsta/cli"
	gorp "github.com/letsencrypt/boulder/Godeps/_workspace/src/gopkg.in/gorp.v1"
	"github.com/letsencrypt/boulder/policy"

	"github.com/letsencrypt/boulder/cmd"
	"github.com/letsencrypt/boulder/core"
	blog "github.com/letsencrypt/boulder/log"
	"github.com/letsencrypt/boulder/sa"
)

const (
	good = "valid"
	bad  = "invalid"
)

type report struct {
	Valid    bool     `json:"validity"`
	Problems []string `json:"problem,omitempty"`
}

type certChecker struct {
	pa           core.PolicyAuthority
	dbMap        *gorp.DbMap
	certs        chan core.Certificate
	sampleReport map[string]report
	goodCerts    int64
	badCerts     int64
}

func newChecker(dbMap *gorp.DbMap) certChecker {
	return certChecker{
		pa:           policy.NewPolicyAuthorityImpl(),
		dbMap:        dbMap,
		sampleReport: make(map[string]report),
	}
}

func (c *certChecker) getCerts() error {
	var certs []core.Certificate
	_, err := c.dbMap.Select(
		&certs,
		"SELECT * FROM certificates WHERE issued > :issued",
		map[string]interface{}{"issued": time.Now().Add(-time.Hour * 24 * 90)},
	)
	if err != nil {
		return err
	}
	c.certs = make(chan core.Certificate, len(certs))
	for _, cert := range certs {
		c.certs <- cert
	}
	// Close channel so range operations won't block when the channel empties out
	close(c.certs)
	return nil
}

func (c *certChecker) processCerts(wg *sync.WaitGroup) {
	for cert := range c.certs {
		// DEBUG
		fmt.Println("CERT:", cert.Serial)

		problems := c.checkCert(cert)
		valid := len(problems) == 0
		c.sampleReport[cert.Serial] = report{Valid: valid, Problems: problems}
		if !valid {
			atomic.AddInt64(&c.badCerts, 1)
		} else {
			atomic.AddInt64(&c.goodCerts, 1)
		}
	}
	wg.Done()
}

func (c *certChecker) checkCert(cert core.Certificate) (problems []string) {
	// Check digests match
	if cert.Digest != core.Fingerprint256(cert.DER) {
		problems = append(problems, "Stored digest doesn't match certificate digest")
	}

	// Parse certificate
	parsedCert, err := x509.ParseCertificate(cert.DER)
	if err != nil {
		problems = append(problems, fmt.Sprintf("Couldn't parse stored certificate: %s", err))
	} else {
		// Check we have the right expiration time
		if parsedCert.NotAfter != cert.Expires {
			problems = append(problems, "Stored expiration doesn't match certificate NotAfter")
		}
		// Check basic constraints are set
		if !parsedCert.BasicConstraintsValid {
			problems = append(problems, "Certificate doesn't have basic constraints set")
		}
		// Check the cert isn't able to sign other certificates
		if parsedCert.IsCA {
			problems = append(problems, "Certificate can sign other certificates")
		}
		// Check the cert has the correct validity period
		if parsedCert.NotAfter.Sub(cert.Issued) > (time.Hour * 24 * 90) {
			problems = append(problems, "Certificate has a validity period longer than 90 days")
		}
		// Check that the PA is still willing to issue for each name in DNSNames + CommonName
		for _, name := range append(parsedCert.DNSNames, parsedCert.Subject.CommonName) {
			if err = c.pa.WillingToIssue(core.AcmeIdentifier{Type: core.IdentifierDNS, Value: name}); err != nil {
				problems = append(problems, fmt.Sprintf("Policy Authority isn't willing to issue for %s: %s", name, err))
			}
		}
		// Check the cert has the correct key usage extensions
		if !core.CmpExtKeyUsageSlice(parsedCert.ExtKeyUsage, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}) {
			problems = append(problems, "Certificate has incorrect key usage extensions")
		}
	}
	return problems
}

func main() {
	app := cmd.NewAppShell("cert-checker")
	app.App.Flags = append(app.App.Flags, cli.IntFlag{
		Name:  "workers",
		Value: 5,
		Usage: "The number of cocurrent workers used to process certificates",
	}, cli.StringFlag{
		Name:  "report-path",
		Usage: "The path to write a JSON report on the certificates checks to (if no path is provided the report will not be written out)",
	}, cli.StringFlag{
		Name:  "sql-uri",
		Usage: "SQL URI if not provided in the configuration file",
	})

	app.Config = func(c *cli.Context, config cmd.Config) cmd.Config {
		config.CertChecker.ReportDirectoryPath = c.GlobalString("report-dir-path")

		if connect := c.GlobalString("sql-uri"); connect != "" {
			config.CertChecker.DBConnect = connect
		}

		return config
	}

	app.Action = func(c cmd.Config) {
		stats, err := statsd.NewClient(c.Statsd.Server, c.Statsd.Prefix)
		cmd.FailOnError(err, "Couldn't connect to statsd")

		auditlogger, err := blog.Dial(c.Syslog.Network, c.Syslog.Server, c.Syslog.Tag, stats)
		cmd.FailOnError(err, "Could not connect to Syslog")

		blog.SetAuditLogger(auditlogger)
		auditlogger.Info(app.VersionString())

		dbMap, err := sa.NewDbMap(c.CertChecker.DBConnect)
		cmd.FailOnError(err, "Could not connect to database")

		checker := newChecker(dbMap)
		auditlogger.Info("# Getting certificates issued in the last 90 days")
		err = checker.getCerts()
		cmd.FailOnError(err, "Failed to get sample certificates")

		if c.CertChecker.Workers > len(checker.certs) {
			c.CertChecker.Workers = len(checker.certs)
		}
		auditlogger.Info(fmt.Sprintf("# Processing sample, %d certificates using %d workers", len(checker.certs), c.CertChecker.Workers))
		wg := new(sync.WaitGroup)
		for i := 0; i < c.CertChecker.Workers; i++ {
			wg.Add(1)
			go checker.processCerts(wg)
		}
		wg.Wait()
		auditlogger.Info(fmt.Sprintf(
			"# Finished processing certificates, sample: %d, good: %d, bad: %d",
			len(checker.sampleReport),
			checker.goodCerts,
			checker.badCerts,
		))
	}

	app.Run()
}
