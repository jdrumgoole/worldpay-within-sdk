package hte
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

type OrderManager interface {

	AddOrder(order types.Order) error
	GetOrder(paymentReference string) (types.Order, error)
}