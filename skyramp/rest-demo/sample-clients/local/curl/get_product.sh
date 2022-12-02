#1/bin/sh


set -e

curl -X 'GET' \
  'http://rest-demo.product-catalog-service-port60000.e2e-target.skyramp.test/get-product?product_id=OLJCESPC7Z' \
  -H 'accept: application/json'

