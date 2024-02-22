package session

import (
	"fmt"
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/developertom01/klaviyo-go/options"
)

const authorizationPrefix = "Klaviyo-API-Key"

type ApiKeySession struct {
	opt      options.Options
	retryOpt common.RetryOptions
}

func NewApiKeySession(opt options.Options, rOpt *common.RetryOptions) Session {
	options := options.NewOptionsWithDefaultValues().WithApiKey(*opt.ApiKey())

	var retryOptions *common.RetryOptions

	if rOpt == nil {
		retryOptions = common.NewRetryOptionsWithDefaultValues()
	} else {
		retryOptions = rOpt
	}

	return &ApiKeySession{
		opt:      options,
		retryOpt: *retryOptions,
	}
}

func (s ApiKeySession) ApplyToRequest(option options.Options, req *http.Request) error {
	if s.opt.ApiKey() == nil {
		return exceptions.NewApiKeyRequiredError("API key not set")
	}
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", authorizationPrefix, *s.opt.ApiKey()))

	return nil
}

func (s ApiKeySession) GetRetryOptions() common.RetryOptions {
	return s.retryOpt
}

func (s ApiKeySession) GetOptions() options.Options {
	return s.opt
}
