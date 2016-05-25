package wpwithin
import (
"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/domain"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/utils"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/core"
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/psp/onlineworldpay"
)

const (

	BROADCAST_STEP_SLEEP = 5000
	BROADCAST_PORT = 8980
	HTE_SVC_URL_PREFIX = ""
	UUID_FILE_PATH = "uuid.txt"
	HTE_SVC_PORT = 8080
	WP_ONLINE_API_ENDPOINT = "https://api.worldpay.com/v1"
)

type WPWithin interface {

	AddService(service domain.Service) error
	RemoveService(service domain.Service) error
	SetHTECredentials(hteCredentials hte.Credential) error
	SetHCECardCredential(hceCardCredential hce.CardCredential) error
	SetHCEClientCredential(hceClientCredential hce.ClientCredential) error
	InitConsumer() error
	InitProducer() (chan bool, error)
	GetDevice() *domain.Device
	StartSvcBroadcast(timeoutMillis int) error
	StopSvcBroadcast()
	ScanServices(timeoutMillis int) ([]servicediscovery.BroadcastMessage, error)
	GetSvcPrices(svc domain.Service) []domain.Price
	SelectSvc(svc domain.Service) domain.PaymentRequest
	MakePayment(payRequest domain.PaymentRequest) domain.PaymentResponse
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

	device, err := domain.NewDevice(name, description, deviceGUID, deviceAddress.String())

	if err != nil {

		return &wpWithinImpl{}, err
	}

	core.Device = device

	// Set up PSP
	psp, err := onlineworldpay.New(WP_ONLINE_API_ENDPOINT)

	if err != nil {

		return &wpWithinImpl{}, err
	}

	// Set up HTE service

	hte, err := hte.NewService(device, psp, device.IPv4Address, HTE_SVC_URL_PREFIX, HTE_SVC_PORT)

	if err != nil {

		return &wpWithinImpl{}, err
	}

	core.HTE = hte

	// Service broadcaster

	svcBroadcaster, err := servicediscovery.NewBroadcaster(core.HTE.IPv4Address, core.HTE.Port, BROADCAST_STEP_SLEEP)

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

	return wp, nil
}

func (wp *wpWithinImpl) AddService(service domain.Service) error {

	if wp.core.Device.Services == nil {

		wp.core.Device.Services = make(map[string]domain.Service, 0)
	}

	wp.core.Device.Services[service.Uid] = service

	return nil
}

func (wp *wpWithinImpl) RemoveService(service domain.Service) error {

	fmt.Println("Remove service..")

	return nil
}

func (wp *wpWithinImpl) InitConsumer() error {

	fmt.Println("init consumer...")

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

func (wp *wpWithinImpl) SetHTECredentials(hteCredentials hte.Credential) error {

	return nil
}

func (wp *wpWithinImpl) SetHCECardCredential(hceCardCredential hce.CardCredential) error {

	return nil
}

func (wp *wpWithinImpl) SetHCEClientCredential(hceClientCredential hce.ClientCredential) error {

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

func (wp *wpWithinImpl) GetSvcPrices(svc domain.Service) []domain.Price {

	return nil
}

func (wp *wpWithinImpl) SelectSvc(svc domain.Service) domain.PaymentRequest {

	return domain.PaymentRequest{}
}

func (wp *wpWithinImpl) MakePayment(payRequest domain.PaymentRequest) domain.PaymentResponse {

	return domain.PaymentResponse{}
}