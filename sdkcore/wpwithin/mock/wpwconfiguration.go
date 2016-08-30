package mock

import "github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/configuration"

// NewWPWConfig provides a new Mock config with default values used for tests
func NewWPWConfig() configuration.WPWithin {

	wpw := configuration.WPWithin{}

	wpw.WSLogEnable = false
	wpw.WSLogLevel = "info"
	wpw.WSLogPort = 9999

	return wpw
}
