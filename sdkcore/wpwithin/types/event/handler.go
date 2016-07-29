package event
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"

type Handler interface {

	BeginServiceDelivery(clientID string, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int)
	EndServiceDelivery(clientID string, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int)
}