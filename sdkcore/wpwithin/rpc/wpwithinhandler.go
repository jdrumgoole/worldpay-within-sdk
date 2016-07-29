package rpc
import (

	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/rpc/wpthrift/wpthrift_types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/utils"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types/event"
)

type WPWithinHandler struct {

	wpwithin wpwithin.WPWithin
	callback event.Handler
}

func NewWPWithinHandler(wpWithin wpwithin.WPWithin, callback event.Handler) *WPWithinHandler {

	log.Debug("Begin RPC.WPWithinHandler.NewWPWithinHander()")

	result := &WPWithinHandler{
		wpwithin: wpWithin,
		callback: callback,
	}

	log.Debug("End RPC.WPWithinHandler.NewWPWithinHander()")

	return result
}

func (wp *WPWithinHandler) Setup(name, description string) (err error) {

	log.Debug("Begin RPC.WPWithinHandler.Setup()")

	wpw, err := wpwithin.Initialise(name, description)

	if err != nil {

		return err
	}

	wp.wpwithin = wpw

	if wp.callback != nil {

		log.Debug("Callback is set within WPWithinHandler - setting in wpwithin instance now")
		
		wp.wpwithin.SetEventHandler(wp.callback)
	} else {

		log.Debug("Callback is now set within WPWithinHandler")
	}

	log.Debug("End RPC.WPWithinHandler.Setup()")

	return nil
}

func (wp *WPWithinHandler) AddService(svc *wpthrift_types.Service) (err error) {

	log.Debug("Begin RPC.WPWithinHandler.AddService()")

	gSvc := &types.Service{
		Id: int(svc.ID),
		Name: svc.Name,
		Description: svc.Description,
	}

	log.Debug("End RPC.WPWithinHandler.AddService()")

	return wp.wpwithin.AddService(gSvc)
}

func (wp *WPWithinHandler) RemoveService(svc *wpthrift_types.Service) (err error) {

	log.Debug("Begin RPC.WPWithinHandler.RemoveService()")

	gSvc := &types.Service{
		Id: int(svc.ID),
		Name: svc.Name,
		Description: svc.Description,
	}

	log.Debug("End RPC.WPWithinHandler.RemoveService()")

	return wp.wpwithin.RemoveService(gSvc)
}

func (wp *WPWithinHandler) InitConsumer(scheme string, hostname string, port int32, urlPrefix string, serviceId string, hceCard *wpthrift_types.HCECard) (err error) {

	log.Debug("RPC.WPWithinHandler.InitConsumer()")

	_hceCard := types.HCECard{
		FirstName: hceCard.FirstName,
		LastName: hceCard.LastName,
		ExpMonth: hceCard.ExpMonth,
		ExpYear: hceCard.ExpYear,
		CardNumber: hceCard.CardNumber,
		Type: hceCard.Type,
		Cvc: hceCard.Cvc,
	}

	return wp.wpwithin.InitConsumer(scheme, hostname, int(port), urlPrefix, serviceId, &_hceCard)
}

func (wp *WPWithinHandler) InitProducer(merchantClientKey string, merchantServiceKey string) (err error) {

	log.Debug("RPC.WPWithinHandler.InitProducer()")


	go func(){

		wp.wpwithin.InitProducer(merchantClientKey, merchantServiceKey)

	}()

	return nil
}

func (wp *WPWithinHandler) GetDevice() (r *wpthrift_types.Device, err error) {

	log.Debug("Begin RPC.WPWithinHandler.GetDevice()")

	device := wp.wpwithin.GetDevice()

	result := &wpthrift_types.Device {

		UID: device.Uid,
		Name: device.Name,
		Description: device.Description,
		Services: make(map[int32]*wpthrift_types.Service, 0),
		Ipv4Address: device.IPv4Address,
		CurrencyCode: device.CurrencyCode,
	}

	log.Debugf("Found %d services for device", len(device.Services))

	if device != nil && len(device.Services) > 0 {

		log.Debug("Begin convert Go Service type to Thrift Service type")

		// Convert the services to Thrift services
		for i, svc := range device.Services {

			// Convert the prices to Thrift prices
			svcPrices := svc.Prices()
			thriftPrices := make(map[int32]wpthrift_types.Price, 0)

			log.Debugf("Found %d prices for service: %s (%d)", len(svcPrices), svc.Id, svc.Name)

			if len(svcPrices) > 0 {

				log.Debug("Begin convert Go price type to Thrift price type")

				for _, svcPrice := range svcPrices {

					thriftPrices[int32(svcPrice.ID)] = wpthrift_types.Price{

						ID: int32(svcPrice.ID),
						Description: svcPrice.Description,
						PricePerUnit: &wpthrift_types.PricePerUnit{
							Amount: int32(svcPrice.PricePerUnit.Amount),
							CurrencyCode: svcPrice.PricePerUnit.CurrencyCode,
						},
						UnitId: int32(svcPrice.UnitID),
						UnitDescription: svcPrice.UnitDescription,
					}
				}

				log.Debug("End convert Go price type to Thrift price type")
			}

			result.Services[int32(i)] = &wpthrift_types.Service{

				ID: int32(svc.Id),
				Name: svc.Name,
				Description: svc.Description,
				Prices: thriftPrices,
			}
		}

		log.Debug("End convert Go Service type to Thrift Service type")
	}

	log.Debug("End RPC.WPWithinHandler.GetDevice()")

	return result, nil
}

func (wp *WPWithinHandler) StartServiceBroadcast(timeoutMillis int32) (err error) {

	log.Debug("RPC.WPWithinHandler.StartServiceBroadcast()")

	return wp.wpwithin.StartServiceBroadcast(int(timeoutMillis))
}

func (wp *WPWithinHandler) StopServiceBroadcast() (err error) {

	log.Debug("Begin RPC.WPWithinHandler.StopServiceBroadcast()")

	wp.wpwithin.StopServiceBroadcast()

	log.Debug("End RPC.WPWithinHandler.StopServiceBroadcast()")

	return nil
}

func (wp *WPWithinHandler) DeviceDiscovery(timeoutMillis int32) (r map[*wpthrift_types.ServiceMessage]bool, err error) {

	log.Debug("Begin RPC.WPWithinHandler.ServiceDiscovery()")

	gSvcMsgs, err := wp.wpwithin.DeviceDiscovery(int(timeoutMillis))

	if err != nil {

		return nil, err
	}

	result := make(map[*wpthrift_types.ServiceMessage]bool, 0)

	for _, gSvcMsg := range gSvcMsgs {

		tmp := &wpthrift_types.ServiceMessage{
			DeviceDescription: gSvcMsg.DeviceDescription,
			Hostname: gSvcMsg.Hostname,
			PortNumber: int32(gSvcMsg.PortNumber),
			ServerId: gSvcMsg.ServerID,
			UrlPrefix: gSvcMsg.UrlPrefix,
		}

		result[tmp] = true
	}

	log.Debug("End RPC.WPWithinHandler.ServiceDiscovery()")

	return result, nil
}

func (wp *WPWithinHandler) RequestServices() (r map[*wpthrift_types.ServiceDetails]bool, err error) {

	log.Debug("Begin RPC.WPWithinHandler.RequestServices()")

	gServices, err := wp.wpwithin.RequestServices()

	if err != nil {

		return nil, err
	}

	result := make(map[*wpthrift_types.ServiceDetails]bool, 0)

	for _, gService := range gServices {

		tmp := &wpthrift_types.ServiceDetails{
			ServiceId: int32(gService.ServiceID),
			ServiceDescription: gService.ServiceDescription,
		}

		result[tmp] = true
	}

	log.Debug("End RPC.WPWithinHandler.RequestServices()")

	return result, nil
}

func (wp *WPWithinHandler) GetServicePrices(serviceId int32) (r map[*wpthrift_types.Price]bool, err error) {

	log.Debug("RPC.WPWithinHandler.GetServicePrices()")

	gSvcPrices, err := wp.wpwithin.GetServicePrices(int(serviceId))

	if err != nil {

		return nil, err
	}

	result := make(map[*wpthrift_types.Price]bool, 0)

	for _, gSvcPrice := range gSvcPrices {

		tmp := &wpthrift_types.Price{
			ID: int32(gSvcPrice.ID),
			Description: gSvcPrice.Description,
			PricePerUnit: &wpthrift_types.PricePerUnit {
				Amount: int32(gSvcPrice.PricePerUnit.Amount),
				CurrencyCode: gSvcPrice.PricePerUnit.CurrencyCode,
			},
			UnitId: int32(gSvcPrice.UnitID),
			UnitDescription: gSvcPrice.UnitDescription,
		}

		result[tmp] = true
	}

	return result, nil
}

func (wp *WPWithinHandler) SelectService(serviceId int32, numberOfUnits int32, priceId int32) (r *wpthrift_types.TotalPriceResponse, err error) {

	log.Debug("Begin RPC.WPWithinHandler.SelectService()")

	gPriceResponse, err := wp.wpwithin.SelectService(int(serviceId), int(numberOfUnits), int(priceId))

	if err != nil {

		return nil, err
	}

	result := &wpthrift_types.TotalPriceResponse{
		ServerId: gPriceResponse.ServerID,
		ClientId: gPriceResponse.ClientID,
		PriceId: int32(gPriceResponse.PriceID),
		UnitsToSupply: int32(gPriceResponse.UnitsToSupply),
		TotalPrice: int32(gPriceResponse.TotalPrice),
		PaymentReferenceId: gPriceResponse.PaymentReferenceID,
		MerchantClientKey: gPriceResponse.MerchantClientKey,
	}

	log.Debug("End RPC.WPWithinHandler.SelectService()")

	return result, nil
}

func (wp *WPWithinHandler) MakePayment(request *wpthrift_types.TotalPriceResponse) (r *wpthrift_types.PaymentResponse, err error) {

	log.Debug("Begin RPC.WPWithinHandler.MakePayment()")

	gRequest := types.TotalPriceResponse{
		ServerID: request.ServerId,
		ClientID: request.ClientId,
		PriceID: int(request.PriceId),
		UnitsToSupply: int(request.UnitsToSupply),
		TotalPrice: int(request.TotalPrice),
		PaymentReferenceID: request.PaymentReferenceId,
		MerchantClientKey: request.MerchantClientKey,
	}

	log.Debug("Finised converting thrift.TotalPriceResponse to go.TotalPriceResponse")

	log.Debug("Proceeding to call MakePayment internally using converted request object")

	gPaymentResponse, err := wp.wpwithin.MakePayment(gRequest)

	if err != nil {

		return nil, err
	}

	// TODO create delivery token manually and assign to paymentresponse - need automatpping
	deliveryToken := &wpthrift_types.ServiceDeliveryToken{

		Key: gPaymentResponse.ServiceDeliveryToken.Key,
		Issued: utils.TimeFormatISO(gPaymentResponse.ServiceDeliveryToken.Issued),
		Expiry: utils.TimeFormatISO(gPaymentResponse.ServiceDeliveryToken.Expiry),
		RefundOnExpiry: gPaymentResponse.ServiceDeliveryToken.RefundOnExpiry,
		Signature: gPaymentResponse.ServiceDeliveryToken.Signature,
	}

	result := &wpthrift_types.PaymentResponse{
		ServerId: gPaymentResponse.ServerID,
		ClientId: gPaymentResponse.ClientID,
		TotalPaid: int32(gPaymentResponse.TotalPaid),
		ServiceDeliveryToken: deliveryToken,
		ClientUUID: gPaymentResponse.ClientUUID,
	}

	log.Debug("End RPC.WPWithinHandler.MakePayment()")

	return result, nil
}

func (wp *WPWithinHandler) BeginServiceDelivery(clientId string, serviceDeliveryToken *wpthrift_types.ServiceDeliveryToken, unitsToSupply int32) (err error) {
	issueTime, err := utils.ParseISOTime(serviceDeliveryToken.Issued)

	if err != nil {

		return err
	}

	expiryTime, err := utils.ParseISOTime(serviceDeliveryToken.Expiry)

	if err != nil {

		return err
	}

	sdt := types.ServiceDeliveryToken{

		Key: serviceDeliveryToken.Key,
		Issued: issueTime,
		Expiry: expiryTime,
		RefundOnExpiry: serviceDeliveryToken.RefundOnExpiry,
		Signature: serviceDeliveryToken.Signature,
	}

	wp.wpwithin.BeginServiceDelivery(clientId, sdt, int(unitsToSupply))

	return nil
}

func (wp *WPWithinHandler) EndServiceDelivery(clientId string, serviceDeliveryToken *wpthrift_types.ServiceDeliveryToken, unitsReceived int32) (err error) {

	issueTime, err := utils.ParseISOTime(serviceDeliveryToken.Issued)

	if err != nil {

		return err
	}

	expiryTime, err := utils.ParseISOTime(serviceDeliveryToken.Expiry)

	if err != nil {

		return err
	}

	sdt := types.ServiceDeliveryToken{

		Key: serviceDeliveryToken.Key,
		Issued: issueTime,
		Expiry: expiryTime,
		RefundOnExpiry: serviceDeliveryToken.RefundOnExpiry,
		Signature: serviceDeliveryToken.Signature,
	}

	wp.wpwithin.EndServiceDelivery(clientId, sdt, int(unitsReceived))

	return nil
}