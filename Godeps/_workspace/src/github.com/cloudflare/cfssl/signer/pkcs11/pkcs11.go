// +build pkcs11

package pkcs11

import (
	"io/ioutil"
	"fmt"

	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/config"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/crypto/pkcs11key"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/errors"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/helpers"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/log"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/signer"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/signer/local"
)

// Config contains configuration information required to use a PKCS
// #11 key.
type Config struct {
	Module string
	Token  string
	PIN    string
	Label  string
}

// Enabled is set to true if PKCS #11 support is present.
const Enabled = true

// New returns a new PKCS #11 signer.
func New(caCertFile string, policy *config.Signing, cfg *Config) (signer.Signer, error) {
	if cfg == nil {
		fmt.Println("A")
		return nil, errors.New(errors.PrivateKeyError, errors.ReadFailed)
	}

	log.Debugf("Loading PKCS #11 module %s", cfg.Module)
	certData, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		fmt.Println("B")
		return nil, errors.New(errors.PrivateKeyError, errors.ReadFailed)
	}

	cert, err := helpers.ParseCertificatePEM(certData)
	if err != nil {
		return nil, err
	}

	priv, err := pkcs11key.New(cfg.Module, cfg.Token, cfg.PIN, cfg.Label)
	if err != nil {
		fmt.Println("C")
		return nil, errors.New(errors.PrivateKeyError, errors.ReadFailed)
	}
	sigAlgo := signer.DefaultSigAlgo(priv)

	return local.NewSigner(priv, cert, sigAlgo, policy)
}
