#!/bin/sh

set -e

curl -X 'DELETE' \
  'http://cart-service-port60000.demo.skyramp.test/cart/user_id/abcde' \
   -H 'accept: application/json'
