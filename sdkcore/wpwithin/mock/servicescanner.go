package mock

import (
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/core"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

type scannerImpl struct {
	core *core.Core
}

func (scanner scannerImpl) ScanForServices(timeout int) (map[string]types.ServiceMessage, error) {

	return nil, nil
}

func (scanner scannerImpl) StopScanner() {

}
