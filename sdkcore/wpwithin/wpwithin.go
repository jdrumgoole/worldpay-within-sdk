package wpwithin
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/utils"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/core"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp/onlineworldpay"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/fsm"
	"errors"
)

const (

	BROADCAST_STEP_SLEEP = 5000
	BROADCAST_PORT = 8980
	HTE_SVC_URL_PREFIX = ""
	UUID_FILE_PATH = "uuid.txt"
	HTE_SVC_PORT = 8080
	WP_ONLINE_API_ENDPOINT = "https://api.worldpay.com/v1"
	HTE_CLIENT_SCHEME = "http://"
)

type WPWithin interface {

	AddService(service *types.Service) error
	RemoveService(service *types.Service) error
	InitHCE(hceCard types.HCECard) error
	InitHTE(merchantClientKey, merchantServiceKey string) error
	InitConsumer(scheme, hostname string, portNumber int, urlPrefix, serverID string) error
	InitProducer() (chan bool, error)
	GetDevice() (*types.Device, error)
	StartServiceBroadcast(timeoutMillis int) error
	StopServiceBroadcast() error
	ServiceDiscovery(timeoutMillis int) ([]types.ServiceMessage, error)
	RequestServices() ([]types.ServiceDetails, error)
	GetServicePrices(serviceId int) ([]types.Price, error)
	SelectService(serviceId, numberOfUnits, priceId int) (types.TotalPriceResponse, error)
	MakePayment(payRequest types.TotalPriceResponse) (types.PaymentResponse, error)

}

type wpWithinImpl struct {

	core *core.Core
}

func Initialise(name, description string) (WPWithin, error) {

	var err error

	// Set up SDK core

	core, err := core.New()

	if err != nil {

		return &wpWithinImpl{}, err
	}

	// Setup Finite State Machine - Begin at device not ready state.

	core.FSM, _ = fsm.Init(fsm.DEV_NOT_READY)
	core.FSMHelper = fsm.NewSDKHelper()

	// Add core and device to WPWithin SDK Implementation
	wp := &wpWithinImpl{}
	wp.core = core

	// Device setup

	var deviceGUID string

	if b, _ := utils.FileExists(UUID_FILE_PATH); b {

		deviceGUID, err = utils.ReadLocalUUID(UUID_FILE_PATH)
	} else {

		deviceGUID, err = utils.NewUUID()

		utils.WriteString(UUID_FILE_PATH, deviceGUID, true)
	}

	if err != nil {

		return &wpWithinImpl{}, err
	}

	deviceAddress, err := utils.ExternalIPv4()

	if err != nil {

		return &wpWithinImpl{}, err
	}

	device, err := types.NewDevice(name, description, deviceGUID, deviceAddress.String(), "GBP")

	if err != nil {

		return &wpWithinImpl{}, err
	}

	core.Device = device

	// Setup Order Manager

	orderManager, err := hte.NewOrderManager()

	if err != nil {

		return &wpWithinImpl{}, err
	}

	core.OrderManager = orderManager

	// Service broadcaster

	svcBroadcaster, err := servicediscovery.NewBroadcaster(core.Device.IPv4Address, BROADCAST_PORT, BROADCAST_STEP_SLEEP)

	if err != nil {

		return &wpWithinImpl{}, err
	}

	core.SvcBroadcaster = svcBroadcaster

	// Service scanner

	svcScanner, err := servicediscovery.NewScanner(BROADCAST_PORT, BROADCAST_STEP_SLEEP)

	if err != nil {

		return &wpWithinImpl{}, err
	}

	core.SvcScanner = svcScanner

	core.HCE = &hce.Manager{}

	// Setup complete - transition to device ready state
	return wp, core.FSM.Transition(fsm.DEV_READY)
}

func (wp *wpWithinImpl) InitHTE(merchantClientKey, merchantServiceKey string) error {

	stateGoal := fsm.PRO_READY

	if valid := wp.core.FSM.Permitted(stateGoal); !valid {

		return errors.New("Invalid state")
	}

	// Set up PSP
	psp, err := onlineworldpay.New(merchantClientKey, merchantServiceKey, WP_ONLINE_API_ENDPOINT)

	if err != nil {

		return err
	}

	wp.core.Psp = psp

	// Set up HTE service

	hteCredential, err := hte.NewHTECredential(merchantClientKey, merchantServiceKey)

	if err != nil {

		return err
	}

	hte, err := hte.NewService(wp.core.Device, psp, wp.core.Device.IPv4Address, HTE_SVC_URL_PREFIX, HTE_SVC_PORT, hteCredential, wp.core.OrderManager)

	if err != nil {

		return err
	}

	wp.core.HTE = hte

	// Update the state machine
	return wp.core.FSM.Transition(stateGoal)

}

func (wp *wpWithinImpl) AddService(service *types.Service) error {

	stateGoal := fsm.PRO_READY

	if valid := wp.core.FSM.Permitted(stateGoal); !valid {

		return errors.New("Invalid state")
	}

	if wp.core.Device.Services == nil {

		wp.core.Device.Services = make(map[int]*types.Service, 0)
	}

	wp.core.Device.Services[service.Id] = service

	return wp.core.FSM.Transition(stateGoal)
}

func (wp *wpWithinImpl) RemoveService(service *types.Service) error {

	stateGoal := fsm.PRO_READY

	if valid := wp.core.FSM.Permitted(stateGoal); !valid {

		return errors.New("Invalid state")
	}

	if wp.core.Device.Services != nil {

		delete(wp.core.Device.Services, service.Id)
	}

	return wp.core.FSM.Transition(stateGoal)
}

func (wp *wpWithinImpl) InitConsumer(scheme, hostname string, portNumber int, urlPrefix, serverID string) error {

	stateGoal := fsm.CON_READY

	if valid := wp.core.FSM.Permitted(stateGoal); !valid {

		return errors.New("Invalid state")
	}

	// Setup HTE Client

	client, err := hte.NewClient(scheme, hostname, portNumber, urlPrefix, serverID)

	if err != nil {

		return err
	}

	wp.core.HTEClient = client

	return wp.core.FSM.Transition(stateGoal)
}

func (wp *wpWithinImpl) InitProducer() (chan bool, error) {

	stateGoal := fsm.PRO_READY

	if valid := wp.core.FSM.Permitted(stateGoal); !valid {

		return nil, errors.New("Invalid state")
	}

	err := wp.core.HTE.Start()

	if err != nil {

		return nil, err
	}

	done := make(chan bool)

	err = wp.core.FSM.Transition(stateGoal)

	return done, err
}

func (wp *wpWithinImpl) GetDevice() (*types.Device, error) {

	stateGoal := fsm.DEV_READY

	if valid := wp.core.FSM.Permitted(stateGoal); !valid {

		return nil, errors.New("Invalid state")
	}

	return wp.core.Device, nil
}

func (wp *wpWithinImpl) InitHCE(hceCardCredential types.HCECard) error {

	stateGoal := fsm.CON_READY

	if valid := wp.core.FSM.Permitted(stateGoal); !valid {

		return errors.New("Invalid state")
	}

	cred := new(types.HCECard)
	cred.FirstName = hceCardCredential.FirstName
	cred.LastName = hceCardCredential.LastName
	cred.ExpMonth = hceCardCredential.ExpMonth
	cred.ExpYear = hceCardCredential.ExpYear
	cred.CardNumber = hceCardCredential.CardNumber
	cred.Type = hceCardCredential.Type
	cred.Cvc = hceCardCredential.Cvc

	wp.core.HCE.HCECard = cred

	return wp.core.FSM.Transition(fsm.CON_READY)
}

func (wp *wpWithinImpl) StartServiceBroadcast(timeoutMillis int) error {

	intermedGoal := fsm.PRO_BROADCAST
	endGoal := fsm.PRO_READY

	if valid := wp.core.FSM.Permitted(intermedGoal); !valid {

		return errors.New("Invalid state")
	}

	// Setup message that is broadcast over network
	msg := types.ServiceMessage{

		DeviceDescription: wp.core.Device.Description,
		Hostname: wp.core.HTE.IPv4Address,
		ServerID: wp.core.Device.Uid,
		UrlPrefix: wp.core.HTE.UrlPrefix,
		PortNumber:wp.core.HTE.Port,
	}

	complete, err := wp.core.SvcBroadcaster.StartBroadcast(msg, timeoutMillis)

	// Intermediate goal
	err = wp.core.FSM.Transition(intermedGoal)

	if err != nil {

		return err
	}

	// Wait for broadcast to complete
	<-complete

	// End goal
	return wp.core.FSM.Transition(endGoal)
}

func (wp *wpWithinImpl) StopServiceBroadcast() error {

	if err := wp.core.FSM.Transition(fsm.PRO_READY); err != nil {

		return err
	}

	return wp.core.SvcBroadcaster.StopBroadcast()
}

func (wp *wpWithinImpl) ServiceDiscovery(timeoutMillis int) ([]types.ServiceMessage, error) {

	intermedState := fsm.CON_DISCOVER_DEV
	notFoundState := fsm.DEV_READY
	FoundState := fsm.CON_DEV_AVAILABLE

	if valid := wp.core.FSM.Permitted(intermedState); !valid {

		return nil, errors.New("Invalid state")
	}

	svcResults := make([]types.ServiceMessage, 0)

	scanResult := wp.core.SvcScanner.ScanForServices(timeoutMillis)

	err := wp.core.FSM.Transition(intermedState)

	if err != nil {

		return nil, err
	}

	// Wait for scanning to complete
	<-scanResult.Complete

	err = wp.core.FSM.Transition(notFoundState)

	if err != nil {

		return nil, err
	}

	if scanResult.Error != nil {

		return nil, scanResult.Error
	}

	if len(scanResult.Services) > 0 {

		// Convert map of services to array
		for _, svc := range scanResult.Services {

			svcResults = append(svcResults, svc)
		}
	}

	err = wp.core.FSM.Transition(FoundState)

	if err != nil {

		return nil, err
	}

	return svcResults, nil
}

func (wp *wpWithinImpl) GetServicePrices(serviceId int) ([]types.Price, error) {

	stateGoal := fsm.CON_SVC_AVAILABLE

	if valid := wp.core.FSM.Permitted(stateGoal); !valid {

		return nil, errors.New("Invalid state")
	}

	result := make([]types.Price, 0)

	priceResponse, err := wp.core.HTEClient.GetPrices(serviceId)

	if err != nil {

		return nil, err
	} else {

		for _, price := range priceResponse.Prices {

			result = append(result, price)
		}
	}

	return result, wp.core.FSM.Transition(stateGoal)
}

func (wp *wpWithinImpl) SelectService(serviceId, numberOfUnits, priceId int) (types.TotalPriceResponse, error) {

	if err := wp.core.FSM.Transition(fsm.CON_SEL_SVC); err != nil {

		return types.TotalPriceResponse{}, errors.New("Invalid state")
	}

	tpr, err := wp.core.HTEClient.NegotiatePrice(serviceId, priceId, numberOfUnits)

	if err != nil {

		wp.core.FSM.Transition(fsm.CON_SVC_AVAILABLE)

	} else {

		wp.core.FSM.Transition(fsm.CON_AWAIT_PAYMENT)
	}

	return tpr, err
}

func (wp *wpWithinImpl) MakePayment(request types.TotalPriceResponse) (types.PaymentResponse, error) {

	if err := wp.core.FSM.Transition(fsm.CON_PROC_PAYMENT); err != nil {

		return types.PaymentResponse{}, errors.New("Invalid state")
	}

	token, err := wp.core.Psp.GetToken(wp.core.HCE.HCECard, false)

	if err != nil {

		return types.PaymentResponse{}, err
	}

	paymentResponse, err := wp.core.HTEClient.MakeHtePayment(request.PaymentReferenceID, request.ClientID, token)

	if err != nil {

		wp.core.FSM.Transition(fsm.CON_SVC_AVAILABLE)
	}

	return paymentResponse, err
}

func (wp *wpWithinImpl) RequestServices() ([]types.ServiceDetails, error) {

	intermedState := fsm.CON_REQ_SVC

	if valid := wp.core.FSM.Permitted(intermedState); !valid {

		return nil, errors.New("Invalid state")
	}

	result := make([]types.ServiceDetails, 0)

	serviceResponse, err := wp.core.HTEClient.DiscoverServices()

	wp.core.FSM.Transition(intermedState)

	if err != nil {

		return nil, err
	} else {

		if len(serviceResponse.Services) > 0 {

			for _, svc := range serviceResponse.Services {

				result = append(result, svc)
			}

			wp.core.FSM.Transition(fsm.CON_SVC_AVAILABLE)
		} else {

			wp.core.FSM.Transition(fsm.CON_READY)
		}
	}

	return result, nil
}