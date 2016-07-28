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

	fmt.Println("Device Info:")

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
			fmt.Printf("      %d: ServiceID: %d ID:%d Description:%s PricePerUnit:%d UnitID:%d UnitDescription:%s\n", j, service.Id, price.ID, price.Description, price.PricePerUnit, price.UnitID, price.UnitDescription)
		}
	}

	fmt.Printf("IPv4Address: %s\n", sdk.GetDevice().IPv4Address)

	return 0, nil
}

func mInitDefaultDevice() (int, error) {

	fmt.Println("Initialising default device...")

	_sdk, err := wpwithin.Initialise(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION)

	if err != nil {

		return 0, err
	}

	sdk = _sdk

	return 0, nil
}

func mInitNewDevice() (int, error) {

	fmt.Println("Initialising new device")

	fmt.Print("Name of device: ")
	var nameOfDevice string
	if _, err := getUserInput(&nameOfDevice); err != nil {
		return 0, err
	}

	fmt.Print("Description: ")
	var description string
	if _, err := getUserInput(&description); err != nil {
		return 0, err
	}

	_sdk, err := wpwithin.Initialise(nameOfDevice, description)

	if err != nil {

		return 0, err
	}

	sdk = _sdk

	return 0, err
}

func mResetSessionState() (int, error) {

	fmt.Println("Resetting session state")

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

	fmt.Println("Starting rpc service...")

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

	// Error channel allows us to get the error out of the go routine
	chErr := make(chan error, 1)

	go func() {
		chErr <- rpc.Start()
	}()

	// Error handling go routine
	go func() {
		err := <-chErr
		if err != nil {
			log.Debug("error ", err)
		}

		close(chErr)
	}()

	// return here (error will be logged if it occurs)
	return 0, nil
}
