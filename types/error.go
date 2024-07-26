package types

type ApiError struct {
	Error string `json:"error"`
}

func NewApiError(errMsg string) ApiError {
	return ApiError{
		Error: errMsg,
	}
}
