package onlineworldpay
import (
	"errors"
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp"
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp/onlineworldpay/types"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"
)

type OnlineWorldpay struct {

	MerchantClientKey string
	MerchantServiceKey string
	apiEndpoint string
}

func New(merchantClientKey, merchantServiceKey, apiEndpoint string) (psp.Psp, error) {

	result := &OnlineWorldpay{
		MerchantClientKey:merchantClientKey,
		MerchantServiceKey:merchantServiceKey,
		apiEndpoint:apiEndpoint,
	}

	return result, nil
}

func (owp *OnlineWorldpay) GetToken(hceCredentials *domain.HCECard, reusableToken bool) (string, error) {

	if(reusableToken) {
		// TODO: CH - support reusable token by storing the value (along with merchant client key so link to a merchant) within the car so that token can be re-used if present, or created if not
		return "", errors.New("Reusable token support not implemented")
	}

	paymentMethod := types.TokenRequestPaymentMethod{
		Name: fmt.Sprintf("%s %s", hceCredentials.FirstName, hceCredentials.LastName),
		ExpiryMonth:hceCredentials.ExpMonth,
		ExpiryYear:hceCredentials.ExpYear,
		CardNumber:hceCredentials.CardNumber,
		Type:hceCredentials.Type,
		Cvc:hceCredentials.Cvc,
		StartMonth:nil,
		StartYear:nil,
	}

	tokenRequest := types.TokenRequest{
		Reusable:reusableToken,
		PaymentMethod:paymentMethod,
		ClientKey:owp.MerchantClientKey,
	}

	bJson, err := json.Marshal(tokenRequest)

	if err != nil {

		return "", err
	}

	log.WithField("TokenRequest", string(bJson)).Debug("POST Request Token.")

	reqUrl := fmt.Sprintf("%s/tokens", owp.apiEndpoint)

	var tokenResponse types.TokenResponse

	log.WithFields(log.Fields{ "Url": reqUrl,
		"RequestJSON": string(bJson) }).Debug("Sending Token POST request.")

	err = post(reqUrl, bJson, make(map[string]string, 0), &tokenResponse)

	if err != nil {

		return "", err
	}

	return tokenResponse.Token, nil
}

func (owp *OnlineWorldpay) MakePayment(amount int, currencyCode, clientToken, orderDescription, customerOrderCode string) (string, error) {

	if clientToken == "" {

		return "", errors.New("clientToken cannot be empty")
	}
	if orderDescription == "" {

		return "", errors.New("orderDescription cannot be empty")
	}
	if customerOrderCode == "" {

		return "", errors.New("customerOrderCode cannot be empty")
	}

	orderRequest := types.OrderRequest{

		Token:clientToken,
		Amount:amount,
		CurrencyCode:currencyCode,
		OrderDescription:orderDescription,
		CustomerOrderCode:customerOrderCode,
	}

	bJson, err := json.Marshal(orderRequest)

	if err != nil {

		return "", err
	}

	log.WithField("JSON", string(bJson)).Debug("JSON form of OrderRequest object.")

	reqUrl := fmt.Sprintf("%s/orders", owp.apiEndpoint)

	var orderResponse types.OrderResponse

	headers := make(map[string]string, 0)

	headers["Authorization"] = owp.MerchantServiceKey

	log.WithFields(log.Fields{ "Url": reqUrl,
		"RequestJSON": string(bJson) }).Debug("Sending Order POST request.")

	err = post(reqUrl, bJson, headers, &orderResponse)

	if err != nil {

		return "", err
	}

	if strings.Compare(orderResponse.PaymentStatus, "SUCCESS") != 0 {

		return "", errors.New("Payment failed.")
	} else {

		return orderResponse.OrderCode, nil
	}
}

func post(url string, requestBody []byte, headers map[string]string, v interface{}) error {

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	for k, v := range headers {

		request.Header.Set(k, v)
	}

	if err != nil {

		return err
	}

	// TODO: CH Add a http client as a dependency during construction to aid testing
	client := &http.Client{}

	resp, err := client.Do(request)

	if err != nil {

		return err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return err
	}

	log.WithField("Response Body", string(respBody)).Debug("Response content")

	err = json.Unmarshal(respBody, &v)

	if err != nil {

		return err
	} else {

		return nil
	}
}