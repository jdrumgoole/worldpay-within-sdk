package psp_tests

import (
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/psp"
)

type owpMock struct {
}

func New() (psp.PSP, error) {

	result := &owpMock{}

	return result, nil
}

func (owp *owpMock) GetToken(reusableToken bool) (string, error) {

	return "", nil
}

func (owp *owpMock) MakePayment(amount int, orderDescription, customerOrderCode string) (string, error) {

	return "", nil
}
