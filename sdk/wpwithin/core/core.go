package core
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/psp"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/domain"
)

type Core struct {

	Device *domain.Device
	Psp psp.Psp
	SvcBroadcaster servicediscovery.Broadcaster
	SvcScanner servicediscovery.Scanner
	HTE *hte.Service
	HCE *hce.Manager
	OrderManager *hte.OrderManager
}

func New() (*Core, error) {

	result := &Core{}

	return result, nil
}