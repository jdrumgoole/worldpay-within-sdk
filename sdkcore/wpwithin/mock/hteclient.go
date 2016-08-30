package mock

import (
	"errors"

	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/core"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

type HTEClientImpl struct {
	core *core.Core
}

func (hteClient HTEClientImpl) DiscoverServices() (types.ServiceListResponse, error) {

	services := hteClient.core.Device.Services

	result := types.ServiceListResponse{}
	result.ServerID = hteClient.core.Device.UID
	result.Services = make([]types.ServiceDetails, 0)

	for _, svc := range services {

		result.Services = append(result.Services, types.ServiceDetails{
			ServiceID:          svc.Id,
			ServiceDescription: svc.Description,
		})
	}

	return result, nil
}

func (hteClient HTEClientImpl) GetPrices(serviceID int) (types.ServicePriceResponse, error) {

	if svc, ok := hteClient.core.Device.Services[serviceID]; ok {

		response := types.ServicePriceResponse{}
		response.ServerID = hteClient.core.Device.UID

		for _, price := range svc.Prices() {

			response.Prices = append(response.Prices, price)
		}

		return response, nil
	}

	return types.ServicePriceResponse{}, errors.New("service does not exist")
}

func (hteClient HTEClientImpl) NegotiatePrice(serviceID, priceID, numberOfUnits int) (types.TotalPriceResponse, error) {
	return types.TotalPriceResponse{}, nil
}

func (hteClient HTEClientImpl) MakeHtePayment(paymentReferenceID, clientID, clientToken string) (types.PaymentResponse, error) {
	return types.PaymentResponse{}, nil
}

func (hteClient HTEClientImpl) StartDelivery(serviceID int, serviceDeliveryToken string, unitsToSupply int) (int, error) {
	return 0, nil
}

func (hteClient HTEClientImpl) EndDelivery(serviceID int, serviceDeliveryToken string, unitsReceived int) (int, error) {
	return 0, nil
}
