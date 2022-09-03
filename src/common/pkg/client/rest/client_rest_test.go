package thrift

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	productId1       = "LS4PSXUNUM"
	productPicture   = "/static/img/products/salt-and-pepper-shakers.jpg"
	productQuery     = "kitchen"
	sendConfirmation = "send-confirmation"
)

var client *RestClient
var tscope TestScope

func TestMain(m *testing.M) {
	tscope.Set(CartService)
	client = NewRestClient()
	client.ProductCatalogService = "productcatalogservice:60000"
	client.CartService = "cartservice:60000"
	client.Emailservice = "emailservice:60000"
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

// ---------- CartService -----------------------
func TestAddCart(t *testing.T) {
	err := client.AddCart("12345", productId1, "5")
	assert.Nil(t, err)
}

func TestGetCart(t *testing.T) {
	data, err := client.GetCart("12345")
	assert.Nil(t, err)
	_ = data
}

func TestEmptyCart(t *testing.T) {
	err := client.EmptyCart("12345")
	assert.Nil(t, err)
}

// ---------- CartService -----------------------
func TestSendEmailConfirmation(t *testing.T) {
	// tscope.testScope(t, Email)
	err := client.SendOrderConfirmation(orderResult)
	assert.Nil(t, err)
}
