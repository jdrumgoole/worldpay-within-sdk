package wpwithin

import (
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/core"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"

	log "github.com/Sirupsen/logrus"
)

// Factory to allow easy creation of
var Factory core.SDKFactory

// WPWithin Worldpay Within SDK
type WPWithin interface {
	AddService(service *types.Service) error
	RemoveService(service *types.Service) error
	InitConsumer(scheme, hostname string, portNumber int, urlPrefix, serverID string, hceCard *types.HCECard) error
	InitProducer(merchantClientKey, merchantServiceKey string) error
	GetDevice() *types.Device
	StartServiceBroadcast(timeoutMillis int) error
	StopServiceBroadcast()
	DeviceDiscovery(timeoutMillis int) ([]types.ServiceMessage, error)
	RequestServices() ([]types.ServiceDetails, error)
	GetServicePrices(serviceID int) ([]types.Price, error)
	SelectService(serviceID, numberOfUnits, priceID int) (types.TotalPriceResponse, error)
	MakePayment(payRequest types.TotalPriceResponse) (types.PaymentResponse, error)
	BeginServiceDelivery(clientID string, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int) error
	EndServiceDelivery(clientID string, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) error
}

// Initialise Initialise the SDK - Returns an implementation of WPWithin
// Must provide a device name and description
func Initialise(name, description string) (WPWithin, error) {

	defer func() {
		if r := recover(); r != nil {

			log.WithFields(log.Fields{"panic_message": r, "name": name, "description": description, "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.Initialise()")
		}
	}()

	// Parameter validation

	if name == "" {

		return nil, errors.New("name should not be empty")

	} else if description == "" {

		return nil, errors.New("description should not be empty")
	}

	// Start initialisation tasks

	if Factory == nil {

		_Factory, err := core.NewSDKFactory()
		Factory = _Factory

		if err != nil {

			return nil, fmt.Errorf("Unable to create SDK Factory: %q", err.Error())
		}
	}

	result := &wpWithinImpl{}

	core, err := core.NewCore()

	if err != nil {

		return result, err
	}

	result.core = core

	dev, err := Factory.GetDevice(name, description)

	if err != nil {

		return result, err
	}

	result.core.Device = dev

	om, err := Factory.GetOrderManager()

	if err != nil {

		return result, err

	}

	result.core.OrderManager = om

	bc, err := Factory.GetSvcBroadcaster(result.core.Device.IPv4Address)

	if err != nil {

		return result, err

	}

	result.core.SvcBroadcaster = bc

	sc, err := Factory.GetSvcScanner()

	if err != nil {

		return result, err

	}

	result.core.SvcScanner = sc

	return result, nil
}

type wpWithinImpl struct {
	core *core.Core
}

func (wp *wpWithinImpl) AddService(service *types.Service) error {

	defer func() {
		if r := recover(); r != nil {

			log.WithFields(log.Fields{"panic_message": r, "service": fmt.Sprintf("%+v", service), "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.AddService()")
		}
	}()

	if wp.core.Device.Services == nil {

		wp.core.Device.Services = make(map[int]*types.Service, 0)
	}

	if _, exists := wp.core.Device.Services[service.Id]; exists {

		return errors.New("Service with that id already exists")
	}

	wp.core.Device.Services[service.Id] = service

	return nil
}

func (wp *wpWithinImpl) RemoveService(service *types.Service) error {

	defer func() {
		if r := recover(); r != nil {

			log.WithFields(log.Fields{"panic_message": r, "service": fmt.Sprintf("%+v", service), "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.RemoveService()")
		}
	}()

	if wp.core.Device.Services != nil {

		delete(wp.core.Device.Services, service.Id)
	}

	return nil
}

func (wp *wpWithinImpl) InitConsumer(scheme, hostname string, portNumber int, urlPrefix, serverID string, hceCard *types.HCECard) error {

	defer func() {
		if r := recover(); r != nil {

			log.WithFields(log.Fields{"panic_message": r, "scheme": scheme, "hostname": hostname, "port": portNumber,
				"urlPrefix": urlPrefix, "serverID": serverID, "hceCard": fmt.Sprintf("%+v", hceCard), "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.InitConsumer()")
		}
	}()

	// Setup PSP as client

	_psp, err := Factory.GetPSPClient()

	if err != nil {

		return err
	}

	wp.core.Psp = _psp

	// Set core HCE Card

	wp.core.HCECard = hceCard

	// Setup HTE Client

	httpHTE, err := Factory.GetHTEClientHTTP()

	if err != nil {

		return err
	}

	client, err := hte.NewClient(scheme, hostname, portNumber, urlPrefix, serverID, httpHTE)

	if err != nil {

		return err
	}

	wp.core.HTEClient = client

	return nil
}

func (wp *wpWithinImpl) InitProducer(merchantClientKey, merchantServiceKey string) error {

	defer func() {
		if r := recover(); r != nil {

			log.WithFields(log.Fields{"panic_message": r, "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.InitProducer()")
		}
	}()

	// Parameter validation

	if merchantClientKey == "" {

		return errors.New("merchant client key should not be empty")
	} else if merchantServiceKey == "" {

		return errors.New("merchant service key should not be empty")
	}

	// Start HTE initialisation tasks

	psp, err := Factory.GetPSPMerchant(merchantClientKey, merchantServiceKey)

	if err != nil {

		return fmt.Errorf("Unable to create psp: %q", err.Error())
	}

	wp.core.Psp = psp

	hteCredential, err := hte.NewHTECredential(merchantClientKey, merchantServiceKey)

	if err != nil {

		return err
	}

	hteSvcHandler := Factory.GetHTEServiceHandler(wp.core.Device, wp.core.Psp, hteCredential, wp.core.OrderManager)

	svc, err := Factory.GetHTE(wp.core.Device, wp.core.Psp, wp.core.Device.IPv4Address, hteCredential, wp.core.OrderManager, hteSvcHandler)

	if err != nil {

		return err
	}

	wp.core.HTE = svc

	// Error channel allows us to get the error out of the go routine
	chStartResult := make(chan error)
	var startErr error

	go func() {

		chStartResult <- wp.core.HTE.Start()

	}()

	// Receive the error from the channel or wait a predefined amount of time
	// TODO CH : Fix this race condition - Matthew B has a solution, find and implement.
	select {

	case res := <-chStartResult:

		startErr = res

	case <-time.After(time.Millisecond * 750):

	}

	return startErr
}

func (wp *wpWithinImpl) GetDevice() *types.Device {

	defer func() {
		if r := recover(); r != nil {

			log.WithField("Stack", string(debug.Stack())).Errorf("Recover: WPWithin.GetDevice()")
		}
	}()

	return wp.core.Device
}

func (wp *wpWithinImpl) StartServiceBroadcast(timeoutMillis int) error {

	defer func() {
		if r := recover(); r != nil {

			fmt.Print(string(debug.Stack()))

			log.WithFields(log.Fields{"panic_message": r, "timeoutMillis": timeoutMillis, "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.StartServiceBroadcast()")
		}
	}()

	// Setup message that is broadcast over network
	msg := types.ServiceMessage{

		DeviceDescription: wp.core.Device.Description,
		Hostname:          wp.core.HTE.IPAddr(),
		ServerID:          wp.core.Device.Uid,
		UrlPrefix:         wp.core.HTE.UrlPrefix(),
		PortNumber:        wp.core.HTE.Port(),
	}

	// Set up a channel to get the error out of the go routine
	chBroadcastErr := make(chan error)
	var errBroadcast error

	go func() {

		chBroadcastErr <- wp.core.SvcBroadcaster.StartBroadcast(msg, timeoutMillis)
	}()

	// Either get the error or wait a small amount of time to give the all clear.
	// This is a race condition - ahhhh! TODO CH : Fix this
	select {

	case res := <-chBroadcastErr:

		errBroadcast = res

	case <-time.After(time.Millisecond * 750):

	}

	return errBroadcast
}

func (wp *wpWithinImpl) StopServiceBroadcast() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithField("Stack", string(debug.Stack())).Errorf("Recover: WPWithin.StopServiceBroadcast()")
		}
	}()

	wp.core.SvcBroadcaster.StopBroadcast()
}

func (wp *wpWithinImpl) DeviceDiscovery(timeoutMillis int) ([]types.ServiceMessage, error) {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithFields(log.Fields{"panic_message": r, "timeoutMillis": timeoutMillis, "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.DeviceDiscovery()")
		}
	}()

	var svcResults []types.ServiceMessage

	if scanResult, err := wp.core.SvcScanner.ScanForServices(timeoutMillis); err != nil {

		return nil, err

	} else if len(scanResult) > 0 {

		// Convert map of services to array
		for _, svc := range scanResult {

			svcResults = append(svcResults, svc)
		}
	}

	return svcResults, nil
}

func (wp *wpWithinImpl) GetServicePrices(serviceID int) ([]types.Price, error) {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithFields(log.Fields{"panic_message": r, "serviceID": serviceID, "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.GetServicePrices()")
		}
	}()

	var result []types.Price

	priceResponse, err := wp.core.HTEClient.GetPrices(serviceID)

	if err != nil {

		return nil, err
	}

	for _, price := range priceResponse.Prices {

		result = append(result, price)
	}

	return result, nil
}

func (wp *wpWithinImpl) SelectService(serviceID, numberOfUnits, priceID int) (types.TotalPriceResponse, error) {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithFields(log.Fields{"panic_message": r, "serviceID": serviceID, "numberOfUnits": numberOfUnits, "priceID": priceID, "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.SelectService()")
		}
	}()

	tpr, err := wp.core.HTEClient.NegotiatePrice(serviceID, priceID, numberOfUnits)

	return tpr, err
}

func (wp *wpWithinImpl) MakePayment(request types.TotalPriceResponse) (types.PaymentResponse, error) {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithFields(log.Fields{"panic_message": r, "price request": fmt.Sprintf("%+v", request), "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.MakePayment()")
		}
	}()

	token, err := wp.core.Psp.GetToken(wp.core.HCECard, request.MerchantClientKey, false)

	if err != nil {

		return types.PaymentResponse{}, err
	}

	paymentResponse, err := wp.core.HTEClient.MakeHtePayment(request.PaymentReferenceID, request.ClientID, token)

	return paymentResponse, err
}

func (wp *wpWithinImpl) RequestServices() ([]types.ServiceDetails, error) {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithField("Stack", string(debug.Stack())).Errorf("Recover: WPWithin.RequestServices()")
		}
	}()

	var result []types.ServiceDetails

	serviceResponse, err := wp.core.HTEClient.DiscoverServices()

	if err != nil {

		return nil, err
	}

	for _, svc := range serviceResponse.Services {

		result = append(result, svc)
	}

	return result, nil
}

func (wp *wpWithinImpl) Core() (*core.Core, error) {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithField("stack", string(debug.Stack())).Errorf("Recover: WPWithin.Core()")
		}
	}()

	return wp.core, nil
}

func (wp *wpWithinImpl) BeginServiceDelivery(clientID string, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int) error {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithFields(log.Fields{"panic_message": r, "clientID": clientID, "unitsToSupply": unitsToSupply,
				"serviceDeliveryToken": fmt.Sprintf("%+v", serviceDeliveryToken), "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.BeginServiceDelivery()")
		}
	}()

	return errors.New("BeginServiceDelivery() not yet implemented..")
}

func (wp *wpWithinImpl) EndServiceDelivery(clientID string, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) error {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithFields(log.Fields{"panic_message": r, "clientID": clientID, "unitsReceived": unitsReceived,
				"serviceDeliveryToken": fmt.Sprintf("%+v", serviceDeliveryToken), "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.EndServiceDelivery()")
		}
	}()

	return errors.New("EndServiceDelivery() not yet implemented..")
}
