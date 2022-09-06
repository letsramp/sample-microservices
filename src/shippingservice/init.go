package main

import (
	"fmt"
)

const (
	thriftPort = "50000"
	restPort   = "60000"
)

func main() {
	fmt.Println("Starting ShippingService")
	opt := NewDefaultOption()
	opt.HttpTransport = true
	opt.HttpUrl = "/ShippingService"

	startThrift(thriftPort, opt)
	log.Infof("ShippingService thrift server started on port :%s", thriftPort)
	startRest()
	log.Infof("ShippingService rest server started on port :%s", restPort)
	startGrpc()
	select {}
}
