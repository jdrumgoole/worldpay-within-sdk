package core
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

type Core struct {

	Device *types.Device
	Psp psp.Psp
	SvcBroadcaster servicediscovery.Broadcaster
	SvcScanner servicediscovery.Scanner
	HTE hte.Service
	HCECard *types.HCECard
	OrderManager hte.OrderManager
	HTEClient hte.Client
}

func NewCore() (*Core, error) {

	result := &Core{}

	return result, nil
}

func (core *Core) SetDevice(device *types.Device) {

	core.Device = device
}

func (core *Core) SetPsp(psp psp.Psp) {

	core.Psp = psp
}

func (core *Core) SetSvcBroadcaster(svcBroadcaster servicediscovery.Broadcaster) {

	core.SvcBroadcaster = svcBroadcaster
}

func (core *Core) SetSvcScanner(serviceScanner servicediscovery.Scanner) {

	core.SvcScanner = serviceScanner
}

func (core *Core) SetHTE(hteService hte.Service) {

	core.HTE = hteService
}

func (core *Core) SetHCECard(hceCard *types.HCECard) {

	core.HCECard = hceCard
}

func (core *Core) SetOrderManager(orderManager hte.OrderManager) {

	core.OrderManager = orderManager
}

func (core *Core) SetHTEClient(hteClient hte.Client) {

	core.HTEClient = hteClient
}