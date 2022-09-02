package thrift

import "net/http"

var (
	ThriftPort = 50000
)

type Client struct {
	restClient            *http.Client
	thriftPort            int
	restPort              int
	secureRest            bool
	secureThrift          bool
	ProductCatalogService string
	RecommendationService string
	CartService           string
	Shippingservice       string
	Currencyservice       string
	Checkoutservice       string
	Adservice             string
	Paymentservice        string
	Emailservice          string
}

func NewClient() *Client {
	return &Client{}
}
