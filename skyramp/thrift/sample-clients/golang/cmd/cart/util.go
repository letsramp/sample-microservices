/*
 * Copyright Skyramp Authors 2022
 */
package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
)

type ThriftProtocolType string

const (
	Binary     ThriftProtocolType = "binary"
	Json       ThriftProtocolType = "json"
	SimpleJson ThriftProtocolType = "simplejson"
	Compact    ThriftProtocolType = "compact"
)

type Resolver struct {
	FQDN         string
	ResolvedAddr string
}

type Option struct {
	HttpTransport bool
	HttpUrl       string
	Protocol      ThriftProtocolType
	Buffered      bool
	resolver      *Resolver
}

func NewDefaultOption() *Option {
	return &Option{
		Protocol:      Binary,
		HttpTransport: true,
		Buffered:      true,
		resolver:      &Resolver{ResolvedAddr: "127.0.0.1:80"},
	}
}

var (
	protocolFactoryMap          = make(map[ThriftProtocolType]thrift.TProtocolFactory)
	bufferedTransportFactoryMap = make(map[bool]thrift.TTransportFactory)
)

func init() {
	protocolFactoryMap[Binary] = thrift.NewTBinaryProtocolFactoryConf(nil)
	protocolFactoryMap[Json] = thrift.NewTJSONProtocolFactory()
	protocolFactoryMap[SimpleJson] = thrift.NewTSimpleJSONProtocolFactoryConf(nil)
	protocolFactoryMap[Compact] = thrift.NewTCompactProtocolFactoryConf(nil)
	bufferedTransportFactoryMap[true] = thrift.NewTBufferedTransportFactory(8192)
	bufferedTransportFactoryMap[false] = thrift.NewTTransportFactory()
}

/*
*  Return a new ThriftClient
 */
func NewThriftClient(hostPort string, opt *Option) (client *thrift.TStandardClient, trans thrift.TTransport, err error) {

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(nil)

	thttpOption := thrift.THttpClientOptions{
		Client: &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
					dialer := &net.Dialer{}
					if opt.resolver.FQDN == strings.Split(addr, ":")[0] {
						addr = opt.resolver.ResolvedAddr
					}
					return dialer.DialContext(ctx, network, addr)
				},
			},
		},
	}

	trans, err = thrift.NewTHttpClientWithOptions(fmt.Sprintf("http://%s%s", hostPort, opt.HttpUrl), thttpOption)
	if opt.Buffered {
		trans = thrift.NewTBufferedTransport(trans, 8192)
	}
	if err != nil {
		return nil, nil, err
	}
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client = thrift.NewTStandardClient(iprot, oprot)
	return
}
