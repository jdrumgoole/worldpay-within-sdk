package wpwithin
import (
"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/domain"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/utils"
)

const (

	BCAST_STEP_SLEEP = 5000
	HTE_PORT = 8980
	SVC_URL_PREFIX = "/services"
)

type WPWithin interface {

	AddService(service domain.Service) error
	RemoveService(service domain.Service) error
	SetHTECredentials(hteCredentials hte.HTECredential) error
	SetHCECardCredential(hceCardCredential hce.HCECardCredential) error
	SetHCEClientCredential(hceClientCredential hce.HCEClientCredential) error
	InitConsumer() error
	InitProducer() error
	GetDevice() (domain.Device, error)
	StartSvcBroadcast(msg servicediscovery.BroadcastMessage, timeoutMillis int)
	StopSvcBroadcast()
	ScanServices() []domain.Service
	GetSvcPrices(svc domain.Service) []domain.Price
	SelectSvc(svc domain.Service) domain.PaymentRequest
	MakePayment(payRequest domain.PaymentRequest) domain.PaymentResponse
}

func Initialise(name, description string) (WPWithin, error) {

	// Device

	// TODO CH Check for GUID file and create if not exist
	deviceUID := "<device_guid>"

	deviceAddress, err := utils.ExternalIPv4()

	if err != nil {

		return WPWithin(), err
	}

	device, err := domain.NewDevice(name, description, deviceUID, deviceAddress)

	if err != nil {

		return domain.Device{}, err
	}

	// Service broadcaster

	svcBroadcaster, err := servicediscovery.NewBroadcaster(device.Description, device.IPv4Address, device.Uid, SVC_URL_PREFIX, HTE_PORT, BCAST_STEP_SLEEP)

	if err != nil {

		return device, err
	}

	device.SvcBroadcaster = svcBroadcaster

	// Service scanner

	svcScanner, err := servicediscovery.NewScanner()

	if err != nil {

		return device, err
	}

	device.SvcScanner = svcScanner

	return device, nil
}