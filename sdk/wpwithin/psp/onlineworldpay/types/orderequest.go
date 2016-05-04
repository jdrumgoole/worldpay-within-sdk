package types

type OrderRequest struct {

	Token string `json:"token"`

	Amount int `json:"amount"`

	CurrencyCode string `json:"currencyCode"`

	OrderDescription string `json:"orderDescription"`

	CustomerOrderCode string `json:"customerOrderCode"`
}