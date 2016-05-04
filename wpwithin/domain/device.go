package domain
import (
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/psp"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/servicediscovery"
)

type Device struct {

	Uid string
	Name string
	Description string
	services []Service
	HTECredential hte.HTECredential
	HCECardCredential hce.HCECardCredential
	HCEClientCredential hce.HCEClientCredential
	Psp psp.Psp
	SvcBroadcaster servicediscovery.Broadcaster
	SvcScanner servicediscovery.Scanner
}

func NewDevice(name, description, uid string) (*Device, error) {

	result := &Device{
		Name:name,
		Description:description,
		Uid:uid,
	}

	return result, nil
}

func (wp Device) AddService(service Service) error {

	fmt.Println("Add service..")

	return nil
}

func (wp Device) RemoveService(service Service) error {

	fmt.Println("Remove service..")

	return nil
}

func (wp Device) InitConsumer() error {

	fmt.Println("init consumer...")

	return nil
}

func (wp Device) InitProducer() error {

	fmt.Println("Init producer..")

	return nil
}

func (wp Device) GetDevice() (Device, error) {

	return wp, nil
}

func (wp Device) SetHTECredentials(hteCredentials hte.HTECredential) error {

	return nil
}

func (wp Device) SetHCECardCredential(hceCardCredential hce.HCECardCredential) error {

	return nil
}

func (wp Device) SetHCEClientCredential(hceClientCredential hce.HCEClientCredential) error {

	return nil
}

func (wp Device) StartSvcBroadcast(timeoutMillis int32) {

	wp.SvcBroadcaster.StartBroadcast(timeoutMillis)
}

func (wp Device) StopSvcBroadcast() {

}

func (wp Device) ScanServices() []Service {

	wp.SvcScanner.ScanForServices(1000)

	return nil
}

func (wp Device) GetSvcPrices(svc Service) []Price {

	return nil
}

func (wp Device) SelectSvc(svc Service) PaymentRequest {

	return PaymentRequest{}
}

func (wp Device) MakePayment(payRequest PaymentRequest) PaymentResponse {

	return PaymentResponse{}
}