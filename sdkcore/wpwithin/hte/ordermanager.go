package hte
import (
"errors"
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"
)

type OrderManager struct {

	orders map[string]domain.Order
}

func NewOrderManager() (*OrderManager, error) {

	result := &OrderManager{}
	result.orders = make(map[string]domain.Order, 0)

	return result, nil
}

func (om *OrderManager) AddOrder(order domain.Order) error {

	if _, ok := om.orders[order.PaymentReference]; ok {

		return errors.New("Order already exists")
	} else {

		om.orders[order.PaymentReference] = order

		return nil
	}
}

func (om *OrderManager) GetOrder(paymentReference string) (domain.Order, error) {

	if order, ok := om.orders[paymentReference]; ok {

		return order, nil
	} else {

		return domain.Order{}, errors.New("Order not found")
	}
}