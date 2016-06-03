package domain

type Order struct {

	ServiceID int
	ClientID string
	SelectedNumberOfUnits int
	SelectedPriceId int
	TotalPrice int
	PaymentReference string
	ClientUUID string
	PSPReference string
	DeliveryToken string
}
