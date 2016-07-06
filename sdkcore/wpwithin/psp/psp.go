package psp
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

type Psp interface {

	GetToken(hceCredentials *types.HCECard, clientKey string, reusableToken bool) (string, error)
	MakePayment(amount int, currencyCode, clientToken, orderDescription, customerOrderCode string) (string, error)
}
