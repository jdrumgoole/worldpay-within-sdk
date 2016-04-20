package onlineworldpay
import (
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/psp"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/hte"
	"errors"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/hce"
	"innovation.worldpay.com/worldpay-within-sdk/wpwithin/psp/onlineworldpay/types"
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
)

type OnlineWorldpay struct {

	hteCredential hte.HTECredential
	apiEndpoint string
}

func New(hteCredential hte.HTECredential, apiEndpoint string) (psp.Psp, error) {

	result := &OnlineWorldpay{
		hteCredential:hteCredential,
		apiEndpoint:apiEndpoint,
	}

	return result, nil
}

func (owp *OnlineWorldpay) GetToken(hceCredentials hce.HCECardCredential, reusableToken bool) (string, error) {

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
		ClientKey:owp.hteCredential.MerchantClientKey,
	}

	bJson, err := json.Marshal(tokenRequest)

	if err != nil {

		return "", err
	}

	reqUrl := fmt.Sprintf("%s/tokens", owp.apiEndpoint)

	var tokenResponse types.TokenResponse

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

	reqUrl := fmt.Sprintf("%s/orders", owp.apiEndpoint)

	var orderResponse types.OrderResponse

	headers := make(map[string]string, 0)

	headers["Authorization"] = owp.hteCredential.MerchantServiceKey

	err = post(reqUrl, bJson, headers, &orderResponse)

	if err != nil {

		return "", err
	}

	if strings.Compare(orderResponse.PaymentStatus, "SUCCESS") != 0 {

		return "", errors.New("Payment failed.")
	} else {

		return orderResponse.PaymentStatus, nil
	}
}

func post(url string, requestBody []byte, headers map[string]string, v interface{}) error {

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	for k, v := range headers {

		fmt.Printf("Adding header: %s => %s\n", k, v)
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
fmt.Printf("Response: %s\n", string(respBody))
	if err != nil {

		return err
	}

	err = json.Unmarshal(respBody, &v)

	if err != nil {

		return err
	} else {

		return nil
	}
}