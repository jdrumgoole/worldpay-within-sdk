package types

type ErrorResponse struct {


	HttpStatusCode int `json:"httpStatusCode"`
	CustomCode string `json:"customCode"`
	Message string `json:"message"`
	Description string `json:"description"`
	ErrorHelpUrl string `json:"errorHelpUrl"`
	OriginalRequest string `json:"originalRequest"`
}