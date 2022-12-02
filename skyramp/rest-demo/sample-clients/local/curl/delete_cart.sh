#!/bin/sh

set -e

curl -X 'DELETE' \
  'http://rest-demo.cart-service-port60000.e2e-target.skyramp.test/cart/user_id/abcde' \
   -H 'accept: application/json'
