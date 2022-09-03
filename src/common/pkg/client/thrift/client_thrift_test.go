package thrift

import (
	"context"
	"os"
	"testing"

	pb "github.com/pellepedro/gcp-microservices/src/common/idl/grpc"
	"github.com/pellepedro/gcp-microservices/src/common/idl/thrift/demo"
	"github.com/stretchr/testify/assert"
)

const (
	errotStartServer = 1
)

var client *Client
var tscope TestScope

func TestMain(m *testing.M) {
	tscope.Set(Email)
	client = NewClient()
	client.ProductCatalogService = "productcatalogservice:5000"
	client.CartService = "cartservice:50000"
	client.Emailservice = "emailService:50000"
	retCode := m.Run()
	os.Exit(retCode)
}

func TestListProducts(t *testing.T) {
	tscope.testScope(t, ProductCatalog)
	p, err := client.ListProducts()
	assert.Nil(t, err)
	_ = p
}

func TestSearchProducts(t *testing.T) {
	tscope.testScope(t, ProductCatalog)
	p, err := client.SearchProducts("summer")
	assert.Nil(t, err)
	_ = p
}

func TestGetProduct(t *testing.T) {
	tscope.testScope(t, ProductCatalog)
	p, err := client.GetProduct("OLJCESPC7Z")
	assert.Nil(t, err)
	_ = p
}

func TestListRecomendations(t *testing.T) {
	tscope.testScope(t, Recommendation)
	p, err := client.ListRecommendations([]string{"OLJCESPC7Z"})
	assert.Nil(t, err)
	_ = p
}

func TestGetAds(t *testing.T) {
	tscope.testScope(t, Ad)
	p, err := client.GetAds(&pb.AdRequest{ContextKeys: []string{"kitchen"}})
	assert.Nil(t, err)
	_ = p
}

func TestAddCart(t *testing.T) {
	tscope.testScope(t, CartService)
	err := client.AddCart("abcdef", "GRAVLAX", 5)
	assert.Nil(t, err)
}

func TestGetCart(t *testing.T) {
	tscope.testScope(t, CartService)
	items, err := client.GetCart("abcdef")
	assert.Nil(t, err)
	_ = items
}

func TestEmptyCart(t *testing.T) {
	tscope.testScope(t, CartService)
	err := client.EmptyCart("abcdef")
	assert.Nil(t, err)
}

func TestSendEmailConfirmation(t *testing.T) {
	tscope.testScope(t, Email)
	email := "user@mail.com"
	orderResult := demo.OrderResult_{
		OrderID:            "order-id-123",
		ShippingTrackingID: "tracking-id-000",
		ShippingCost: &demo.Money{
			CurrencyCode: "USD",
			Units:        10,
			Nanos:        10,
		},
		ShippingAddress: &demo.Address{
			StreetAddress: "Main Street 101",
			City:          " San Fransisco",
			State:         "CA",
			Country:       "US",
			ZipCode:       95000,
		},
		Items: []*demo.OrderItem{
			{
				Item: &demo.CartItem{ProductID: "OLJCESPC7Z", Quantity: 5},
				Cost: &demo.Money{
					CurrencyCode: "USD",
					Units:        95,
					Nanos:        500000000,
				},
			},
			{
				Item: &demo.CartItem{ProductID: "66VCHSJNUP", Quantity: 2},
				Cost: &demo.Money{
					CurrencyCode: "USD",
					Units:        36,
					Nanos:        200000000,
				},
			},
		},
	}
	if err := client.SendOrderConfirmation(context.Background(), email, &orderResult); err != nil {
		t.Logf("Failed to call SendOrderConfirmation : %v", err)
		t.FailNow()
	}
}
