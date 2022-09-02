package thrift

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	productId        = "LS4PSXUNUM"
	productPicture   = "/static/img/products/salt-and-pepper-shakers.jpg"
	productQuery     = "kitchen"
	sendConfirmation = "send-confirmation"
)

var client *RestClient
var tscope TestScope

func TestMain(m *testing.M) {
	tscope.Set(Email)
	client = NewRestClient()
	client.ProductCatalogService = "productcatalogservice:60000"
	client.Emailservice = "dev:60000"
	retCode := m.Run()
	os.Exit(retCode)
}

func TestRestListProducts(t *testing.T) {
	tscope.testScope(t, ProductCatalog)
	products, err := client.ListProducts()
	assert.Nil(t, err)
	assert.Equal(t, 9, len(products))
}

func TestRestSearchProducts(t *testing.T) {
	tscope.testScope(t, ProductCatalog)
	products, err := client.SearchProducts(productQuery)
	assert.Nil(t, err)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
}

func TestSendEmailConfirmation(t *testing.T) {
	// tscope.testScope(t, Email)
	err := client.SendOrderConfirmation(orderResult)
	assert.Nil(t, err)
}
