package core
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/fsm"
)

type Core struct {

	Device *types.Device
	Psp psp.Psp
	SvcBroadcaster servicediscovery.Broadcaster
	SvcScanner servicediscovery.Scanner
	HTE *hte.Service
	HCE *hce.Manager
	OrderManager *hte.OrderManager
	HTEClient hte.Client
	FSM *fsm.FSM
	FSMHelper fsm.SDKHelper
}

func New() (*Core, error) {

	result := &Core{}

	return result, nil
}