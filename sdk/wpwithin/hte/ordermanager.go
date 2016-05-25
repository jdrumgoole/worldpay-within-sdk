package hte
import (
"errors"
)

type OrderManager struct {

	orders map[string]Order
}

func NewOrderManager() (*OrderManager, error) {

	result := &OrderManager{}
	result.orders = make(map[string]Order, 0)

	return result, nil
}

func (om *OrderManager) AddOrder(order Order) error {

	if _, ok := om.orders[order.PaymentReference]; ok {

		return errors.New("Order already exists")
	} else {

		om.orders[order.PaymentReference] = order

		return nil
	}
}

func (om *OrderManager) GetOrder(paymentReference string) (Order, error) {

	if order, ok := om.orders[paymentReference]; ok {

		return order, nil
	} else {

		return Order{}, errors.New("Order not found")
	}
}