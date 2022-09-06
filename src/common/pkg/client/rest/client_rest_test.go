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
	tscope.Set(ProductCatalog)
	client = NewRestClient()
	client.ProductCatalogService = "productcatalogservice:60000"
	client.ProductCatalogService = "productcatalogservice:6000"
	client.CartService = "cartservice:60000"
	client.Emailservice = "emailservice:60000"
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetProduct(t *testing.T) {
	const productId = "OLJCESPC7Z"
	tscope.testScope(t, ProductCatalog)
	product, err := client.GetProduct(productId)
	assert.Nil(t, err)
	assert.Equal(t, productId, product.Id)
}
func TestGetProducts(t *testing.T) {
	// tscope.testScope(t, ProductCatalog)
	products, err := client.GetProducts()
	assert.Nil(t, err)
	assert.Equal(t, 9, len(products))
}

func TestRestSearchProducts(t *testing.T) {
	// tscope.testScope(t, ProductCatalog)
	products, err := client.SearchProducts(productQuery)
	assert.Nil(t, err)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
}

// --------- RecommendationService
func TestListRecommendation(t *testing.T) {
	var (
		cart_id = []string{"LS4PSXUNUM"}
	)
	client.RecommendationService = "recommendationservice:60000"
	client.ProductCatalogService = "productcatalogservice:60000"
	r, err := client.ListRecommendations(cart_id)
	assert.Nil(t, err)
	_ = r
}

// ---------- CartService -----------------------
func TestAddCart(t *testing.T) {
	client.CartService = "cartservice:60000"
	err := client.AddCart("12345", productId1, 5)
	assert.Nil(t, err)
}

func TestGetCart(t *testing.T) {
	client.CartService = "cartservice:60000"
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
