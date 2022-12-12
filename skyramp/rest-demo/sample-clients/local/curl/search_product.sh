#!/bin/sh

set -e

curl -X 'GET' \
  'http://rest-demo.product-catalog-service-port60000.e2e-target.skyramp.test/search-products?query=kitchen' \
  -H 'accept: application/json'

