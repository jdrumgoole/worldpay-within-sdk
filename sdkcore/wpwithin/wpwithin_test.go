package wpwithin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/core"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/mock"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

var wpw *wpWithinImpl
var factory core.SDKFactory

func setupWPWConsumer() error {

	// Setup PSP as client

	_psp, err := factory.GetPSPClient()

	if err != nil {

		return err
	}

	wpw.core.Psp = _psp

	wpw.core.HCECard = &types.HCECard{
		FirstName:  "John",
		LastName:   "Smyth",
		ExpMonth:   12,
		ExpYear:    2025,
		CardNumber: "4444333322221111",
		Type:       "Card",
		Cvc:        "123",
	}

	// Setup HTE Client

	client, _ := factory.GetHTEClient()

	wpw.core.HTEClient = client

	return nil
}

func setupWPW() error {

	_wpw := &wpWithinImpl{}
	wpw = _wpw

	core, err := core.NewCore()

	if err != nil {

		return err
	}

	wpw.core = core

	_factory, err := mock.NewSDKFactory(core)
	factory = _factory

	if err != nil {

		return fmt.Errorf("Unable to create SDK Factory: %q", err.Error())
	}

	// Mock configuration
	core.Configuration = mock.NewWPWConfig()

	dev, err := factory.GetDevice("test-device-name", "test-device-description")

	if err != nil {

		return err
	}

	wpw.core.Device = dev

	om, err := factory.GetOrderManager()

	if err != nil {

		return err

	}

	wpw.core.OrderManager = om

	bc, err := factory.GetSvcBroadcaster(wpw.core.Device.IPv4Address)

	if err != nil {

		return err
	}

	wpw.core.SvcBroadcaster = bc

	sc, err := factory.GetSvcScanner()

	if err != nil {

		return err

	}

	wpw.core.SvcScanner = sc

	return nil
}

func TestProducerAddService(t *testing.T) {

	setupWPW()
	setupWPWConsumer()

	// wpw, err := Initialise("test", "description")
	//
	// if err != nil {
	//
	// 	t.Error(err.Error())
	// }

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

	assert.Equal(t, 1, len(services), "Should have received one service")

	svcDetails := services[0]

	assert.Equal(t, "test-service-description", svcDetails.ServiceDescription, "Service description should match")
	assert.Equal(t, 1, svcDetails.ServiceID, "Service id should be 1")

	svcPrices, err := wpw.GetServicePrices(svcDetails.ServiceID)

	if err != nil {

		t.Error(err.Error())
	}

	assert.Equal(t, 2, len(svcPrices), "Should be equal - 2 prices added previously")
}
