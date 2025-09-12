#!/bin/bash

# Monitoring Microservice test script
# Replace these variables with your actual values
MICROSERVICE_URL="your-microservice-url"
HOST_FQDN="your-host-fqdn"

echo "Testing Monitoring Microservice..."
echo "Microservice URL: $MICROSERVICE_URL"
echo "Host FQDN: $HOST_FQDN"
echo "----------------------------------------"

# GET pinghost
echo "Testing GET pinghost..."
curl -X GET "http://${MICROSERVICE_URL}/${HOST_FQDN}/pinghost"
sleep 1

echo "----------------------------------------"
echo "All tests completed."
