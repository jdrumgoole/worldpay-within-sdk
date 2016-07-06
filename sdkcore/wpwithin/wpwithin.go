package wpwithin
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hte"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/core"
	"time"
	"errors"
	"fmt"
)

// Factory to allow easy creation of
var Factory core.SDKFactory

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

	if Factory == nil {

		_Factory, err := core.NewSDKFactory()
		Factory = _Factory

		if err != nil {

			return nil, errors.New(fmt.Sprintf("Unable to create SDK Factory: %q", err.Error()))
		}
	}

	result := &wpWithinImpl{}

	if core, err := core.NewCore(); err != nil {

		return result, err

	} else {

		result.core = core
	}

	if dev, err := Factory.GetDevice(name, description); err != nil {

		return result, err
	} else {

		result.core.Device = dev
	}

	if om, err := Factory.GetOrderManager(); err != nil {

		return result, err

	} else {

		result.core.OrderManager = om;
	}

	if bc, err := Factory.GetSvcBroadcaster(result.core.Device.IPv4Address); err != nil {

		return result, err

	} else {

		result.core.SvcBroadcaster = bc
	}

	if sc, err := Factory.GetSvcScanner(); err != nil {

		return result, err

	} else {

		result.core.SvcScanner = sc
	}

	return result, nil
}

func (wp *wpWithinImpl) InitHTE(merchantClientKey, merchantServiceKey string) error {

	if psp, err := Factory.GetPSPMerchant(merchantClientKey, merchantServiceKey); err != nil {

		return errors.New(fmt.Sprintf("Unable to create psp", err.Error()))
	} else {

		wp.core.Psp = psp
	}

	hteCredential, err := hte.NewHTECredential(merchantClientKey, merchantServiceKey)

	if err != nil {

		return err
	}

	hteSvcHandler := Factory.GetHTEServiceHandler(wp.core.Device, wp.core.Psp, hteCredential, wp.core.OrderManager)

	if svc, err := Factory.GetHTE(wp.core.Device, wp.core.Psp, wp.core.Device.IPv4Address, hteCredential, wp.core.OrderManager, hteSvcHandler); err != nil {

		return err

	} else {

		wp.core.HTE = svc
	}

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

	httpHTE, err := Factory.GetHTEClientHTTP()

	if err != nil {

		return err
	}

	client, err := hte.NewClient(scheme, hostname, portNumber, urlPrefix, serverID, httpHTE)

	if err != nil {

		return err
	}

	wp.core.HTEClient = client

	return nil
}

func (wp *wpWithinImpl) InitProducer() (chan bool, error) {

	// Error channel allows us to get the error out of the go routine
	chStartResult := make(chan error)
	var startErr error

	go func() {

		chStartResult <- wp.core.HTE.Start()

	}()

	// Receive the error from the channel or wait a predefined amount of time
	// TODO CH : Fix this race condition
	select {

	case res := <-chStartResult:

		startErr = res

	case <-time.After(time.Millisecond * 750):

	}

	return nil, startErr
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

	wp.core.HCECard = cred

	if psp, err := Factory.GetPSPClient(); err != nil {

		return errors.New(fmt.Sprintf("Unable to create psp", err.Error()))
	} else {

		wp.core.Psp = psp
	}

	return nil
}

func (wp *wpWithinImpl) StartServiceBroadcast(timeoutMillis int) error {

	// Setup message that is broadcast over network
	msg := types.ServiceMessage{

		DeviceDescription: wp.core.Device.Description,
		Hostname: wp.core.HTE.IPAddr(),
		ServerID: wp.core.Device.Uid,
		UrlPrefix: wp.core.HTE.UrlPrefix(),
		PortNumber:wp.core.HTE.Port(),
	}

	// Set up a channel to get the error out of the go routine
	chBroadcastErr := make(chan error)
	var errBroadcast error

	go func() {

		chBroadcastErr <- wp.core.SvcBroadcaster.StartBroadcast(msg, timeoutMillis)
	}()

	// Either get the error or wait a small amount of time to give the all clear.
	// This is a race condition - ahhhh! TODO CH : Fix this
	select {

	case res := <- chBroadcastErr:

		errBroadcast = res

	case <- time.After(time.Millisecond * 750):

	}

	return errBroadcast
}

func (wp *wpWithinImpl) StopServiceBroadcast() {

	wp.core.SvcBroadcaster.StopBroadcast()
}

func (wp *wpWithinImpl) ServiceDiscovery(timeoutMillis int) ([]types.ServiceMessage, error) {

	svcResults := make([]types.ServiceMessage, 0)

	if scanResult, err := wp.core.SvcScanner.ScanForServices(timeoutMillis); err != nil {

		return nil, err

	} else if len(scanResult) > 0 {

		// Convert map of services to array
		for _, svc := range scanResult {

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

	token, err := wp.core.Psp.GetToken(wp.core.HCECard, request.MerchantClientKey, false)

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

func (wp *wpWithinImpl) Core() (*core.Core, error) {

	return wp.core, nil
}