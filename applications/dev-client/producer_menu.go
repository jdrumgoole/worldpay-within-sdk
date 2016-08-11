package main

import (
	"errors"
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/applications/dev-client/dev-client-defaults"
	"innovation.worldpay.com/worldpay-within-sdk/applications/dev-client/dev-client-errors"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func mDefaultProducer() error {

	if err := mInitDefaultDevice(); err != nil {
		return err
	}

	//if err := mDefaultHTECredentials(); err != nil {
	//	return err
	//}

	// Disabling as this just calls initProducer inside, which is called below.
	// Was causing issue with the HTE port already being bound.
	//	if _, err := mDefaultHTECredentials(); err != nil {
	//		return 0, err
	//	}

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	if err := sdk.InitProducer(devclientdefaults.DEFAULT_HTE_MERCHANT_CLIENT_KEY, devclientdefaults.DEFAULT_HTE_MERCHANT_SERVICE_KEY); err != nil {
		return err
	}

	fmt.Println("Initialised default producer")

	return nil
}

func mNewProducer() error {

	if err := mInitNewDevice(); err != nil {
		return err
	}

	if err := mNewHTECredentials(); err != nil {
		return err
	}

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Println("Initialised new producer")

	return nil
}

func mDefaultHTECredentials() error {

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Println("Added default HTE credentials")

	return sdk.InitProducer(devclientdefaults.DEFAULT_HTE_MERCHANT_CLIENT_KEY, devclientdefaults.DEFAULT_HTE_MERCHANT_SERVICE_KEY)
}

func mNewHTECredentials() error {

	fmt.Println("Add new HTE credentials")

	fmt.Print("Merchant Client Key: ")
	var merchantClientKey string
	if err := getUserInput(&merchantClientKey); err != nil {
		return err
	}

	fmt.Print("Merchant Service Key: ")
	var merchantServiceKey string
	if err := getUserInput(&merchantServiceKey); err != nil {
		return err
	}

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	return sdk.InitProducer(merchantClientKey, merchantServiceKey)
}

func mAddRoboWashService() error {

	roboWash, _ := types.NewService()
	roboWash.Name = "RoboWash"
	roboWash.Description = "Car washed by robot"
	roboWash.Id = 1

	washPriceCar := types.Price{

		UnitID:          1,
		ID:              1,
		Description:     "Car wash",
		UnitDescription: "Single wash",
		PricePerUnit: &types.PricePerUnit{
			Amount:       500,
			CurrencyCode: "GBP",
		},
	}

	washPriceSUV := types.Price{

		UnitID:          1,
		ID:              2,
		Description:     "SUV Wash",
		UnitDescription: "Single wash",
		PricePerUnit: &types.PricePerUnit{
			Amount:       650,
			CurrencyCode: "GBP",
		},
	}

	roboWash.AddPrice(washPriceCar)
	roboWash.AddPrice(washPriceSUV)

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	if err := sdk.AddService(roboWash); err != nil {

		return err
	}

	fmt.Println("Added robowash service")

	return nil
}

func mAddRoboAirService() error {

	roboAir, _ := types.NewService()
	roboAir.Name = "RoboAir"
	roboAir.Description = "Car tyre pressure checked and topped up by robot"
	roboAir.Id = 2

	airSinglePrice := types.Price{

		UnitID:          1,
		ID:              1,
		Description:     "Measure and adjust pressure",
		UnitDescription: "Tyre",
		PricePerUnit: &types.PricePerUnit{
			Amount:       25,
			CurrencyCode: "GBP",
		},
	}

	airFourPrice := types.Price{

		UnitID:          2,
		ID:              2,
		Description:     "Measure and adjust pressure - four tyres for the price of three",
		UnitDescription: "4 Tyre",
		PricePerUnit: &types.PricePerUnit{
			Amount:       airSinglePrice.PricePerUnit.Amount * 3,
			CurrencyCode: "GBP",
		},
	}

	roboAir.AddPrice(airSinglePrice)
	roboAir.AddPrice(airFourPrice)

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	if err := sdk.AddService(roboAir); err != nil {

		return err
	}

	fmt.Println("Added roboair service")

	return nil
}

func mStartBroadcast() error {

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Print("Broadcast timeout in milliseconds: ")
	var timeout int
	if err := getUserInput(&timeout); err != nil {
		return err
	}

	if err := sdk.StartServiceBroadcast(timeout); err != nil {
		return err
	}

	fmt.Println("Broadcast started...")
	return nil
}

func mStopBroadcast() error {

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	sdk.StopServiceBroadcast()

	fmt.Println("Broadcast stopped")
	return nil
}

func mCarWashDemoProducer() error {

	fmt.Println("Starting car wash demo (Producer)")

	if err := mDefaultProducer(); err != nil {
		return err
	}

	if err := mAddRoboWashService(); err != nil {
		return err
	}

	if err := mAddRoboAirService(); err != nil {
		return err
	}

	if err := sdk.StartServiceBroadcast(20000); err != nil {
		return err
	}

	return nil
}
