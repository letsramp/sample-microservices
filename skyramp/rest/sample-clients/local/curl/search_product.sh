#!/bin/sh

set -e

curl -X 'GET' \
  'http://product-catalog-service-port60000.demo.skyramp.test/search-products?query=kitchen' \
    --resolve product-catalog-service-port60000.demo.skyramp.test:80:127.0.0.1 \
  -H 'accept: application/json'
