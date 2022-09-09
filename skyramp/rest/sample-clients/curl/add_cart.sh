#!/bin/sh

set -e

curl -X 'POST' 'http://cart-service-port60000.demo.skyramp.test/cart/abcde' \
  --resolve cart-service-port60000.demo.skyramp.test:80:127.0.0.1 \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -d '{
  "product_id": "ls4psxunum",
  "quantity": 8
}'
