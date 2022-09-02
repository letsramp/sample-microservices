package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func httpGetCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := r.URL.Query().Get("user_id")
		switch r.Method {
		case "GET":
			cart := GetCart(user_id)
			jsonCart, err := json.Marshal(cart)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if _, err := w.Write(jsonCart); err != nil {
				fmt.Println("ERROR")
			}
		}
	}
}

func httpAddCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := r.URL.Query().Get("user_id")
		product_id := r.URL.Query().Get("product_id")
		q := r.URL.Query().Get("quantity")
		quantity, err := strconv.Atoi(q)
		if err != nil {
			log.Error("failed to convert quantity to integer")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		AddItem(user_id, product_id, int32(quantity))
		w.WriteHeader(http.StatusOK)
	}
}

func httpEmptyCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := r.URL.Query().Get("user_id")
		EmtyCart(user_id)
		w.WriteHeader(http.StatusOK)
	}
}

func runRest(port string) {
	mux := http.NewServeMux()
	mux.Handle("/get-cart", httpGetCart())
	mux.Handle("/add-cart", httpAddCart())
	mux.Handle("/empty-cart", httpEmptyCart())

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	go func() {
		log.Infof("rest server started on port %s", port)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
		log.Info("rest server terminated")
	}()
}
