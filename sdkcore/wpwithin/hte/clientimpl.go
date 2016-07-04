package hte
import (
	"errors"
	"fmt"
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
"encoding/json"
)

// Concrete implementation of HTEClient
type clientImpl struct {

	scheme string
	host string
	port int
	urlPrefix string
	clientId string
	baseUrl string
	httpClient HTEClientHTTP
}

// Initialise a new HTE Client
func NewClient(scheme, host string, port int, urlPrefix string, clientId string, httpClient HTEClientHTTP) (Client, error) {

	if host == "" {

		return nil, errors.New("host cannot be empty")
	}

	if port < PORT_RANGE_MIN || port > PORT_RANGE_MAX {

		return nil, errors.New(fmt.Sprintf("Port number cannot exceed range [%d - %d]", PORT_RANGE_MIN, PORT_RANGE_MAX))
	}

	// Do not need to validate against empty urlPrefix as it can actually be empty

	if clientId == "" {

		return nil, errors.New("clientId cannot be empty")
	}

	result := &clientImpl{}
	result.host = host
	result.port = port
	result.urlPrefix = urlPrefix
	result.clientId = clientId
	result.httpClient = httpClient

	// Compose base url and ensure there are no duplicate slashes from passed in urlPrefix
	result.baseUrl = fmt.Sprintf("%s%s:%d%s", scheme, host, port, urlPrefix)

	return result, nil
}

func (client *clientImpl) DiscoverServices() (types.ServiceListResponse, error) {

	url := fmt.Sprintf("%s/service/discover", client.baseUrl)

	response, err := client.httpClient.Get(url)

	if err != nil {

		return types.ServiceListResponse{}, err
	}

	svcDetails := types.ServiceListResponse {}

	err = json.Unmarshal(response, &svcDetails)

	if err != nil {

		return types.ServiceListResponse{}, err
	}

	return svcDetails, nil
}

func (client *clientImpl) GetPrices(serviceId int) (types.ServicePriceResponse, error) {

	url := fmt.Sprintf("%s/service/%d/prices", client.baseUrl, serviceId)

	response, err := client.httpClient.Get(url)

	if err != nil {

		return types.ServicePriceResponse{}, err
	}

	svcPriceResponse := types.ServicePriceResponse{}

	err = json.Unmarshal(response, &svcPriceResponse)

	if err != nil {

		return types.ServicePriceResponse{}, err
	}

	return svcPriceResponse, nil
}

func (client *clientImpl) NegotiatePrice(serviceId, priceId, numberOfUnits int) (types.TotalPriceResponse, error) {

	url := fmt.Sprintf("%s/service/%d/requestTotal", client.baseUrl, serviceId)

	req := types.TotalPriceRequest{
		ClientID: client.clientId,
		SelectedNumberOfUnits: numberOfUnits,
		SelectedPriceId: priceId,
	}

	jsonReq, err := json.Marshal(req)

	if err != nil {

		return types.TotalPriceResponse{}, err
	}

	bytesResp, err := client.httpClient.PostJSON(url, jsonReq)

	if err != nil {

		return types.TotalPriceResponse{}, err
	}

	priceResp := types.TotalPriceResponse{}

	err = json.Unmarshal(bytesResp, &priceResp)

	if err != nil {

		return types.TotalPriceResponse{}, err
	}

	return priceResp, nil
}

func (client *clientImpl) MakeHtePayment(paymentReferenceId, clientId, clientToken string) (types.PaymentResponse, error) {

	url := fmt.Sprintf("%s/payment", client.baseUrl)

	requestBody := types.PaymentRequest{
		ClientID:clientId,
		ClientToken: clientToken,
		PaymentReferenceID: paymentReferenceId,
	}

	jsonBody, err := json.Marshal(requestBody)

	if err != nil {

		return types.PaymentResponse{}, err
	}

	byteResp, err := client.httpClient.PostJSON(url, jsonBody)

	if err != nil {

		return types.PaymentResponse{}, err
	}

	paymentResponse := types.PaymentResponse{}

	err = json.Unmarshal(byteResp, &paymentResponse)

	if err != nil {

		return types.PaymentResponse{}, err
	}

	return paymentResponse, nil
}

func (client *clientImpl) StartDelivery(serviceId int, serviceDeliveryToken string, unitsToSupply int) (int, error) {

	return 0, errors.New("Not implemented..")
}

func (client *clientImpl) EndDelivery(serviceId int, serviceDeliveryToken string, unitsReceived int) (int, error) {

	return 0, errors.New("Not implemented..")
}