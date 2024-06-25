#!/bin/sh

set -e

curl -X 'POST' 'http://localhost:60003/cart/user_id/abcde' \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -d '{
  "product_id": "1YMWWN1N4O",
  "quantity": 8
}'; echo


curl -X 'POST' \
  'http://localhost:60001/checkout' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": "abcde",
  "user_currency": "USD",
  "address": {
    "street_address": "1600 Amp street",
    "city": "Mountain View",
    "state": "CA",
    "country": "USA",
    "zip_code": "94043"
  },
  "email": "someone@example.com",
  "credit_card": {
    "credit_card_number": "4432-8015-6251-0454",
    "credit_card_cvv": 672,
    "credit_card_expiration_year": 28,
    "credit_card_expiration_month": 1
  }
}'; echo

