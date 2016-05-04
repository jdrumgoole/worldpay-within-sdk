package psp
import (
	"innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/hce"
)

type Psp interface {

	GetToken(hceCredentials hce.HCECardCredential, reusableToken bool) (string, error)
	MakePayment(amount int, currencyCode, clientToken, orderDescription, customerOrderCode string) (string, error)
}
