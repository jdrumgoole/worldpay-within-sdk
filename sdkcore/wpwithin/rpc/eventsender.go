package rpc

import (
	"crypto/tls"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	log "github.com/Sirupsen/logrus"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/rpc/wpthrift"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/rpc/wpthrift/wpthrift_types"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types/event"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/utils"
)

type EventSenderImpl struct {
	client          wpthrift.WPWithinCallback
	connected       bool
	protocolFactory thrift.TProtocolFactory
	transport       thrift.TTransport
}

func (cb *EventSenderImpl) BeginServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int) {

	log.WithFields(log.Fields{"serviceID": serviceID, "serviceDeliveryToken": fmt.Sprintf("%+v", serviceDeliveryToken), "unitsToSupply": unitsToSupply}).Debug("begin EventSenderImpl.BeginServiceDelivery()")

	defer log.Debug("end EventSenderImpl.BeginServiceDelivery()")

	cb.connectCallbackIfNotConnected()

	sdt := wpthrift_types.ServiceDeliveryToken{

		Key:            serviceDeliveryToken.Key,
		Issued:         utils.TimeFormatISO(serviceDeliveryToken.Issued),
		Expiry:         utils.TimeFormatISO(serviceDeliveryToken.Expiry),
		RefundOnExpiry: serviceDeliveryToken.RefundOnExpiry,
		Signature:      serviceDeliveryToken.Signature,
	}

	err := cb.client.BeginServiceDelivery(int32(serviceID), &sdt, int32(unitsToSupply))

	if err != nil {

		fmt.Println(err.Error())
		log.WithField("error", err.Error()).Error("error calling BeginServiceDelivery using thrift callback client.")
	}
}

func (cb *EventSenderImpl) EndServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) {

	log.WithFields(log.Fields{"serviceID": serviceID, "serviceDeliveryToken": fmt.Sprintf("%+v", serviceDeliveryToken), "unitsReceived": unitsReceived}).Debug("begin EventSenderImpl.EndServiceDelivery()")

	defer log.Debug("end EventSenderImpl.EndServiceDelivery()")

	cb.connectCallbackIfNotConnected()

	sdt := wpthrift_types.ServiceDeliveryToken{

		Key:            serviceDeliveryToken.Key,
		Issued:         utils.TimeFormatISO(serviceDeliveryToken.Issued),
		Expiry:         utils.TimeFormatISO(serviceDeliveryToken.Expiry),
		RefundOnExpiry: serviceDeliveryToken.RefundOnExpiry,
		Signature:      serviceDeliveryToken.Signature,
	}

	err := cb.client.EndServiceDelivery(int32(serviceID), &sdt, int32(unitsReceived))

	if err != nil {

		fmt.Println(err.Error())
		log.WithField("error", err.Error()).Error("error calling EndServiceDelivery using thrift callback client.")
	}
}

func NewEventSender(cfg Configuration) (event.Handler, error) {

	log.WithField("config", fmt.Sprintf("%+v", cfg)).Debug("begin rpc.EventSenderImpl.NewEventSender()")
	defer log.Debug("end rpc.EventSenderImpl.NewEventSender()")

	protocolFactory := thrift.NewTBinaryProtocolFactory(true, true)
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	var transport thrift.TTransport
	var err error

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.CallbackPort)

	log.Debugf("Will use callback connection string = %s", addr)

	if cfg.Secure {
		tlsCfg := new(tls.Config)
		tlsCfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, tlsCfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {

		log.Errorf("Error opening socket. Error = %s", err.Error())

		return nil, err
	}
	transport = transportFactory.GetTransport(transport)
	log.Warn("TODO: Transport not going to close..")
	// TODO - How to close this later?
	//	defer transport.Close()
	// if err := transport.Open(); err != nil {
	// 	return nil, err
	// }

	result := &EventSenderImpl{
		connected:       false,
		transport:       transport,
		protocolFactory: protocolFactory,
	}

	return result, nil
}

func (cb *EventSenderImpl) connectCallbackIfNotConnected() error {

	log.Debug("begin EventSenderImpl.connectCallbackIfNotConnected()")

	defer log.Debug("end EventSenderImpl.connectCallbackIfNotConnected()")

	log.Debugf("cb.connected: %t", cb.connected)

	if !cb.connected {

		log.Debug("cb is not connected, attempting to connect now.")

		cb.client = wpthrift.NewWPWithinCallbackClientFactory(cb.transport, cb.protocolFactory)

		if err := cb.transport.Open(); err != nil {

			log.Errorf("Cannot connect to callback RPC server.. did you forget to restart this RPC service? Error = %s", err.Error())

			return err
		}

		log.Debug("Did connect to cb.")

		cb.connected = true
	}

	return nil
}
