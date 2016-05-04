package psp_tests

import (
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/psp"
)

type owpMock struct {


}

func New() (psp.Psp, error) {

	result := &owpMock{}

	return result, nil
}

func (owp *owpMock) GetToken(reusableToken bool) (string, error) {

	return "", nil
}

func (owp *owpMock) MakePayment(amount int, orderDescription, customerOrderCode string) (string, error) {

	return "", nil
}