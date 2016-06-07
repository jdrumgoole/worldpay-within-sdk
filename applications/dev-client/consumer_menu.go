package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func mScanService() error {

	log.Debug("testDiscoveryAndNegotiation")

	if err := mInitDefaultDevice(); err != nil {
		return err
	}

	if err := mDefaultHTECredentials(); err != nil {
		return err
	}

	if err := mDefaultHCECredential(); err != nil {
		return err
	}

	if sdk == nil {
		return errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	log.Debug("pre scan for services")
	services, err := sdk.ServiceDiscovery(20000)
	log.Debug("end scan for services")

	if err != nil {
		return err
	}

	for _, svc := range services {

		fmt.Printf("(%s:%d/%s) - %s", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)
	}
	return nil
}

func mDefaultHCECredential() error {

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
		return errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	return sdk.InitHCE(card)
}

func mDiscoverSvcs() error {

	return errors.New("Not implemented yet..")
}

func mGetSvcPrices() error {

	return errors.New("Not implemented yet..")
}

func mSelectService() error {

	return errors.New("Not implemented yet..")
}

func mMakePayment() error {

	return errors.New("Not implemented yet..")
}

func mConsumerStatus() error {

	return errors.New("Not implemented yet..")
}
