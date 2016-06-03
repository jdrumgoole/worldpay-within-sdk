package servicediscovery
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"

type Broadcaster interface {

	StartBroadcast(msg domain.ServiceMessage, timeoutMillis int) (chan bool, error)
	StopBroadcast() error
}