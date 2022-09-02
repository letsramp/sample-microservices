package thrift

import (
	"context"
	"fmt"

	pb "github.com/pellepedro/gcp-microservices/src/common/idl/grpc"
	"github.com/pellepedro/gcp-microservices/src/common/idl/thrift/demo"
	thrift "github.com/pellepedro/gcp-microservices/src/common/pkg/thrift"
	"github.com/pkg/errors"
	_ "github.com/sirupsen/logrus"
)

// ------- ProductCatalogService
func (tc *Client) ListProducts() ([]*pb.Product, error) {
	opt := thrift.NewDefaultOption()
	client, transport, err := thrift.NewThriftClient(tc.ProductCatalogService, opt)
	if err != nil {
		return nil, fmt.Errorf("failed creating productcatalogservice client")
	}
	err = transport.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	c := demo.NewProductCatalogServiceClient(client)
	ctx := context.TODO()
	if products, err := c.ListProducts(ctx); err != nil {
		return nil, err
	} else {
		result := make([]*pb.Product, 0, len(products))
		for _, p := range products {
			product := &pb.Product{
				Id:          p.ID,
				Categories:  p.Categories,
				Description: p.Description,
				Name:        p.Name,
				Picture:     p.Picture,
				PriceUsd: &pb.Money{
					CurrencyCode: p.PriceUsd.CurrencyCode,
					Units:        p.PriceUsd.Units,
					Nanos:        p.PriceUsd.Nanos,
				},
			}
			result = append(result, product)
		}
		return result, nil
	}
}
func (tc *Client) GetProduct(product_id string) (*pb.Product, error) {
	opt := thrift.NewDefaultOption()
	client, transport, err := thrift.NewThriftClient(tc.ProductCatalogService, opt)
	if err != nil {
		return nil, fmt.Errorf("failed creating productcatalogservice client")
	}

	err = transport.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	c := demo.NewProductCatalogServiceClient(client)
	ctx := context.TODO()
	if p, err := c.GetProduct(ctx, product_id); err != nil {
		return nil, err
	} else {
		product := &pb.Product{
			Id:          p.ID,
			Categories:  p.Categories,
			Description: p.Description,
			Name:        p.Name,
			Picture:     p.Picture,
			PriceUsd: &pb.Money{
				CurrencyCode: p.PriceUsd.CurrencyCode,
				Units:        p.PriceUsd.Units,
				Nanos:        p.PriceUsd.Nanos,
			},
		}
		return product, nil
	}
}
func (tc *Client) SearchProducts(query string) ([]*pb.Product, error) {
	opt := thrift.NewDefaultOption()
	client, transport, err := thrift.NewThriftClient(tc.ProductCatalogService, opt)
	if err != nil {
		return nil, fmt.Errorf("failed creating productcatalogservice client")
	}

	err = transport.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	c := demo.NewProductCatalogServiceClient(client)
	ctx := context.TODO()
	if products, err := c.SearchProducts(ctx, query); err != nil {
		return nil, err
	} else {
		result := make([]*pb.Product, 0, len(products))
		for _, p := range products {
			product := &pb.Product{
				Id:          p.ID,
				Categories:  p.Categories,
				Description: p.Description,
				Name:        p.Name,
				Picture:     p.Picture,
				PriceUsd: &pb.Money{
					CurrencyCode: p.PriceUsd.CurrencyCode,
					Units:        p.PriceUsd.Units,
					Nanos:        p.PriceUsd.Nanos,
				},
			}
			result = append(result, product)
		}
		return result, nil
	}
}

// ------- RecommendationService
func (tc *Client) ListRecommendations(cart []string) ([]*pb.Product, error) {
	opt := thrift.NewDefaultOption()
	client, transport, err := thrift.NewThriftClient(tc.RecommendationService, opt)
	if err != nil {
		return nil, fmt.Errorf("failed creating recommendationservice client")
	}

	err = transport.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	c := demo.NewRecommendationServiceClient(client)
	ctx := context.TODO()

	productIds, err2 := c.ListRecommendations(ctx, cart)
	if err2 != nil {
		return nil, err2
	}
	out := make([]*pb.Product, len(productIds))
	for i, v := range productIds {
		p, err := tc.GetProduct(v)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get recommended product info (%s)", v)
		}
		out[i] = p
	}
	if len(out) > 4 {
		out = out[:4] // take only first four to fit the UI
	}
	return out, nil
}

// ------- AD Service
func (tc *Client) GetAds(adReq *pb.AdRequest) ([]*pb.Ad, error) {
	opt := thrift.NewDefaultOption()
	client, transport, err := thrift.NewThriftClient(tc.Adservice, opt)
	if err != nil {
		return nil, fmt.Errorf("failed creating adservice client")
	}

	err = transport.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	c := demo.NewAdServiceClient(client)
	ctx := context.TODO()
	ads, err := c.GetAds(ctx, adReq.ContextKeys)
	if err != nil {
		return nil, fmt.Errorf("failed creating adservice client")
	}
	result := make([]*pb.Ad, len(ads), len(ads))
	for _, ad := range ads {
		a := pb.Ad{RedirectUrl: ad.RedirectURL, Text: ad.Text}
		result = append(result, &a)
	}
	return result, nil
}

// ------- Cart Service
func (tc *Client) AddCart(userID, productID string, quantity int32) error {
	opt := thrift.NewDefaultOption()
	client, transport, err := thrift.NewThriftClient(tc.CartService, opt)
	if err != nil {
		return fmt.Errorf("failed creating cartservice client")
	}
	err = transport.Open()
	if err != nil {
		return fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	c := demo.NewCartServiceClient(client)
	return c.AddItem(context.Background(), userID, &demo.CartItem{ProductID: productID, Quantity: quantity})
}
func (tc *Client) EmptyCart(userID string) error {
	opt := thrift.NewDefaultOption()
	client, transport, err := thrift.NewThriftClient(tc.CartService, opt)
	if err != nil {
		return fmt.Errorf("failed creating cartservice client")
	}
	err = transport.Open()
	if err != nil {
		return fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	c := demo.NewCartServiceClient(client)
	return c.EmptyCart(context.Background(), userID)
}
func (tc *Client) GetCart(userID string) ([]*pb.CartItem, error) {
	opt := thrift.NewDefaultOption()
	client, transport, err := thrift.NewThriftClient(tc.CartService, opt)
	if err != nil {
		return nil, fmt.Errorf("failed creating cartservice client")
	}
	err = transport.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	c := demo.NewCartServiceClient(client)
	cart, err := c.GetCart(context.Background(), userID)
	if err != nil {
		return nil, fmt.Errorf("failed calling thrift API GetCart()")
	}
	cartItems := make([]*pb.CartItem, 0, len(cart.Items))
	for _, item := range cart.Items {
		c := pb.CartItem{ProductId: item.GetProductID(), Quantity: item.GetQuantity()}
		cartItems = append(cartItems, &c)
	}
	return cartItems, nil
}

// ------- Payment Service
// ------- Shipping Service
func (tc *Client) SendOrderConfirmation(ctx context.Context, address string, items *demo.OrderResult_) error {
	opt := thrift.NewDefaultOption()
	opt.HttpTransport = false
	opt.Buffered = true
	opt.Secure = false
	opt.Framed = false

	client, transport, err := thrift.NewThriftClient(tc.Emailservice, opt)
	if err != nil {
		return fmt.Errorf("failed email client")
	}

	err = transport.Open()
	if err != nil {
		return fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	c := demo.NewEmailServiceClient(client)
	if err := c.SendOrderConfirmation(ctx, address, items); err != nil {
		return err
	}
	return nil
}

// ------- Email
func (tc *Client) ShipOrder(ctx context.Context, address *demo.Address, items []*demo.CartItem) error {
	opt := thrift.NewDefaultOption()
	client, transport, err := thrift.NewThriftClient(tc.Shippingservice, opt)
	if err != nil {
		return fmt.Errorf("failed creating ShippingService client")
	}
	err = transport.Open()
	if err != nil {
		return fmt.Errorf("failed to open transport %w", err)
	}
	defer transport.Close()
	_ = client
	return nil
}

// ------- Concurrency
