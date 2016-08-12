package onlineworldpay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/psp/onlineworldpay/types"
	wpwithin_types "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

type OnlineWorldpay struct {
	MerchantClientKey  string
	MerchantServiceKey string
	apiEndpoint        string
}

func NewMerchant(merchantClientKey, merchantServiceKey, apiEndpoint string) (psp.Psp, error) {

	result := &OnlineWorldpay{
		MerchantClientKey:  merchantClientKey,
		MerchantServiceKey: merchantServiceKey,
		apiEndpoint:        apiEndpoint,
	}

	return result, nil
}

func NewClient(apiEndpoint string) (psp.Psp, error) {

	result := &OnlineWorldpay{
		apiEndpoint: apiEndpoint,
	}

	return result, nil
}

func (owp *OnlineWorldpay) GetToken(hceCredentials *wpwithin_types.HCECard, clientKey string, reusableToken bool) (string, error) {

	if reusableToken {
		// TODO: CH - support reusable token by storing the value (along with merchant client key so link to a merchant) within the car so that token can be re-used if present, or created if not
		return "", errors.New("Reusable token support not implemented")
	}

	paymentMethod := types.TokenRequestPaymentMethod{
		Name:        fmt.Sprintf("%s %s", hceCredentials.FirstName, hceCredentials.LastName),
		ExpiryMonth: hceCredentials.ExpMonth,
		ExpiryYear:  hceCredentials.ExpYear,
		CardNumber:  hceCredentials.CardNumber,
		Type:        hceCredentials.Type,
		Cvc:         hceCredentials.Cvc,
		StartMonth:  nil,
		StartYear:   nil,
	}

	tokenRequest := types.TokenRequest{
		Reusable:      reusableToken,
		PaymentMethod: paymentMethod,
		ClientKey:     clientKey,
	}

	bJSON, err := json.Marshal(tokenRequest)

	if err != nil {

		return "", err
	}

	log.WithField("TokenRequest", string(bJSON)).Debug("POST Request Token.")

	reqURL := fmt.Sprintf("%s/tokens", owp.apiEndpoint)

	var tokenResponse types.TokenResponse

	log.WithFields(log.Fields{"Url": reqURL,
		"RequestJSON": string(bJSON)}).Debug("Sending Token POST request.")

	err = post(reqURL, bJSON, make(map[string]string, 0), &tokenResponse)

	return tokenResponse.Token, err
}

func (owp *OnlineWorldpay) MakePayment(amount int, currencyCode, clientToken, orderDescription, customerOrderCode string) (string, error) {

	log.WithFields(log.Fields{"Amount": strconv.Itoa(amount), "CurrencyCode": currencyCode, "ClientToken": clientToken,
		"OrderDescription": orderDescription, "CustomerOrderCode": customerOrderCode}).Debug("Begin OWP MakePayment")

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

		Token:             clientToken,
		Amount:            amount,
		CurrencyCode:      currencyCode,
		OrderDescription:  orderDescription,
		CustomerOrderCode: customerOrderCode,
	}

	bJSON, err := json.Marshal(orderRequest)

	if err != nil {

		return "", err
	}

	log.WithField("JSON", string(bJSON)).Debug("JSON form of OrderRequest object.")

	reqURL := fmt.Sprintf("%s/orders", owp.apiEndpoint)

	log.WithFields(log.Fields{"Request URL": reqURL, "MerchantSvcKey": owp.MerchantServiceKey}).Debug("Using OWP parameters")

	var orderResponse types.OrderResponse

	headers := make(map[string]string, 0)

	headers["Authorization"] = owp.MerchantServiceKey

	log.WithFields(log.Fields{"Url": reqURL,
		"RequestJSON": string(bJSON)}).Debug("Sending Order POST request.")

	err = post(reqURL, bJSON, headers, &orderResponse)

	if err != nil {

		return "", err
	}

	if strings.EqualFold(orderResponse.PaymentStatus, "SUCCESS") {

		return orderResponse.OrderCode, nil
	}

	return "", fmt.Errorf("Payment failed for customer order %s ", orderResponse.CustomerOrderCode)
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

	if resp.StatusCode == HTTP_OK {

		log.Debug(fmt.Sprintf("Response body: %s", string(respBody)))

		return json.Unmarshal(respBody, &v)
	} else {

		wpErr := types.ErrorResponse{}

		if err := json.Unmarshal(respBody, &wpErr); err == nil {

			log.WithFields(log.Fields{"Message": wpErr.Message, "Description": wpErr.Description, "CustomCode": wpErr.CustomCode, "HTTP Status Code": wpErr.HTTPStatusCode, "HelpUrl": wpErr.ErrorHelpURL}).Debug("** POST Response")

			return fmt.Errorf("HTTP Status: %d - CustomCode: %s - Message: %s", wpErr.HTTPStatusCode, wpErr.CustomCode, wpErr.Message)
		}
	}

	return nil
}
