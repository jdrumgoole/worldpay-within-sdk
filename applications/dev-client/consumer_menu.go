package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func mScanService() error {

	log.Debug("testDiscoveryAndNegotiation")

	sdk, err := wpwithin.Initialise("conorhwp-macbook", "Conor H WP - Raspberry Pi")

	if err != nil {

		return err
	}

	err = sdk.InitHTE("T_C_c93d7723-2b1c-4dd2-bfb7-58dd48cd093e", "T_S_6ec32d94-77fa-42ff-bede-de487d643793")

	if err != nil {

		return err
	}

	card := types.HCECard{

		FirstName:  "Bilbo",
		LastName:   "Baggins",
		ExpMonth:   11,
		ExpYear:    2018,
		CardNumber: "5555555555554444",
		Type:       "Card",
		Cvc:        "113",
	}

	err = sdk.InitHCE(card)

	if err != nil {

		return err
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

	return errors.New("Not implemented yet..")
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
