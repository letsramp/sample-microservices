
## Send Order Confirmation
```
Curl

curl -X 'POST' \
  'http://emailservice:60000/send-order-confirmation' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "user@mail.com",
  "order": {
    "order_id": "aa-bb-cc",
    "shipping_tracking_id": "qwe-rty",
    "shipping_cost": {
      "currency_code": "USD",
      "units": 10,
      "nanos": 900
    },
    "shipping_address": {
      "street_address": "Main Street 1",
      "city": "Springfield",
      "state": "CA",
      "country": "US",
      "zip_code": "95000"
    },
    "items": [
      {
        "item": {
          "product_id": "L9ECAV7KIM",
          "quantity": 2
        },
        "cost": {
          "currency_code": "USD",
          "units": 89,
          "nanos": 990000000
        }
      },
      {
        "item": {
          "product_id": "2ZYFJ3GM2N",
          "quantity": 1
        },
        "cost": {
          "currency_code": "USD",
          "units": 24,
          "nanos": 990000000
        }
      }
    ]
  }
}
'


## expected result
```
{ "status" : "200 OK" }
```