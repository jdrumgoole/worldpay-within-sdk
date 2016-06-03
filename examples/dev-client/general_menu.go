package main
import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func mGetDeviceInfo() {

	fmt.Println("Not implemented yet..")
}

func mInitDefaultDevice() {

	_sdk, err := wpwithin.Initialise("conorhwp-macbook", "Conor H WP - Raspberry Pi")

	if err != nil {

		fmt.Println(err)
		return
	}

	_sdk.InitHTE("T_C_c93d7723-2b1c-4dd2-bfb7-58dd48cd093e", "T_S_6ec32d94-77fa-42ff-bede-de487d643793")

	card := types.HCECard{

		FirstName:"Bilbo",
		LastName:"Baggins",
		ExpMonth:11,
		ExpYear:2018,
		CardNumber:"5555555555554444",
		Type:"Card",
		Cvc:"113",
	}

	err = _sdk.InitHCE(card)

	if err != nil {

		fmt.Println(err)
		return
	}

	sdk = _sdk

}

func mInitNewDevice() {

	fmt.Println("Not implemented yet..")
}

func mCarWashDemoConsumer() {

	log.Debug("testDiscoveryAndNegotiation")

	sdk, err := wpwithin.Initialise("conorhwp-macbook", "Conor H WP - Raspberry Pi")

	if err != nil {

		fmt.Println(err)
	}

	err = sdk.InitHTE("T_C_c93d7723-2b1c-4dd2-bfb7-58dd48cd093e", "T_S_6ec32d94-77fa-42ff-bede-de487d643793")

	if err != nil {

		fmt.Println(err)
	}

	card := types.HCECard{

		FirstName:"Bilbo",
		LastName:"Baggins",
		ExpMonth:11,
		ExpYear:2018,
		CardNumber:"5555555555554444",
		Type:"Card",
		Cvc:"113",
	}

	err = sdk.InitHCE(card)

	if err != nil {

		fmt.Printf("%q\n", err.Error())
		return
	}

	log.Debug("pre scan for services")
	services, err := sdk.ServiceDiscovery(20000)
	log.Debug("end scan for services")


	if err != nil {

		fmt.Println(err)
		return
	}

	if len(services) >= 1 {

		svc := services[0]

		fmt.Println("# Service:: (%s:%d/%s) - %s", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)

		log.Debug("Init consumer")
		err := sdk.InitConsumer("http://", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.ServerID)

		if err != nil {

			fmt.Println(err.Error())
		} else {

			log.Debug("Client created..")

			serviceDetails, err := sdk.RequestServices()

			if err != nil {

				fmt.Println(err.Error())
			} else {

				if len(serviceDetails) >= 1 {

					svcDetails := serviceDetails[0]

					fmt.Printf("%d - %s\n", svcDetails.ServiceID, svcDetails.ServiceDescription)

					prices, err := sdk.GetServicePrices(svcDetails.ServiceID)

					if err != nil {

						fmt.Println(err.Error())
					} else {

						fmt.Printf("------- Prices -------\n")
						if len(prices) >= 1 {

							price := prices[0]

							fmt.Printf("(%d) %s @ %d, %s (Unit id = %d)\n", price.ID, price.Description, price.PricePerUnit, price.UnitDescription, price.UnitID)

							tpr, err := sdk.SelectService(price.ServiceID, 2, price.ID)

							if err != nil {

								fmt.Printf("%q\n", err.Error())


							} else {

								fmt.Println("#Begin Request#")
								fmt.Printf("ServerID: %s\n", tpr.ServerID)
								fmt.Printf("PriceID = %d - %d units = %d\n", tpr.PriceID, tpr.UnitsToSupply, tpr.TotalPrice)
								fmt.Printf("ClientID: %s, MerchantClientKey: %s, PaymentRef: %s\n", tpr.ClientID, tpr.MerchantClientKey, tpr.PaymentReferenceID)
								fmt.Println("#End Request#")

								log.Debug("Making payment of %d\n", tpr.TotalPrice)

								payResp, err := sdk.MakePayment(tpr)

								if err != nil {

									fmt.Printf("Error making payment: %s\n", err)
								} else {

									fmt.Printf("Payment of %d made successfully\n", payResp.TotalPaid)

									fmt.Printf("Service delivery token: %s\n", payResp.ServiceDeliveryToken)
								}
							}
						}
					}
				}
			}
		}
	}
}

func mResetSessionState() {

	fmt.Println("Not implemented yet..")
}

func mLoadConfig() {

	// Ask user for path to config file
	// (And password if secured)

	fmt.Println("Not implemented yet..")
}

func mReadConfig() {

	// Print out loaded configuration
	// Print out the path to file that was loaded (Need to keep reference during load stage)

	fmt.Println("Not implemented yet..")
}