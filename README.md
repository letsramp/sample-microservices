#**Skyramp** sample microservices are based on GCP **Online Boutique** with extension
to communicate using rest and thrift.


## Getting Started

### install tools pre-requisit
- docker
- compose
- jq
- curl


## Copy Code for sample services
```
git clone https://github.com/letsramp/sample-services.git
cd sample-services/src
```

## Building Sample services from source

 <!-- TODO Update service images with local registry  -->

## Push Services to container registry
```
docker compose push
```

## Use Skyramo to Create a local kubernetes cluster
```
skyramo config local apply
```


## Use skyramp to deploy Online Boutique with grpc

```
cd skyramp-projects/grpc

$ skyramp up demo
```

Expected Result:
```
product-catalog-service-556fd4b54f-8jjqf      ready
recommendation-service-7fb4598577-spj2t       ready
cart-service-68fc59dc8c-lbqht                 ready
currency-service-69dbcd889d-f4v9z             ready
frontend-6b5ddd8f69-ssj4c                     ready
redis-7f6445f856-sgk9m                        ready
ad-service-5b56d86b5f-jt8pn                   ready
checkout-service-57549d999c-pd22c             ready
email-service-db4c8f558-vh7w7                 ready
payment-service-568974bcd9-vth2t              ready
shipping-service-554b6c8757-ngrqm             ready
skyramp-debug-worker-7cd4d58c6b-dx8bb         ready
All pods are ready.

```

<!-- TODO Update DNS resolver or /etc/hosts   -->


Open a browser on url http://frontend-port8080.demo.skyramp.test to access
the online store experience

<br/><br/>
![Online Boutique](docs/img/online-boutique.jpg)


## Undeploy Services
```
$ skyramp down demo

```


## Mocking Services with Skyramp Mock Worker

Here we will explore the power of Skyramp Mockworker to mock two services.
- payment-service
- shipping-service

Note: inspect the Container description of the payment-service and pay attention to the endpoint
and signature definition that declares the behaviour.

```
$ vi skyramp-projects/grpc/containers/payment-service.yaml
$ vi skyramp-projects/grpc/containers/shipping-service.yaml

```

Update the target description to indicate that Skyramp Mockworker is to mock the payment and shipping service.

```
vi targets/demo.yaml
```
Add "mock: true" to define that the service is to be mocked
```
containers:
  - container: ad-service
  - container: frontend
  - container: cart-service
  - container: checkout-service
  - container: payment-service
    mock: true
  - container: product-catalog-service
  - container: recommendation-service
  - container: email-service
  - container: shipping-service
    mock: true
  - container: currency-service

```

## Use Skyramp to deploy the new configuration

```
$ skyramp up demo
```

Expected Result:
```
product-catalog-service-556fd4b54f-8jjqf      ready
recommendation-service-7fb4598577-spj2t       ready
cart-service-68fc59dc8c-lbqht                 ready
currency-service-69dbcd889d-f4v9z             ready
frontend-6b5ddd8f69-ssj4c                     ready
redis-7f6445f856-sgk9m                        ready
ad-service-5b56d86b5f-jt8pn                   ready
checkout-service-57549d999c-pd22c             ready
email-service-db4c8f558-vh7w7                 ready
skyramp-worker-7cd4d58c6b-dx8bb               ready
All pods are ready.

```

Note: The shipping and payment service pods are not deployed,
as the Skyramp Mockworker Mocks their behaviour.


# API

## Rest

### Cartservice Operation

**Add Product to Shopping cart**
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

example result
```
{"success":"200 OK"}
```

**Get Shopping cart**
```
curl -X 'GET' \
  'http://cartservice:60000/cart/abcde' \
   -H 'accept: application/json' | jq .
```

example
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


**Delete Shopping Cart**
```
curl -X 'DELETE' \
  'http://cartservice:60000/cart/abcde' \
   -H 'accept: application/json'
```


## Product Catalog Service

### Get Products
```
curl -X 'GET' \
  'http://productcatalogservice:60000/get-products' \
  -H 'accept: application/json' | jq .
```

example result
```
[
  {
    "id": "OLJCESPC7Z",
    "name": "Sunglasses",
    "description": "Add a modern touch to your outfits with these sleek aviator sunglasses.",
    "picture": "/static/img/products/sunglasses.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 19,
      "nanos": 990000000
    },
    "categories": [
      "accessories"
    ]
  },
  {
    "id": "66VCHSJNUP",
    "name": "Tank Top",
    "description": "Perfectly cropped cotton tank, with a scooped neckline.",
    "picture": "/static/img/products/tank-top.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 18,
      "nanos": 990000000
    },
    "categories": [
      "clothing",
      "tops"
    ]
  },
  {
    "id": "1YMWWN1N4O",
    "name": "Watch",
    "description": "This gold-tone stainless steel watch will work with most of your outfits.",
    "picture": "/static/img/products/watch.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 109,
      "nanos": 990000000
    },
    "categories": [
      "accessories"
    ]
  },
  {
    "id": "L9ECAV7KIM",
    "name": "Loafers",
    "description": "A neat addition to your summer wardrobe.",
    "picture": "/static/img/products/loafers.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 89,
      "nanos": 990000000
    },
    "categories": [
      "footwear"
    ]
  },
  {
    "id": "2ZYFJ3GM2N",
    "name": "Hairdryer",
    "description": "This lightweight hairdryer has 3 heat and speed settings. It's perfect for travel.",
    "picture": "/static/img/products/hairdryer.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 24,
      "nanos": 990000000
    },
    "categories": [
      "hair",
      "beauty"
    ]
  },
  {
    "id": "0PUK6V6EV0",
    "name": "Candle Holder",
    "description": "This small but intricate candle holder is an excellent gift.",
    "picture": "/static/img/products/candle-holder.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 18,
      "nanos": 990000000
    },
    "categories": [
      "decor",
      "home"
    ]
  },
  {
    "id": "LS4PSXUNUM",
    "name": "Salt & Pepper Shakers",
    "description": "Add some flavor to your kitchen.",
    "picture": "/static/img/products/salt-and-pepper-shakers.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 18,
      "nanos": 490000000
    },
    "categories": [
      "kitchen"
    ]
  },
  {
    "id": "9SIQT8TOJO",
    "name": "Bamboo Glass Jar",
    "description": "This bamboo glass jar can hold 57 oz (1.7 l) and is perfect for any kitchen.",
    "picture": "/static/img/products/bamboo-glass-jar.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 5,
      "nanos": 490000000
    },
    "categories": [
      "kitchen"
    ]
  },
  {
    "id": "6E92ZMYYFZ",
    "name": "Mug",
    "description": "A simple mug with a mustard interior.",
    "picture": "/static/img/products/mug.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 8,
      "nanos": 990000000
    },
    "categories": [
      "kitchen"
    ]
  }
]
```

### Get Product
```
curl -X 'GET' \
  'http://productcatalogservice:60000/get-product?product_id=OLJCESPC7Z' \
  -H 'accept: application/json' | jq .
```

example result
```
{
  "id": "OLJCESPC7Z",
  "name": "Sunglasses",
  "description": "Add a modern touch to your outfits with these sleek aviator sunglasses.",
  "picture": "/static/img/products/sunglasses.jpg",
  "priceUsd": {
    "currencyCode": "USD",
    "units": 19,
    "nanos": 990000000
  },
  "categories": [
    "accessories"
  ]
}
```

### Search Product
```
curl -X 'GET' \
  'http://productcatalogservice:60000/search-products?query=kitchen' \
  -H 'accept: application/json' | jq .
```

example result
```
[
  {
    "id": "LS4PSXUNUM",
    "name": "Salt & Pepper Shakers",
    "description": "Add some flavor to your kitchen.",
    "picture": "/static/img/products/salt-and-pepper-shakers.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 18,
      "nanos": 490000000
    },
    "categories": [
      "kitchen"
    ]
  },
  {
    "id": "9SIQT8TOJO",
    "name": "Bamboo Glass Jar",
    "description": "This bamboo glass jar can hold 57 oz (1.7 l) and is perfect for any kitchen.",
    "picture": "/static/img/products/bamboo-glass-jar.jpg",
    "priceUsd": {
      "currencyCode": "USD",
      "units": 5,
      "nanos": 490000000
    },
    "categories": [
      "kitchen"
    ]
  }
]
```

## Checkout Service
**Checkout Order**
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

example results
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
    "zip_code": "94043"
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
`

## Recommendation Service

**List Recommendation**
```
curl -X 'GET' \
  'http://recommendationservice:60000/list-recommendations?product_id=LS4PSXUNUM%22' \
  -H 'accept: application/json'
```

example results
```
[
  "2ZYFJ3GM2N",
  "OLJCESPC7Z",
  "LS4PSXUNUM",
  "9SIQT8TOJO",
  "66VCHSJNUP"
]
```

## Email Service
**Send Order Confirmation**
```
curl -X 'POST' \
  'http://emailservice:60000/send-order-confirmation' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "someone@example.com",
  "order": {
    "order_id": "99ec9ebd-13fc-411d-8997-075d9a5cdb9d",
    "shipping_tracking_id": "85f19309-cdd6-4b42-abba-42f9d8d1bd60",
    "shipping_cost": {
      "currency_code": "USD",
      "units": 10,
      "nanos": 900
    },
    "shipping_address": {
      "street_address": "1600 Amp street",
      "city": "Mountain View",
      "state": "CA",
      "country": "USA",
      "zip_code": "94043"
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
```

## Payment Service
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

## Shipping Serviece


**Get Quote**
```
curl -X 'PUT' \
  'http://shippingservice:60000/get-quote' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "address": {
    "street_address": "1600 Amp street",
    "city": "Mountain View",
    "state": "CA",
    "country": "USA",
    "zip_code": "94043"
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
}' | jq .

```

expected result
```
{
  "cost_usd": {
    "currency_code": "USD",
    "units": 8,
    "nanos": 990000000
  }
}
```

---

### Python
The python test clients are available in the project folder and sub directory /files/python

1. Change working directory to the python test clients
```
cd <project folder>/files/python
```

2. Create a python virtual environment
```
$ python3 -m venv env
```

3. Activate the virtual environment
```
$ source ./venv/bin/activate
```

4. Install the python dependencies
```
$ pip install -r requirements.txt
```

4. Add item to the CartService
Inspect the file addCartLocal.py and note that user_id 'abcde' is adding one unit of product_id 'OLJCESPC7Z' to the cart service

```
$  python3 addCartLocal.py
Successfully added item to cart.
```
5. Place the order, and expect the successful result.

```
$  python placeOrderLocal.py
Successfully placed order.

order {
  order_id: "9e5ef048-2994-11ed-bdca-56f53d787cb9"
  shipping_tracking_id: "ET-34675-173898815"
  shipping_cost {
    currency_code: "USD"
    units: 8
    nanos: 990000000
  }
  shipping_address {
    street_address: "1600 Amp street"
    city: "Mountain View"
    state: "CA"
    country: "USA"
    zip_code: 94043
  }
  items {
    item {
      product_id: "OLJCESPC7Z"
      quantity: 1
    }
    cost {
      currency_code: "USD"
      units: 19
      nanos: 990000000
    }
  }
}
```
---



# Openapi
<details>
  <summary>Rest API</summary>
  ```
{
  "openapi": "3.0.2",
  "info": {
    "title": "Skyramp Sample Microservices",
    "version": "0.0.1"
  },
  "paths": {
    "/get-ad": {
      "get": {
        "tags": [
          "ad"
        ],
        "summary": "Get Ad",
        "operationId": "get_ad_get_ad_get",
        "parameters": [
          {
            "required": false,
            "schema": {
              "title": "Category",
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "name": "category",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "title": "Response Get Ad Get Ad Get",
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/Ad"
                  }
                }
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    },
    "/cart/{user_id}": {
      "get": {
        "tags": [
          "cart"
        ],
        "summary": "Get Cart",
        "operationId": "get_cart_cart__user_id__get",
        "parameters": [
          {
            "required": true,
            "schema": {
              "title": "User Id",
              "type": "string"
            },
            "name": "user_id",
            "in": "path"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "title": "Response Get Cart Cart  User Id  Get",
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/Cart"
                  }
                }
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "cart"
        ],
        "summary": "Add To Cart",
        "operationId": "add_to_cart_cart__user_id__post",
        "parameters": [
          {
            "required": true,
            "schema": {
              "title": "User Id",
              "type": "string"
            },
            "name": "user_id",
            "in": "path"
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/Cart"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {}
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      },
      "delete": {
        "tags": [
          "cart"
        ],
        "summary": "Delete Cart",
        "operationId": "delete_cart_cart__user_id__delete",
        "parameters": [
          {
            "required": true,
            "schema": {
              "title": "User Id",
              "type": "string"
            },
            "name": "user_id",
            "in": "path"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {}
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    },
    "/checkout/": {
      "post": {
        "tags": [
          "checkout"
        ],
        "summary": "Checkout Order",
        "operationId": "checkout_order_checkout__post",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/Order"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/OrderResult"
                }
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    },
    "/send-order-confirmation": {
      "post": {
        "tags": [
          "confirmation"
        ],
        "summary": "Send Order Confirmation",
        "operationId": "send_order_confirmation_send_order_confirmation_post",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/OrderConfirmation"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {}
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    },
    "/charge": {
      "put": {
        "tags": [
          "payment"
        ],
        "summary": "Charge",
        "operationId": "charge_charge_put",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/Payment"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/PaymentResponse"
                }
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    },
    "/get-quote": {
      "put": {
        "tags": [
          "shipping"
        ],
        "summary": "Get Shipping Quote",
        "operationId": "get_shipping_quote_get_quote_put",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/GetQuoteRequest"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ShippingQuote"
                }
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    },
    "/ship-order": {
      "put": {
        "tags": [
          "shipping"
        ],
        "summary": "Ship Order",
        "operationId": "ship_order_ship_order_put",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ShipOrderRequest"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ShipOrderResponse"
                }
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    },
    "/list-recommendation": {
      "get": {
        "tags": [
          "recommendation"
        ],
        "summary": "Listrecommendations",
        "operationId": "listRecommendations_list_recommendation_get",
        "parameters": [
          {
            "required": false,
            "schema": {
              "title": "Product Id",
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "name": "product_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/RecommendationResponse"
                }
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    },
    "/get-product": {
      "get": {
        "tags": [
          "productcatalog"
        ],
        "summary": "Get Product",
        "operationId": "get_product_get_product_get",
        "parameters": [
          {
            "required": true,
            "schema": {
              "title": "Product Id",
              "type": "string"
            },
            "name": "product_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Product"
                }
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    },
    "/get-products": {
      "get": {
        "tags": [
          "productcatalog"
        ],
        "summary": "Get Products",
        "operationId": "get_products_get_products_get",
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "title": "Response Get Products Get Products Get",
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/Product"
                  }
                }
              }
            }
          }
        }
      }
    },
    "/search-product": {
      "get": {
        "tags": [
          "productcatalog"
        ],
        "summary": "Search Product",
        "operationId": "search_product_search_product_get",
        "parameters": [
          {
            "required": true,
            "schema": {
              "title": "Query",
              "type": "string"
            },
            "name": "query",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Response",
            "content": {
              "application/json": {
                "schema": {
                  "title": "Response Search Product Search Product Get",
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/Product"
                  }
                }
              }
            }
          },
          "422": {
            "description": "Validation Error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HTTPValidationError"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "Ad": {
        "title": "Ad",
        "required": [
          "redirect_url",
          "text"
        ],
        "type": "object",
        "properties": {
          "redirect_url": {
            "title": "Redirect Url",
            "type": "string"
          },
          "text": {
            "title": "Text",
            "type": "string"
          }
        }
      },
      "Address": {
        "title": "Address",
        "required": [
          "street_address",
          "city",
          "state",
          "country",
          "zip_code"
        ],
        "type": "object",
        "properties": {
          "street_address": {
            "title": "Street Address",
            "type": "string"
          },
          "city": {
            "title": "City",
            "type": "string"
          },
          "state": {
            "title": "State",
            "type": "string"
          },
          "country": {
            "title": "Country",
            "type": "string"
          },
          "zip_code": {
            "title": "Zip Code",
            "type": "string"
          }
        }
      },
      "Cart": {
        "title": "Cart",
        "required": [
          "user_id",
          "items"
        ],
        "type": "object",
        "properties": {
          "user_id": {
            "title": "User Id",
            "type": "string"
          },
          "items": {
            "title": "Items",
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/CartItem"
            }
          }
        }
      },
      "CartItem": {
        "title": "CartItem",
        "required": [
          "product_id",
          "quantity"
        ],
        "type": "object",
        "properties": {
          "product_id": {
            "title": "Product Id",
            "type": "string"
          },
          "quantity": {
            "title": "Quantity",
            "type": "integer"
          }
        }
      },
      "CreditCardInfo": {
        "title": "CreditCardInfo",
        "required": [
          "credit_card_number",
          "credit_card_cvv",
          "credit_card_expiration_year",
          "credit_card_expiration_month"
        ],
        "type": "object",
        "properties": {
          "credit_card_number": {
            "title": "Credit Card Number",
            "type": "string"
          },
          "credit_card_cvv": {
            "title": "Credit Card Cvv",
            "type": "integer"
          },
          "credit_card_expiration_year": {
            "title": "Credit Card Expiration Year",
            "type": "integer"
          },
          "credit_card_expiration_month": {
            "title": "Credit Card Expiration Month",
            "type": "integer"
          }
        }
      },
      "GetQuoteRequest": {
        "title": "GetQuoteRequest",
        "required": [
          "address",
          "items"
        ],
        "type": "object",
        "properties": {
          "address": {
            "$ref": "#/components/schemas/Address"
          },
          "items": {
            "title": "Items",
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/CartItem"
            }
          }
        }
      },
      "HTTPValidationError": {
        "title": "HTTPValidationError",
        "type": "object",
        "properties": {
          "detail": {
            "title": "Detail",
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/ValidationError"
            }
          }
        }
      },
      "Money": {
        "title": "Money",
        "required": [
          "currency_code",
          "units",
          "nanos"
        ],
        "type": "object",
        "properties": {
          "currency_code": {
            "title": "Currency Code",
            "type": "string"
          },
          "units": {
            "title": "Units",
            "type": "integer"
          },
          "nanos": {
            "title": "Nanos",
            "type": "integer"
          }
        }
      },
      "Order": {
        "title": "Order",
        "required": [
          "user_id",
          "user_currency",
          "address",
          "email",
          "credit_card"
        ],
        "type": "object",
        "properties": {
          "user_id": {
            "title": "User Id",
            "type": "string"
          },
          "user_currency": {
            "title": "User Currency",
            "type": "string"
          },
          "address": {
            "$ref": "#/components/schemas/Address"
          },
          "email": {
            "title": "Email",
            "type": "string"
          },
          "credit_card": {
            "$ref": "#/components/schemas/CreditCardInfo"
          }
        }
      },
      "OrderConfirmation": {
        "title": "OrderConfirmation",
        "required": [
          "email",
          "order"
        ],
        "type": "object",
        "properties": {
          "email": {
            "title": "Email",
            "type": "string"
          },
          "order": {
            "$ref": "#/components/schemas/OrderResult"
          }
        }
      },
      "OrderItem": {
        "title": "OrderItem",
        "required": [
          "item",
          "cost"
        ],
        "type": "object",
        "properties": {
          "item": {
            "$ref": "#/components/schemas/CartItem"
          },
          "cost": {
            "$ref": "#/components/schemas/Money"
          }
        }
      },
      "OrderResult": {
        "title": "OrderResult",
        "required": [
          "order_id",
          "shipping_tracking_id",
          "shipping_cost",
          "shipping_address",
          "items"
        ],
        "type": "object",
        "properties": {
          "order_id": {
            "title": "Order Id",
            "type": "string"
          },
          "shipping_tracking_id": {
            "title": "Shipping Tracking Id",
            "type": "string"
          },
          "shipping_cost": {
            "$ref": "#/components/schemas/Money"
          },
          "shipping_address": {
            "$ref": "#/components/schemas/Address"
          },
          "items": {
            "title": "Items",
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/OrderItem"
            }
          }
        }
      },
      "Payment": {
        "title": "Payment",
        "required": [
          "amount",
          "credit_card"
        ],
        "type": "object",
        "properties": {
          "amount": {
            "$ref": "#/components/schemas/Money"
          },
          "credit_card": {
            "$ref": "#/components/schemas/CreditCardInfo"
          }
        }
      },
      "PaymentResponse": {
        "title": "PaymentResponse",
        "required": [
          "transaction_id"
        ],
        "type": "object",
        "properties": {
          "transaction_id": {
            "title": "Transaction Id",
            "type": "string"
          }
        }
      },
      "Product": {
        "title": "Product",
        "required": [
          "id",
          "name",
          "description",
          "picture",
          "price_usd",
          "categories"
        ],
        "type": "object",
        "properties": {
          "id": {
            "title": "Id",
            "type": "string"
          },
          "name": {
            "title": "Name",
            "type": "string"
          },
          "description": {
            "title": "Description",
            "type": "string"
          },
          "picture": {
            "title": "Picture",
            "type": "string"
          },
          "price_usd": {
            "$ref": "#/components/schemas/Money"
          },
          "categories": {
            "title": "Categories",
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      },
      "RecommendationResponse": {
        "title": "RecommendationResponse",
        "required": [
          "product_id"
        ],
        "type": "object",
        "properties": {
          "product_id": {
            "title": "Product Id",
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      },
      "ShipOrderRequest": {
        "title": "ShipOrderRequest",
        "required": [
          "address",
          "items"
        ],
        "type": "object",
        "properties": {
          "address": {
            "$ref": "#/components/schemas/Address"
          },
          "items": {
            "title": "Items",
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/CartItem"
            }
          }
        }
      },
      "ShipOrderResponse": {
        "title": "ShipOrderResponse",
        "required": [
          "tracking_id"
        ],
        "type": "object",
        "properties": {
          "tracking_id": {
            "title": "Tracking Id",
            "type": "string"
          }
        }
      },
      "ShippingQuote": {
        "title": "ShippingQuote",
        "required": [
          "cost_usd"
        ],
        "type": "object",
        "properties": {
          "cost_usd": {
            "$ref": "#/components/schemas/Money"
          }
        }
      },
      "ValidationError": {
        "title": "ValidationError",
        "required": [
          "loc",
          "msg",
          "type"
        ],
        "type": "object",
        "properties": {
          "loc": {
            "title": "Location",
            "type": "array",
            "items": {
              "anyOf": [
                {
                  "type": "string"
                },
                {
                  "type": "integer"
                }
              ]
            }
          },
          "msg": {
            "title": "Message",
            "type": "string"
          },
          "type": {
            "title": "Error Type",
            "type": "string"
          }
        }
      }
    }
  },
  "tags": [
    {
      "name": "ad",
      "description": "**AdService operation**"
    },
    {
      "name": "cart",
      "description": "**Cartservice operation**"
    },
    {
      "name": "checkout",
      "description": "**Checkout Service operation**"
    },
    {
      "name": "payment",
      "description": "**Payment Service operation**"
    },
    {
      "name": "shipping",
      "description": "**Shipping Service operation**"
    },
    {
      "name": "confirmation",
      "description": "**Email Service operation**"
    },
    {
      "name": "recommendation",
      "description": "**RecommendationService operation**"
    },
    {
      "name": "productcatalog",
      "description": "**Product Catalog Service operation**"
    }
  ]
}
  ```
</details>
