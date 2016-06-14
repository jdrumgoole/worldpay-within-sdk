package main

import (
	"errors"
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// TODO: put this somewhere sensible
var ERR_DEVICE_NOT_INITIALISED = "Error: Device not initialised"

func mBroadcast() (int, error) {

	fmt.Print("Broadcast timeout in milliseconds: ")
	var input int
	_, err := fmt.Scanln(&input)

	if err != nil {

		return 0, err
	}

	return 0, nil
}

func mProducerStatus() (int, error) {

	// Show all services
	// Show all prices
	// Status of broadcast

	return 0, errors.New("Not implemented yet..")
}

func mDefaultProducer() (int, error) {

	return 0, errors.New("Not implemented yet..")
}

func mNewProducer() (int, error) {

	return 0, errors.New("Not implemented yet..")
}

func mDefaultHTECredentials() (int, error) {

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	return 0, sdk.InitHTE("T_C_c93d7723-2b1c-4dd2-bfb7-58dd48cd093e", "T_S_6ec32d94-77fa-42ff-bede-de487d643793")
}

func mNewHTECredentials() (int, error) {

	fmt.Print("Merchant Client Key: ")
	var merchantClientKey string
	_, err := fmt.Scanf("%s", &merchantClientKey)

	if err != nil {
		return 0, err
	}

	fmt.Print("Merchant Service Key: ")
	var merchantServiceKey string
	_, err = fmt.Scanf("%s", &merchantServiceKey)

	if err != nil {
		return 0, err
	}

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	return 0, sdk.InitHTE(merchantClientKey, merchantServiceKey)
}

func mAddRoboWashService() (int, error) {

	roboWash, _ := types.NewService()
	roboWash.Name = "RoboWash"
	roboWash.Description = "Car washed by robot"
	roboWash.Id = 1

	washPriceCar := types.Price{

		ServiceID:       roboWash.Id,
		UnitID:          1,
		ID:              1,
		Description:     "Car wash",
		UnitDescription: "Single wash",
		PricePerUnit:    500,
	}

	washPriceSUV := types.Price{

		ServiceID:       roboWash.Id,
		UnitID:          1,
		ID:              2,
		Description:     "SUV Wash",
		UnitDescription: "Single wash",
		PricePerUnit:    650,
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
		ServiceID:       roboAir.Id,
		UnitID:          1,
		ID:              1,
		Description:     "Measure and adjust pressure",
		UnitDescription: "Tyre",
		PricePerUnit:    25,
	}

	airFourPrice := types.Price{
		ServiceID:       roboAir.Id,
		UnitID:          2,
		ID:              2,
		Description:     "Measure and adjust pressure - four tyres for the price of three",
		UnitDescription: "4 Tyre",
		PricePerUnit:    airSinglePrice.PricePerUnit * 3,
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

	return 0, errors.New("Not implemented yet..")
}

func mStopBroadcast() (int, error) {

	return 0, errors.New("Not implemented yet..")
}

func mCarWashDemoProducer() (int, error) {

	if _, err := mInitDefaultDevice(); err != nil {
		return 0, err
	}

	if _, err := mDefaultHTECredentials(); err != nil {
		return 0, err
	}

	if _, err := mAddRoboWashService(); err != nil {
		return 0, err
	}

	if _, err := mAddRoboAirService(); err != nil {
		return 0, err
	}

	_, err := sdk.InitProducer()

	if err != nil {

		return 0, err
	}

	if err := sdk.StartServiceBroadcast(20000); err != nil {

		return 0, err
	}

	return 0, nil
}
