#!/bin/sh

set -e

curl -X 'GET' \
  'http://product-catalog-service:60000/search-products?query=kitchen' \
  -H 'accept: application/json' ; echo
