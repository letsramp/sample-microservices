// Copyright 2022 Skyramp, Inc.
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	pb "checkoutservice/genproto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	CART_SERVICE_ADDR            = "CART_SERVICE_ADDR"
	CURRENCY_SERVICE_ADDR        = "CURRENCY_SERVICE_ADDR"
	EMAIL_SERVICE_ADDR           = "EMAIL_SERVICE_ADDR"
	PAYMENT_SERVICE_ADDR         = "PAYMENT_SERVICE_ADDR"
	PRODUCT_CATALOG_SERVICE_ADDR = "PRODUCT_CATALOG_SERVICE_ADDR"
	SHIPPING_SERVICE_ADDR        = "SHIPPING_SERVICE_ADDR"
)

var defaultServiceName = map[string]string{
	PRODUCT_CATALOG_SERVICE_ADDR: "product-catalog-service",
	CART_SERVICE_ADDR:            "cart-service",
	SHIPPING_SERVICE_ADDR:        "shipping-service",
	CURRENCY_SERVICE_ADDR:        "currency-service",
	PAYMENT_SERVICE_ADDR:         "payment-service",
	EMAIL_SERVICE_ADDR:           "email-service-addr",
}

var client = NewRestClient()

func getService(serviceEnv string, port int) string {
	serviceHost := defaultServiceName[serviceEnv]
	if os.Getenv(serviceEnv) != "" {
		serviceUrlEnv := os.Getenv(serviceEnv)
		serviceHost = strings.Split(serviceUrlEnv, ":")[0]
	}
	service := fmt.Sprintf("%s:%d", serviceHost, port)
	return service
}

func init() {
	client.CartService = getService(CART_SERVICE_ADDR, 60000)
	client.ProductCatalogService = getService(PRODUCT_CATALOG_SERVICE_ADDR, 60000)
	client.Currencyservice = getService(CURRENCY_SERVICE_ADDR, 60000)
	client.Paymentservice = getService(PAYMENT_SERVICE_ADDR, 60000)
	client.Emailservice = getService(EMAIL_SERVICE_ADDR, 60000)
	client.Shippingservice = getService(SHIPPING_SERVICE_ADDR, 6000)
}

func checkout(c *gin.Context) {
	in, err := c.GetRawData()
	if err != nil {
		c.JSON(501, gin.H{"error": "failed to read data"})
	}
	data := strings.ReplaceAll(string(in), "zip_code", "zipCode")
	order := &pb.PlaceOrderRequest{}
	if err := json.Unmarshal([]byte(data), order); err != nil {
		c.JSON(501, gin.H{"error": "failed to parse PlaceOrderRequest"})
		return
	}

	expYear := order.CreditCard.CreditCardExpirationYear
	// if credit card is given with 2 digit, then add centuary:w
	if expYear < 100 {
		order.CreditCard.CreditCardExpirationYear = expYear + 2000
	}

	result := pb.OrderResult{
		ShippingAddress: order.Address,
		Items:           make([]*pb.OrderItem, 0),
		ShippingCost:    &pb.Money{CurrencyCode: "USD", Units: 10, Nanos: 100},
		OrderId:         uuid.NewString(),
	}

	// Get Cart
	cart, err := client.GetCart(order.UserId)
	if err != nil {
		c.JSON(501, "failed to get user cart")
		return
	}

	var orderCost float32
	// Update order with items
	for _, item := range cart.Items {
		if product, err := client.GetProduct(item.ProductId); err != nil {
			c.JSON(501, fmt.Sprintf("failed to get product[%s] from the Product Catalog Cervice, %v", item.ProductId, err))
			return
		} else {
			result.Items = append(result.Items, &pb.OrderItem{Item: item, Cost: product.PriceUsd})
			orderCost += float32(product.PriceUsd.GetUnits()) * float32(item.Quantity)
		}
	}
	chargeRequest := pb.ChargeRequest{
		Amount: &pb.Money{
			CurrencyCode: "USD",
			Units:        int64(orderCost),
			Nanos:        0,
		},
		CreditCard: order.CreditCard,
	}
	cResult, err := client.charge(chargeRequest)
	if err != nil {
		c.JSON(501, fmt.Sprintf("failed to charge credit card, %v", err))
		return
	}
	result.ShippingTrackingId = cResult.TransactionId

	c.JSON(200, result)
}

func startRest() {
	go func() {
		port := "60000"
		if os.Getenv("REST_PORT") != "" {
			port = os.Getenv("REST_PORT")
		}
		router := gin.Default()
		router.POST("/checkout", checkout)
		router.Run(fmt.Sprintf("0.0.0.0:%s", port))
	}()
}
