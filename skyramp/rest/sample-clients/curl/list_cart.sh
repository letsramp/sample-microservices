#!/bin/sh

set -e

curl -X 'GET' \
  'http://cart-service-port60000.demo.skyramp.test/cart/abcde' \
   --resolve cart-service-port60000.demo.skyramp.test:80:127.0.0.1 \
   -H 'accept: application/json'
