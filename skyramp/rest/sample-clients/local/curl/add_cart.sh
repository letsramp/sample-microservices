#!/bin/sh

set -e

curl -X 'POST' 'http://cart-service-port60000.rest.skyramp.test/cart/user_id/abcde' \
  --resolve cart-service-port60000.rest.skyramp.test:80:127.0.0.1 \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -d '{
  "product_id": "L9ECAV7KIM",
  "quantity": 8
}'
