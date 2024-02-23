package common

import (
	"fmt"
	"net/http"

	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/developertom01/klaviyo-go/options"
)

type Session interface {
	//Apply Api key header
	ApplyToRequest(option options.Options, req *http.Request) error
	GetRetryOptions() RetryOptions
	GetOptions() options.Options
}

const authorizationPrefix = "Klaviyo-API-Key"

type ApiKeySession struct {
	opt      options.Options
	retryOpt RetryOptions
}

func NewApiKeySession(opt options.Options, rOpt *RetryOptions) Session {
	options := options.NewOptionsWithDefaultValues().WithApiKey(*opt.ApiKey())

	var retryOptions *RetryOptions

	if rOpt == nil {
		retryOptions = NewRetryOptionsWithDefaultValues()
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

func (s ApiKeySession) GetRetryOptions() RetryOptions {
	return s.retryOpt
}

func (s ApiKeySession) GetOptions() options.Options {
	return s.opt
}
