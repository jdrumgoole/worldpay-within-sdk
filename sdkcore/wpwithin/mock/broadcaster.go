package mock

import "github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"

type broadcasterImpl struct {
}

func (bcast broadcasterImpl) StartBroadcast(msg types.ServiceMessage, timeoutMillis int) error {

	return nil
}

func (bcast broadcasterImpl) StopBroadcast() error {

	return nil
}
