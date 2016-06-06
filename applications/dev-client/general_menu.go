package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/rpc"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func mGetDeviceInfo() error {

	return errors.New("Not implemented yet..")
}

func mInitDefaultDevice() error {

	_sdk, err := wpwithin.Initialise("conorhwp-macbook", "Conor H WP - Raspberry Pi")

	if err != nil {

		return err
	}

	_sdk.InitHTE("T_C_c93d7723-2b1c-4dd2-bfb7-58dd48cd093e", "T_S_6ec32d94-77fa-42ff-bede-de487d643793")

	card := types.HCECard{

		FirstName:  "Bilbo",
		LastName:   "Baggins",
		ExpMonth:   11,
		ExpYear:    2018,
		CardNumber: "5555555555554444",
		Type:       "Card",
		Cvc:        "113",
	}

	err = _sdk.InitHCE(card)

	if err != nil {

		return err
	}

	sdk = _sdk

	return nil
}

func mInitNewDevice() error {

	//fmt.Println("Not implemented yet..")

	fmt.Print("Name of device: ")
	var nameOfDevice string
	_, err := fmt.Scanf("%s", &nameOfDevice)

	if err != nil {

		return err
	}

	fmt.Print("Description: ")
	var description string
	_, err = fmt.Scanf("%s", &description)

	if err != nil {

		return err
	}

	_sdk, err := wpwithin.Initialise(nameOfDevice, description)

	if err != nil {

		return err
	}

	sdk = _sdk

	return nil
}

func mCarWashDemoConsumer() error {

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

	if len(services) >= 1 {

		svc := services[0]

		fmt.Println("# Service:: (%s:%d/%s) - %s", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)

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

				fmt.Printf("(%d) %s @ %d, %s (Unit id = %d)\n", price.ID, price.Description, price.PricePerUnit, price.UnitDescription, price.UnitID)

				tpr, err := sdk.SelectService(price.ServiceID, 2, price.ID)

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

func mResetSessionState() error {

	return errors.New("Not implemented yet..")
}

func mLoadConfig() error {

	// Ask user for path to config file
	// (And password if secured)

	return errors.New("Not implemented yet..")
}

func mReadConfig() error {

	// Print out loaded configuration
	// Print out the path to file that was loaded (Need to keep reference during load stage)

	return errors.New("Not implemented yet..")
}

func mStartRPCService() error {

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

		return err
	}

	if err := rpc.Start(); err != nil {

		return err
	}

	return nil
}
