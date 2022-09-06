
## Payment Service

Charge Creditcard
```
curl -X 'PUT' \
  'http://paymentservice:60000/charge' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "amount": {
    "currency_code": "USD",
    "units": 245,
    "nanos": 9900000
  },
  "credit_card": {
    "credit_card_number": "4432-8015-6152-0454",
    "credit_card_cvv": 672,
    "credit_card_expiration_year": 24,
    "credit_card_expiration_month": 1
  }
}'

```


Expected Result
```
{
  "transaction_id": "cb93ecd4-2cc3-11ed-a261-0242ac120002"
}

```