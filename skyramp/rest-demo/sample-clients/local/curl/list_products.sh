#!/bin/sh

set -e

curl -X 'GET' \
  'http://product-catalog-service:60000/get-products' \
  -H 'accept: application/json' ; echo
