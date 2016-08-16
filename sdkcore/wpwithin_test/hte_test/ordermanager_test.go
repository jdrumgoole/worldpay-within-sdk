package hte_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
	"testing"
)

func TestAddOrder(t *testing.T) {

	if om, err := hte.NewOrderManager(); err == nil {

		tmpOrder := types.Order{
			ServiceID:             1,
			ClientID:              "client_id_123",
			SelectedNumberOfUnits: 2,
			SelectedPriceId:       3,
			TotalPrice:            199,
			PaymentReference:      "pay-ref",
			ClientUUID:            "client-uuid-here",
			PSPReference:          "pay-ref",
			DeliveryToken:         "delivery-token-here",
		}

		if err := om.AddOrder(tmpOrder); err == nil {

			if !om.OrderExists(tmpOrder.PaymentReference) {

				assert.Fail(t, "Order does not exist in order manager")
			}

		} else {

			assert.Fail(t, "Error adding order to order manager", err.Error())
		}

	} else {

		assert.Fail(t, "Error creating order manager", err.Error())
	}
}

func TestGetOrder(t *testing.T) {

	if om, err := hte.NewOrderManager(); err == nil {

		tmpOrder := types.Order{
			ServiceID:             1,
			ClientID:              "client_id_123",
			SelectedNumberOfUnits: 2,
			SelectedPriceId:       3,
			TotalPrice:            199,
			PaymentReference:      "pay-ref",
			ClientUUID:            "client-uuid-here",
			PSPReference:          "pay-ref",
			DeliveryToken:         "delivery-token-here",
		}

		err := om.AddOrder(tmpOrder)

		if err != nil {

			assert.Fail(t, "Error adding order ", err.Error())
		}

		if order, err := om.GetOrder(tmpOrder.PaymentReference); err == nil {

			assert.Equal(t, 1, order.ServiceID)
			assert.Equal(t, "client_id_123", order.ClientID)
			assert.Equal(t, 2, order.SelectedNumberOfUnits)
			assert.Equal(t, 3, order.SelectedPriceId)
			assert.Equal(t, 199, order.TotalPrice)
			assert.Equal(t, "pay-ref", order.PaymentReference)
			assert.Equal(t, "client-uuid-here", order.ClientUUID)
			assert.Equal(t, "pay-ref", order.PSPReference)
			assert.Equal(t, "delivery-token-here", order.DeliveryToken)

		} else {

			assert.Fail(t, "Order does not exist in order manager", err.Error())
		}

	} else {

		assert.Fail(t, "Error creating order manager", err.Error())
	}
}
