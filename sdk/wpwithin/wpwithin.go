package wpwithin
import (
"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/domain"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/utils"
)

const (

	BROADCAST_STEP_SLEEP = 5000
	BROADCAST_PORT = 8980
	HTE_SVC_URL_PREFIX = "/services"
	UUID_FILE_PATH = "uuid.txt"
	HTE_SVC_PORT = 8080
)

type WPWithin interface {

	AddService(service domain.Service) error
	RemoveService(service domain.Service) error
	SetHTECredentials(hteCredentials hte.HTECredential) error
	SetHCECardCredential(hceCardCredential hce.HCECardCredential) error
	SetHCEClientCredential(hceClientCredential hce.HCEClientCredential) error
	InitConsumer() error
	InitProducer() error
	GetDevice() domain.Device
	StartSvcBroadcast(timeoutMillis int) error
	StopSvcBroadcast()
	ScanServices(timeoutMillis int) ([]servicediscovery.BroadcastMessage, error)
	GetSvcPrices(svc domain.Service) []domain.Price
	SelectSvc(svc domain.Service) domain.PaymentRequest
	MakePayment(payRequest domain.PaymentRequest) domain.PaymentResponse
}

func Initialise(name, description string) (WPWithin, error) {

	// Device

	var deviceGUID string
	var err error

	if b, _ := utils.FileExists(UUID_FILE_PATH); b {

		deviceGUID, err = utils.ReadLocalUUID(UUID_FILE_PATH)
	} else {

		deviceGUID, err = utils.NewUUID()

		utils.WriteString(UUID_FILE_PATH, deviceGUID, true)
	}

	if err != nil {

		return domain.Device{}, err
	}

	deviceAddress, err := utils.ExternalIPv4()

	if err != nil {

		return domain.Device{}, err
	}

	device, err := domain.NewDevice(name, description, deviceGUID, deviceAddress.String(), HTE_SVC_URL_PREFIX, HTE_SVC_PORT)

	if err != nil {

		return domain.Device{}, err
	}

	// Service broadcaster

	svcBroadcaster, err := servicediscovery.NewBroadcaster(device.HTEIPv4Address, BROADCAST_PORT, BROADCAST_STEP_SLEEP)

	if err != nil {

		return device, err
	}

	device.SvcBroadcaster = svcBroadcaster

	// Service scanner

	svcScanner, err := servicediscovery.NewScanner(BROADCAST_PORT, BROADCAST_STEP_SLEEP)

	if err != nil {

		return device, err
	}

	device.SvcScanner = svcScanner

	return device, nil
}