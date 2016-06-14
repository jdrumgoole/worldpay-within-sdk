package servicediscovery
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"

type Broadcaster interface {

	StartBroadcast(msg types.ServiceMessage, timeoutMillis int) error
	StopBroadcast() error
}