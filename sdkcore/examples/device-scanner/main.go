package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin"
)

var wpw wpwithin.WPWithin

func main() {

	initLog()

	wp, err := wpwithin.Initialise("conorhwp-pi", "Conor Hacketts Raspberry Pi - DEVICE SCANNER")
	wpw = wp

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	done := make(chan bool)
	fnForever := func() {
		for {

			fmt.Println("Scanning network for devices now...")
			devices, err := wpw.DeviceDiscovery(5000)

			if err != nil {
				fmt.Println("Error finding devices..")
				fmt.Println(err.Error())
			} else if len(devices) > 0 {

				fmt.Println("------------------------------------------------------------")
				fmt.Printf("Found %d devices\n", len(devices))

				for _, dev := range devices {

					fmt.Printf("[%s] %s @ %s:%d%s\n", dev.ServerID, dev.DeviceDescription, dev.Hostname, dev.PortNumber, dev.UrlPrefix)
				}
				fmt.Println("------------------------------------------------------------")
			}

			// time.Sleep(time.Second * 5)
		}
	}

	go fnForever()

	<-done // Block forever
}

func initLog() {

	log.SetFormatter(&log.JSONFormatter{})

	f, err := os.OpenFile("wpwithin.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {

		fmt.Println(err.Error())
	}

	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	log.Debug("Log is initialised.")
}
