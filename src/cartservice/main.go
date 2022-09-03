package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Items struct {
	ProductId string `json:"product_id" yaml:"product_id"`
	Quantity  int    `json:"quantity" yaml:"quantity"`
}
type Cart struct {
	UserId string  `json:"user_id" yaml:"user_name"`
	Items  []Items `json:"items" yaml:"items"`
}

const (
	serverCrt         = "server.crt"
	serverKey         = "server.key"
	defaultThriftPort = "50000"
	defaultRestPort   = "60000"
)

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
	log.Out = os.Stdout
}

func main() {
	var (
		port  string
		found bool
	)
	port, found = os.LookupEnv("REST_PORT")
	if !found {
		port = defaultRestPort
		log.Infof("env REST_PORT is not se, using default %s", defaultRestPort)
	}
	runGrpc()
	runRest(port)
	opt := NewDefaultOption()
	opt.HttpUrl = "/CartService"
	port, found = os.LookupEnv("THRIFT_PORT")
	if !found {
		port = defaultThriftPort
		log.Infof("env THRIFT_PORT is not set, using default %s", defaultThriftPort)
	}
	runThrift(port, opt)
	select {}
}
