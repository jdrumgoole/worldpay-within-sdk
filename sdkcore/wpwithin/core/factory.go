package core
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/utils"
	"fmt"
	"errors"
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp/onlineworldpay"
)

const (

	// The amount of time to sleep between sending each broadcast message (Milliseconds)
	BROADCAST_STEP_SLEEP = 5000
	// The port to broadcast messages on
	BROADCAST_PORT = 8980
	// HTE REST API Url prefix - can be empty
	HTE_SVC_URL_PREFIX = ""
	// Path to store devie UUID once created
	UUID_FILE_PATH = "uuid.txt"
	// Port that the HTE REST API listens on
	HTE_SVC_PORT = 8080
	// Worldpay online API endpoint
	WP_ONLINE_API_ENDPOINT = "https://api.worldpay.com/v1"
	// HTE REST API Scheme typically http:// or https://
	HTE_CLIENT_SCHEME = "http://"
)

type SDKFactory interface {

	GetDevice(name, description string) (*types.Device, error)
	GetPSPMerchant(merchantClientKey, merchantServiceKey string) (psp.Psp, error)
	GetPSPClient() (psp.Psp, error)
	GetSvcBroadcaster(ipv4Address string) (servicediscovery.Broadcaster, error)
	GetSvcScanner() (servicediscovery.Scanner, error)
	GetHTE(device *types.Device, psp psp.Psp, ipv4Address string, hteCredential *hte.Credential, om hte.OrderManager, hteSvcHandler *hte.ServiceHandler) (hte.Service, error)
	GetOrderManager() (hte.OrderManager, error)
	GetHTEClient() (hte.Client, error)
	GetHTEClientHTTP() (hte.HTEClientHTTP, error)
	GetHTEServiceHandler(device *types.Device, psp psp.Psp, credential *hte.Credential, orderManager hte.OrderManager) *hte.ServiceHandler
}

type SDKFactoryImpl struct {}

func NewSDKFactory() (SDKFactory, error) {

	return &SDKFactoryImpl{}, nil
}

func (factory *SDKFactoryImpl) GetDevice(name, description string) (*types.Device, error) {

	var deviceGUID string

	if b, _ := utils.FileExists(UUID_FILE_PATH); b {

		if _deviceGUID, err := utils.ReadLocalUUID(UUID_FILE_PATH); err != nil {

			return nil, errors.New(fmt.Sprintf("Could not read UUID file (%s). Try deleting it. %q", UUID_FILE_PATH, err.Error()))

		} else {

			deviceGUID = _deviceGUID
		}
	} else {

		if _deviceGUID, err := utils.NewUUID(); err != nil {

			return nil, errors.New(fmt.Sprintf("Unable to create new UUID: %q", err.Error()))

		} else {

			deviceGUID = _deviceGUID
		}

		if err := utils.WriteString(UUID_FILE_PATH, deviceGUID, true); err != nil {

			return nil, errors.New(fmt.Sprintf("Could not save UUID to %s", UUID_FILE_PATH))
		}
	}

	if deviceAddress, err := utils.ExternalIPv4(); err != nil {

		return nil, errors.New(fmt.Sprintf("Unable to get IP address: %q", err.Error()))
	} else {

		d, e := types.NewDevice(name, description, deviceGUID, deviceAddress.String(), "GBP")

		return d, e
	}
}

func (factory *SDKFactoryImpl) GetPSPMerchant(merchantClientKey, merchantServiceKey string) (psp.Psp, error) {

	return onlineworldpay.NewMerchant(merchantClientKey, merchantServiceKey, WP_ONLINE_API_ENDPOINT)
}

func (factory *SDKFactoryImpl) GetPSPClient() (psp.Psp, error) {

	return onlineworldpay.NewClient(WP_ONLINE_API_ENDPOINT)
}

func (factory *SDKFactoryImpl) GetSvcBroadcaster(ipv4Address string) (servicediscovery.Broadcaster, error) {

	return servicediscovery.NewBroadcaster(ipv4Address, BROADCAST_PORT, BROADCAST_STEP_SLEEP)
}

func (factory *SDKFactoryImpl) GetSvcScanner() (servicediscovery.Scanner, error) {

	return servicediscovery.NewScanner(BROADCAST_PORT, BROADCAST_STEP_SLEEP)
}

func (factory *SDKFactoryImpl) GetHTE(device *types.Device, psp psp.Psp, ipv4Address string, hteCredential *hte.Credential, om hte.OrderManager, hteSvcHandler *hte.ServiceHandler) (hte.Service, error) {

	return hte.NewService(device, psp, ipv4Address, HTE_SVC_URL_PREFIX, HTE_SVC_PORT, hteCredential, om, hteSvcHandler)
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

func (factory *SDKFactoryImpl) GetHTEServiceHandler(device *types.Device, psp psp.Psp, credential *hte.Credential, orderManager hte.OrderManager) *hte.ServiceHandler {

	return hte.NewServiceHandler(device, psp, credential, orderManager)
}