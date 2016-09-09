package types

// HCECard represents details of a payment card
type HCECard struct {
	FirstName  string
	LastName   string
	ExpMonth   int32
	ExpYear    int32
	CardNumber string
	Type       string
	Cvc        string
}
