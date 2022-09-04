package main

import (
	pb "cartservice/genproto"
	"encoding/json"
	"fmt"
	"net/http"
)

func httpCart() http.HandlerFunc {
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
		case "POST":
			cartItem := &pb.CartItem{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(cartItem)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			AddItem(user_id, cartItem.ProductId, cartItem.Quantity)
			w.WriteHeader(http.StatusOK)
		case "DELETE":
			EmtyCart(user_id)
		}
	}
}

func runRest(port string) {
	mux := http.NewServeMux()
	mux.Handle("/cart", httpCart())

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
