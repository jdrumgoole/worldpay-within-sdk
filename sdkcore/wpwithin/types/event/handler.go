package event
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"

type Handler struct {

	BeginServiceDelivery func(clientID string, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int)
	EndServiceDelivery func(clientID string, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int)
}