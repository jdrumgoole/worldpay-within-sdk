package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/rpc"
)

// TODO: put these somewhere sensible
var DEFAULT_DEVICE_NAME = "conorhwp-macbook"
var DEFAULT_DEVICE_DESCRIPTION = "Conor H WP - Raspberry Pi"

func mGetDeviceInfo() (int, error) {

	//return 0, errors.New("Not implemented yet..")

	if sdk == nil {
		return 0, errors.New(ERR_DEVICE_NOT_INITIALISED)
	}

	fmt.Printf("Uid of device: %s\n", sdk.GetDevice().Uid)
	fmt.Printf("Name of device: %s\n", sdk.GetDevice().Name)
	fmt.Printf("Description: %s\n", sdk.GetDevice().Description)
	fmt.Printf("Services: \n")

	for i, service := range sdk.GetDevice().Services {
		fmt.Printf("   %d: Id:%d Name:%s Description:%s\n", i, service.Id, service.Name, service.Description)
		fmt.Printf("   Prices: \n")
		for j, price := range service.Prices() {
			fmt.Printf("      %d: ServiceID: %d ID:%d Description:%s PricePerUnit:%d UnitID:%d UnitDescription:%s\n", j, price.ServiceID, price.ID, price.Description, price.PricePerUnit, price.UnitID, price.UnitDescription)
		}
	}

	fmt.Printf("IPv4Address: %s\n", sdk.GetDevice().IPv4Address)
	fmt.Printf("CurrencyCode: %s\n", sdk.GetDevice().CurrencyCode)

	return 0, nil
}

func mInitDefaultDevice() (int, error) {

	_sdk, err := wpwithin.Initialise(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION)

	if err != nil {

		return 0, err
	}

	sdk = _sdk

	return 0, nil
}

func mInitNewDevice() (int, error) {

	fmt.Print("Name of device: ")
	var nameOfDevice string
	if _, err := mGetUserInput(&nameOfDevice); err != nil {
		return 0, err
	}

	fmt.Print("Description: ")
	var description string
	if _, err := mGetUserInput(&description); err != nil {
		return 0, err
	}

	_sdk, err := wpwithin.Initialise(nameOfDevice, description)

	if err != nil {

		return 0, err
	}

	sdk = _sdk

	return 0, err
}

func mCarWashDemoConsumer() (int, error) {

	log.Debug("testDiscoveryAndNegotiation")

	if _, err := mInitDefaultDevice(); err != nil {
		return 0, err
	}

//	if _, err := mDefaultHTECredentials(); err != nil {
//		return 0, err
//	}

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

				tpr, err := sdk.SelectService(price.ServiceID, 2, price.ID)

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

func mResetSessionState() (int, error) {

	sdk = nil

	return 0, nil
}

func mLoadConfig() (int, error) {

	// Ask user for path to config file
	// (And password if secured)

	return 0, errors.New("Not implemented yet..")
}

func mReadConfig() (int, error) {

	// Print out loaded configuration
	// Print out the path to file that was loaded (Need to keep reference during load stage)

	return 0, errors.New("Not implemented yet..")
}

func mStartRPCService() (int, error) {

	config := rpc.Configuration{
		Protocol:   "binary",
		Framed:     false,
		Buffered:   false,
		Host:       "127.0.0.1",
		Port:       9091,
		Secure:     false,
		BufferSize: 8192,
	}

	rpc, err := rpc.NewService(config, sdk)

	if err != nil {

		return 0, err
	}

	if err := rpc.Start(); err != nil {

		return 0, err
	}

	return 0, nil
}
