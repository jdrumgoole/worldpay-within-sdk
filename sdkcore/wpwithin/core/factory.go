package core

import (
	"fmt"

	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/psp"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/psp/onlineworldpay"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/servicediscovery"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types/event"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/utils"
)

const (

	// BroadcastStepSleep The amount of time to sleep between sending each broadcast message (Milliseconds)
	BroadcastStepSleep = 5000
	// BroadcastPort The port to broadcast messages on
	BroadcastPort = 8980
	// HteSvcURLPrefix HTE REST API Url prefix - can be empty
	HteSvcURLPrefix = ""
	// UUIDFilePath Path to store devie UUID once created
	UUIDFilePath = "uuid.txt"
	// HteSvcPort Port that the HTE REST API listens on
	HteSvcPort = 64521
	// WPOnlineAPIEndpoint Worldpay online API endpoint
	WPOnlineAPIEndpoint = "https://api.worldpay.com/v1"
	// HteClientScheme HTE REST API Scheme typically http:// or https://
	HteClientScheme = "http://"
)

// SDKFactory for creating WPWithin instances. // TODO Needs to be reworked so can be partially implemented.
type SDKFactory interface {
	GetDevice(name, description string) (*types.Device, error)
	GetPSPMerchant(merchantClientKey, merchantServiceKey string) (psp.Psp, error)
	GetPSPClient() (psp.Psp, error)
	GetSvcBroadcaster(ipv4Address string) (servicediscovery.Broadcaster, error)
	GetSvcScanner() (servicediscovery.Scanner, error)
	GetHTE(device *types.Device, psp psp.Psp, ipv4Address, scheme string, hteCredential *hte.Credential, om hte.OrderManager, hteSvcHandler *hte.ServiceHandler) (hte.Service, error)
	GetOrderManager() (hte.OrderManager, error)
	GetHTEClient() (hte.Client, error)
	GetHTEClientHTTP() (hte.HTEClientHTTP, error)
	GetHTEServiceHandler(device *types.Device, psp psp.Psp, credential *hte.Credential, orderManager hte.OrderManager, eventHandler event.Handler) *hte.ServiceHandler
}

// SDKFactoryImpl implementation of SDKFactory
type SDKFactoryImpl struct{}

// NewSDKFactory create a new SDKFactory
func NewSDKFactory() (SDKFactory, error) {

	return &SDKFactoryImpl{}, nil
}

func (factory *SDKFactoryImpl) GetDevice(name, description string) (*types.Device, error) {

	var deviceGUID string

	if b, _ := utils.FileExists(UUIDFilePath); b {

		if _deviceGUID, err := utils.ReadLocalUUID(UUIDFilePath); err != nil {

			return nil, fmt.Errorf("Could not read UUID file (%s). Try deleting it. %q", UUIDFilePath, err.Error())

		} else {

			deviceGUID = _deviceGUID
		}
	} else {

		if _deviceGUID, err := utils.NewUUID(); err != nil {

			return nil, fmt.Errorf("Unable to create new UUID: %q", err.Error())

		} else {

			deviceGUID = _deviceGUID
		}

		if err := utils.WriteString(UUIDFilePath, deviceGUID, true); err != nil {

			return nil, fmt.Errorf("Could not save UUID to %s", UUIDFilePath)
		}
	}

	if deviceAddress, err := utils.ExternalIPv4(); err != nil {

		return nil, fmt.Errorf("Unable to get IP address: %q", err.Error())
	} else {

		d, e := types.NewDevice(name, description, deviceGUID, deviceAddress.String(), "GBP")

		return d, e
	}
}

func (factory *SDKFactoryImpl) GetPSPMerchant(merchantClientKey, merchantServiceKey string) (psp.Psp, error) {

	return onlineworldpay.NewMerchant(merchantClientKey, merchantServiceKey, WPOnlineAPIEndpoint)
}

func (factory *SDKFactoryImpl) GetPSPClient() (psp.Psp, error) {

	return onlineworldpay.NewClient(WPOnlineAPIEndpoint)
}

func (factory *SDKFactoryImpl) GetSvcBroadcaster(ipv4Address string) (servicediscovery.Broadcaster, error) {

	return servicediscovery.NewBroadcaster(ipv4Address, BroadcastPort, BroadcastStepSleep)
}

func (factory *SDKFactoryImpl) GetSvcScanner() (servicediscovery.Scanner, error) {

	return servicediscovery.NewScanner(BroadcastPort, BroadcastStepSleep)
}

func (factory *SDKFactoryImpl) GetHTE(device *types.Device, psp psp.Psp, ipv4Address, scheme string, hteCredential *hte.Credential, om hte.OrderManager, hteSvcHandler *hte.ServiceHandler) (hte.Service, error) {

	return hte.NewService(device, psp, ipv4Address, HteSvcURLPrefix, scheme, HteSvcPort, hteCredential, om, hteSvcHandler)
}

func (factory *SDKFactoryImpl) GetOrderManager() (hte.OrderManager, error) {

	return hte.NewOrderManager()
}

func (factory *SDKFactoryImpl) GetHTEClient() (hte.Client, error) {

	return nil, nil
}

func (factory *SDKFactoryImpl) GetHTEClientHTTP() (hte.HTEClientHTTP, error) {

	return hte.NewHTEClientHTTP()
}

func (factory *SDKFactoryImpl) GetHTEServiceHandler(device *types.Device, psp psp.Psp, credential *hte.Credential, orderManager hte.OrderManager, eventHandler event.Handler) *hte.ServiceHandler {

	return hte.NewServiceHandler(device, psp, credential, orderManager, eventHandler)
}
