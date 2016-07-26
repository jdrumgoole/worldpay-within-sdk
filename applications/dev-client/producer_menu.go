package main

import (
	"errors"
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// TODO: put this somewhere sensible
var ERR_DEVICE_NOT_INITIALISED = "Error: Device not initialised"
var DEFAULT_HTE_MERCHANT_CLIENT_KEY = "T_C_03eaa1d3-4642-4079-b030-b543ee04b5af"
var DEFAULT_HTE_MERCHANT_SERVICE_KEY = "T_S_f50ecb46-ca82-44a7-9c40-421818af5996"

/*
func mBroadcast() (int, error) {

	fmt.Print("Broadcast timeout in milliseconds: ")
	var input int
	if _, err := mGetUserInput(&input); err != nil {
		return 0, err
	}

	return 0, nil
}
*/

func mProducerStatus() (int, error) {

	// Show all services
	// Show all prices
	// Status of broadcast

	return 0, errors.New("Not implemented yet..")
}

func mDefaultProducer() (int, error) {

	if _, err := mInitDefaultDevice(); err != nil {
		return 0, err
	}

	// Disabling as this just calls initProducer inside, which is called below.
	// Was causing issue with the HTE port already being bound.
//	if _, err := mDefaultHTECredentials(); err != nil {
//		return 0, err
//	}

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	if err := sdk.InitProducer(DEFAULT_HTE_MERCHANT_CLIENT_KEY, DEFAULT_HTE_MERCHANT_SERVICE_KEY); err != nil {
		return 0, err
	}

	return 0, nil
}

func mNewProducer() (int, error) {

	if _, err := mInitNewDevice(); err != nil {
		return 0, err
	}

	if _, err := mNewHTECredentials(); err != nil {
		return 0, err
	}

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	return 0, nil
}

func mDefaultHTECredentials() (int, error) {

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	return 0, sdk.InitProducer(DEFAULT_HTE_MERCHANT_CLIENT_KEY, DEFAULT_HTE_MERCHANT_SERVICE_KEY)
}

func mNewHTECredentials() (int, error) {

	fmt.Print("Merchant Client Key: ")
	var merchantClientKey string
	if _, err := mGetUserInput(&merchantClientKey); err != nil {
		return 0, err
	}

	fmt.Print("Merchant Service Key: ")
	var merchantServiceKey string
	if _, err := mGetUserInput(&merchantServiceKey); err != nil {
		return 0, err
	}

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	return 0, sdk.InitProducer(merchantClientKey, merchantServiceKey)
}

func mAddRoboWashService() (int, error) {

	roboWash, _ := types.NewService()
	roboWash.Name = "RoboWash"
	roboWash.Description = "Car washed by robot"
	roboWash.Id = 1

	washPriceCar := types.Price{

		UnitID:          1,
		ID:              1,
		Description:     "Car wash",
		UnitDescription: "Single wash",
		PricePerUnit: &types.PricePerUnit {
			Amount: 500,
			CurrencyCode: "GBP",
		},
	}

	washPriceSUV := types.Price{

		UnitID:          1,
		ID:              2,
		Description:     "SUV Wash",
		UnitDescription: "Single wash",
		PricePerUnit: &types.PricePerUnit {
			Amount: 650,
			CurrencyCode: "GBP",
		},
	}

	roboWash.AddPrice(washPriceCar)
	roboWash.AddPrice(washPriceSUV)

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	if err := sdk.AddService(roboWash); err != nil {

		return 0, err
	}

	return 0, nil
}

func mAddRoboAirService() (int, error) {

	roboAir, _ := types.NewService()
	roboAir.Name = "RoboAir"
	roboAir.Description = "Car tyre pressure checked and topped up by robot"
	roboAir.Id = 2

	airSinglePrice := types.Price{

		UnitID:          1,
		ID:              1,
		Description:     "Measure and adjust pressure",
		UnitDescription: "Tyre",
		PricePerUnit: &types.PricePerUnit {
			Amount: 25,
			CurrencyCode: "GBP",
		},
	}

	airFourPrice := types.Price{

		UnitID:          2,
		ID:              2,
		Description:     "Measure and adjust pressure - four tyres for the price of three",
		UnitDescription: "4 Tyre",
		PricePerUnit: &types.PricePerUnit {
			Amount: airSinglePrice.PricePerUnit.Amount * 3,
			CurrencyCode: "GBP",
		},
	}

	roboAir.AddPrice(airSinglePrice)
	roboAir.AddPrice(airFourPrice)

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	if err := sdk.AddService(roboAir); err != nil {

		return 0, err
	}

	return 0, nil
}

func mStartBroadcast() (int, error) {

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Print("Broadcast timeout in milliseconds: ")
	var timeout int
	if _, err := mGetUserInput(&timeout); err != nil {
		return 0, err
	}

	if err := sdk.StartServiceBroadcast(timeout); err != nil {
		return 0, err
	}

	return 0, nil
}

func mStopBroadcast() (int, error) {

	return 0, errors.New("Not implemented yet..")
}

func mCarWashDemoProducer() (int, error) {

	if _, err := mDefaultProducer(); err != nil {
		return 0, err
	}

	if _, err := mAddRoboWashService(); err != nil {
		return 0, err
	}

	if _, err := mAddRoboAirService(); err != nil {
		return 0, err
	}

	if err := sdk.StartServiceBroadcast(20000); err != nil {
		return 0, err
	}

	return 0, nil
}
