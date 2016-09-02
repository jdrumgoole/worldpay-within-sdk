package hte

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/psp"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types/event"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/utils"
)

// Coordinate requests between RPC interface and internal SDK interface
type ServiceHandler struct {
	device       *types.Device
	psp          psp.Psp
	credential   *Credential
	orderManager OrderManager
	eventHandler event.Handler
}

// Create a new Service Handler
func NewServiceHandler(device *types.Device, psp psp.Psp, credential *Credential, orderManager OrderManager, eventHandler event.Handler) *ServiceHandler {

	result := &ServiceHandler{
		device:       device,
		psp:          psp,
		credential:   credential,
		orderManager: orderManager,
		eventHandler: eventHandler,
	}

	return result
}

// List all the services available on the current device
func (srv *ServiceHandler) ServiceDiscovery(w http.ResponseWriter, r *http.Request) {

	// GET

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	var responseServices []types.ServiceDetails

	for _, srv := range srv.device.Services {

		responseServices = append(responseServices, types.ServiceDetails{
			ServiceID:          srv.Id,
			ServiceDescription: srv.Description,
		})
	}

	response := types.ServiceListResponse{
		Services: responseServices,
		ServerID: srv.device.UID,
	}

	returnMessage(w, http.StatusOK, response)
}

// List all the price variants for a specified service
func (srv *ServiceHandler) ServicePrices(w http.ResponseWriter, r *http.Request) {

	// GET

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	// Parse variables from request
	reqVars := mux.Vars(r)
	svcID, err := strconv.Atoi(reqVars["service_id"])

	if err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to parse input service id",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if svc, ok := srv.device.Services[svcID]; ok {

		response := types.ServicePriceResponse{}
		response.ServerID = srv.device.UID

		for _, price := range svc.Prices() {

			response.Prices = append(response.Prices, price)
		}

		returnMessage(w, http.StatusOK, response)

	} else {

		errorResponse := types.ErrorResponse{
			Message: fmt.Sprintf("Service not found for id %d", svcID),
		}

		returnMessage(w, http.StatusNotFound, errorResponse)
	}
}

// Get the total price for a current service selection
func (srv *ServiceHandler) ServiceTotalPrice(w http.ResponseWriter, r *http.Request) {

	// POST

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	// Parse variables from URI
	reqVars := mux.Vars(r)
	svcID, err := strconv.Atoi(reqVars["service_id"])

	if err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to parse input service id",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	// Parse message body (POST)
	var totalPriceRequest types.TotalPriceRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to read POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := r.Body.Close(); err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to close POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := json.Unmarshal(body, &totalPriceRequest); err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to parse POST body",
		}

		returnMessage(w, 422 /*Unprocessable Entity*/, errorResponse)
		return
	}

	if svc, ok := srv.device.Services[svcID]; ok {

		if price, ok := svc.Prices()[totalPriceRequest.SelectedPriceID]; ok {

			response := types.TotalPriceResponse{}
			response.ServerID = srv.device.UID
			response.ClientID = totalPriceRequest.ClientID
			response.PriceID = totalPriceRequest.SelectedPriceID
			response.UnitsToSupply = totalPriceRequest.SelectedNumberOfUnits
			response.TotalPrice = price.PricePerUnit.Amount * totalPriceRequest.SelectedNumberOfUnits
			response.CurrencyCode = price.PricePerUnit.CurrencyCode
			response.MerchantClientKey = srv.credential.MerchantClientKey

			payRef, err := utils.NewUUID()
			if err != nil {

				errorResponse := types.ErrorResponse{
					Message: "Internal error [payment-ref]",
				}

				returnMessage(w, http.StatusInternalServerError, errorResponse)
			}
			response.PaymentReferenceID = payRef

			order := types.Order{
				PaymentReference:      payRef,
				TotalPrice:            response.TotalPrice,
				ClientID:              response.ClientID,
				SelectedNumberOfUnits: response.UnitsToSupply,
				SelectedPriceId:       response.PriceID,
				ServiceID:             svcID,
				ClientUUID:            totalPriceRequest.ClientUUID,
			}

			err = srv.orderManager.AddOrder(order)

			if err != nil {

				errorResponse := types.ErrorResponse{
					Message: "Unable to add order to local store",
				}

				returnMessage(w, http.StatusInternalServerError, errorResponse)

			} else {

				returnMessage(w, http.StatusOK, response)
			}

		} else {

			errorResponse := types.ErrorResponse{
				Message: fmt.Sprintf("Price not found for id %d", totalPriceRequest.SelectedPriceID),
			}

			returnMessage(w, http.StatusNotFound, errorResponse)
		}

	} else {

		errorResponse := types.ErrorResponse{
			Message: fmt.Sprintf("Service not found for id %d", svcID),
		}

		returnMessage(w, http.StatusNotFound, errorResponse)
	}
}

// Make a payment for a service
func (srv *ServiceHandler) Payment(w http.ResponseWriter, r *http.Request) {

	// POST

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	// Parse message body (POST)
	var paymentRequest types.PaymentRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to read POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := r.Body.Close(); err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to close POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := json.Unmarshal(body, &paymentRequest); err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to parse POST body",
		}

		returnMessage(w, 422 /*HTTP Status Code: Unprocessable Entity*/, errorResponse)
		return
	}

	order, err := srv.orderManager.GetOrder(paymentRequest.PaymentReferenceID)

	orderCurrency := srv.device.Services[order.ServiceID].Prices()[order.SelectedPriceId].PricePerUnit.CurrencyCode

	if err != nil {

		errorResponse := types.ErrorResponse{
			Message: fmt.Sprintf("Unable to find order for payment ref %s", paymentRequest.PaymentReferenceID),
		}

		returnMessage(w, http.StatusNotFound, errorResponse)
	} else {

		// Some quick checks to compare validity of incoming request
		if strings.Compare(order.ClientID, paymentRequest.ClientID) != 0 {

			errorResponse := types.ErrorResponse{
				Message: "Client ID does not match Order Client ID",
			}

			returnMessage(w, http.StatusBadRequest, errorResponse)
		} else {

			// TODO CH - Fix order description
			paymentOrderCode, err := srv.psp.MakePayment(order.TotalPrice, orderCurrency, paymentRequest.ClientToken, "fixme", order.PaymentReference)

			if err != nil {

				errorResponse := types.ErrorResponse{
					Message: "Unable to process payment with gateway at this time",
				}

				returnMessage(w, http.StatusInternalServerError, errorResponse)

			} else {

				// TODO currently creating delivery token manually and assign to paymentresponse - this should be automapped.
				deliveryToken := &types.ServiceDeliveryToken{

					Key:            paymentOrderCode, // TODO cryptographically secure, random key generation.
					Issued:         time.Now(),
					Expiry:         time.Now().Add(time.Duration(24 * time.Hour)), // TODO add a value DeliveryToken lifetime (MINS) into the Service struct. 0 should be unlimited.
					RefundOnExpiry: false,                                         // TODO Map this into the Service Struct
					Signature:      nil,                                           // TODO implement HMAC generation scheme.
				}

				paymentResponse := types.PaymentResponse{
					ClientID:             order.ClientID,
					ServerID:             srv.device.UID,
					TotalPaid:            order.TotalPrice,
					ServiceDeliveryToken: deliveryToken,
					ClientUUID:           order.ClientUUID,
				}

				// TODO CH - Save paymentOrderCode to the order

				returnMessage(w, http.StatusOK, paymentResponse)
			}
		}
	}
}

// Begin delivery of a purchased service
func (srv *ServiceHandler) ServiceDeliveryBegin(w http.ResponseWriter, r *http.Request) {

	// POST

	log.Debug("begin hte.ServiceHandlerImpl.ServiceDeliveryBegin()")

	defer log.Debug("end hte.ServiceHandlerImpl.ServiceDeliveryBegin()")

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	// Parse variables from request
	reqVars := mux.Vars(r)
	svcID, err := strconv.Atoi(reqVars["service_id"])

	if err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to parse input service id",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	// Parse message body (POST)
	var deliveryRequest types.BeginServiceDeliveryRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to read POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := r.Body.Close(); err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to close POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := json.Unmarshal(body, &deliveryRequest); err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to parse POST body",
		}

		returnMessage(w, 422 /*Unprocessable Entity*/, errorResponse)
		return
	}

	if _, ok := srv.device.Services[svcID]; ok {

		response := types.BeginServiceDeliveryResponse{}
		response.ClientID = deliveryRequest.ClientID
		response.ServiceDeliveryToken = deliveryRequest.ServiceDeliveryToken
		response.ServerID = srv.device.UID
		response.UnitsToSupply = deliveryRequest.UnitsToSupply

		if srv.eventHandler != nil {

			log.Debug("Core event handler is set, calling event in core EventHandler")

			go srv.eventHandler.BeginServiceDelivery(svcID, deliveryRequest.ServiceDeliveryToken, deliveryRequest.UnitsToSupply)
		}

		returnMessage(w, http.StatusOK, response)

	} else {

		errorResponse := types.ErrorResponse{
			Message: fmt.Sprintf("Service not found for id %d", svcID),
		}

		returnMessage(w, http.StatusNotFound, errorResponse)
	}
}

// End delivery of a purchased service
func (srv *ServiceHandler) ServiceDeliveryEnd(w http.ResponseWriter, r *http.Request) {

	// POST

	log.Debug("begin hte.ServiceHandlerImpl.ServiceDeliveryEnd()")

	defer log.Debug("end hte.ServiceHandlerImpl.ServiceDeliveryEnd()")

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	// Parse variables from request
	reqVars := mux.Vars(r)
	svcID, err := strconv.Atoi(reqVars["service_id"])

	if err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to parse input service id",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	// Parse message body (POST)
	var deliveryRequest types.EndServiceDeliveryRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to read POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := r.Body.Close(); err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to close POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := json.Unmarshal(body, &deliveryRequest); err != nil {

		errorResponse := types.ErrorResponse{
			Message: "Unable to parse POST body",
		}

		returnMessage(w, 422 /*Unprocessable Entity*/, errorResponse)
		return
	}

	if _, ok := srv.device.Services[svcID]; ok {

		response := types.EndServiceDeliveryResponse{}
		response.ClientID = deliveryRequest.ClientID
		response.ServiceDeliveryToken = deliveryRequest.ServiceDeliveryToken
		response.ServerID = srv.device.UID
		response.UnitsJustSupplied = deliveryRequest.UnitsReceived
		response.UnitsRemaining = 9999

		if srv.eventHandler != nil {

			log.Debug("Core event handler is set, calling event in core EventHandler")

			go srv.eventHandler.EndServiceDelivery(svcID, deliveryRequest.ServiceDeliveryToken, deliveryRequest.UnitsReceived)
		}

		returnMessage(w, http.StatusOK, response)

	} else {

		errorResponse := types.ErrorResponse{
			Message: fmt.Sprintf("Service not found for id %d", svcID),
		}

		returnMessage(w, http.StatusNotFound, errorResponse)
	}
}

// Helper function for returning HTTP responses
func returnMessage(w http.ResponseWriter, statusCode int, message interface{}) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(message); err != nil {

		panic(err)
	}
}
