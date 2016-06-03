package psp
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/hce"

type Psp interface {

	GetToken(hceCredentials *hce.CardCredential, reusableToken bool) (string, error)
	MakePayment(amount int, currencyCode, clientToken, orderDescription, customerOrderCode string) (string, error)
}
