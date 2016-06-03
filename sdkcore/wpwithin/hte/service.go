package hte
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp"
)

type Service struct {

	router *mux.Router
	IPv4Address string
	Port int
	UrlPrefix string
	routes []Route
	handler *ServiceHandler
	HTECredential *Credential
	orderManager *OrderManager
}

type Route struct {

	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

func NewService(device *types.Device, psp psp.Psp, ip, prefix string, port int, hteCredential *Credential, orderManager *OrderManager) (*Service, error) {

	service := &Service{}

	service.handler = NewServiceHandler(device, psp, hteCredential, orderManager)

	service.UrlPrefix = prefix
	service.Port = port
	service.IPv4Address = ip
	service.orderManager = orderManager

	initRoutes(service)

	service.router = mux.NewRouter().StrictSlash(true)

	for _, route := range service.routes {

		service.router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
	}

	return service, nil
}

func (service *Service) Start() error {

	return http.ListenAndServe(fmt.Sprintf(":%d", service.Port), service.router)
}

func initRoutes(srv *Service) {

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