package hte
import (
	"net/http"
"encoding/json"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/domain"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/psp"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
	"io/ioutil"
	"io"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/utils"
	"strings"
)

type ServiceHandler struct {

	device *domain.Device
	psp psp.Psp
	credential *Credential
	orderManager *OrderManager
}

func NewServiceHandler(device *domain.Device, psp psp.Psp, credential *Credential, orderManager *OrderManager) *ServiceHandler {

	result := &ServiceHandler{
		device: device,
		psp: psp,
		credential: credential,
		orderManager: orderManager,
	}

	return result
}

// List all the services available from this thing
func (srv *ServiceHandler) ServiceDiscovery(w http.ResponseWriter, r *http.Request) {

	// GET

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	responseServices := make([]ServiceDetails, 0)

	for _, srv := range srv.device.Services {

		responseServices = append(responseServices, ServiceDetails{
			ServiceID:srv.Id,
			ServiceDescription:srv.Description,
		})
	}

	response := ServiceListResponse{
		Services:responseServices,
		ServerID:srv.device.Uid,
	}

	returnMessage(w, http.StatusOK, response)
}

func (srv *ServiceHandler) ServicePrices(w http.ResponseWriter, r *http.Request) {

	// GET

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	// Parse variables from request
	reqVars := mux.Vars(r)
	svcId, err := strconv.Atoi(reqVars["service_id"])

	if err != nil {

		errorResponse := ErrorResponse{
			Message: "Unable to parse input service id",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if svc, ok := srv.device.Services[svcId]; ok {

		response := ServicePriceResponse{}
		response.ServerID = srv.device.Uid

		for _, price := range svc.Prices() {

			response.Prices = append(response.Prices, price)
		}

		returnMessage(w, http.StatusOK, response)

	} else {

		errorResponse := ErrorResponse{
			Message: fmt.Sprintf("Service not found for id %d", svcId),
		}

		returnMessage(w, http.StatusNotFound, errorResponse)
	}
}

func (srv *ServiceHandler) ServiceTotalPrice(w http.ResponseWriter, r *http.Request) {

	// POST

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	// Parse variables from URI
	reqVars := mux.Vars(r)
	svcId, err := strconv.Atoi(reqVars["service_id"])

	if err != nil {

		errorResponse := ErrorResponse{
			Message: "Unable to parse input service id",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	// Parse message body (POST)
	var totalPriceRequest TotalPriceRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {

		errorResponse := ErrorResponse{
			Message: "Unable to read POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := r.Body.Close(); err != nil {

		errorResponse := ErrorResponse{
			Message: "Unable to close POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := json.Unmarshal(body, &totalPriceRequest); err != nil {

		errorResponse := ErrorResponse{
			Message: "Unable to parse POST body",
		}

		returnMessage(w, 422/*Unprocessable Entity*/, errorResponse)
		return
	}

	if svc, ok := srv.device.Services[svcId]; ok {

		if price, ok := svc.Prices()[totalPriceRequest.SelectedPriceId]; ok {

			response := TotalPriceResponse{}
			response.ServerID = srv.device.Uid
			response.ClientID = totalPriceRequest.ClientID
			response.PriceID = totalPriceRequest.SelectedPriceId
			response.UnitsToSupply = totalPriceRequest.SelectedNumberOfUnits
			response.TotalPrice = price.PricePerUnit * totalPriceRequest.SelectedNumberOfUnits
			response.MerchantClientKey = srv.credential.MerchantClientKey

			payRef, err := utils.NewUUID()
			if err != nil {

				errorResponse := ErrorResponse{
					Message: "Internal error [payment-ref]",
				}

				returnMessage(w, http.StatusInternalServerError, errorResponse)
			}
			response.PaymentReferenceID = payRef

			order := Order{
				PaymentReference:payRef,
				TotalPrice:response.TotalPrice,
				ClientID:response.ClientID,
				SelectedNumberOfUnits:response.UnitsToSupply,
				SelectedPriceId:response.PriceID,
				ServiceID:svcId,
				ClientUUID:totalPriceRequest.ClientUUID,
			}

			err = srv.orderManager.AddOrder(order)

			if err != nil {

				errorResponse := ErrorResponse{
					Message:"Unable to add order to local store",
				}

				returnMessage(w, http.StatusInternalServerError, errorResponse)

			} else {

				returnMessage(w, http.StatusOK, response)
			}

		} else {

			errorResponse := ErrorResponse{
				Message: fmt.Sprintf("Price not found for id %d", totalPriceRequest.SelectedPriceId),
			}

			returnMessage(w, http.StatusNotFound, errorResponse)
		}

	} else {

		errorResponse := ErrorResponse{
			Message: fmt.Sprintf("Service not found for id %d", svcId),
		}

		returnMessage(w, http.StatusNotFound, errorResponse)
	}
}

func (srv *ServiceHandler) Payment(w http.ResponseWriter, r *http.Request) {

	// POST

	defer func() {
		if err := recover(); err != nil {

			returnMessage(w, http.StatusInternalServerError, err)
		}
	}()

	// Parse message body (POST)
	var paymentRequest PaymentRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {

		errorResponse := ErrorResponse{
			Message: "Unable to read POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := r.Body.Close(); err != nil {

		errorResponse := ErrorResponse{
			Message: "Unable to close POST body",
		}

		returnMessage(w, http.StatusBadRequest, errorResponse)
		return
	}

	if err := json.Unmarshal(body, &paymentRequest); err != nil {

		errorResponse := ErrorResponse{
			Message: "Unable to parse POST body",
		}

		returnMessage(w, 422/*HTTP Status Code: Unprocessable Entity*/, errorResponse)
		return
	}

	order, err := srv.orderManager.GetOrder(paymentRequest.PaymentReferenceID)

	if err != nil {

		errorResponse := ErrorResponse{
			Message: fmt.Sprintf("Unable to find order for payment ref %s", paymentRequest.PaymentReferenceID),
		}

		returnMessage(w, http.StatusNotFound, errorResponse)
	} else {

		// Some quick checks to compare validity of incoming request
		if strings.Compare(order.ClientID, paymentRequest.ClientID) != 0 {

			errorResponse := ErrorResponse{
				Message: "Client ID does not match Order Client ID",
			}

			returnMessage(w, http.StatusBadRequest, errorResponse)
		} else {

			paymentOrderCode, err := srv.psp.MakePayment(order.TotalPrice, srv.device, srv.device.CurrencyCode, paymentRequest.ClientToken, order.PaymentReference)

			if err != nil {

				errorResponse := ErrorResponse{
					Message:"Unable to process payment with gateway at this time",
				}

				returnMessage(w, http.StatusInternalServerError, errorResponse)

			} else {

				paymentResponse := PaymentResponse{
					ClientID: order.ClientID,
					ServerID: srv.device.Uid,
					TotalPaid: order.TotalPrice,
					ServiceDeliveryToken: paymentOrderCode, // TODO CH - Use generated GUID instead of payment ref and persist in order manager for validation later
					ClientUUID: order.ClientUUID,
				}

				// TODO CH - Save paymentOrderCode to the order

				returnMessage(w, http.StatusOK, paymentResponse)
			}
		}
	}
}

func (srv *ServiceHandler) ServiceDeliveryBegin(w http.ResponseWriter, r *http.Request) {

	// POST

	returnMessage(w, 200, "Service delivery begin")
}

func (srv *ServiceHandler) ServiceDeliveryEnd(w http.ResponseWriter, r *http.Request) {

	// POST

	returnMessage(w, 200, "Service delivery end")
}

func returnMessage(w http.ResponseWriter, statusCode int, message interface{}) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(message); err != nil {

		panic(err)
	}
}