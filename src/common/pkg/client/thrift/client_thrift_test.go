package thrift

import (
	"context"
	"fmt"
	"os"
	"testing"

	pb "github.com/pellepedro/gcp-microservices/src/common/idl/grpc"
	"github.com/pellepedro/gcp-microservices/src/common/idl/thrift/demo"
	"github.com/stretchr/testify/assert"
)

const (
	errotStartServer = 1
)

var (
	defaultThriftPort     = 50000
	productcatalogservice = fmt.Sprintf("127.0.0.1:%d", defaultThriftPort)
	emailService          = fmt.Sprintf("%s:%d", "dev", defaultThriftPort)
	//"productcatalogservice:5000"
)

var client *Client
var tscope TestScope

func TestMain(m *testing.M) {
	// server, err := startProductCatalogMock(productCatalogServicePort)
	// if err != nil {
	// 	os.Exit(errotStartServer)
	// }
	// go server.Serve()
	tscope.Set(Email)
	// time.Sleep(5 * time.Second)
	client = NewClient()
	client.ProductCatalogService = productcatalogservice
	client.Emailservice = emailService
	retCode := m.Run()
	// server.Stop()

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
