package hte
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

const PORT_RANGE_MIN = 1
const PORT_RANGE_MAX = 65535

// HTE Client - Allow interaction with a HTE service
type Client interface {

	// Retrieve services available from a HTE service
	DiscoverServices() (types.ServiceListResponse, error)
	// Get the price variants of a particular service
	GetPrices(serviceId int) (types.ServicePriceResponse, error)
	// For a given service, price and quantity, get the price requested by the service
	NegotiatePrice(serviceId, priceId, numberOfUnits int) (types.TotalPriceResponse, error)
	// Given a negotiated price, make a payment for that order
	MakeHtePayment(paymentReferenceId, clientId, clientToken string) (types.PaymentResponse, error)
	// For a purchased service, begin delivery of that product/service
	StartDelivery(serviceId int, serviceDeliveryToken string, unitsToSupply int) (int, error)
	// For a purchased service, end delivery of that product/service
	EndDelivery(serviceId int, serviceDeliveryToken string, unitsReceived int) (int, error)
}