package exceptions

// Error when status code returns 5xx or 4xx
type ErrorResponse struct {
	err ApiErrorResponse
}

func (ErrorResponse) Error() string {
	return "Api call error"
}

func NewResponseError(err ApiErrorResponse) ErrorResponse {
	return ErrorResponse{
		err: err,
	}
}
