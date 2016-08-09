package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/applications/dev-client/dev-client-errors"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func mDefaultConsumer() error {

	if err := mInitDefaultDevice(); err != nil {
		return err
	}

	if err := mDefaultHCECredential(); err != nil {
		return err
	}

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Println("Initialised default consumer")

	return nil
}

func mNewConsumer() error {

	fmt.Println("Initialising new consumer")

	if err := mInitNewDevice(); err != nil {
		return err
	}

	if err := mNewHCECredential(); err != nil {
		return err
	}

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	return nil
}

func mScanService() error {

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Print("Scan timeout in milliseconds: ")
	var timeout int
	if err := getUserInput(&timeout); err != nil {
		return err
	}

	log.Debug("pre scan for services")
	services, err := sdk.ServiceDiscovery(timeout)
	log.Debug("end scan for services")

	if err != nil {
		return err
	}

	for _, svc := range services {
		log.Debug("(%s:%d/%s) - %s", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)
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
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Println("Added default HCE credential")

	return sdk.InitHCE(card)
}

func mNewHCECredential() error {

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Print("First Name: ")
	var firstName string
	if err := getUserInput(&firstName); err != nil {
		return err
	}

	fmt.Print("Last Name: ")
	var lastName string
	if err := getUserInput(&lastName); err != nil {
		return err
	}

	fmt.Print("Expiry month: ")
	var expMonth int32
	if err := getUserInput(&expMonth); err != nil {
		return err
	}

	fmt.Print("Expiry year: ")
	var expYear int32
	if err := getUserInput(&expYear); err != nil {
		return err
	}

	fmt.Print("CardNumber: ")
	var cardNumber string
	if err := getUserInput(&cardNumber); err != nil {
		return err
	}

	fmt.Print("Type: ")
	var cardType string
	if err := getUserInput(&cardType); err != nil {
		return err
	}

	fmt.Print("CVC: ")
	var cvc string
	if err := getUserInput(&cvc); err != nil {
		return err
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

	return sdk.InitHCE(card)
}

func mAutoConsume() error {

	fmt.Println("Starting auto consume...")

	log.Debug("pre scan for services")
	services, err := sdk.ServiceDiscovery(deviceProfile.DeviceEntity.Consumer.ConsumerConfig.DeviceDiscoveryTimeout)
	log.Debug("end scan for services")

	if err != nil {

		return err
	}

	if len(services) >= 1 {

		var foundServiceIdx int = -1
		for i, service := range services {
			if service.ServerID == deviceProfile.DeviceEntity.Consumer.AutoConsume.DeviceUid {
				foundServiceIdx = i
				break
			}
		}

		if foundServiceIdx == -1 {
			fmt.Println("Could not find service - is the device id in the autoconsume section correct?")
		} else {

			fmt.Printf("Found Service:: (%s:%d/%s) - %s\n", services[foundServiceIdx].Hostname, services[foundServiceIdx].PortNumber, services[foundServiceIdx].UrlPrefix, services[foundServiceIdx].DeviceDescription)

			log.Debug("Init consumer")
			err := sdk.InitConsumer("http://", services[foundServiceIdx].Hostname, services[foundServiceIdx].PortNumber, services[foundServiceIdx].UrlPrefix, services[foundServiceIdx].ServerID)

			if err != nil {

				return err
			}

			log.Debug("Client created..")

			serviceDetails, err := sdk.RequestServices()

			if err != nil {

				return err
			}

			if len(serviceDetails) >= 1 {

				var foundDetailsIdx int = -1
				for i, serviceDetail := range serviceDetails {
					fmt.Printf("%d - %s\n", serviceDetail.ServiceID, serviceDetail.ServiceDescription)
					if serviceDetail.ServiceID == deviceProfile.DeviceEntity.Consumer.AutoConsume.ServiceID {
						foundDetailsIdx = i
						break
					}
				}

				if foundDetailsIdx == -1 {
					fmt.Println("Could not find service details - is the service id in the autoconsume section correct?")
				} else {
					fmt.Printf("Selecting service: %d - %s\n", serviceDetails[foundDetailsIdx].ServiceID, serviceDetails[foundDetailsIdx].ServiceDescription)

					prices, err := sdk.GetServicePrices(serviceDetails[foundDetailsIdx].ServiceID)

					if err != nil {

						return err
					}

					if len(prices) >= 1 {

						var foundUnitIdIdx int = -1
						for i, price := range prices {
							fmt.Printf("(%d) %s @ %d%s, %s (Unit id = %d)\n", price.ID, price.Description, price.PricePerUnit.Amount, price.PricePerUnit.CurrencyCode, price.UnitDescription, price.UnitID)
							if price.UnitID == deviceProfile.DeviceEntity.Consumer.AutoConsume.UnitID {
								foundUnitIdIdx = i
								break
							}
						}

						if foundUnitIdIdx == -1 {
							fmt.Println("Could not find unit id - is the unit id in the autoconsume section correct?")
						} else {

							fmt.Printf("Selecting price: (%d) %s @ %d%s, %s (Unit id = %d)\n", prices[foundUnitIdIdx].ID, prices[foundUnitIdIdx].Description, prices[foundUnitIdIdx].PricePerUnit.Amount, prices[foundUnitIdIdx].PricePerUnit.CurrencyCode, prices[foundUnitIdIdx].UnitDescription, prices[foundUnitIdIdx].UnitID)

							tpr, err := sdk.SelectService(serviceDetails[foundDetailsIdx].ServiceID, 1, prices[foundUnitIdIdx].ID)

							if err != nil {

								return err
							}

							log.Debug("Making payment of %d\n", tpr.TotalPrice)

							payResp, err := sdk.MakePayment(tpr)

							if err != nil {

								return err
							}

							fmt.Printf("Payment of %d made successfully\n", payResp.TotalPaid)
						}
					}
				}
			}
		}
	}
	return nil

}

func mCarWashDemoConsumer() error {

	fmt.Println("Starting car wash demo (Consumer)")

	if err := mInitDefaultDevice(); err != nil {
		return err
	}

	if err := mDefaultHCECredential(); err != nil {
		return err
	}

	if sdk == nil {
		return errors.New(devclienterrors.ERR_DEVICE_NOT_INITIALISED)
	}

	log.Debug("pre scan for services")
	services, err := sdk.ServiceDiscovery(20000)
	log.Debug("end scan for services")

	if err != nil {

		return err
	}

	if len(services) >= 1 {

		svc := services[0]

		fmt.Printf("# Service:: (%s:%d/%s) - %s\n", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)

		log.Debug("Init consumer")
		err := sdk.InitConsumer("http://", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.ServerID)

		if err != nil {

			return err
		}

		log.Debug("Client created..")

		serviceDetails, err := sdk.RequestServices()

		if err != nil {

			return err
		}

		if len(serviceDetails) >= 1 {

			svcDetails := serviceDetails[0]

			fmt.Printf("%d - %s\n", svcDetails.ServiceID, svcDetails.ServiceDescription)

			prices, err := sdk.GetServicePrices(svcDetails.ServiceID)

			if err != nil {

				return err
			}

			fmt.Printf("------- Prices -------\n")
			if len(prices) >= 1 {

				price := prices[0]

				fmt.Printf("(%d) %s @ %d%s, %s (Unit id = %d)\n", price.ID, price.Description, price.PricePerUnit.Amount, price.PricePerUnit.CurrencyCode, price.UnitDescription, price.UnitID)

				tpr, err := sdk.SelectService(svcDetails.ServiceID, 2, price.ID)

				if err != nil {

					return err
				}

				fmt.Println("#Begin Request#")
				fmt.Printf("ServerID: %s\n", tpr.ServerID)
				fmt.Printf("PriceID = %d - %d units = %d\n", tpr.PriceID, tpr.UnitsToSupply, tpr.TotalPrice)
				fmt.Printf("ClientID: %s, MerchantClientKey: %s, PaymentRef: %s\n", tpr.ClientID, tpr.MerchantClientKey, tpr.PaymentReferenceID)
				fmt.Println("#End Request#")

				log.Debug("Making payment of %d\n", tpr.TotalPrice)

				payResp, err := sdk.MakePayment(tpr)

				if err != nil {

					return err
				}

				fmt.Printf("Payment of %d made successfully\n", payResp.TotalPaid)

				fmt.Printf("Service delivery token: %s\n", payResp.ServiceDeliveryToken)

			}
		}
	}
	return nil
}
