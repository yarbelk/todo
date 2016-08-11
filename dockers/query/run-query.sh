#!/bin/sh
echo "sleeping for 1 seconds"
sleep 1
query-srv --registry_address=consul --broker_address=0.0.0.0:9000 --server_address=0.0.0.0:9001
