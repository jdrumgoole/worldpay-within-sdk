package hte
import (
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"
)

const PORT_RANGE_MIN = 1
const PORT_RANGE_MAX = 65535

type Client interface {

	DiscoverServices() (domain.ServiceListResponse, error)
	GetPrices(serviceId int) (domain.ServicePriceResponse, error)
	NegotiatePrice(serviceId, priceId, numberOfUnits int) (domain.TotalPriceResponse, error)
	MakeHtePayment(paymentReferenceId, clientId, clientToken string) (domain.PaymentResponse, error)
	StartDelivery(serviceId int, serviceDeliveryToken string, unitsToSupply int) (int, error)
	EndDelivery(serviceId int, serviceDeliveryToken string, unitsReceived int) (int, error)
}

type clientImpl struct {

	scheme string
	host string
	port int
	urlPrefix string
	clientId string
	baseUrl string
}

func NewClient(scheme, host string, port int, urlPrefix string, clientId string) (Client, error) {

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

	// Compose base url and ensure there are no duplicate slashes from passed in urlPrefix
	result.baseUrl = fmt.Sprintf("%s%s:%d%s", scheme, host, port, urlPrefix)

	return result, nil
}

func (client *clientImpl) DiscoverServices() (domain.ServiceListResponse, error) {

	url := fmt.Sprintf("%s/service/discover", client.baseUrl)

	response, err := client.get(url)

	if err != nil {

		return domain.ServiceListResponse{}, err
	}

	svcDetails := domain.ServiceListResponse {}

	err = json.Unmarshal(response, &svcDetails)

	if err != nil {

		return domain.ServiceListResponse{}, err
	}

	return svcDetails, nil
}

func (client *clientImpl) GetPrices(serviceId int) (domain.ServicePriceResponse, error) {

	url := fmt.Sprintf("%s/service/%d/prices", client.baseUrl, serviceId)

	response, err := client.get(url)

	if err != nil {

		return domain.ServicePriceResponse{}, err
	}

	svcPriceResponse := domain.ServicePriceResponse{}

	err = json.Unmarshal(response, &svcPriceResponse)

	if err != nil {

		return domain.ServicePriceResponse{}, err
	}

	return svcPriceResponse, nil
}

func (client *clientImpl) NegotiatePrice(serviceId, priceId, numberOfUnits int) (domain.TotalPriceResponse, error) {

	url := fmt.Sprintf("%s/service/%d/requestTotal", client.baseUrl, serviceId)

	req := domain.TotalPriceRequest{
		ClientID: client.clientId,
		SelectedNumberOfUnits: numberOfUnits,
		SelectedPriceId: priceId,
	}

	jsonReq, err := json.Marshal(req)

	if err != nil {

		return domain.TotalPriceResponse{}, err
	}

	bytesResp, err := client.postJSON(url, jsonReq)

	if err != nil {

		return domain.TotalPriceResponse{}, err
	}

	priceResp := domain.TotalPriceResponse{}

	err = json.Unmarshal(bytesResp, &priceResp)

	if err != nil {

		return domain.TotalPriceResponse{}, err
	}

	return priceResp, nil
}

func (client *clientImpl) MakeHtePayment(paymentReferenceId, clientId, clientToken string) (domain.PaymentResponse, error) {

	url := fmt.Sprintf("%s/payment", client.baseUrl)

	requestBody := domain.PaymentRequest{
		ClientID:clientId,
		ClientToken: clientToken,
		PaymentReferenceID: paymentReferenceId,
	}

	jsonBody, err := json.Marshal(requestBody)

	if err != nil {

		return domain.PaymentResponse{}, err
	}

	byteResp, err := client.postJSON(url, jsonBody)

	if err != nil {

		return domain.PaymentResponse{}, err
	}

	paymentResponse := domain.PaymentResponse{}

	err = json.Unmarshal(byteResp, &paymentResponse)

	if err != nil {

		return domain.PaymentResponse{}, err
	}

	return paymentResponse, nil
}

func (client *clientImpl) StartDelivery(serviceId int, serviceDeliveryToken string, unitsToSupply int) (int, error) {

	return 0, errors.New("Not implemented..")
}

func (client *clientImpl) EndDelivery(serviceId int, serviceDeliveryToken string, unitsReceived int) (int, error) {

	return 0, errors.New("Not implemented..")
}

func (client *clientImpl) get(url string) ([]byte, error) {

	response, err := http.Get(url)

	if err != nil {

		return nil, err
	}

	byteResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {

		return nil, err
	}

	return byteResponse, nil
}

func (_client *clientImpl) postJSON(url string, postBody []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))

	if err != nil {

		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return nil, err
	}

	return bodyBytes, nil
}