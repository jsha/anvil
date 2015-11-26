// Copyright 2014 ISRG.  All rights reserved
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

// This command does a one-time setup of the RabbitMQ exchange and the Activity
// Monitor queue, suitable for setting up a dev environment or Travis.

import (
	"fmt"
	"os"

	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/streadway/amqp"

	"github.com/letsencrypt/boulder/cmd"
)

// Constants for AMQP
const (
	monitorQueueName = "Monitor"
	amqpExchange     = "boulder"
	amqpExchangeType = "topic"
	amqpInternal     = false
	amqpDurable      = false
	amqpDeleteUnused = false
	amqpExclusive    = false
	amqpNoWait       = false
)

func setup(c *cli.Context) {
	server := c.GlobalString("server")
	conn, err := amqp.Dial(server)
	cmd.FailOnError(err, "Could not connect to AMQP")
	ch, err := conn.Channel()
	cmd.FailOnError(err, "Could not connect to AMQP")

	err = ch.ExchangeDeclare(
		amqpExchange,
		amqpExchangeType,
		amqpDurable,
		amqpDeleteUnused,
		amqpInternal,
		amqpNoWait,
		nil)
	cmd.FailOnError(err, "Declaring exchange")

	_, err = ch.QueueDeclare(
		monitorQueueName,
		amqpDurable,
		amqpDeleteUnused,
		amqpExclusive,
		amqpNoWait,
		nil)
	if err != nil {
		cmd.FailOnError(err, "Could not declare queue")
	}

	routingKey := "#" //wildcard

	err = ch.QueueBind(
		monitorQueueName,
		routingKey,
		amqpExchange,
		false,
		nil)
	if err != nil {
		txt := fmt.Sprintf("Could not bind to queue [%s]. NOTE: You may need to delete %s to re-trigger the bind attempt after fixing permissions, or manually bind the queue to %s.", monitorQueueName, monitorQueueName, routingKey)
		cmd.FailOnError(err, txt)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "rabbitmq-setup"
	app.Usage = "Sets up rabbitmq"
	app.Version = cmd.Version()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "server",
			Value: "",
			Usage: "RabbitMQ server URL",
		},
	}

	app.Action = setup
	app.Run(os.Args)
}
