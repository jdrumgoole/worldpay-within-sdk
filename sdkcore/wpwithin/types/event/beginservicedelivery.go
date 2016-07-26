package event
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"

type BeginServiceDelivery struct {

	DeliveryToken types.ServiceDeliveryToken
	Name string
}
