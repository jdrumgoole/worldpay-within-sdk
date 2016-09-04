package hte

import (
	"errors"
	"fmt"

	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
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

	if _, ok := om.orders[order.UUID]; ok {

		return errors.New("Order already exists")
	} else {

		om.orders[order.UUID] = order

		return nil
	}
}

func (om *OrderManagerImpl) GetOrder(paymentReference string) (*types.Order, error) {

	if order, ok := om.orders[paymentReference]; ok {

		return &order, nil
	}

	return nil, errors.New("Order not found")
}

func (om *OrderManagerImpl) OrderExists(uuid string) bool {

	if _, found := om.orders[uuid]; found {

		return true
	}

	return false

}

func (om *OrderManagerImpl) UpdateOrder(order types.Order) error {

	if om.OrderExists(order.UUID) {

		om.orders[order.UUID] = order

		return nil
	}

	return fmt.Errorf("Cannot update order %s as it does not exist", order.UUID)
}
