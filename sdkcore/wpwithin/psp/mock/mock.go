package mock
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"

type MockPSP struct {}

func (psp *MockPSP) GetToken(hceCredentials *types.HCECard, clientKey string, reusableToken bool) (string, error) {

	return "mock-token-here", nil
}

func (psp *MockPSP) MakePayment(amount int, currencyCode, clientToken, orderDescription, customerOrderCode string) (string, error) {

	return "mock-psp-reference", nil
}