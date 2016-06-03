package hce
import "errors"

type ClientCredential struct {

	ClientKey string
}

func NewHCEClientCredential(clientKey string) (ClientCredential, error) {

	if clientKey == "" {

		return ClientCredential{}, errors.New("ClientKey cannot be empty")
	}

	result := ClientCredential{
		ClientKey:clientKey,
	}

	return result, nil
}