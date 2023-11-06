- [Introduction](#introduction)
  - [Use Cases](#use-cases)
  - [Implementation](#implementation)
- [Usage](#usage)
  - [1. Install pre-requisite tools](#1-install-pre-requisite-tools)
  - [2. Clone sample services](#2-clone-sample-services)
  - [3. Build and Push Images (Optional)](#3-build-and-push-images-optional)
    - [Build images from source](#build-images-from-source)
    - [Push services to container registry](#push-services-to-container-registry)
  - [4. Deploy](#4-deploy)
  - [5. Interacting with the microservices](#5-interacting-with-the-microservices)
    - [Add to cart using gRPC API](#add-to-cart-using-grpc-api)
    - [Add to cart using REST API](#add-to-cart-using-rest-api)
    - [Add to cart using Thrift API](#add-to-cart-using-thrift-api)
    - [Verify contents of cart](#verify-contents-of-cart)
  - [6. Exploring the pre-generated tests](#6-exploring-the-pre-generated-tests)
    - [Running the tests from the Skyramp CLI](#running-the-tests-from-the-skyramp-cli)
    - [Exploring the demo tests](#6-exploring-the-demo-tests)
- [Contributing Code](#contributing-code)


# Introduction

This is a demo project based on [GCP Online Boutique](https://github.com/GoogleCloudPlatform/microservices-demo) with added support for REST and Thrift APIs. It is designed to be used during cloud-native development for testing across different APIs–REST, gRPC, and Thrift. It is extensible in support of future functionality (e.g. GraphQL).

## Use Cases

1. Use as a test project for developing cloud-native dev tools.
2. Learn how to implement Thrift (and REST) support for an existing project.

Implemented by the [Skyramp](https://skyramp.dev) team.

## Implementation
For each microservice in the repo, REST and Thrift implementations open up separate ports to listen to the corresponding traffic. Implementations are in the `rest.go` and `thrift.go` files in the top-level directory for each service. 

Default ports for REST and Thrift are hardcoded in the `main.go` file in each service directory. For example, the `carts` microservice has the following ports defined:

```
const (
	serverCrt         = "server.crt"
	serverKey         = "server.key"
	defaultThriftPort = "50000"
	defaultRestPort   = "60000"
)
```
# Usage

## 1. Install pre-requisite tools
- docker
- curl (required for REST)
- jq (optional)
- kubectl (optional)


## 2. Clone sample services
```
git clone https://github.com/letsramp/sample-microservices.git
cd sample-microservices/src
```

## 3. Build and Push Images (Optional)

Images for sample-microservices are already available in a publically accessible registry. To use your own registry, follow these steps.

### Build images from source 
```
docker compose build
```

### Push services to container registry
Update the host path/s in the docke-compose.yaml file to point to your own docker registry. Push images by running:

```
docker compose push
```

## 4. Deploy

Deploy Online Boutique by running the following command:

```
docker compose up
```
Now, familiarize yourself with the application by navigating to http://127.0.0.1:8080 on your browser.

<br/><br/>
<img width="1728" alt="Online Boutique" src="https://user-images.githubusercontent.com/1672645/217123094-00e455d5-316d-44f3-8e80-b56da07b668d.png">
<br/><br/>

> 📝 **NOTE:** If you are using Docker Desktop, you can look at the logs for `cart service` to see that we've opened up separate ports for gRPC(7070), REST(60000), and Thrift(50000) traffice to the service. 


<img width="1127" alt="carts-ports" src="https://user-images.githubusercontent.com/1672645/217123495-9a516fe5-3bf1-4e97-bd46-2270ae130df6.png">


## 5. Interacting with the microservices

To demonstrate the REST and Thrift implementations, we've created simple clients (`src/clients`) for the "add to cart" scenario for each of the APIs. The client code is accessible from the `clients` container in the cluster.

To connect to the `clients` container, run the following command in your terminal:

```
docker compose exec -it clients ash
```

Optionally, you can download and install [Docker Desktop](https://www.docker.com/products/docker-desktop/) to follow along.

If you are using Docker Desktop, click on the `clients_1` container and go to the CLI from there.

<img width="1122" alt="CLI for the clients container" src="https://user-images.githubusercontent.com/1672645/217331154-3be0e78b-3c22-43c3-bdb2-ac5b5365c50b.png">


### Add to cart using gRPC API

1. In the `clients` container, navigate to the grpc/golang folder and download the required `go` modules.
```
cd /grpc/golang
go mod download
```

2. In the `addCart.go` file, you will notice that we open a connection to port 7070 of `cart` microservice which listens to grpc traffic. 

```
conn, err := grpc.Dial("cartservice:7070", grpc.WithInsecure())
```

3. Run the code in the file to add an item to the cart.

```
go run ./cmd/cart
```
Expected result
```
Successfully added the item to the cart.
```

Having seen how to successfully add an item to the cart with a gRPC client calling the gRPC endpoint, we can see how to do the same through the REST and Thrift endpoints.

### Add to cart using REST API

Since cURL is already installed in the `clients` container, you can issue a request directly from the CLI of the container to add an item to the cart.

```
curl -X 'POST' 'http://cartservice:60000/cart/user_id/abcde' \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -d '{
  "product_id": "OLJCESPC7Z",
  "quantity": 1
}'
```

Again, notice that the request is being sent to port `60000` of `cart` microservice which listens to REST traffic.

Result

```
{"success":"200 OK"}
```

### Add to cart using Thrift API

1. In the `clients` container, navigate to the thrift/golang folder and download the required `go` modules for the Thrift client.

Setup
```
cd /thrift/golang
go mod download
```

2. In the `addCart.go` file, you will notice that we open a connection to port 50000 of `cart` microservice which listens to Thrift traffic. 

```
clientAddr := "cartservice:50000"
```

3. Now, run the code in the file to add an item to the cart.

```
go run ./cmd/cart
```

Result:

```
"Successfully added the item to the cart."
```


### Verify contents of cart

You can now fetch the contents of the cart using REST API to see that a total of 3 items were added (one using each supported API).

```
curl -X 'GET' \
  'http://cartservice:60000/cart/user_id/abcde' \
   -H 'accept: application/json'
```

Result:

```
{"user_id":"abcde","items":[{"product_id":"OLJCESPC7Z","quantity":3}]}
```

## 6. Exploring the demo tests

This repo contains a number of test scenarios used for demo purposes. These test scenarios are located under the `skyramp` folder in `grpc-demo`, `rest-demo`, and `thrift-demo`. The test scenarios can be run using the Skyramp client. Visit the [Skyramp Docs](https://skyramp.dev/docs/get-started/install-client/) for instructions on installing the Skyramp client.

### Running the tests from the Skyramp CLI

Once the microservices have been deployed to a cluster, you can run the demo tests with the Skyramp CLI. See the [Walkthrough](https://skyramp.dev/docs/get-started/walkthrough/) in the Skyramp Docs for one method of deployment with Skyramp Deployer.

To run the tests for the REST demo, as an example, navigate to the `rest-demo` folder in a terminal:
```bash
cd sample-microservices/skyramp/rest-demo
```
You can see the existing test description files under `tests`:
```bash
checkout-test.yaml
test-load.yaml       
test-trace-bOl4.yaml
```
Running the `checkout-test` test with Skyramp Tester...
```bash
skyramp tester checkout-test -n test-rest-demo
```
Will produce output similar to this:
```bash
Starting tests
Tester finished
Test [REST] Checkout system testcase------
 [Status: finished] [Started at: 2023-11-04 16:41:21 PDT] [End: 2023-11-04 16:41:22 PDT] [Duration: 1s]
  - pattern0.scenario1
    [Status: finished] [Started at: 2023-11-04 16:41:22 PDT] [Duration: 0s]
  - pattern0.scenario1.0.addCartRequest
    [Status: finished] [Started at: 2023-11-04 16:41:22 PDT] [Duration: 0s]
    Executed: {"success":"200 OK"}
  - pattern0.scenario1.1.getCartRequest
    [Status: finished] [Started at: 2023-11-04 16:41:22 PDT] [Duration: 0s]
    Executed: {"user_id":"abcde","items":[{"product_id":"OLJCESPC7Z","quantity":2}]}
  - pattern0.scenario1.2.assert
    [Status: finished] [Started at: N/A]
    Assert: requests.getCartRequest.res.user_id == "abcde"
    Passed: true
  - pattern0.scenario1.3.checkoutRequest
    [Status: finished] [Started at: 2023-11-04 16:41:22 PDT] [Duration: 0s]
    Executed: {"order_id":"1cae99ba-f22e-42cd-8b4d-79160ef3ae72","shipping_tracking_id":"2c287fac-104a-4a4b-a3fc-32f1ec3a7495","shipping_cost":{"currency_code":"USD","units":10,"nanos":100},"shipping_address":{"street_address":"1600 Amp street","city":"Mountain View","state":"CA","country":"USA"},"items":[{"item":{"product_id":"OLJCESPC7Z","quantity":2}}]}
  - pattern0.scenario1.4.assert
    [Status: finished] [Started at: N/A]
    Assert: requests.checkoutRequest.res.items[0].item.product_id == "OLJCESPC7Z"
    Passed: true
```
### Creating or generating tests

Test files can be created manually or generated using the `skyramp generate` command. The test files can then be edited and customized for your specific testing requirements. For more information on creating or generating tests, see the Skyramp Docs covering the [Test Description](https://www.skyramp.dev/docs/tester/test-description/) and the specific [CLI command - tester generate](https://www.skyramp.dev/docs/reference/cli-commands/tester/generate/).

Notice from the Skyramp docs that there are a number of possible options for generating tests. Tests can be generated using an API schema definition with [OpenAPI](https://www.openapis.org) or [Protocol Buffers](https://protobuf.dev) as well as with telemetry trace data from an observability provider like [Pixie](https://px.dev).

You will see in the `skyramp/rest-demo` folder in this repo that there are `openapi` and `trace` folders containing example input files for generating tests.  The `test-trace-bOl4.yaml` demo test description was generated from trace data. Here are a few examples to illustrate test generation options with various inputs:

```bash
skyramp tester generate \
--trace-file trace/trace.json
```
```bash
skyramp tester generate \
--protocol openapi \
--api-schema openapi/demo.yaml \
--alias test-rest-demo \
--port 60000
```
```bash
skyramp tester generate \
--telemetry-provider pixie \
--cluster-id cluster-id \
--namespace namespace \
--start-time "-5m" 
```
```bash
skyramp tester generate grpc \
--api-schema file.proto \
--alias namespace \
--port port \
--service service
```

# Contributing Code

PRs are welcome!  

This project follows the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/main/code-of-conduct.md) .

<br>
