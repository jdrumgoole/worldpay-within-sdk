package domain

type Device struct {

	Uid string
	Name string
	Description string
	Services map[string]Service
}

func NewDevice(name, description, uid string) (*Device, error) {

	result := &Device{
		Name:name,
		Description:description,
		Uid:uid,
	}

	return result, nil
}