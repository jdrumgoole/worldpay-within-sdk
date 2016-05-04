package hce

type HCECardCredential struct {

	FirstName string
	LastName string
	ExpMonth int32
	ExpYear int32
	CardNumber string
	Type string
	Cvc string
}

func NewHCECardCredential() (*HCECardCredential, error) {

	result := &HCECardCredential{}

	return result, nil
}