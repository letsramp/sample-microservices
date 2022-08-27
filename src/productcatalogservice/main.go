package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Money struct {
	CurrencyCode string `json:"currencyCode,omitempty"`
	Units        int64  `json:"units,omitempty"`
	Nanos        int32  `json:"nanos,omitempty"`
}
type Product struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Picture     string   `json:"picture,omitempty"`
	PriceUsd    Money    `json:"priceUsd,omitempty"`
	Categories  []string `json:"categories,omitempty"`
}

const (
	getProductUrl  = "get-product"
	listProductUrl = "list-products"
	searchProduct  = "search-products"
	productId      = "product_id"
	query          = "query"
	serverCrt      = "server.crt"
	serverKey      = "server.key"
)

var (
	catalog  map[string][]Product
	products []Product
)

var (
	thriftPort = "50000"
	restPort   = "60000"
)

func init() {
	// Logger
	log = logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout
	if _, err := os.Stat(serverCrt); errors.Is(err, os.ErrNotExist) {
		cert, key, err := GenerateTLS()
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(serverCrt, cert, 0600)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(serverKey, key, 0600)
		if err != nil {
			log.Fatal(err)
		}
	}
	catalogJSON, err := ioutil.ReadFile("products.json")
	if err != nil {
		panic(fmt.Sprintf("failed to open product catalog json file: %v", err))
	}
	if err := json.Unmarshal(catalogJSON, &catalog); err != nil {
		panic(fmt.Sprintf("failed to parse the catalog JSON: %v", err))
	}
	fmt.Println("successfully parsed product catalog json")
	products = catalog["products"]
}

func main() {

	if os.Getenv("THRIFT_PORT") != "" {
		thriftPort = os.Getenv("THRIFT_PORT")
	}
	if os.Getenv("REST_PORT") != "" {
		restPort = os.Getenv("REST_PORT")
	}

	go func() {
		startGrpc()
	}()
	go func() {
		opt := NewDefaultOption()
		startThrift(thriftPort, opt)
	}()
	go func() {
		runRest(restPort)
	}()
	select {}
}
