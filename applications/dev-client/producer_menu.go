package main

import (
	"errors"
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func mBroadcast() error {

	fmt.Print("Broadcast timeout in milliseconds: ")
	var input int
	_, err := fmt.Scanln(&input)

	if err != nil {

		return err
	}

	return nil
}

func mProducerStatus() error {

	// Show all services
	// Show all prices
	// Status of broadcast

	return errors.New("Not implemented yet..")
}

func mDefaultProducer() error {

	return errors.New("Not implemented yet..")
}

func mNewProducer() error {

	return errors.New("Not implemented yet..")
}

func mDefaultHTECredentials() error {

	return errors.New("Not implemented yet..")
}

func mNewHTECredentials() error {

	fmt.Print("Merchant Client Key: ")
	var merchantClientKey string
	_, err := fmt.Scanf("%s", &merchantClientKey)

	if err != nil {

		return err
	}

	fmt.Print("Merchant Service Key: ")
	var merchantServiceKey string
	_, err = fmt.Scanf("%s", &merchantServiceKey)

	if err != nil {

		return err
	}

	if sdk != nil {

		err = sdk.InitHTE(merchantClientKey, merchantServiceKey)
	} else {

		err = errors.New("Error: Must initialise the device first")
	}

	return err
}

func mStartBroadcast() error {

	return errors.New("Not implemented yet..")
}

func mStopBroadcast() error {

	return errors.New("Not implemented yet..")
}

func mCarWashDemoProducer() error {

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
	sdk.AddService(roboWash)

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
	sdk.AddService(roboAir)

	prodDone := make(chan bool)

	go func() error {

		_, err := sdk.InitProducer()

		if err != nil {

			return err
		}

		return nil
	}()

	bcastDone := make(chan bool)

	go func() error {

		sdk.StartServiceBroadcast(20000)

		return nil
	}()

	<-prodDone
	<-bcastDone

	return nil
}
