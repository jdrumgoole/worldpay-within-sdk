package hte
import (
	"net/http"
"encoding/json"
)

type ServiceHandler struct {

}

func NewServiceHandler() *ServiceHandler {

	result := &ServiceHandler{}

	return result
}

func (srv *ServiceHandler) ServiceDiscovery(w http.ResponseWriter, r *http.Request) {

	// GET

	returnMessage(w, 200, "Service discovert")
}

func (srv *ServiceHandler) ServicePrices(w http.ResponseWriter, r *http.Request) {

	// GET

	returnMessage(w, 200, "Service prices..")
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