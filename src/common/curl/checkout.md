

## Checklout
```
curl -X 'POST' \
  'http://checkoutservice:60000/checkout' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": "abcde",
  "user_currency": "UDS",
  "address": {
    "street_address": "1600 Amp street",
    "city": "Mountain View",
    "state": "CA",
    "country": "USA",
    "zip_code": "94043"
  },
  "email": "someone@example.com",
  "credit_card": {
    "credit_card_number": "4432-8015-6251-0454",
    "credit_card_cvv": 672,
    "credit_card_expiration_year": 24,
    "credit_card_expiration_month": 1
  }
}' | jq .
```


Expected Result
```
{
  "order_id": "99ec9ebd-13fc-411d-8997-075d9a5cdb9d",
  "shipping_tracking_id": "85f19309-cdd6-4b42-abba-42f9d8d1bd60",
  "shipping_cost": {
    "currency_code": "USD",
    "units": 10,
    "nanos": 100
  },
  "shipping_address": {
    "street_address": "1600 Amp street",
    "city": "Mountain View",
    "state": "CA",
    "country": "USA",
    zip_code": "94043"
  },
  "items": [
    {
      "item": {
        "product_id": "LS4PSXUNUM",
        "quantity": 8
      }
    }
  ]
}



curl -X 'PUT' \
  'http://shippingservice:60000/get-quote' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "address": {
    "street_address": "1600 Amp street",
    "city": "Mountain View",
    "state": "CA",`
    "country": "USA",
     zip_code": "94043"
  },
  "items": [
    {
        "product_id": "L9ECAV7KIM",
        "quantity": 2
    },
    {
        "product_id": "2ZYFJ3GM2N",
        "quantity": 1
    }
  ]
}'
```
