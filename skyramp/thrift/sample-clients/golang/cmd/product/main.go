package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	api "sample-thrift/demo"
)

const (
	thriftHttpPath = "/ProductCatalogService"
)

func main() {
	clientAddr := fmt.Sprintf("product-catalog-service-port50000.demo.skyramp.test")
	opt := NewDefaultOption()
	opt.HttpUrl = thriftHttpPath

	c, trans, err := NewThriftClient(clientAddr, opt)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
		os.Exit(1)
	}
	err = trans.Open()
	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
		os.Exit(1)
	}
	fmt.Println("Sucessfully Connected to Product Catalog Server")
	client := api.NewProductCatalogServiceClient(c)
	product_id := "OLJCESPC7Z"
	fmt.Printf("Trying to get product with id[%s] \n", product_id)
	p, err := client.GetProduct(context.Background(), product_id)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
		os.Exit(1)
	}
	if p == nil {
		fmt.Println("No Products found", err)
		os.Exit(1)
	}
	jsonProd, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		fmt.Printf("Failed to Marshal the response from ProductCatalogService: %v", err)
		os.Exit(1)
	}
	fmt.Println("Result:")
	fmt.Println(string(jsonProd))

	fmt.Printf("\nTrying to get all products \n")
	products, err := client.ListProducts(context.Background())
	if err != nil {
		fmt.Printf("Failed to list products: %v ", err)
		os.Exit(1)
	}
	if p == nil {
		fmt.Println("No Products found", err)
		os.Exit(1)
	}
	jsonProd, err = json.MarshalIndent(products, "", "\t")
	if err != nil {
		fmt.Printf("Failed to Marshal the response from ProductCatalogService: %v", err)
		os.Exit(1)
	}
	fmt.Println("Result:")
	fmt.Println(string(jsonProd))
}
