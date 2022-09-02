


PAYMENT_SERVICE=localhost:60000
curl -d @payment_charge http://$PAYMENT_SERVICE/charge




curl -X POST -H "Content-Type: application/json"  -d @./charge.json -k https://localhost:60000/charge
```
{
  	"amount": {
		"currency_code": "USD",
		"units": 100,
		"nanos": 50
	},
	"credit_card": {
		"credit_card_number": "4432-8015-6152-0454",
		"credit_card_cvv": 672,
		"credit_card_expiration_year": 2024,
		"credit_card_expiration_month": 2
	}
}
```

```
{
  "transaction_id": "5ec76179-b4d6-4047-af03-02629933c826"
}
```
cat < EOF |
{
  	"amount": {
		"currency_code": "USD",
		"units": 100,
		"nanos": 50
	},
	"credit_card": {
		"credit_card_number": "4432-8015-6152-0454",
		"credit_card_cvv": 672,
		"credit_card_expiration_year": 2024,
		"credit_card_expiration_month": 2
	}
}
EOF
