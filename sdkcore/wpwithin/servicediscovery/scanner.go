package servicediscovery
import "github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"

type Scanner interface {

	ScanForServices(timeout int) (map[string]types.ServiceMessage, error)

	StopScanner()
}