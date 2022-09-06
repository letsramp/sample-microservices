
# Add Item to Cart
```
curl -X 'POST' \
  'http://cartservice:60000/cart/abcde' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "product_id": "LS4PSXUNUM",
  "quantity": 8
}'
```
Expected Result
```
{"success":"200 OK"}
```

# Get Cart
```
curl -X 'GET' \
  'http://cartservice:60000/cart/abcde' \
   -H 'accept: application/json' | jq .
```

Example Result
```
{
  "user_id": "abcde",
  "items": [
    {
      "product_id": "LS4PSXUNUM",
      "quantity": 8
    }
  ]
}
```

# Delete Cart
```
curl -X 'DELETE' \
  'http://cartservice:60000/cart/abcde' \
   -H 'accept: application/json'
```

Example Result
```
{"success":"200 OK"}
```