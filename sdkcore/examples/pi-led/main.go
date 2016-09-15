package main

import (
	"fmt"
	"os"
	"time"

	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

var wpw wpwithin.WPWithin
var wpwHandler Handler

func main() {

	_wpw, err := wpwithin.Initialise("wpw-pi-led-box", "Worldpay Within LED Demo")
	wpw = _wpw

	errCheck(err, "WorldpayWithin Initialise")

	doSetupServices()

	err = wpwHandler.setup()
	errCheck(err, "wpwHandler setup")
	wpw.SetEventHandler(&wpwHandler)

	err = wpw.InitProducer("T_C_03eaa1d3-4642-4079-b030-b543ee04b5af", "T_S_f50ecb46-ca82-44a7-9c40-421818af5996")

	errCheck(err, "Init producer")

	err = wpw.StartServiceBroadcast(0)

	errCheck(err, "start service broadcast")

	// Do forever

	done := make(chan bool)
	fnForever := func() {
		for {
			time.Sleep(time.Second * 10)
		}
	}

	go fnForever()

	<-done // Block forever
}

func doSetupServices() {

	// Green LED

	svcGreenLed, err := types.NewService()
	errCheck(err, "Create new service - Green LED")
	svcGreenLed.ID = 1
	svcGreenLed.Name = "Big LED"
	svcGreenLed.Description = "Turn on the green LED"

	priceGreenLedSecond, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceGreenLedSecond.Description = "Turn on the green LED"
	priceGreenLedSecond.ID = 1
	priceGreenLedSecond.UnitDescription = "One second"
	priceGreenLedSecond.UnitID = 1
	priceGreenLedSecond.PricePerUnit = &types.PricePerUnit{
		Amount:       10,
		CurrencyCode: "GBP",
	}

	svcGreenLed.AddPrice(*priceGreenLedSecond)

	priceGreenLedMinute, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceGreenLedMinute.Description = "Turn on the green LED"
	priceGreenLedMinute.ID = 2
	priceGreenLedMinute.UnitDescription = "One minute"
	priceGreenLedMinute.UnitID = 2
	priceGreenLedMinute.PricePerUnit = &types.PricePerUnit{
		Amount:       40,
		CurrencyCode: "GBP",
	}

	svcGreenLed.AddPrice(*priceGreenLedMinute)

	err = wpw.AddService(svcGreenLed)
	errCheck(err, "Add service - green led")

	// Red LED

	svcRedLed, err := types.NewService()
	errCheck(err, "New service - red led")

	svcRedLed.ID = 2
	svcRedLed.Name = "Red LED"
	svcRedLed.Description = "Turn on the red LED"

	priceRedLedSecond, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceRedLedSecond.Description = "Turn on the red LED"
	priceRedLedSecond.ID = 3
	priceRedLedSecond.UnitDescription = "One second"
	priceRedLedSecond.UnitID = 1
	priceRedLedSecond.PricePerUnit = &types.PricePerUnit{
		Amount:       5,
		CurrencyCode: "GBP",
	}

	svcRedLed.AddPrice(*priceRedLedSecond)

	priceRedLedMinute, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceRedLedMinute.Description = "Turn on the red LED"
	priceRedLedMinute.ID = 4
	priceRedLedMinute.UnitDescription = "One minute"
	priceRedLedMinute.UnitID = 2
	priceRedLedMinute.PricePerUnit = &types.PricePerUnit{
		Amount:       20,
		CurrencyCode: "GBP",
	}

	svcRedLed.AddPrice(*priceRedLedMinute)

	err = wpw.AddService(svcRedLed)
	errCheck(err, "Add service - red led")

	// Yellow LED

	svcYellowLed, err := types.NewService()
	errCheck(err, "New service - yellow led")

	svcYellowLed.ID = 3
	svcYellowLed.Name = "Yellow LED"
	svcYellowLed.Description = "Turn on the yellow LED"

	priceYellowLedSecond, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceYellowLedSecond.Description = "Turn on the yellow LED"
	priceYellowLedSecond.ID = 1
	priceYellowLedSecond.UnitDescription = "One second"
	priceYellowLedSecond.UnitID = 1
	priceYellowLedSecond.PricePerUnit = &types.PricePerUnit{
		Amount:       5,
		CurrencyCode: "GBP",
	}

	svcYellowLed.AddPrice(*priceYellowLedSecond)

	priceYellowLedMinute, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceYellowLedMinute.Description = "Turn on the yellow LED"
	priceYellowLedMinute.ID = 2
	priceYellowLedMinute.UnitDescription = "One minute"
	priceYellowLedMinute.UnitID = 2
	priceYellowLedMinute.PricePerUnit = &types.PricePerUnit{
		Amount:       20,
		CurrencyCode: "GBP",
	}

	svcYellowLed.AddPrice(*priceYellowLedMinute)

	err = wpw.AddService(svcYellowLed)
	errCheck(err, "Add service - yellow led")
}

func errCheck(err error, hint string) {

	if err != nil {
		fmt.Printf("Did encounter error during: %s", hint)
		fmt.Println(err.Error())
		fmt.Println("Quitting...")
		os.Exit(1)
	}
}
