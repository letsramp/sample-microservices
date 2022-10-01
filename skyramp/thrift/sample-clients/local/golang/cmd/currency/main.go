package main

import (
	"context"
	"fmt"
	"os"

	api "sample-thrift/demo"
)

const (
	thriftHttpPath = "/CurrencyService"
)

func main() {
	clientAddr := "currency-service-port50000.demo.skyramp.test"
	opt := NewDefaultOption()
	opt.HttpUrl = thriftHttpPath

	c, trans, err := NewThriftClient(clientAddr, opt)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
		os.Exit(1)
	}
	err = trans.Open()
	if err != nil {
		fmt.Printf("Failed to connect to Currency Service: %v", err)
		os.Exit(1)
	}
	fmt.Println("connected to the currency service")
	client := api.NewCurrencyServiceClient(c)
	fmt.Println("About to get list of currencies")
	currencies, err := client.GetSupportedCurrencies(context.Background())
	if err != nil {
		fmt.Printf("Failed to get list of supported currencies: %v", err)
		os.Exit(1)
	}
	fmt.Printf("\nsupported currencies\n%v", currencies)

	fmt.Println("\nAbout to convert USD to CAD")
	result, err := client.Convert(context.Background(),
		&api.Money{
			CurrencyCode: "USD",
			Units:        int64(100),
			Nanos:        int32(99000000),
		}, "CAD")
	if err != nil {
		fmt.Printf("Failed to convert currencies: %v", err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully converted USD to CAD: %v\n", result)
}
