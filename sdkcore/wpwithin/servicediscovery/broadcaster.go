package servicediscovery
import "github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"

type Broadcaster interface {

	StartBroadcast(msg types.ServiceMessage, timeoutMillis int) error
	StopBroadcast() error
}