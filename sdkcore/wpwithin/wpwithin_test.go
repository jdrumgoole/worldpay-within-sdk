package wpwithin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func TestInitialise(t *testing.T) {

}

func TestAddService(t *testing.T) {

	wpw, err := Initialise("test", "description")

	if err != nil {

		t.Error(err.Error())
	}

	svc, err := types.NewService()
	if err != nil {
		t.Error(err.Error())
	}

	svc.Id = 1
	svc.Name = "test-service"
	svc.Description = "test-service-description"

	price1, err := types.NewPrice()

	if err != nil {
		t.Error(err.Error())
	}

	price1.ID = 1
	price1.Description = "test-service-price"
	price1.UnitDescription = "test-service-price-description"
	price1.UnitID = 1
	price1.PricePerUnit = &types.PricePerUnit{}
	price1.PricePerUnit.Amount = 100
	price1.PricePerUnit.CurrencyCode = "GBP"

	price2, err := types.NewPrice()

	if err != nil {
		t.Error(err.Error())
	}

	price2.ID = 2
	price2.Description = "test-service-price-in-EUR"
	price2.UnitDescription = "test-service-price-description-in-EUR"
	price2.UnitID = 1
	price2.PricePerUnit = &types.PricePerUnit{}
	price2.PricePerUnit.Amount = 999
	price2.PricePerUnit.CurrencyCode = "EUR"

	err = svc.AddPrice(*price1)
	if err != nil {
		t.Error(err.Error())
	}

	err = svc.AddPrice(*price2)
	if err != nil {
		t.Error(err.Error())
	}

	err = wpw.AddService(svc)
	if err != nil {
		t.Error(err.Error())
	}

	// Now assert that everything was added correctly by reading them back
	services, err := wpw.RequestServices()

	if err != nil {

		t.Error(err)
	}

	assert.Equal(t, 1, len(services), "Should have only received one service")

	svcDetails := services[0]

	assert.Equal(t, "test-service-description", svcDetails.ServiceDescription, "Service description should match")
	assert.Equal(t, 1, svcDetails.ServiceID, "Service id should be 1")

	svcPrices, err := wpw.GetServicePrices(svcDetails.ServiceID)

	if err != nil {

		t.Error(err.Error())
	}

	assert.Equal(t, 2, len(svcPrices), "Should be equal - 2 prices added previously")

	// svcPriceA := svcPrices[0]
	// svcPriceB := svcPrices[1]

}
