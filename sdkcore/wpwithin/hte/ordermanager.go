package hte
import (
"errors"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

type OrderManager struct {

	orders map[string]types.Order
}

func NewOrderManager() (*OrderManager, error) {

	result := &OrderManager{}
	result.orders = make(map[string]types.Order, 0)

	return result, nil
}

func (om *OrderManager) AddOrder(order types.Order) error {

	if _, ok := om.orders[order.PaymentReference]; ok {

		return errors.New("Order already exists")
	} else {

		om.orders[order.PaymentReference] = order

		return nil
	}
}

func (om *OrderManager) GetOrder(paymentReference string) (types.Order, error) {

	if order, ok := om.orders[paymentReference]; ok {

		return order, nil
	} else {

		return types.Order{}, errors.New("Order not found")
	}
}