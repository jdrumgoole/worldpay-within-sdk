package mock

import (
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/core"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/psp"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/servicediscovery"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// SDKFactoryImpl implementation of SDKFactory
type SDKFactoryImpl struct {
	core *core.Core
}

// NewSDKFactory create a new SDKFactory
func NewSDKFactory(core *core.Core) (core.SDKFactory, error) {

	return &SDKFactoryImpl{
		core: core,
	}, nil
}

func (factory *SDKFactoryImpl) GetDevice(name, description string) (*types.Device, error) {

	return types.NewDevice(name, description, "test-device-guid", "127.0.0.1", "GBP")

}

func (factory *SDKFactoryImpl) GetPSPMerchant(merchantClientKey, merchantServiceKey string) (psp.Psp, error) {

	return PSP{}, nil
}

func (factory *SDKFactoryImpl) GetPSPClient() (psp.Psp, error) {

	return PSP{}, nil
}

func (factory *SDKFactoryImpl) GetSvcBroadcaster(ipv4Address string) (servicediscovery.Broadcaster, error) {

	return broadcasterImpl{}, nil
}

func (factory *SDKFactoryImpl) GetSvcScanner() (servicediscovery.Scanner, error) {

	return scannerImpl{
		core: factory.core,
	}, nil
}

func (factory *SDKFactoryImpl) GetHTE(device *types.Device, psp psp.Psp, ipv4Address string, hteCredential *hte.Credential, om hte.OrderManager, hteSvcHandler *hte.ServiceHandler) (hte.Service, error) {

	return nil, nil
}

func (factory *SDKFactoryImpl) GetOrderManager() (hte.OrderManager, error) {

	return hte.NewOrderManager()
}

func (factory *SDKFactoryImpl) GetHTEClient() (hte.Client, error) {

	return HTEClientImpl{
		core: factory.core,
	}, nil
}

func (factory *SDKFactoryImpl) GetHTEClientHTTP() (hte.HTEClientHTTP, error) {

	return nil, nil
}

func (factory *SDKFactoryImpl) GetHTEServiceHandler(device *types.Device, psp psp.Psp, credential *hte.Credential, orderManager hte.OrderManager) *hte.ServiceHandler {

	return nil
}
