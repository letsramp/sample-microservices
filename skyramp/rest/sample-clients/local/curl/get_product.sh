#1/bin/sh


set -e

curl -X 'GET' \
  'http://product-catalog-service-port60000.demo.skyramp.test/get-product?product_id=OLJCESPC7Z' \
  -H 'accept: application/json'

