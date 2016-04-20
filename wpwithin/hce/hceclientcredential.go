package hce
import "errors"

type HCEClientCredential struct {

	ClientKey string
}

func NewHCEClientCredential(clientKey string) (HCEClientCredential, error) {

	if clientKey == "" {

		return HCEClientCredential{}, errors.New("ClientKey cannot be empty")
	}

	result := HCEClientCredential{
		ClientKey:clientKey,
	}

	return result, nil
}