package types

type PaymentRequest struct {

	ClientID string `json:"clientID"`
	ClientToken string `json:"clientToken"`
	PaymentReferenceID string `json:"paymentReferenceID"`
}
