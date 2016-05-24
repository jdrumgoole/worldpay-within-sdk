package core
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/psp"
"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/domain"
)

type Core struct {

	Psp psp.Psp
	SvcBroadcaster servicediscovery.Broadcaster
	SvcScanner servicediscovery.Scanner
	HTE *hte.Service
	HCE *hce.Manager
}

func New(device *domain.Device, hteIpv4 string, htePrefix string, htePort int) (*Core, error) {

	hte, err := hte.NewService(hteIpv4, htePrefix, htePort)

	if err != nil {

		return nil, err
	}

	result := &Core{
		HTE: hte,
	}

	return result, nil
}