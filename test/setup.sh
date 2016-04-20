#!/bin/bash
#
# Fetch dependencies of Boulderthat are necessary for development or testing,
# and configure database and RabbitMQ.
#

set -ev

go get \
  github.com/golang/lint/golint \
  github.com/golang/mock/mockgen \
  github.com/golang/protobuf/proto \
  github.com/google/protobuf/protoc-gen-go \
  github.com/jcjones/github-pr-status \
  github.com/jsha/listenbuddy \
  github.com/kisielk/errcheck \
  github.com/mattn/goveralls \
  github.com/modocache/gover \
  github.com/tools/godep \
  golang.org/x/tools/cmd/stringer \
  golang.org/x/tools/cover &

(wget https://github.com/jsha/boulder-tools/raw/master/goose.gz &&
 mkdir -p $GOPATH/bin &&
 zcat goose.gz > $GOPATH/bin/goose &&
 chmod +x $GOPATH/bin/goose &&
 ./test/create_db.sh) &

# Set up rabbitmq exchange
go run cmd/rabbitmq-setup/main.go -server amqp://boulder-rabbitmq &

# Wait for all the background commands to finish.
for pid in $(jobs -p); do
  wait $pid || exit 1
done
