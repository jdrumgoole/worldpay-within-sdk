package main
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"
)

func mScanService() {

	log.Debug("testDiscoveryAndNegotiation")

	sdk, err := wpwithin.Initialise("conorhwp-macbook", "Conor H WP - Raspberry Pi")

	if err != nil {

		fmt.Println(err)
		return
	}

	err = sdk.InitHTE("T_C_c93d7723-2b1c-4dd2-bfb7-58dd48cd093e", "T_S_6ec32d94-77fa-42ff-bede-de487d643793")

	if err != nil {

		fmt.Println(err)
		return
	}

	card := domain.HCECard{

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

	for _, svc := range services {

		fmt.Printf("(%s:%d/%s) - %s", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)
	}

}

func mDefaultHCECredential() {

	fmt.Println("Not implemented yet..")
}

func mDiscoverSvcs() {

	fmt.Println("Not implemented yet..")
}

func mGetSvcPrices() {

	fmt.Println("Not implemented yet..")
}

func mSelectService() {

	fmt.Println("Not implemented yet..")
}

func mMakePayment() {

	fmt.Println("Not implemented yet..")
}

func mConsumerStatus() {

	fmt.Println("Not implemented yet..")
}