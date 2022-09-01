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

## Test the online butique experience with browser.

## Testing API interaction with a GRPC client.
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

## Use skyramp to deploy Online Boutique with rest
TBD:

## Use skyramp to deploy Online Boutique with thrift
TBD:



## Exploring API and Mocking Services

### REST
```
TBD
```
### GRPC
```
TBD
```
### Thrift
```
TBD
```



Appendix: Example exploringAPI calls with curl and jq

```
curl -k http://product-catalog-service-port60000.demo.skyramp.test/list-products | jq .
curl -k http://product-catalog-service-port60000.demo.skyramp.test/search-products\?\query\=kitchen | jq .
curl -k http://product-catalog-service-port60000.demo.skyramp.test/get-product\?\product_id\=LS4PSXUNUM | jq .
```

#### Response
list-products
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

**Search Products**
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

Get Product
```
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
}
```

