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

	// Big LED

	svcBigLed, err := types.NewService()
	errCheck(err, "Create new service - Big LED")
	svcBigLed.ID = 1
	svcBigLed.Name = "Big LED"
	svcBigLed.Description = "Turn on the big LED"

	priceBigLedSecond, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceBigLedSecond.Description = "Turn on the big LED"
	priceBigLedSecond.ID = 1
	priceBigLedSecond.UnitDescription = "One second"
	priceBigLedSecond.UnitID = 1
	priceBigLedSecond.PricePerUnit = &types.PricePerUnit{
		Amount:       10,
		CurrencyCode: "GBP",
	}

	svcBigLed.AddPrice(*priceBigLedSecond)

	priceBigLedMinute, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceBigLedMinute.Description = "Turn on the big LED"
	priceBigLedMinute.ID = 2
	priceBigLedMinute.UnitDescription = "One minute"
	priceBigLedMinute.UnitID = 2
	priceBigLedMinute.PricePerUnit = &types.PricePerUnit{
		Amount:       40,
		CurrencyCode: "GBP",
	}

	svcBigLed.AddPrice(*priceBigLedMinute)

	err = wpw.AddService(svcBigLed)
	errCheck(err, "Add service - big led")

	// Small LED

	svcSmallLed, err := types.NewService()
	errCheck(err, "New service - small led")

	svcSmallLed.ID = 2
	svcSmallLed.Name = "Small LED"
	svcSmallLed.Description = "Turn on the small LED"

	priceSmallLedSecond, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceSmallLedSecond.Description = "Turn on the small LED"
	priceSmallLedSecond.ID = 3
	priceSmallLedSecond.UnitDescription = "One second"
	priceSmallLedSecond.UnitID = 1
	priceSmallLedSecond.PricePerUnit = &types.PricePerUnit{
		Amount:       5,
		CurrencyCode: "GBP",
	}

	svcSmallLed.AddPrice(*priceBigLedSecond)

	priceSmallLedMinute, err := types.NewPrice()
	errCheck(err, "Create new price")

	priceSmallLedMinute.Description = "Turn on the small LED"
	priceSmallLedMinute.ID = 4
	priceSmallLedMinute.UnitDescription = "One minute"
	priceSmallLedMinute.UnitID = 2
	priceSmallLedMinute.PricePerUnit = &types.PricePerUnit{
		Amount:       20,
		CurrencyCode: "GBP",
	}

	svcSmallLed.AddPrice(*priceSmallLedMinute)

	err = wpw.AddService(svcSmallLed)
	errCheck(err, "Add service - small led")
}

func errCheck(err error, hint string) {

	if err != nil {
		fmt.Printf("Did encounter error during: %s", hint)
		fmt.Println(err.Error())
		fmt.Println("Quitting...")
		os.Exit(1)
	}
}
