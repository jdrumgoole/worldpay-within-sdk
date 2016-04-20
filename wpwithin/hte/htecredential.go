package hte
import "errors"

type HTECredential struct {

	MerchantClientKey string
	MerchantServiceKey string
}

func NewHTECredential(MerchantClientKey, MerchantServiceKey string) (HTECredential, error) {

	if(MerchantClientKey == "") {

		return HTECredential{}, errors.New("MerchantClientKey required, cannot be empty")

	} else if(MerchantServiceKey == "") {

		return HTECredential{}, errors.New("MerchantServiceKey required, cannot be empty")
	}

	result := HTECredential{
		MerchantClientKey:MerchantClientKey,
		MerchantServiceKey:MerchantServiceKey,
	}

	return result, nil
}