package thrift

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
	demo "github.com/pellepedro/gcp-microservices/src/common/idl/thrift/demo"
	server "github.com/pellepedro/gcp-microservices/src/common/pkg/thrift"
)

type Handler struct {
}

func (h *Handler) ListProducts(ctx context.Context) (products []*demo.Product, err error) {
	err = json.Unmarshal([]byte(jsonProducts), &products)
	return
}
func (h *Handler) GetProduct(ctx context.Context, product_id string) (product *demo.Product, err error) {
	products := []*demo.Product{}
	if err = json.Unmarshal([]byte(jsonProducts), &products); err != nil {
		return nil, err
	}
	return products[0], nil
}

func (h *Handler) SearchProducts(ctx context.Context, query string) (searchResult []*demo.Product, err error) {
	products := []*demo.Product{}
	if err = json.Unmarshal([]byte(jsonProducts), &products); err != nil {
		return nil, err
	}
	for _, product := range products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(product.Description), strings.ToLower(query)) {
			searchResult = append(searchResult, product)
		}
	}
	return
}

func startProductCatalogMock(port int) (*thrift.TSimpleServer, error) {
	opt := server.NewDefaultOption()
	_ = opt
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	processor := demo.NewProductCatalogServiceProcessor(&Handler{})
	return server.NewStandardThriftServer(addr, opt, processor)
}

const jsonProducts = `[
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
       "categories": ["accessories"]
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
       "categories": ["clothing", "tops"]
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
       "categories": ["accessories"]
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
       "categories": ["footwear"]
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
       "categories": ["hair", "beauty"]
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
       "categories": ["decor", "home"]
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
       "categories": ["kitchen"]
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
       "categories": ["kitchen"]
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
       "categories": ["kitchen"]
   }
]`
