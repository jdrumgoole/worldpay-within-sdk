package servicediscovery
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"

type ScanResult struct {

	Complete chan bool
	Services map[string]domain.ServiceMessage
	Error error
}
