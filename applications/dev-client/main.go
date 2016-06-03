package main
import
(

	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
"github.com/gorilla/websocket"
	"os/signal"
	"time"
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp/onlineworldpay"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hte"
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/utils"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/utils/wslog"
)

func main() {

	initLog()

//	testWPPay()

//	testWPTokenise()

//	testBroadcast()

//	testDiscovery()

//	calculateBcastIP()

//	testUUID()

//	externalIP()

//	testWebSocketClient()

//	testWebSocketServer()

//	testHTEService()

//	testFun()

//	testHTEandBroadcast()

//	testDiscoveryAndNegotiation()

	doUI()

//	doWebSocketLogger()
}

func initLog() {

	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	log.Debug("Log is initialised.")
}

func testUUID() {

	device, err := wpwithin.Initialise("Conor-Macbook", "Macbook Pro laptop computer")

	if err != nil {

		fmt.Printf("%q", err.Error())

		return
	}

	fmt.Printf("UID: %s", device.GetDevice().Uid)

}

func testWPPay() {

	hteCred, err := hte.NewHTECredential("T_C_c93d7723-2b1c-4dd2-bfb7-58dd48cd093e", "T_S_6ec32d94-77fa-42ff-bede-de487d643793")

	if err != nil {

		fmt.Printf("%q", err.Error())
		return
	}

//	hceClient, err := hce.NewHCEClientCredential("T_C_c93d7723-2b1c-4dd2-bfb7-58dd48cd093e")
//
//	if err != nil {
//
//		fmt.Printf("%q", err.Error())
//		return
//	}

	psp, err := onlineworldpay.New(hteCred.MerchantClientKey, hteCred.MerchantServiceKey, "https://api.worldpay.com/v1")
//	psp, err := onlineworldpay.New(hteCred, "https://127.0.0.1:9000")

	if err != nil {

		fmt.Printf("%q", err.Error())
		return
	}

	card := &types.HCECard{

		FirstName:"Bilbo",
		LastName:"Baggins",
		ExpMonth:11,
		ExpYear:2018,
		CardNumber:"5555555555554444",
		Type:"Card",
		Cvc:"113",
	}

	token, err := psp.GetToken(card, false)

	if err != nil {

		fmt.Printf("%q", err.Error())
		return
	}

	fmt.Printf("Token: %s", token)

	payResponse, err := psp.MakePayment(199, "GBP", token, "Test txn", "12345")

	if err != nil {

		fmt.Printf("%q", err.Error())
		return
	}

	fmt.Printf("Payment result: %s", payResponse)


}

//func testWPSDK() {
//
//	fmt.Println("Hello WP Within..")
//	testChannel()
//	sdk, _ := wpwithin.Initialise("Conor-Macbook", "Macbook Pro laptop computer", hteCred)
//
//	svc, _ := types.NewService()
//
//	price1 := types.Price{
//
//		UnitID:1,
//		ID:123,
//		Description:"dasd",
//		UnitDescription:"sadas",
//		PricePerUnit:21,
//
//	}
//
//	svc.AddPrice(price1)
//
//	sdk.AddService(svc)
//
//	sdk.InitConsumer()
//
////	dev := sdk.GetDevice()
//
////	fmt.Printf("Device name: %s", dev.Name)
//}

func testDiscovery() {

//	sdk, _ := wpwithin.Initialise("Conor-Macbook", "Macbook Pro laptop computer", hteCred)
//
//	sdk.InitConsumer()
//
//	result, err := sdk.ScanServices(15000)
//
//	if err != nil {
//
//		fmt.Printf("%q", err.Error())
//	}
//
//	// Wait for scanning to complete
//	fmt.Println("************ Finished Scanning ************")
//
//	if err != nil {
//
//		fmt.Printf("Error scanning: %q", err)
//
//	} else {
//
//		fmt.Printf("Found %d services:\n", len(result))
//
//		if len(result) > 0 {
//
//			for i, service := range result {
//
//				fmt.Printf("%d - %#v\n", i, service)
//			}
//		}
//	}
}

func testBroadcast() {

	sdk, _ := wpwithin.Initialise("Conor-Macbook", "Macbook Pro laptop computer")

	svc, _ := types.NewService()

	price1 := types.Price{

		UnitID:1,
		ID:123,
		Description:"dasd",
		UnitDescription:"sadas",
		PricePerUnit:21,

	}

	svc.AddPrice(price1)

	sdk.AddService(svc)

	err := sdk.StartServiceBroadcast(21000)

	if err != nil {

		fmt.Printf("%q", err.Error())
	}

	fmt.Println("Finished broadcast..")
}

func testWebSocketClient() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial("ws://echo.websocket.org/", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {

		c.WriteMessage(websocket.TextMessage, []byte("Hello."))

		select {
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
		// To cleanly close a connection, a client should send a close
		// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
			c.Close()
			return
		}
	}
}

func calculateBcastIP() {

	addr, err := utils.ExternalIPv4()

	if err != nil {

		fmt.Printf("%q\n", err.Error())
	}

	nm, err := utils.NetMask()

	if err != nil {

		fmt.Printf("%q\n", err.Error())
	}

	fmt.Printf("IP: %s\n", addr.String())
	fmt.Printf("NM: %s\n", nm.String())
}

func testHTEService() {

	sdk, _ := wpwithin.Initialise("conorhwp-pi", "Conor H WP - Raspberry Pi")

	roboWash, _ := types.NewService()
	roboWash.Name = "RoboWash"
	roboWash.Description = "Car washed by robot"
	roboWash.Id = 1

	washPriceCar := types.Price{

		ServiceID:roboWash.Id,
		UnitID:1,
		ID:1,
		Description:"Car wash",
		UnitDescription:"Single wash",
		PricePerUnit:500,
	}

	washPriceSUV := types.Price{

		ServiceID:roboWash.Id,
		UnitID:1,
		ID:2,
		Description:"SUV Wash",
		UnitDescription:"Single wash",
		PricePerUnit:650,
	}

	roboWash.AddPrice(washPriceCar)
	roboWash.AddPrice(washPriceSUV)
	sdk.AddService(roboWash)

	roboAir, _ := types.NewService()
	roboAir.Name = "RoboAir"
	roboAir.Description = "Car tyre pressure checked and topped up by robot"
	roboAir.Id = 2

	airSinglePrice := types.Price{
		ServiceID: roboAir.Id,
		UnitID: 1,
		ID: 1,
		Description: "Measure and adjust pressure",
		UnitDescription:"Tyre",
		PricePerUnit:25,
	}

	airFourPrice := types.Price{
		ServiceID: roboAir.Id,
		UnitID: 2,
		ID: 2,
		Description: "Measure and adjust pressure - four tyres for the price of three",
		UnitDescription:"4 Tyre",
		PricePerUnit:airSinglePrice.PricePerUnit * 3,
	}

	roboAir.AddPrice(airSinglePrice)
	roboAir.AddPrice(airFourPrice)
	sdk.AddService(roboAir)

	_, err := sdk.InitProducer()

	if err != nil {

		fmt.Printf(err.Error())

		return
	}
	fmt.Printf("End.")
//	} else {
//
//		time.Sleep(time.Duration(100000) * time.Millisecond)
//		fmt.Printf("Stopped server")
//		go func(){done<-true}()
////		done<-true
//
//
//		time.Sleep(time.Duration(150000) * time.Millisecond)
//	}
//
//	fmt.Printf("finished..")
}

func testFun() {

	funs := make(map[string]func())

	funs["hello"] = printHello
	funs["world"] = printWorld

	funs["hello"]()
}

func printHello() {

	fmt.Println("Hello")
}

func printWorld() {

	fmt.Println("World")
}

func testWPTokenise() {

	hteCred, err := hte.NewHTECredential("T_C_c93d7723-2b1c-4dd2-bfb7-58dd48cd093e", "T_S_6ec32d94-77fa-42ff-bede-de487d643793")

	if err != nil {

		fmt.Printf("%q", err.Error())
		return
	}

	psp, err := onlineworldpay.New(hteCred.MerchantClientKey, hteCred.MerchantServiceKey, "https://api.worldpay.com/v1")

	if err != nil {

		fmt.Printf("%q", err.Error())
		return
	}

	card := &types.HCECard{

		FirstName:"Bilbo",
		LastName:"Baggins",
		ExpMonth:11,
		ExpYear:2018,
		CardNumber:"5555555555554444",
		Type:"Card",
		Cvc:"113",
	}

	token, err := psp.GetToken(card, false)

	if err != nil {

		fmt.Printf("%q", err.Error())
		return
	}

	fmt.Printf("Token: %s", token)
}

func testHTEandBroadcast() {

	sdk, _ := wpwithin.Initialise("conorhwp-pi", "Conor H WP - Raspberry Pi")

	roboWash, _ := types.NewService()
	roboWash.Name = "RoboWash"
	roboWash.Description = "Car washed by robot"
	roboWash.Id = 1

	washPriceCar := types.Price{

		ServiceID:roboWash.Id,
		UnitID:1,
		ID:1,
		Description:"Car wash",
		UnitDescription:"Single wash",
		PricePerUnit:500,
	}

	washPriceSUV := types.Price{

		ServiceID:roboWash.Id,
		UnitID:1,
		ID:2,
		Description:"SUV Wash",
		UnitDescription:"Single wash",
		PricePerUnit:650,
	}

	roboWash.AddPrice(washPriceCar)
	roboWash.AddPrice(washPriceSUV)
	sdk.AddService(roboWash)

	roboAir, _ := types.NewService()
	roboAir.Name = "RoboAir"
	roboAir.Description = "Car tyre pressure checked and topped up by robot"
	roboAir.Id = 2

	airSinglePrice := types.Price{
		ServiceID: roboAir.Id,
		UnitID: 1,
		ID: 1,
		Description: "Measure and adjust pressure",
		UnitDescription:"Tyre",
		PricePerUnit:25,
	}

	airFourPrice := types.Price{
		ServiceID: roboAir.Id,
		UnitID: 2,
		ID: 2,
		Description: "Measure and adjust pressure - four tyres for the price of three",
		UnitDescription:"4 Tyre",
		PricePerUnit:airSinglePrice.PricePerUnit * 3,
	}

	roboAir.AddPrice(airSinglePrice)
	roboAir.AddPrice(airFourPrice)
	sdk.AddService(roboAir)

	prodDone := make(chan bool)

	go func() {

		_, err := sdk.InitProducer()

		if err != nil {

			fmt.Printf(err.Error())

			return
		}
	}()

	bcastDone := make(chan bool)

	go func() {

		sdk.StartServiceBroadcast(20000)
	}()

	<-prodDone
	<-bcastDone

	fmt.Printf("End.")
}

func testDiscoveryAndNegotiation() {

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

	for a, svc := range services {

		fmt.Println("# Service:: (%s:%d/%s) - %s", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.DeviceDescription)

		log.Debug("Init consumer")
		err := sdk.InitConsumer("http://", svc.Hostname, svc.PortNumber, svc.UrlPrefix, svc.ServerID)

		if err != nil {

			fmt.Println(err.Error())
			continue
		} else {

			log.Debug("Client created..")

			serviceDetails, err := sdk.RequestServices()

			if err != nil {

				fmt.Println(err.Error())
				continue
			} else {

				for b, svc := range serviceDetails {

					fmt.Printf("%d - %s\n", svc.ServiceID, svc.ServiceDescription)

					prices, err := sdk.GetServicePrices(svc.ServiceID)

					if err != nil {

						fmt.Println(err.Error())
					} else {

						fmt.Printf("------- Prices -------\n")
						for c, price := range prices {

							fmt.Printf("(%d) %s @ %d, %s (Unit id = %d)\n", price.ID, price.Description, price.PricePerUnit, price.UnitDescription, price.UnitID)

							tpr, err := sdk.SelectService(price.ServiceID, 2, price.ID)

							if err != nil {

								fmt.Printf("%q\n", err.Error())

								continue
							}

							fmt.Println("#Begin Request#")
							fmt.Printf("ServerID: %s\n", tpr.ServerID)
							fmt.Printf("PriceID = %d - %d units = %d\n", tpr.PriceID, tpr.UnitsToSupply, tpr.TotalPrice)
							fmt.Printf("ClientID: %s, MerchantClientKey: %s, PaymentRef: %s\n", tpr.ClientID, tpr.MerchantClientKey, tpr.PaymentReferenceID)
							fmt.Println("#End Request#")

							if a == 0 && b == 0 && c == 0 {

								log.Debug("Making payment of %d", tpr.TotalPrice)

//								payResp, err := sdk.MakePayment(tpr)
//
//								if err != nil {
//
//									fmt.Printf("Error making payment: %s", err)
//								} else {
//
//									fmt.Printf("Payment of %d made successfully", payResp.TotalPaid)
//
//									fmt.Printf("Service delivery token: %s", payResp.ServiceDeliveryToken)
//								}
							}
						}
						fmt.Printf("----- End prices -----\n")
					}
				}
			}
		}
	}
}

func doWebSocketLogger() {

	// Support all levels
	levels := make([]log.Level, 0)
	levels = append(levels, log.PanicLevel)
	levels = append(levels, log.FatalLevel)
	levels = append(levels, log.ErrorLevel)
	levels = append(levels, log.WarnLevel)
	levels = append(levels, log.InfoLevel)
	levels = append(levels, log.DebugLevel)

	err := wslog.Initialise("127.0.0.1", 8181, levels)

	if err != nil {

		fmt.Println(err.Error())

		return
	}
	
	time.Sleep(time.Duration(10) * time.Second)

	log.Debug("This is debug :)")
}