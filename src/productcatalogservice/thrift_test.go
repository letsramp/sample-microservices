package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/microservices-demo/src/productcatalogservice/thriftgo/demo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	product_id = "OLJCESPC7Z"
)

func TestA(t *testing.T) {
	opt := NewDefaultOption()
	opt.HttpTransport = true
	opt.HttpUrl = "/ProductCatalogService"
	startThrift("50001", opt)
	time.Sleep(5 * time.Second)
	hostPort := "127.0.0.1:50001"
	client, trans, err := NewThriftClient(hostPort, opt)
	assert.Nil(t, err)
	c := demo.NewProductCatalogServiceClient(client)
	err = trans.Open()
	assert.Nil(t, err)
	allproducts, err := c.ListProducts(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, 9, len(allproducts))
	jsonProducts, err := json.Marshal(allproducts)
	ioutil.WriteFile("products.txt", []byte(jsonProducts), 0644)
	_ = jsonProducts
	products, err := c.SearchProducts(context.TODO(), "kitchen")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
	product, err := c.GetProduct(context.TODO(), product_id)
	assert.Nil(t, err)
	_ = product

}
