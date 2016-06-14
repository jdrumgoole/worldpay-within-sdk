package hte
import (
"github.com/gorilla/mux"
"net/http"
	"fmt"
)

// Concrete Implementation of HTE service
type ServiceImpl struct {

	router *mux.Router
	_IPv4Address string
	_Port int
	_UrlPrefix string
	routes []Route
	handler *ServiceHandler
	HTECredential *Credential
	orderManager OrderManager
}

func (service *ServiceImpl) Start() error {

	return http.ListenAndServe(fmt.Sprintf(":%d", service.Port), service.router)
}

func (srv *ServiceImpl) setupRoutes() {

	srv.routes = []Route{
		Route{
			"Service Discovery",
			"GET",
			fmt.Sprintf("%s/service/discover", srv.UrlPrefix),
			srv.handler.ServiceDiscovery,
		},
		Route{
			"Service Price Request",
			"GET",
			fmt.Sprintf("%s/service/{service_id}/prices", srv.UrlPrefix),
			srv.handler.ServicePrices,
		},
		Route{
			"Service Total Price Request",
			"POST",
			fmt.Sprintf("%s/service/{service_id}/requestTotal", srv.UrlPrefix),
			srv.handler.ServiceTotalPrice,
		},
		Route{
			"Payment",
			"POST",
			fmt.Sprintf("%s/payment", srv.UrlPrefix),
			srv.handler.Payment,
		},
		Route{
			"Service Delivery Begin",
			"POST",
			fmt.Sprintf("%s/service/{service_id}/delivery/begin", srv.UrlPrefix),
			srv.handler.ServiceDeliveryBegin,
		},
		Route{
			"Service Delivery End",
			"POST",
			fmt.Sprintf("%s/service/{service_id}/delivery/end", srv.UrlPrefix),
			srv.handler.ServiceDeliveryEnd,
		},
	}
}

func (srv *ServiceImpl) IPAddr() string {

	return srv._IPv4Address
}

func (srv *ServiceImpl) Port() int {

	return srv._Port
}

func (srv *ServiceImpl) UrlPrefix() string {

	return srv._UrlPrefix
}