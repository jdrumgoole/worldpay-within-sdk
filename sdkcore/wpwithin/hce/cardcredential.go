package hce

type CardCredential struct {

	FirstName string
	LastName string
	ExpMonth int32
	ExpYear int32
	CardNumber string
	Type string
	Cvc string
}

func NewHCECardCredential() (*CardCredential, error) {

	result := &CardCredential{}

	return result, nil
}