#!/bin/sh
echo "sleeping for 10 seconds"
sleep 10
store-srv --registry_address=$CONSUL_PORT_8500_TCP_ADDR:8500 --broker_address=127.0.0.1:9000 --server_address=127.0.0.1:9001
