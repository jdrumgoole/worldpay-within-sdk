package wpwithin
import (
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/domain"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/servicediscovery"
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
	StartSvcBroadcast(timeoutMillis int32)
	StopSvcBroadcast()
	ScanServices() []domain.Service
	GetSvcPrices(svc domain.Service) []domain.Price
	SelectSvc(svc domain.Service) domain.PaymentRequest
	MakePayment(payRequest domain.PaymentRequest) domain.PaymentResponse
}

func Initialise(deviceName, deviceDescription string) (WPWithin, error) {

	// TODO CH Check for GUID file and create if not exist

	uid := "<device_guid>"

	device, err := domain.NewDevice(deviceName, deviceDescription, uid)

	if err != nil {

		return domain.Device{}, err
	}

	svcBroadcaster, err := servicediscovery.NewBroadcaster(deviceDescription, "127.0.0.1", uid, "/services", 8980)

	if err != nil {

		return device, err
	}

	device.SvcBroadcaster = svcBroadcaster

	svcScanner, err := servicediscovery.NewScanner()

	if err != nil {

		return device, err
	}

	device.SvcScanner = svcScanner

	return device, nil
}