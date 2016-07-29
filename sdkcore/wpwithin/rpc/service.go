package rpc
import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/rpc/wpthrift"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"crypto/tls"
	"errors"
)

type ServiceImpl struct {

	wpWithin wpwithin.WPWithin
	transportFactory thrift.TTransportFactory
	protocolFactory thrift.TProtocolFactory
	host string
	port int
	secure bool
	callback CallbackClient
}

func NewService(config Configuration, wpWithin wpwithin.WPWithin) (Service, error) {


	var protocolFactory thrift.TProtocolFactory
	switch config.Protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		return nil, errors.New(fmt.Sprintf("Invalid protocol specified: %s\n", config.Protocol))
	}

	var transportFactory thrift.TTransportFactory
	if config.Buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(config.BufferSize)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if config.Framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}


	result := new(ServiceImpl)
	result.transportFactory = transportFactory
	result.protocolFactory = protocolFactory
	result.host = config.Host
	result.port = config.Port
	result.secure = config.Secure
	result.wpWithin = wpWithin

	// Only setup callbacks if there is a callback port specified
	if config.CallbackPort > 0 {

		if cb, err := NewCallback(config); err != nil {

			return nil, err
		} else {

			result.callback = cb
		}
	}

	return result, nil
}

func (svc *ServiceImpl) Start() error {

	strAddr := fmt.Sprintf("%s:%d", svc.host, svc.port)

	var transport thrift.TServerTransport
	var err error
		if svc.secure {
			cfg := new(tls.Config)
			if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
				cfg.Certificates = append(cfg.Certificates, cert)
			} else {
				return err
			}
			transport, err = thrift.NewTSSLServerSocket(strAddr, cfg)
		} else {
			transport, err = thrift.NewTServerSocket(strAddr)
		}

	if err != nil {
		return err
	}
	fmt.Printf("Transport: %T\n", transport)
	handler := NewWPWithinHandler(svc.wpWithin, svc.callback)
	processor := wpthrift.NewWPWithinProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, svc.transportFactory, svc.protocolFactory)

	fmt.Printf("Starting the rpc server on...: %s\n", strAddr)

	return server.Serve()
}
