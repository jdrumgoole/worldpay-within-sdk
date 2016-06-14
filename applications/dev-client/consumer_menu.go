package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

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
