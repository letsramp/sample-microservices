package thrift

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pb "github.com/pellepedro/gcp-microservices/src/common/idl/grpc"
	"github.com/pellepedro/gcp-microservices/src/common/pkg/client/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	APPLICATION_JSON = "application/json"
)

type RestClient struct {
	restClient            *http.Client
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

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
}

func NewRestClient() *RestClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &RestClient{
		restClient: &http.Client{Transport: tr},
	}
}

func (c *RestClient) GetProduct(productID string) (*pb.Product, error) {
	url := fmt.Sprintf("http://%s/%s?product_id=%s", c.ProductCatalogService, "get-product", productID)
	log.Debugf("--- REST GetProduct for id %s", url)

	res, err := c.restClient.Get(url)
	if err != nil {
		error := fmt.Sprintf("error sending get: url [%s], error:  %v", url, err)
		return nil, fmt.Errorf(error)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var p types.Product
	if err := json.Unmarshal(out, &p); err != nil {
		error := fmt.Sprintf("error mmarshaling response: %v", err)
		return nil, fmt.Errorf(error)
	}
	proto := pb.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Picture,
		Categories:  p.Categories,
	}
	proto.PriceUsd = &pb.Money{}
	proto.PriceUsd.CurrencyCode = p.PriceUsd.CurrencyCode
	proto.PriceUsd.Units = int64(p.PriceUsd.Units)
	proto.PriceUsd.Nanos = int32(p.PriceUsd.Nanos)
	return &proto, nil
}

func (c *RestClient) GetProducts() ([]*pb.Product, error) {
	url := fmt.Sprintf("http://%s/%s", c.ProductCatalogService, "get-products")
	log.Debugf("--- REST GetProducts")

	res, err := c.restClient.Get(url)
	if err != nil {
		error := fmt.Sprintf("error sending get: url [%s], error:  %v", url, err)
		return nil, fmt.Errorf(error)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var p []types.Product
	if err := json.Unmarshal(out, &p); err != nil {
		error := fmt.Sprintf("error mmarshaling response: %v", err)
		return nil, fmt.Errorf(error)
	}
	products := make([]*pb.Product, 0)
	for _, item := range p {
		proto := pb.Product{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			Categories:  item.Categories,
		}
		proto.PriceUsd = &pb.Money{}
		proto.PriceUsd.CurrencyCode = item.PriceUsd.CurrencyCode
		proto.PriceUsd.Units = int64(item.PriceUsd.Units)
		proto.PriceUsd.Nanos = int32(item.PriceUsd.Nanos)
		products = append(products, &proto)
	}
	return products, nil
}

func (c *RestClient) SearchProducts(query string) ([]*pb.Product, error) {
	url := fmt.Sprintf("http://%s/%s?query=%s", c.ProductCatalogService, "search-products", query)
	log.Debugf("--- REST GetProducts")

	res, err := c.restClient.Get(url)
	if err != nil {
		error := fmt.Sprintf("error sending get: url [%s], error:  %v", url, err)
		return nil, fmt.Errorf(error)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var p []types.Product
	if err := json.Unmarshal(out, &p); err != nil {
		error := fmt.Sprintf("error mmarshaling response: %v", err)
		return nil, fmt.Errorf(error)
	}
	products := make([]*pb.Product, 0)
	for _, item := range p {
		proto := pb.Product{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			Categories:  item.Categories,
		}
		proto.PriceUsd = &pb.Money{}
		proto.PriceUsd.CurrencyCode = item.PriceUsd.CurrencyCode
		proto.PriceUsd.Units = int64(item.PriceUsd.Units)
		proto.PriceUsd.Nanos = int32(item.PriceUsd.Nanos)
		products = append(products, &proto)
	}
	return products, nil
}

func (c *RestClient) ListRecommendations(cartitems []string) ([]*pb.Product, error) {
	qparam := ""
	for index, param := range cartitems {
		if index == 0 {
			qparam = "?"
		}
		qparam = qparam + fmt.Sprintf("product_id=%s", param)
		if index < len(cartitems)-2 {
			qparam = qparam + "&"
		}
	}
	url := fmt.Sprintf("http://%s/%s%s", c.RecommendationService, "list-recommendations", qparam)
	log.Debugf("REST endpoint %s", url)

	res, err := c.restClient.Get(url)
	if err != nil {
		error := fmt.Sprintf("error sending get: url [%s], error:  %v", url, err)
		return nil, fmt.Errorf(error)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var productIds []string
	if err := json.Unmarshal(data, &productIds); err != nil {
		error := fmt.Sprintf("error mmarshaling response: %v", err)
		return nil, fmt.Errorf(error)
	}
	out := make([]*pb.Product, len(productIds))
	for i, v := range productIds {
		p, err := c.GetProduct(v)
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

// Cartservice
func (c *RestClient) AddCart(user_id, productId string, quantity int32) error {
	url := fmt.Sprintf("http://%s/%s/%s", c.CartService, "cart", user_id)
	cartItem := pb.CartItem{ProductId: productId, Quantity: quantity}
	jsonStr, err := json.Marshal(cartItem)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	res, err := c.restClient.Do(request)
	if err != nil {
		return err
	}
	if res.Status != "200 OK" {
		return fmt.Errorf("Expected 200 OK, but received %s", res.Status)
	}
	return nil
}

func (c *RestClient) EmptyCart(user_id string) error {
	url := fmt.Sprintf("http://%s/%s/%s", c.CartService, "cart", user_id)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	res, err := c.restClient.Do(request)
	if err != nil {
		return err
	}
	if res.Status != "200 OK" {
		return fmt.Errorf("Expected 200 OK, but received %s", res.Status)
	}
	return nil
}

func (c *RestClient) GetCart(user_id string) (string, error) {
	url := fmt.Sprintf("http://%s/%s/%s", c.CartService, "cart", user_id)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	res, err := c.restClient.Do(request)
	if err != nil {
		return "", err
	}
	if res.Status != "200 OK" {
		return "", fmt.Errorf("Expected 200 OK, but received %s", res.Status)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading GetCart response: %v", err)
		return "", fmt.Errorf(error)
	}
	return string(data), nil
}

func (c *RestClient) ConvertCurrency(money *pb.Money, toCode string) (*pb.Money, error) {
	url := fmt.Sprintf("https://%s/%s?from=%s&units=%d&nanos=%d&to_code=%s",
		c.Currencyservice, "convert", money.CurrencyCode, money.Units, money.Nanos, toCode)
	log.Debugf("calling rest endpoint %s", url)
	res, err := c.restClient.Get(url)
	if err != nil || res.Status != "200 OK" {
		fmt.Printf("error sending get: url [%s], status %s, error:  %v", url, res.Status, err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading Convert response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var m pb.Money
	if err := json.Unmarshal(data, &m); err != nil {
		error := fmt.Sprintf("error marshaling Convert response: %v", err)
		return nil, fmt.Errorf(error)
	}

	return &m, nil
}

func (c *RestClient) GetSupportedCurrencies() ([]string, error) {
	url := fmt.Sprintf("https://%s/%s",
		c.Currencyservice, "list-supported-currencies")
	log.Debugf("calling rest endpoint %s", url)
	res, err := c.restClient.Get(url)
	if err != nil {
		fmt.Printf("error sending get: url [%s], status %s, error:  %v", url, res.Status, err)
		return nil, nil
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading GetSupportedCurrencies response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var curr []string
	if err := json.Unmarshal(data, &curr); err != nil {
		error := fmt.Sprintf("error marshaling GetSupportedCurrencies response: %v", err)
		return nil, fmt.Errorf(error)
	}
	return curr, nil
}

func (c *RestClient) GetShippingQuote() (*pb.Money, error) {
	url := fmt.Sprintf("https://%s/%s",
		c.Shippingservice, "get-quote")
	log.Debugf("calling rest endpoint %s", url)
	res, err := c.restClient.Get(url)
	if err != nil {
		fmt.Printf("error sending get: url [%s], status %s, error:  %v", url, res.Status, err)
		return nil, nil
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading GetQuote response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var money pb.Money
	if err := json.Unmarshal(data, &money); err != nil {
		error := fmt.Sprintf("error marshaling GetQuote response: %v", err)
		return nil, fmt.Errorf(error)
	}
	return &money, nil
}

func (c *RestClient) SendOrderConfirmation(data string) error {
	url := fmt.Sprintf("http://%s/%s", c.Emailservice, "send-confirmation")
	log.Debugf("calling rest endpoint %s", url)
	res, err := c.restClient.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Printf("error sending get: url [%s], status %s, error:  %v", url, res.Status, err)
	}
	return err
}

func (c *RestClient) PlaceOrder(orderReq *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	url := fmt.Sprintf("https://%s/%s", c.Checkoutservice, "place-order")
	log.Debugf("calling rest endpoint %s", url)
	data, err := json.Marshal(orderReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal PlaceOrderRequest")
	}
	res, err := c.restClient.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("error sending post: url [%s], status %s, error:  %v", url, res.Status, err)
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading PlaceOrderResponse response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var result pb.OrderResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to populate PlaceOrderResponse")
	}
	response := pb.PlaceOrderResponse{
		Order: &result,
	}

	return &response, nil
}

func (c *RestClient) GetAds(adReq *pb.AdRequest) ([]*pb.Ad, error) {
	url := fmt.Sprintf("https://%s/%s", c.Adservice, "get-ads")
	log.Debugf("calling rest endpoint %s", url)
	data, err := json.Marshal(adReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal PlaceOrderRequest")
	}
	res, err := c.restClient.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		error := fmt.Sprintf("error sending post: url [%s],  error:  %v", url, err)
		return nil, fmt.Errorf(error)
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading GetAdsResponse response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var ad []pb.Ad
	err = json.Unmarshal(data, &ad)
	if err != nil {
		return nil, fmt.Errorf("failed to populate GetAdResponse")
	}

	result := make([]*pb.Ad, len(ad), len(ad))
	for i := range ad {
		result[i] = &ad[i]
	}
	return result, nil
}
