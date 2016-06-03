package wpwithin
import (
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"
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

	AddService(service *domain.Service) error
	RemoveService(service *domain.Service) error
	InitHCE(hceCardCredential *hce.CardCredential) error
	InitHTE(hteCredential *hte.Credential) error
	InitConsumer(scheme, hostname string, portNumber int, urlPrefix, serverID string) error
	InitProducer() (chan bool, error)
	GetDevice() *domain.Device
	StartSvcBroadcast(timeoutMillis int) error
	StopSvcBroadcast()
	ScanServices(timeoutMillis int) ([]servicediscovery.BroadcastMessage, error)
	DiscoverServices() ([]hte.ServiceDetails, error)
	GetSvcPrices(serviceId int) ([]domain.Price, error)
	SelectSvc(serviceId, numberOfUnits, priceId int) (hte.TotalPriceResponse, error)
	MakePayment(payRequest hte.TotalPriceResponse) (hte.PaymentResponse, error)

}

func (wp *wpWithinImpl) InitHTE(hteCredential *hte.Credential) error {

	// Set up PSP
	psp, err := onlineworldpay.New(hteCredential.MerchantClientKey, hteCredential.MerchantServiceKey, WP_ONLINE_API_ENDPOINT)

	if err != nil {

		return err
	}

	wp.core.Psp = psp

	// Set up HTE service

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

	device, err := domain.NewDevice(name, description, deviceGUID, deviceAddress.String(), "GBP")

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

func (wp *wpWithinImpl) AddService(service *domain.Service) error {

	if wp.core.Device.Services == nil {

		wp.core.Device.Services = make(map[int]*domain.Service, 0)
	}

	wp.core.Device.Services[service.Id] = service

	return nil
}

func (wp *wpWithinImpl) RemoveService(service *domain.Service) error {

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

func (wp *wpWithinImpl) GetDevice() *domain.Device {

	return wp.core.Device
}

func (wp *wpWithinImpl) InitHCE(hceCardCredential *hce.CardCredential) error {

	wp.core.HCE.HCECardCredential = hceCardCredential

	return nil
}

func (wp *wpWithinImpl) StartSvcBroadcast(timeoutMillis int) error {

	// Setup message that is broadcast over network
	msg := servicediscovery.BroadcastMessage{

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

func (wp *wpWithinImpl) StopSvcBroadcast() {

	wp.core.SvcBroadcaster.StopBroadcast()
}

func (wp *wpWithinImpl) ScanServices(timeoutMillis int) ([]servicediscovery.BroadcastMessage, error) {

	svcResults := make([]servicediscovery.BroadcastMessage, 0)

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

func (wp *wpWithinImpl) GetSvcPrices(serviceId int) ([]domain.Price, error) {

	result := make([]domain.Price, 0)

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

func (wp *wpWithinImpl) SelectSvc(serviceId, numberOfUnits, priceId int) (hte.TotalPriceResponse, error) {

	tpr, err := wp.core.HTEClient.NegotiatePrice(serviceId, priceId, numberOfUnits)

	// TODO CH - Should we be returning a hte.TotalPriceResponse here ?

	return tpr, err
}

func (wp *wpWithinImpl) MakePayment(request hte.TotalPriceResponse) (hte.PaymentResponse, error) {

	token, err := wp.core.Psp.GetToken(wp.core.HCE.HCECardCredential, false)

	if err != nil {

		return hte.PaymentResponse{}, err
	}

	paymentResponse, err := wp.core.HTEClient.MakeHtePayment(request.PaymentReferenceID, request.ClientID, token)

	// TODO CH - Should we be returning the hte.PaymentResponse here ?

	return paymentResponse, err
}

func (wp *wpWithinImpl) DiscoverServices() ([]hte.ServiceDetails, error) {

	// TODO CH - Should we be returning hte.ServiceListResponse here ?

	result := make([]hte.ServiceDetails, 0)

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