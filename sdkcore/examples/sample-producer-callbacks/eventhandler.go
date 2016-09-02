package main

import (
	"fmt"

	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

type EventHandlerImpl struct{}

func (eh *EventHandlerImpl) BeginServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int) {

	fmt.Println("go event from core - begin service delivery")
}

func (eh *EventHandlerImpl) EndServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) {

	fmt.Println("go event from core - end service delivery")
}
