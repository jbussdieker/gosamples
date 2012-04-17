#!/bin/bash -e
HOST=ubuntu@$1
RECIPE=$2

scp -o StrictHostKeyChecking=no -i master.pem authorized_keys ${HOST}:.ssh/authorized_keys > /dev/null
rsync --verbose -r -d --delete bootstrap ${HOST}: > /dev/null
ssh ${HOST} sudo ./bootstrap/setup.sh $RECIPE || exit 1

