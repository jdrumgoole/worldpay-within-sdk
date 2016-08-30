package mock

import "github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"

type PSP struct{}

func (psp PSP) GetToken(hceCredentials *types.HCECard, clientKey string, reusableToken bool) (string, error) {

	return "test-payment-token", nil
}
func (psp PSP) MakePayment(amount int, currencyCode, clientToken, orderDescription, customerOrderCode string) (string, error) {

	return "mock-payment-reference", nil
}
