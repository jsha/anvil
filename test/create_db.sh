#!/bin/bash
set -o errexit
set -o xtrace
cd $(dirname $0)/..
source test/db-common.sh

for svc in $SERVICES; do
  for dbenv in $DBENVS; do
    db="boulder_${svc}_${dbenv}"

    mysql -u root -e "drop database if exists \`${db}\`; create database if not exists \`${db}\`; grant all privileges on ${db}.* to 'boulder'@'localhost'" || die "unable to create ${db}"
    echo "created empty ${db} database"

    goose -path=./$svc/_db/ -env=$dbenv up || die "unable to migrate ${db}"
    echo "migrated ${db} database"
  done
done
echo "created all databases"
