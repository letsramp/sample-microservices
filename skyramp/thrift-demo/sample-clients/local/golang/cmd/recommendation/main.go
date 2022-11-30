package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	api "sample-thrift/demo"
)

const (
	thriftHttpPath = "/RecommendationService"
)

func main() {
	clientAddr := "thrift-demo.recommendation-service-port50000.e2e-target.skyramp.test"
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
	fmt.Println("Successfully Connected to Product Catalog Server")
	client := api.NewRecommendationServiceClient(c)

	selected_ids := []string{"OLJCESPC7Z"}
	fmt.Printf("Trying to get recommendations for product with ids[%v] \n", selected_ids)
	p, err := client.ListRecommendations(context.Background(), selected_ids)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
		os.Exit(1)
	}
	if p == nil {
		fmt.Println("No Recommendations found", err)
		os.Exit(1)
	}
	jsonProd, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		fmt.Printf("Failed to Marshal the response from ProductCatalogService: %v", err)
		os.Exit(1)
	}
	fmt.Println("Result:")
	fmt.Println(string(jsonProd))
}
