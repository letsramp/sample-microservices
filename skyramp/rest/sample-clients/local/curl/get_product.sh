#1/bin/sh


set -e

curl -X 'GET' \
  'http://product-catalog-service-port60000.rest.skyramp.test/get-product?product_id=OLJCESPC7Z' \
    --resolve product-catalog-service-port60000.rest.skyramp.test:80:127.0.0.1 \
  -H 'accept: application/json'

