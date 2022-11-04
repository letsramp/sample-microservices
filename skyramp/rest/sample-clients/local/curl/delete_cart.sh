#!/bin/sh

set -e

curl -X 'DELETE' \
   --resolve cart-service-port60000.rest.skyramp.test:80:127.0.0.1 \
  'http://cart-service-port60000.rest.skyramp.test/cart/user_id/abcde' \
   -H 'accept: application/json'
