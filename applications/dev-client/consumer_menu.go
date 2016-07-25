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

	fmt.Println("Initialised default consumer")

	return 0, nil
}

func mNewConsumer() (int, error) {

	fmt.Println("Initialising new consumer")

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

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Print("Scan timeout in milliseconds: ")
	var timeout int
	if _, err := getUserInput(&timeout); err != nil {
		return 0, err
	}

	log.Debug("pre scan for services")
	services, err := sdk.ServiceDiscovery(timeout)
	log.Debug("end scan for services")

	if err != nil {
		return 0, err
	}

	for _, svc := range services {
		log.Debug("(%s:%d/%s) - %s", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)
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

	fmt.Println("Added default HCE credential")

	return 0, sdk.InitHCE(card)
}

func mNewHCECredential() (int, error) {

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Print("First Name: ")
	var firstName string
	if _, err := getUserInput(&firstName); err != nil {
		return 0, err
	}

	fmt.Print("Last Name: ")
	var lastName string
	if _, err := getUserInput(&lastName); err != nil {
		return 0, err
	}

	fmt.Print("Expiry month: ")
	var expMonth int32
	if _, err := getUserInput(&expMonth); err != nil {
		return 0, err
	}

	fmt.Print("Expiry year: ")
	var expYear int32
	if _, err := getUserInput(&expYear); err != nil {
		return 0, err
	}

	fmt.Print("CardNumber: ")
	var cardNumber string
	if _, err := getUserInput(&cardNumber); err != nil {
		return 0, err
	}

	fmt.Print("Type: ")
	var cardType string
	if _, err := getUserInput(&cardType); err != nil {
		return 0, err
	}

	fmt.Print("CVC: ")
	var cvc string
	if _, err := getUserInput(&cvc); err != nil {
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

	fmt.Println("Added HCE credential")

	return 0, sdk.InitHCE(card)
}

func mCarWashDemoConsumer() (int, error) {

	fmt.Println("Starting car wash demo (Consumer)")

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

	if len(services) >= 1 {

		svc := services[0]

		fmt.Println("# Service:: (%s:%d/%s) - %s", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)

		log.Debug("Init consumer")
		err := sdk.InitConsumer("http://", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.ServerID)

		if err != nil {

			return 0, err
		}

		log.Debug("Client created..")

		serviceDetails, err := sdk.RequestServices()

		if err != nil {

			return 0, err
		}

		if len(serviceDetails) >= 1 {

			svcDetails := serviceDetails[0]

			fmt.Printf("%d - %s\n", svcDetails.ServiceID, svcDetails.ServiceDescription)

			prices, err := sdk.GetServicePrices(svcDetails.ServiceID)

			if err != nil {

				return 0, err
			}

			fmt.Printf("------- Prices -------\n")
			if len(prices) >= 1 {

				price := prices[0]

				fmt.Printf("(%d) %s @ %d, %s (Unit id = %d)\n", price.ID, price.Description, price.PricePerUnit, price.UnitDescription, price.UnitID)

				tpr, err := sdk.SelectService(svcDetails.ServiceID, 2, price.ID)

				if err != nil {

					return 0, err
				}

				fmt.Println("#Begin Request#")
				fmt.Printf("ServerID: %s\n", tpr.ServerID)
				fmt.Printf("PriceID = %d - %d units = %d\n", tpr.PriceID, tpr.UnitsToSupply, tpr.TotalPrice)
				fmt.Printf("ClientID: %s, MerchantClientKey: %s, PaymentRef: %s\n", tpr.ClientID, tpr.MerchantClientKey, tpr.PaymentReferenceID)
				fmt.Println("#End Request#")

				log.Debug("Making payment of %d\n", tpr.TotalPrice)

				payResp, err := sdk.MakePayment(tpr)

				if err != nil {

					return 0, err
				}

				fmt.Printf("Payment of %d made successfully\n", payResp.TotalPaid)

				fmt.Printf("Service delivery token: %s\n", payResp.ServiceDeliveryToken)

			}
		}
	}
	return 0, nil
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
