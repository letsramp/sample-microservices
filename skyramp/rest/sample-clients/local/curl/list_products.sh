#!/bin/sh

set -e

curl -X 'GET' \
  'http://product-catalog-service-port60000.demo.skyramp.test/get-products' \
  -H 'accept: application/json'
