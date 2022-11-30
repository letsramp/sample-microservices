package main

import (
	"context"
	"fmt"
	"os"

	api "sample-thrift/demo"
)

const (
	thriftHttpPath = "/EmailService"
)

func main() {
	clientAddr := "thrift-demo.email-service-port50000.e2e-target.skyramp.test"
	opt := NewDefaultOption()
	opt.HttpUrl = thriftHttpPath

	c, trans, err := NewThriftClient(clientAddr, opt)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
		os.Exit(1)
	}
	err = trans.Open()
	if err != nil {
		fmt.Printf("Failed to connect to Email service: %v", err)
		os.Exit(1)
	}
	fmt.Println("connected to the email service")
	client := api.NewEmailServiceClient(c)
	fmt.Println("About to call SendOrderConfirmation")
	email := "someone@example.c om"
	orderResult := api.OrderResult_{
		OrderID:            "0a8e4c64-7fb3-4579-b05b-65cacdab69a1",
		ShippingTrackingID: "7b16c3c4-41ab-11ed-b878-0242ac120002",
		ShippingCost:       &api.Money{CurrencyCode: "USD", Units: int64(100), Nanos: int32(99000000)},
		ShippingAddress: &api.Address{
			StreetAddress: "1600 Amp street",
			City:          "Mountain View",
			State:         "CA",
			Country:       "USA",
			ZipCode:       94043,
		},
	}
	err = client.SendOrderConfirmation(context.Background(), email, &orderResult)
	if err != nil {
		fmt.Printf("Failed to call SendOrderConfirmation: %v", err)
		os.Exit(1)
	}

	fmt.Println("\nSuccessfully executed email service scenario")
}
