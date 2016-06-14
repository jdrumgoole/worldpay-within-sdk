package hte
import (
	"net/http"
	"github.com/gorilla/mux"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp"
)

type Service interface {

	Start() error
	setupRoutes()
	IPAddr() string
	Port() int
	UrlPrefix() string
}

type Route struct {

	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

func NewService(device *types.Device, psp psp.Psp, ip, prefix string, port int, hteCredential *Credential, orderManager OrderManager) (Service, error) {

	service := &ServiceImpl{}

	service.handler = NewServiceHandler(device, psp, hteCredential, orderManager)

	service._UrlPrefix = prefix
	service._Port = port
	service._IPv4Address = ip
	service.orderManager = orderManager

	service.setupRoutes()

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