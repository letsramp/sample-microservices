#!/bin/sh

set -e

curl -X 'GET' \
  'http://product-catalog-service-port60000.demo.skyramp.test/search-products?query=kitchen' \
  -H 'accept: application/json'
