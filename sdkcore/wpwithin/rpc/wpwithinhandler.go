package rpc
import (
"errors"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/rpc/wpthrift/wpthrift_types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"fmt"
)

type WPWithinHandler struct {

	wpwithin wpwithin.WPWithin
}

func NewWPWithinHandler(wpWithin wpwithin.WPWithin) *WPWithinHandler {

	result := &WPWithinHandler{}
	result.wpwithin = wpWithin

	return result
}

func (wp *WPWithinHandler) AddService(svc *wpthrift_types.Service) (err error) {

	return errors.New("Not implemented..")
}

func (wp *WPWithinHandler) RemoveService(svc *wpthrift_types.Service) (err error) {

	return errors.New("Not implemented..")
}

func (wp *WPWithinHandler) InitHCE(hceCard *wpthrift_types.HCECard) (err error) {

	return errors.New("Not implemented..")
}

func (wp *WPWithinHandler) InitHTE(merchantClientKey string, merchantServiceKey string) (err error) {

	return errors.New("Not implemented..")
}

func (wp *WPWithinHandler) InitConsumer(scheme string, hostname string, port int32, urlPrefix string, serviceId string) (err error) {

	return errors.New("Not implemented..")
}

func (wp *WPWithinHandler) InitProducer() (err error) {

	return errors.New("Not implemented..")
}

func (wp *WPWithinHandler) GetDevice() (r *wpthrift_types.Device, err error) {

	fmt.Println("rpc wpwithin handler - getDevice()")

	device := wp.wpwithin.GetDevice()

	result := &wpthrift_types.Device {

		UID: device.Uid,
		Name: device.Name,
		Description: device.Description,
		//Services: device.Services, /* TODO CH - Convert types */
		Ipv4Address: device.IPv4Address,
		CurrencyCode: device.CurrencyCode,
	}

	return result, nil
}

func (wp *WPWithinHandler) StartServiceBroadcast(timeoutMillis int32) (err error) {

	return errors.New("Not implemented..")
}

func (wp *WPWithinHandler) StopServiceBroadcast() (err error) {

	return errors.New("Not implemented..")
}

func (wp *WPWithinHandler) ServiceDiscovery(timeoutMillis int32) (r map[*wpthrift_types.ServiceMessage]bool, err error) {

	return nil, errors.New("Not implemented..")
}

func (wp *WPWithinHandler) RequestServices() (r map[*wpthrift_types.ServiceDetails]bool, err error) {

	return nil, errors.New("Not implemented..")
}

func (wp *WPWithinHandler) GetServicePrices(serviceId int32) (r map[*wpthrift_types.Price]bool, err error) {

	return nil, errors.New("Not implemented..")
}

func (wp *WPWithinHandler) SelectService(serviceId int32, numberOfUnits int32, priceId int32) (r *wpthrift_types.TotalPriceResponse, err error) {

	return nil, errors.New("Not implemented..")
}

func (wp *WPWithinHandler) MakePayment(request *wpthrift_types.TotalPriceResponse) (r *wpthrift_types.PaymentResponse, err error) {

	return nil, errors.New("Not implemented..")
}
