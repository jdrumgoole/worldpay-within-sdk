package domain
import (
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/servicediscovery"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/psp"
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
	HTEIPv4Address string
	HTEPrefix string
	HTEPort int
}

func NewDevice(name, description, uid string, hteIpv4 string, htePrefix string, htePort int) (*Device, error) {

	result := &Device{
		Name:name,
		Description:description,
		Uid:uid,
		HTEIPv4Address: hteIpv4,
		HTEPrefix: htePrefix,
		HTEPort: htePort,
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

func (wp Device) GetDevice() Device {

	return wp
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

func (wp Device) StartSvcBroadcast(timeoutMillis int) (chan bool, error) {

	done := make(chan bool)

	go func() {

		msg := servicediscovery.BroadcastMessage{

			DeviceDescription: wp.Description,
			Hostname: wp.HTEIPv4Address,
			ServerID: wp.Uid,
			UrlPrefix: wp.HTEPrefix,
			PortNumber:wp.HTEPort,
		}

		wp.SvcBroadcaster.StartBroadcast(msg, timeoutMillis)

		done <- true
	}()

	return done, nil
}

func (wp Device) StopSvcBroadcast() {

	wp.SvcBroadcaster.StopBroadcast()
}

func (wp Device) ScanServices(timeoutMillis int) ([]servicediscovery.BroadcastMessage, error) {

	svcResults := make([]servicediscovery.BroadcastMessage, 0)

	scanResult := wp.SvcScanner.ScanForServices(timeoutMillis)

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

func (wp Device) GetSvcPrices(svc Service) []Price {

	return nil
}

func (wp Device) SelectSvc(svc Service) PaymentRequest {

	return PaymentRequest{}
}

func (wp Device) MakePayment(payRequest PaymentRequest) PaymentResponse {

	return PaymentResponse{}
}