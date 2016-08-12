package hte

import (
	"errors"

	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// Concrete implementation of order manager.. uses an in memory persistence for orders
type OrderManagerImpl struct {
	orders map[string]types.Order
}

func NewOrderManager() (OrderManager, error) {

	result := &OrderManagerImpl{}
	result.orders = make(map[string]types.Order, 0)

	return result, nil
}

func (om *OrderManagerImpl) AddOrder(order types.Order) error {

	if _, ok := om.orders[order.PaymentReference]; ok {

		return errors.New("Order already exists")
	} else {

		om.orders[order.PaymentReference] = order

		return nil
	}
}

func (om *OrderManagerImpl) GetOrder(paymentReference string) (types.Order, error) {

	if order, ok := om.orders[paymentReference]; ok {

		return order, nil
	}

	return types.Order{}, errors.New("Order not found")
}

func (om *OrderManagerImpl) OrderExists(paymentReference string) bool {

	if _, found := om.orders[paymentReference]; found {

		return true
	}

	return false

}
