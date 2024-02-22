package session

import (
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/options"
)

type Session interface {
	//Apply Api key header
	ApplyToRequest(option options.Options, req *http.Request) error
	GetRetryOptions() common.RetryOptions
}
