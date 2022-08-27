package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func httpListProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("received a http request to list products")
		data, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func httpGetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		product_id := r.URL.Query().Get("product_id")
		log.Debugf("received a http request to get product for id[%s]", product_id)
		for _, product := range products {
			if product_id == product.Id {
				data, err := json.Marshal(product)
				if err == nil {
					w.WriteHeader(http.StatusOK)
					w.Write(data)
					return
				}
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}
}

func httpSearchProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := make([]Product, 0)
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Name), strings.ToLower(query)) ||
				strings.Contains(strings.ToLower(product.Description), strings.ToLower(query)) {
				result = append(result, product)
			}
		}
		data, err := json.Marshal(result)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}
}

func runRest(port string) error {
	mux := http.NewServeMux()
	mux.Handle("/get-product", httpGetProduct())
	mux.Handle("/list-products", httpListProducts())
	mux.Handle("/search-products", httpSearchProduct())

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	log.Infof("rest server started on port %s", port)
	if err := httpServer.ListenAndServe(); err != nil {
		// if err := httpServer.ListenAndServeTLS(serverCrt, serverKey); err != nil {
		log.Fatal(err)
		return (err)
	}
	return nil
}
