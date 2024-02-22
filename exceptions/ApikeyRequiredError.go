package exceptions

type ApiKeyRequiredError struct {
	message string
}

func (e ApiKeyRequiredError) Error() string {
	return e.message
}

func NewApiKeyRequiredError(msg string) ApiKeyRequiredError {
	return ApiKeyRequiredError{
		message: msg,
	}
}
