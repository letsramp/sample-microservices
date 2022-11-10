#!/bin/sh

set -e

curl -X 'POST' 'http://cart-service-port60000.demo.skyramp.test/cart/user_id/abcde' \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -d '{
  "product_id": "L9ECAV7KIM",
  "quantity": 8
}'
