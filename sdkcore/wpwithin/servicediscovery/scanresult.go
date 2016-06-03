package servicediscovery
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"

type ScanResult struct {

	Complete chan bool
	Services map[string]types.ServiceMessage
	Error error
}
