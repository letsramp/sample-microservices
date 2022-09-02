package thrift

var orderResult = `
  {
    "order_id": "order-id-123",
    "shipping_tracking_id": "tracking-id-000",
    "shipping_cost": {
        "currency_code": "USD",
        "units": 10,
        "nanos": 10
    },
    "shipping_address": {
        "street_address": "Main Street 101",
        "city": " San Fransisco",
        "state": "CA",
        "country": "US",
        "zip_code": 95000
    },
    "items": [
        {
            "item": {
                "product_id": "OLJCESPC7Z",
                "quantity": 5
            },
            "cost": {
                "currency_code": "USD",
                "units": 95,
                "nanos": 500000000
            }
        },
        {
            "item": {
                "product_id": "66VCHSJNUP",
                "quantity": 2
            },
            "cost": {
                "currency_code": "USD",
                "units": 36,
                "nanos": 200000000
            }
        }
    ]
}
`
