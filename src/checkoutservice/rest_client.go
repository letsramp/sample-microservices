// Copyright 2022 Skyramp, Inc.
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.
package main

import (
	"bytes"
	api "checkoutservice/genproto"
	pb "checkoutservice/genproto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func NewRestClient() *RestClient {
	return &RestClient{restClient: &http.Client{}}
}

func (c *RestClient) GetProduct(productID string) (*api.Product, error) {
	url := fmt.Sprintf("http://%s/%s?product_id=%s", c.ProductCatalogService, "get-product", productID)

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

	var p api.Product
	if err := json.Unmarshal(out, &p); err != nil {
		error := fmt.Sprintf("error mmarshaling response: %v", err)
		return nil, fmt.Errorf(error)
	}
	return &p, nil
}

func (c *RestClient) GetCart(user_id string) (*api.Cart, error) {
	url := fmt.Sprintf("http://%s/%s/user_id/%s", c.CartService, "cart", user_id)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.restClient.Do(request)
	if err != nil {
		return nil, err
	}
	if res.Status != "200 OK" {
		return nil, fmt.Errorf("Expected 200 OK, but received %s", res.Status)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading GetCart response: %v", err)
		return nil, fmt.Errorf(error)
	}
	cart := &api.Cart{}
	if err := json.Unmarshal(out, cart); err != nil {
		error := fmt.Sprintf("error mmarshaling response: %v", err)
		return nil, fmt.Errorf(error)
	}
	return cart, nil
}

func (c *RestClient) charge(chargeReq pb.ChargeRequest) (*pb.ChargeResponse, error) {
	url := fmt.Sprintf("http://%s/%s", c.Paymentservice, "charge")
	log.Debugf("calling rest endpoint %s", url)
	data, err := json.Marshal(chargeReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ChargeRequest")
	}
	res, err := c.restClient.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("error sending post: url [%s], error:  %v", url, err)
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		error := fmt.Sprintf("error reading ChargeResponse response: %v", err)
		return nil, fmt.Errorf(error)
	}

	var chargeResponse pb.ChargeResponse
	err = json.Unmarshal(data, &chargeResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to populate ChargeResponse")
	}

	return &chargeResponse, nil
}
