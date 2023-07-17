#!/bin/sh

set -e

curl -X 'DELETE' \
  'http://cart-service:60000/cart/user_id/abcde' \
   -H 'accept: application/json' ; echo
