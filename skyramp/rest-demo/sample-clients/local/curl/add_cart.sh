#!/bin/sh

set -e

curl -X 'POST' 'http://rest-demo.cart-service-port60000.e2e-target.skyramp.test/cart/user_id/abcde' \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -d '{
  "product_id": "L9ECAV7KIM",
  "quantity": 8
}'

