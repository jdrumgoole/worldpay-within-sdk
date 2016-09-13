package main

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// Handler handles the events coming from Worldpay Within
type Handler struct {
	ledBig   rpio.Pin
	ledSmall rpio.Pin
}

func (handler *Handler) setup() error {

	handler.ledBig = rpio.Pin(3)
	handler.ledSmall = rpio.Pin(4)

	if err := rpio.Open(); err != nil {

		return err
	}

	// Cleanup (defer until end)
	// rpio.Close()

	// Ensure pins are in output mode
	handler.ledBig.Output()
	handler.ledSmall.Output()

	// Turn on both LEDs, set the pins to high.
	handler.ledBig.Low()
	handler.ledSmall.Low()

	return nil
}

// BeginServiceDelivery is called by Worldpay Within when a consumer wish to begin delivery of a service
func (handler *Handler) BeginServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int) {

	fmt.Printf("BeginServiceDelivery. ServiceID = %d\n", serviceID)

	if serviceID == 1 {
		handler.ledBig.High()

	} else if serviceID == 2 {

		handler.ledSmall.High()
	}
}

// EndServiceDelivery is called by Worldpay Within when a consumer wish to end delivery of a service
func (handler *Handler) EndServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) {

	fmt.Printf("EndServiceDelivery. ServiceID = %d\n", serviceID)

	if serviceID == 1 {

		handler.ledBig.Low()

	} else if serviceID == 2 {

		handler.ledSmall.Low()
	}
}
