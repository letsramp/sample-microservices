package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/microservices-demo/src/shippingservice/thriftgo/demo"
	"github.com/sirupsen/logrus"
)

const (
	getQuote       = "get-quote"
	listProductUrl = "ship-order"
	serverCrt      = "server.crt"
	serverKey      = "server.key"
)

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
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
}

// // ---------------Shipping Service----------
// service ShippingService {
//     Money GetQuote(1: Address address, 2: list<CartItem> items)
//     string ShipOrder(1: Address address, 2: list<CartItem> items)
// }
//
// struct Address {
//     1: string street_address,
//     2: string city,
//     3: string state,
//     4: string country,
//     5: i32 zip_code,
// }
//
// struct CartItem {
//     1: string product_id,
//     2: i32  quantity,
// }
//
// struct Money {
//     1: string currency_code,
//     2: i64 units,
//     3: i32 nanos,
// }

func HttpGetQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("received a http request to get shipping Quote")

		log.Info("[GetQuote] received request")
		defer log.Info("[GetQuote] completed request")

		// 1. Generate a quote based on the total number of items to be shipped.
		quote := CreateQuoteFromCount(0)
		shippingQuite := &demo.Money{
			CurrencyCode: "USD",
			Units:        int64(quote.Dollars),
			Nanos:        int32(quote.Cents * 10000000),
		}
		data, err := json.Marshal(shippingQuite)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		}
	}
}

func HttpShipOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			in := &ShipOrderRequest{}
			err := json.NewDecoder(r.Body).Decode(in)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer log.Info("[ShipOrder] completed request")
			// 1. Create a Tracking ID
			baseAddress := fmt.Sprintf("%s, %s, %s", in.Address.StreetAddress, in.Address.City, in.Address.State)
			id := CreateTrackingId(baseAddress)
			// 2. Generate a response.
			out := ShipOrderResponse{
				TrackingId: id,
			}
			jsonOut, err := json.Marshal(out)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonOut)
		}
	}
}

func startRest(port string) {
	mux := http.NewServeMux()
	mux.Handle("/get-quote", HttpGetQuote())
	mux.Handle("/ship-order", HttpShipOrder())

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	go func() {
		log.Infof("Rest server started on port %s", port)
		if err := httpServer.ListenAndServeTLS(serverCrt, serverKey); err != nil {
			log.Fatal(err)
		}
		log.Info("Rest server terminated")
	}()
}
