package core
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/psp"
)

type Core struct {

	Psp psp.Psp
	SvcBroadcaster servicediscovery.Broadcaster
	SvcScanner servicediscovery.Scanner
	HTE *hte.Service
	HCE *hce.Manager
}

func New() (*Core, error) {

	result := &Core{}

	return result, nil
}