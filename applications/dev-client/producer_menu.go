package main
import (
"fmt"
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func mBroadcast() {

	fmt.Print("Broadcast timeout in milliseconds: ")
	var input int
	_, err := fmt.Scanln(&input)

	if err != nil {

		fmt.Println(err)
		return
	}
}

func mProducerStatus() {

	// Show all services
	// Show all prices
	// Status of broadcast
}

func mDefaultProducer() {

	fmt.Println("Not implemented yet..")
}

func mNewProducer() {

	fmt.Println("Not implemented yet..")
}

func mDefaultHTECredentials() {

	fmt.Println("Not implemented yet..")
}

func mNewHTECredentials() {

	fmt.Println("Not implemented yet..")
}

func mStartBroadcast() {

	fmt.Println("Not implemented yet..")
}

func mStopBroadcast() {

	fmt.Println("Not implemented yet..")
}

func mCarWashDemoProducer() {

	roboWash, _ := types.NewService()
	roboWash.Name = "RoboWash"
	roboWash.Description = "Car washed by robot"
	roboWash.Id = 1

	washPriceCar := types.Price{

		ServiceID:roboWash.Id,
		UnitID:1,
		ID:1,
		Description:"Car wash",
		UnitDescription:"Single wash",
		PricePerUnit:500,
	}

	washPriceSUV := types.Price{

		ServiceID:roboWash.Id,
		UnitID:1,
		ID:2,
		Description:"SUV Wash",
		UnitDescription:"Single wash",
		PricePerUnit:650,
	}

	roboWash.AddPrice(washPriceCar)
	roboWash.AddPrice(washPriceSUV)
	sdk.AddService(roboWash)

	roboAir, _ := types.NewService()
	roboAir.Name = "RoboAir"
	roboAir.Description = "Car tyre pressure checked and topped up by robot"
	roboAir.Id = 2

	airSinglePrice := types.Price{
		ServiceID: roboAir.Id,
		UnitID: 1,
		ID: 1,
		Description: "Measure and adjust pressure",
		UnitDescription:"Tyre",
		PricePerUnit:25,
	}

	airFourPrice := types.Price{
		ServiceID: roboAir.Id,
		UnitID: 2,
		ID: 2,
		Description: "Measure and adjust pressure - four tyres for the price of three",
		UnitDescription:"4 Tyre",
		PricePerUnit:airSinglePrice.PricePerUnit * 3,
	}

	roboAir.AddPrice(airSinglePrice)
	roboAir.AddPrice(airFourPrice)
	sdk.AddService(roboAir)

	prodDone := make(chan bool)

	go func() {

		_, err := sdk.InitProducer()

		if err != nil {

			fmt.Printf(err.Error())

			return
		}
	}()

	bcastDone := make(chan bool)

	go func() {

		sdk.StartServiceBroadcast(20000)
	}()

	<-prodDone
	<-bcastDone
}