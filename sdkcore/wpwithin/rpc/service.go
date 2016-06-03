package rpc
import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/rpc/wpthrift"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
)

type ServiceImpl struct {

	wpWithin wpwithin.WPWithin
	transportFactory thrift.TTransportFactory
	protocolFactory thrift.TProtocolFactory
	addr string
	secure bool
}

func NewService(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, wpWithin wpwithin.WPWithin) (Service, error) {

	result := new(ServiceImpl)
	result.transportFactory = transportFactory
	result.protocolFactory = protocolFactory
	result.addr = addr
	result.secure = secure
	result.wpWithin = wpWithin

	return result, nil
}

func (svc *ServiceImpl) Start() error {

	var transport thrift.TServerTransport
	var err error
	//	if secure {
	//		cfg := new(tls.Config)
	//		if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
	//			cfg.Certificates = append(cfg.Certificates, cert)
	//		} else {
	//			return err
	//		}
	//		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	//	} else {
	transport, err = thrift.NewTServerSocket(svc.addr)
	//	}

	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	handler := NewWPWithinHandler(svc.wpWithin)
	processor := wpthrift.NewWPWithinProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, svc.transportFactory, svc.protocolFactory)

	fmt.Println("Starting the simple server... on ", svc.addr)

	return server.Serve()
}
