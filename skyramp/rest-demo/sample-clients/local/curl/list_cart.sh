#!/bin/sh

set -e

curl -X 'GET' \
  'http://cart-service:60000/cart/user_id/abcde' \
   -H 'accept: application/json' ; echo
