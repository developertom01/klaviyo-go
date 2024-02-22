package exceptions

type (
	ApiErrorSource struct {
		Pointer   *string `json:"pinter,omitempty"`
		Parameter *string `json:"parameter,omitempty"`
	}

	ApiError struct {
		Id     string          `json:"id"`
		Code   string          `json:"code"`
		Title  string          `json:"title"`
		Detail string          `json:"detail"`
		Source *ApiErrorSource `json:"source,omitempty"`
	}

	ApiErrorResponse struct {
		Errors []ApiError `json:"errors"`
	}
)

// 2xx codes
func IsHttpCodeOk(code int) bool {
	return code >= 200 && code < 300
}

// 4xx errors
func IsHttpCodeClientError(code int) bool {
	return code >= 400 && code < 500
}

// 5xx errors
func IsHttpCodeServerError(code int) bool {
	return code >= 500
}
