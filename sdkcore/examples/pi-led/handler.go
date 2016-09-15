package main

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// Handler handles the events coming from Worldpay Within
type Handler struct {
	ledGreen  rpio.Pin
	ledRed    rpio.Pin
	ledYellow rpio.Pin
}

func (handler *Handler) setup() error {

	handler.ledGreen = rpio.Pin(2)
	handler.ledRed = rpio.Pin(3)
	handler.ledYellow = rpio.Pin(4)

	if err := rpio.Open(); err != nil {

		return err
	}

	// Cleanup (defer until end)
	// rpio.Close()

	// Ensure pins are in output mode
	handler.ledGreen.Output()
	handler.ledRed.Output()
	handler.ledYellow.Output()

	// Turn of both LEDs, set the pins to low.
	handler.ledGreen.Low()
	handler.ledRed.Low()
	handler.ledYellow.Low()

	return nil
}

// BeginServiceDelivery is called by Worldpay Within when a consumer wish to begin delivery of a service
func (handler *Handler) BeginServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int) {

	fmt.Printf("BeginServiceDelivery. ServiceID = %d\n", serviceID)

	if serviceID == 1 {

		handler.ledGreen.High()

	} else if serviceID == 2 {

		handler.ledRed.High()
	} else if serviceID == 3 {

		handler.ledYellow.High()
	}
}

// EndServiceDelivery is called by Worldpay Within when a consumer wish to end delivery of a service
func (handler *Handler) EndServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) {

	fmt.Printf("EndServiceDelivery. ServiceID = %d\n", serviceID)

	if serviceID == 1 {

		handler.ledGreen.Low()

	} else if serviceID == 2 {

		handler.ledRed.Low()
	} else if serviceID == 3 {

		handler.ledYellow.Low()
	}
}
