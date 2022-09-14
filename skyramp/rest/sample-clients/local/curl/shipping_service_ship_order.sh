#!/bin/sh

set -e

curl -X 'PUT' \
  'http://shipping-service-port60000.demo.skyramp.test/ship-order' \
  --resolve shipping-service-port60000.demo.skyramp.test:80:127.0.0.1 \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "address": {
    "street_address": "1600 Amp street",
    "city": "Mountain View",
    "state": "CA",
    "country": "USA",
    "zip_code": "94043"
  },
  "items": [
    {
        "product_id": "L9ECAV7KIM",
        "quantity": 2
    },
    {
        "product_id": "2ZYFJ3GM2N",
        "quantity": 1
    }
  ]
}' | jq .

