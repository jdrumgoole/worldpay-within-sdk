package hte
import (
	"net/http"
"encoding/json"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/domain"
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/psp"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
)

type ServiceHandler struct {

	device *domain.Device
	psp psp.Psp
}

func NewServiceHandler(device *domain.Device, psp psp.Psp) *ServiceHandler {

	result := &ServiceHandler{
		device: device,
		psp: psp,
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

	returnMessage(w, 200, "Service total price")
}

func (srv *ServiceHandler) Payment(w http.ResponseWriter, r *http.Request) {

	// POST

	returnMessage(w, 200, "Payment")
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