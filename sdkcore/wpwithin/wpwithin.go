package wpwithin

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/configuration"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/core"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types/event"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/utils"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/utils/wslog"

	log "github.com/Sirupsen/logrus"
)

// Factory to allow easy creation of
var Factory core.SDKFactory

// WPWithin Worldpay Within SDK
type WPWithin interface {
	AddService(service *types.Service) error
	RemoveService(service *types.Service) error
	InitConsumer(scheme, hostname string, portNumber int, urlPrefix, clientID string, hceCard *types.HCECard) error
	InitProducer(merchantClientKey, merchantServiceKey string) error
	GetDevice() *types.Device
	StartServiceBroadcast(timeoutMillis int) error
	StopServiceBroadcast()
	DeviceDiscovery(timeoutMillis int) ([]types.BroadcastMessage, error)
	RequestServices() ([]types.ServiceDetails, error)
	GetServicePrices(serviceID int) ([]types.Price, error)
	SelectService(serviceID, numberOfUnits, priceID int) (types.TotalPriceResponse, error)
	MakePayment(payRequest types.TotalPriceResponse) (types.PaymentResponse, error)
	BeginServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int) (types.ServiceDeliveryToken, error)
	EndServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) (types.ServiceDeliveryToken, error)
	SetEventHandler(handler event.Handler) error
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

	// Parse configuration
	rawCfg, err := configuration.Load("wpwconfig.json")
	wpwConfig := configuration.WPWithin{}
	wpwConfig.ParseConfig(rawCfg)
	core.Configuration = wpwConfig

	doWebSocketLogSetup(core.Configuration)

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


	// MongoDB Diff

	err = result.core.Logger.Initialise()
	
	if err != nil {
		return result, err
	}

	result.core.Logger.LogEventDoc( "WPWithin.initialise", "result", result.core ) 
	
	return result, nil
}

type wpWithinImpl struct {
	core *core.Core
}

func (wp *wpWithinImpl) AddService(service *types.Service) error {

	fmt.Println( "in AddService" ) 
	
	defer func() {
		if r := recover(); r != nil {

			log.WithFields(log.Fields{"panic_message": r, "service": fmt.Sprintf("%+v", service), "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.AddService()")
		}
	}()

	if wp.core.Device.Services == nil {

		wp.core.Device.Services = make(map[int]*types.Service, 0)
	}

	if _, exists := wp.core.Device.Services[service.ID]; exists {

		return errors.New("Service with that id already exists")
	}

	wp.core.Device.Services[service.ID] = service
	
	wp.core.Logger.LogEventDoc( "WPWithin.AddService", "Service", service )
	
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

		delete(wp.core.Device.Services, service.ID)
	}
	
	wp.core.Logger.LogEventDoc( "WPWithin.RemoveService", "Service", service )
	return nil
}

func (wp *wpWithinImpl) InitConsumer(scheme, hostname string, portNumber int, urlPrefix, clientID string, hceCard *types.HCECard) error {

	defer func() {
		if r := recover(); r != nil {

			log.WithFields(log.Fields{"panic_message": r, "scheme": scheme, "hostname": hostname, "port": portNumber,
				"urlPrefix": urlPrefix, "clientID": clientID, "hceCard": fmt.Sprintf("%+v", hceCard), "stack": fmt.Sprintf("%s", debug.Stack())}).
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

	client, err := hte.NewClient(scheme, hostname, portNumber, urlPrefix, clientID, httpHTE)

	if err != nil {

		return err
	}

	wp.core.HTEClient = client

	// MongoDB Diff

	wp.core.Logger.LogEventDoc( "WPWithin.InitConsumer", "Service", clientID )
	
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

	hteSvcHandler := Factory.GetHTEServiceHandler(wp.core.Device, wp.core.Psp, hteCredential, wp.core.OrderManager, wp.core.EventHandler)

	svc, err := Factory.GetHTE(wp.core.Device, wp.core.Psp, wp.core.Device.IPv4Address, "http://", hteCredential, wp.core.OrderManager, hteSvcHandler)

	if err != nil {

		return err
	}

	wp.core.HTE = svc

	// MongoDB Diff

	wp.core.Logger.LogEventDoc( "WPWithin.initProducer",
		                        "producer", svc )

	if err != nil {
		fmt.Errorf("Unable to insert into MongoDB : %q", err.Error())
		return err
	}
	// end MongoDB Diff
	
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
//
//	wp.core.Logger.LogEventDoc( "WPWithin.GetDevice",
//		                        "device", wp.core.Device )
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
	msg := types.BroadcastMessage{

		DeviceDescription: wp.core.Device.Description,
		Hostname:          wp.core.HTE.IPAddr(),
		ServerID:          wp.core.Device.UID,
		URLPrefix:         wp.core.HTE.URLPrefix(),
		PortNumber:        wp.core.HTE.Port(),
		Scheme:            wp.core.HTE.Scheme(),
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
	wp.core.Logger.LogEventDoc( "WPWithin.StartServiceBroadcast",
		                        "message", msg )
	
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
	
	wp.core.Logger.LogEventStr( "WPWithin.StopServiceBroadcast",
		                        "stopping" )
}

func (wp *wpWithinImpl) DeviceDiscovery(timeoutMillis int) ([]types.BroadcastMessage, error) {

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithFields(log.Fields{"panic_message": r, "timeoutMillis": timeoutMillis, "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.DeviceDiscovery()")
		}
	}()

	var svcResults []types.BroadcastMessage

	if scanResult, err := wp.core.SvcScanner.ScanForServices(timeoutMillis); err != nil {

		return nil, err

	} else if len(scanResult) > 0 {

		// Convert map of services to array
		for _, svc := range scanResult {

			svcResults = append(svcResults, svc)
		}
	}

	wp.core.Logger.LogEventDoc( "WPWithin.DeviceDiscovery",
		                        "devices", svcResults )
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

	wp.core.Logger.LogEventDoc( "WPWithin.GetServicePrices",
		                        "prices", result )
	
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

	wp.core.Logger.LogEventDoc( "WPWithin.SelectService",
		                        "Total Price Response", tpr  )
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

	if err != nil {
		return types.PaymentResponse{}, err
	}
	
	wp.core.Logger.LogEventDoc( "WPWithin.MakePayment", "Payment Request", paymentResponse )
	wp.core.Logger.LogEventDoc( "WPWithin.MakePayment", "Payment Response", paymentResponse )

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
	
	wp.core.Logger.LogEventDoc( "WPWithin.RequestServices", "result", result )
	
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

func (wp *wpWithinImpl) BeginServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int) (types.ServiceDeliveryToken, error) {

	log.WithFields(log.Fields{"serviceID": serviceID, "serviceDeliveryToken": fmt.Sprintf("%+v", serviceDeliveryToken), "unitsToSupply": unitsToSupply}).Debug("begin wpwithin.wpwithinimpl.BeginServiceDelivery()")

	defer log.Debug("end wpwithin.wpwithinimpl.BeginServiceDelivery()")

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithFields(log.Fields{"panic_message": r, "serviceID": serviceID, "unitsToSupply": unitsToSupply,
				"serviceDeliveryToken": fmt.Sprintf("%+v", serviceDeliveryToken), "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.BeginServiceDelivery()")
		}
	}()

	deliveryResponse, err := wp.core.HTEClient.StartDelivery(serviceID, serviceDeliveryToken, unitsToSupply)

	if err != nil {

		log.Errorf("Error calling beginServiceDelivery. Error: %s", err.Error())
		return types.ServiceDeliveryToken{}, err
	}

	log.WithFields(log.Fields{"UnitsToSupply": deliveryResponse.UnitsToSupply}).Info("EndDeliveryResponse")

	wp.core.Logger.LogEventDoc( "WPWithin.BeginServiceDelivery", "DeliveryResponse", deliveryResponse )
	return deliveryResponse.ServiceDeliveryToken, nil
}

func (wp *wpWithinImpl) EndServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) (types.ServiceDeliveryToken, error) {

	log.WithFields(log.Fields{"serviceID": serviceID, "serviceDeliveryToken": fmt.Sprintf("%+v", serviceDeliveryToken), "unitsReceived": unitsReceived}).Debug("begin wpwithin.wpwithinimpl.EndServiceDelivery()")

	defer log.Debug("end wpwithin.wpwithinimpl.EndServiceDelivery()")

	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("%s", debug.Stack())

			log.WithFields(log.Fields{"panic_message": r, "serviceID": serviceID, "unitsReceived": unitsReceived,
				"serviceDeliveryToken": fmt.Sprintf("%+v", serviceDeliveryToken), "stack": fmt.Sprintf("%s", debug.Stack())}).
				Errorf("Recover: WPWithin.EndServiceDelivery()")
		}
	}()

	deliveryResponse, err := wp.core.HTEClient.EndDelivery(serviceID, serviceDeliveryToken, unitsReceived)

	wp.core.Logger.LogEventDoc( "WPWithin.EndServiceDelivery", "DeliveryResponse", deliveryResponse )
	
	if err != nil {

		log.Errorf("Error calling endServiceDelivery. Error: %s", err.Error())

		return types.ServiceDeliveryToken{}, err
	}

	log.WithFields(log.Fields{"UnitsJustSupplied": deliveryResponse.UnitsJustSupplied, "UnitsRemaining": deliveryResponse.UnitsRemaining}).Info("EndDeliveryResponse")

	return deliveryResponse.ServiceDeliveryToken, nil
}

func (wp *wpWithinImpl) SetEventHandler(handler event.Handler) error {

	log.Debug("wpwithin.wpwithinimpl setting core event handler")

	wp.core.EventHandler = handler

	return nil
}

func doWebSocketLogSetup(cfg configuration.WPWithin) {

	if cfg.WSLogEnable {

		// Clean up the levels config value - just in case.
		strLevels := strings.Replace(cfg.WSLogLevel, " ", "", -1)
		logLevels := strings.Split(strLevels, ",")

		// Support all levels
		var levels []log.Level

		for _, level := range logLevels {

			switch strings.ToLower(level) {

			case "panic":
				levels = append(levels, log.PanicLevel)
			case "fatal":
				levels = append(levels, log.FatalLevel)
			case "error":
				levels = append(levels, log.ErrorLevel)
			case "warn":
				levels = append(levels, log.WarnLevel)
			case "info":
				levels = append(levels, log.InfoLevel)
			case "debug":
				levels = append(levels, log.DebugLevel)
			}
		}
		ip, err := utils.ExternalIPv4()
		strIP := ""

		if err == nil {

			strIP = ip.String()
		} else {

			fmt.Printf("Error getting ExternalIPv4: %s\n", err.Error())
		}

		err = wslog.Initialise(strIP, cfg.WSLogPort, levels)

		if err != nil {

			fmt.Printf("Error initialising WebSocket logger: %s\n", err.Error())
		}
	}
}
