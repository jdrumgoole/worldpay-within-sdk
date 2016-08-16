package hte

import (
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// OrderManager coordinates during negotitation/payment/delivery flows
type OrderManager interface {
	AddOrder(order types.Order) error
	GetOrder(paymentReference string) (types.Order, error)
	OrderExists(paymentReference string) bool
}
