package core
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"
)

type Core struct {

	Device *domain.Device
	Psp psp.Psp
	SvcBroadcaster servicediscovery.Broadcaster
	SvcScanner servicediscovery.Scanner
	HTE *hte.Service
	HCE *hce.Manager
	OrderManager *hte.OrderManager
	HTEClient hte.Client
}

func New() (*Core, error) {

	result := &Core{}

	return result, nil
}