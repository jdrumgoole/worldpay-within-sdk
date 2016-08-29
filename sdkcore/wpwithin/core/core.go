package core

import (
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/configuration"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/psp"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/servicediscovery"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// SDK Core - This acts as a container for dependencies of the SDK
type Core struct {
	Device         *types.Device
	Psp            psp.Psp
	SvcBroadcaster servicediscovery.Broadcaster
	SvcScanner     servicediscovery.Scanner
	HTE            hte.Service
	HCECard        *types.HCECard
	OrderManager   hte.OrderManager
	HTEClient      hte.Client
	Configuration  configuration.WPWithin
}

// Create a new Core
func NewCore() (*Core, error) {

	result := &Core{}

	return result, nil
}

// Device setter
func (core *Core) SetDevice(device *types.Device) {

	core.Device = device
}

// PSP setter
func (core *Core) SetPsp(psp psp.Psp) {

	core.Psp = psp
}

// Service Broadcaster setter
func (core *Core) SetSvcBroadcaster(svcBroadcaster servicediscovery.Broadcaster) {

	core.SvcBroadcaster = svcBroadcaster
}

// Service Scanner setter
func (core *Core) SetSvcScanner(serviceScanner servicediscovery.Scanner) {

	core.SvcScanner = serviceScanner
}

// HTE Service setter
func (core *Core) SetHTE(hteService hte.Service) {

	core.HTE = hteService
}

// HCE Card setter
func (core *Core) SetHCECard(hceCard *types.HCECard) {

	core.HCECard = hceCard
}

// Order Manager setter
func (core *Core) SetOrderManager(orderManager hte.OrderManager) {

	core.OrderManager = orderManager
}

// HTE Client setter
func (core *Core) SetHTEClient(hteClient hte.Client) {

	core.HTEClient = hteClient
}
