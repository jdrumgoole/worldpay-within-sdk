package types

type ErrorResponse struct {

	HTTPStatusCode int `json:"httpStatusCode"`
	Message string `json:"message"`
	ErrorCode int `json:errorCode`
}
