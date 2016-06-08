package wpwithin
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/utils"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/core"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp/onlineworldpay"
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
	GetDevice() *types.Device
	StartServiceBroadcast(timeoutMillis int) error
	StopServiceBroadcast()
	ServiceDiscovery(timeoutMillis int) ([]types.ServiceMessage, error)
	RequestServices() ([]types.ServiceDetails, error)
	GetServicePrices(serviceId int) ([]types.Price, error)
	SelectService(serviceId, numberOfUnits, priceId int) (types.TotalPriceResponse, error)
	MakePayment(payRequest types.TotalPriceResponse) (types.PaymentResponse, error)

}

func Initialise(name, description string) (WPWithin, error) {

	var err error

	// Set up SDK core

	core, err := core.New()

	if err != nil {

		return &wpWithinImpl{}, err
	}

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

	return wp, nil
}

func (wp *wpWithinImpl) InitHTE(merchantClientKey, merchantServiceKey string) error {

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

	return nil
}

type wpWithinImpl struct {

	core *core.Core
}

func (wp *wpWithinImpl) AddService(service *types.Service) error {

	if wp.core.Device.Services == nil {

		wp.core.Device.Services = make(map[int]*types.Service, 0)
	}

	wp.core.Device.Services[service.Id] = service

	return nil
}

func (wp *wpWithinImpl) RemoveService(service *types.Service) error {

	if wp.core.Device.Services != nil {

		delete(wp.core.Device.Services, service.Id)
	}

	return nil
}

func (wp *wpWithinImpl) InitConsumer(scheme, hostname string, portNumber int, urlPrefix, serverID string) error {

	// Setup HTE Client

	client, err := hte.NewClient(scheme, hostname, portNumber, urlPrefix, serverID)

	if err != nil {

		return err
	}

	wp.core.HTEClient = client

	return nil
}

func (wp *wpWithinImpl) InitProducer() (chan bool, error) {

	err := wp.core.HTE.Start()

	if err != nil {

		return nil, err
	}

	done := make(chan bool)

	return done, nil
}

func (wp *wpWithinImpl) GetDevice() *types.Device {

	return wp.core.Device
}

func (wp *wpWithinImpl) InitHCE(hceCardCredential types.HCECard) error {

	cred := new(types.HCECard)
	cred.FirstName = hceCardCredential.FirstName
	cred.LastName = hceCardCredential.LastName
	cred.ExpMonth = hceCardCredential.ExpMonth
	cred.ExpYear = hceCardCredential.ExpYear
	cred.CardNumber = hceCardCredential.CardNumber
	cred.Type = hceCardCredential.Type
	cred.Cvc = hceCardCredential.Cvc

	wp.core.HCE.HCECard = cred

	return nil
}

func (wp *wpWithinImpl) StartServiceBroadcast(timeoutMillis int) error {

	// Setup message that is broadcast over network
	msg := types.ServiceMessage{

		DeviceDescription: wp.core.Device.Description,
		Hostname: wp.core.HTE.IPv4Address,
		ServerID: wp.core.Device.Uid,
		UrlPrefix: wp.core.HTE.UrlPrefix,
		PortNumber:wp.core.HTE.Port,
	}

	complete, err := wp.core.SvcBroadcaster.StartBroadcast(msg, timeoutMillis)

	if err != nil {

		return err
	}

	// Wait for broadcast to complete
	<-complete

	return nil
}

func (wp *wpWithinImpl) StopServiceBroadcast() {

	wp.core.SvcBroadcaster.StopBroadcast()
}

func (wp *wpWithinImpl) ServiceDiscovery(timeoutMillis int) ([]types.ServiceMessage, error) {

	svcResults := make([]types.ServiceMessage, 0)

	scanResult := wp.core.SvcScanner.ScanForServices(timeoutMillis)

	// Wait for scanning to complete
	<-scanResult.Complete

	if scanResult.Error != nil {

		return nil, scanResult.Error
	} else if len(scanResult.Services) > 0 {

		// Convert map of services to array
		for _, svc := range scanResult.Services {

			svcResults = append(svcResults, svc)
		}
	}

	return svcResults, nil
}

func (wp *wpWithinImpl) GetServicePrices(serviceId int) ([]types.Price, error) {

	result := make([]types.Price, 0)

	priceResponse, err := wp.core.HTEClient.GetPrices(serviceId)

	if err != nil {

		return nil, err
	} else {

		for _, price := range priceResponse.Prices {

			result = append(result, price)
		}
	}

	return result, nil
}

func (wp *wpWithinImpl) SelectService(serviceId, numberOfUnits, priceId int) (types.TotalPriceResponse, error) {

	tpr, err := wp.core.HTEClient.NegotiatePrice(serviceId, priceId, numberOfUnits)

	return tpr, err
}

func (wp *wpWithinImpl) MakePayment(request types.TotalPriceResponse) (types.PaymentResponse, error) {

	token, err := wp.core.Psp.GetToken(wp.core.HCE.HCECard, false)

	if err != nil {

		return types.PaymentResponse{}, err
	}

	paymentResponse, err := wp.core.HTEClient.MakeHtePayment(request.PaymentReferenceID, request.ClientID, token)

	return paymentResponse, err
}

func (wp *wpWithinImpl) RequestServices() ([]types.ServiceDetails, error) {

	result := make([]types.ServiceDetails, 0)

	serviceResponse, err := wp.core.HTEClient.DiscoverServices()

	if err != nil {

		return nil, err
	} else {

		for _, svc := range serviceResponse.Services {

			result = append(result, svc)
		}
	}

	return result, nil
}