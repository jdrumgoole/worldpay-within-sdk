package psp
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"
)

type Psp interface {

	GetToken(hceCredentials *domain.HCECard, reusableToken bool) (string, error)
	MakePayment(amount int, currencyCode, clientToken, orderDescription, customerOrderCode string) (string, error)
}
