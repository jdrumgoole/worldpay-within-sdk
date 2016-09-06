package rpc

import (
	"crypto/tls"
	"errors"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	log "github.com/Sirupsen/logrus"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/rpc/wpthrift"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types/event"
)

type ServiceImpl struct {
	wpWithin         wpwithin.WPWithin
	transportFactory thrift.TTransportFactory
	protocolFactory  thrift.TProtocolFactory
	host             string
	port             int
	secure           bool
	callbackClient   event.Handler
}

func NewService(config Configuration, wpWithin wpwithin.WPWithin) (Service, error) {

	log.WithField("Config", fmt.Sprintf("%+v", config)).Debug("begin rpc.ServiceImpl.NewService()")

	defer log.Debug("end rpc.ServiceImpl.NewService()")

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

		log.Debug("CallbackPort > 0, Setting up RPC Callback")
		log.Debug("Will attempt to create new EventSender")

		if cb, err := NewEventSender(config); err != nil {

			return nil, err
		} else {

			log.Debug("Did create new EventSender.")
			result.callbackClient = cb
		}
	}

	return result, nil
}

func (svc *ServiceImpl) Start() error {

	log.Debug("begin rpc.ServiceImpl.start()")

	defer log.Debug("end rpc.ServiceImpl.start()")

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

	handler := NewWPWithinHandler(svc.wpWithin, svc.callbackClient)
	processor := wpthrift.NewWPWithinProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, svc.transportFactory, svc.protocolFactory)

	log.Debugf("Starting the rpc server on...: %s", strAddr)

	return server.Serve()
}
