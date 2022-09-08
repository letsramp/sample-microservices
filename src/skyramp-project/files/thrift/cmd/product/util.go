/*
 * Copyright Skyramp Authors 2022
 */
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/apache/thrift/lib/go/thrift"
)

type ThriftProtocolType string

const (
	Binary     ThriftProtocolType = "binary"
	Json       ThriftProtocolType = "json"
	SimpleJson ThriftProtocolType = "simplejson"
	Compact    ThriftProtocolType = "compact"
)

type Option struct {
	HttpTransport bool
	HttpUrl       string
	Protocol      ThriftProtocolType
	Secure        bool
	Buffered      bool
	Framed        bool
}

func NewDefaultOption() *Option {
	return &Option{
		Protocol:      Binary,
		HttpTransport: true,
		Secure:        false,
		Buffered:      true,
		Framed:        false,
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

	cfg := &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	protocolFactory := protocolFactoryMap[opt.Protocol]
	if opt.Secure {
		trans = thrift.NewTSSLSocketConf(hostPort, cfg)
	} else {
		trans = thrift.NewTSocketConf(hostPort, nil)
	}
	if err != nil {
		return nil, nil, err
	}
	if opt.HttpTransport {
		if opt.Secure {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			trans, err = thrift.NewTHttpClientWithOptions(fmt.Sprintf("https://%s%s", hostPort, opt.HttpUrl), thrift.THttpClientOptions{Client: client})
		} else {
			trans, err = thrift.NewTHttpClient(fmt.Sprintf("http://%s%s", hostPort, opt.HttpUrl))
		}
	} else {
		if opt.Buffered {
			trans = thrift.NewTBufferedTransport(trans, 8192)
		} else {
			trans = thrift.NewTFramedTransportConf(trans, cfg)
		}
	}
	if err != nil {
		return nil, nil, err
	}
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client = thrift.NewTStandardClient(iprot, oprot)
	return
}
