package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func mDefaultConsumer() (int, error) {

	if _, err := mInitDefaultDevice(); err != nil {
		return 0, err
	}

	if _, err := mDefaultHCECredential(); err != nil {
		return 0, err
	}

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	return 0, nil
}

func mNewConsumer() (int, error) {

	if _, err := mInitNewDevice(); err != nil {
		return 0, err
	}

	if _, err := mNewHCECredential(); err != nil {
		return 0, err
	}

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	return 0, nil
}

func mScanService() (int, error) {

	log.Debug("testDiscoveryAndNegotiation")

	if _, err := mInitDefaultDevice(); err != nil {
		return 0, err
	}

	if _, err := mDefaultHCECredential(); err != nil {
		return 0, err
	}

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	log.Debug("pre scan for services")
	services, err := sdk.ServiceDiscovery(20000)
	log.Debug("end scan for services")

	if err != nil {
		return 0, err
	}

	for _, svc := range services {

		fmt.Printf("(%s:%d/%s) - %s", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)
	}
	return 0, nil
}

func mDefaultHCECredential() (int, error) {

	card := types.HCECard{

		FirstName:  "Bilbo",
		LastName:   "Baggins",
		ExpMonth:   11,
		ExpYear:    2018,
		CardNumber: "5555555555554444",
		Type:       "Card",
		Cvc:        "113",
	}

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	return 0, sdk.InitHCE(card)
}

func mGetUserInput(input interface{}) (int, error) {

	var err error

	switch t := input.(type) {
	case *int:
		_, err = fmt.Scanf("%d", input)
	case *int32:
		_, err = fmt.Scanf("%d", input)
	case *string:
		_, err = fmt.Scanf("%s", input)
	default:
		fmt.Printf("unexpected type %T", t)
	}

	if err != nil {
		return 0, err
	}

	return 0, nil
}

func mNewHCECredential() (int, error) {

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Print("First Name: ")
	var firstName string
	if _, err := mGetUserInput(&firstName); err != nil {
		return 0, err
	}

	fmt.Print("Last Name: ")
	var lastName string
	if _, err := mGetUserInput(&lastName); err != nil {
		return 0, err
	}

	fmt.Print("Expiry month: ")
	var expMonth int32
	if _, err := mGetUserInput(&expMonth); err != nil {
		return 0, err
	}

	fmt.Print("Expiry year: ")
	var expYear int32
	if _, err := mGetUserInput(&expYear); err != nil {
		return 0, err
	}

	fmt.Print("CardNumber: ")
	var cardNumber string
	if _, err := mGetUserInput(&cardNumber); err != nil {
		return 0, err
	}

	fmt.Print("Type: ")
	var cardType string
	if _, err := mGetUserInput(&cardType); err != nil {
		return 0, err
	}

	fmt.Print("CVC: ")
	var cvc string
	if _, err := mGetUserInput(&cvc); err != nil {
		return 0, err
	}

	card := types.HCECard{

		FirstName:  firstName,
		LastName:   lastName,
		ExpMonth:   expMonth,
		ExpYear:    expYear,
		CardNumber: cardNumber,
		Type:       cardType,
		Cvc:        cvc,
	}

	return 0, sdk.InitHCE(card)
}

func mDiscoverSvcs() (int, error) {

	return 0, errors.New("Not implemented yet..")
}

func mGetSvcPrices() (int, error) {

	return 0, errors.New("Not implemented yet..")
}

func mSelectService() (int, error) {

	return 0, errors.New("Not implemented yet..")
}

func mMakePayment() (int, error) {

	return 0, errors.New("Not implemented yet..")
}

func mConsumerStatus() (int, error) {

	return 0, errors.New("Not implemented yet..")
}
