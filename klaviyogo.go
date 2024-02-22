package klaviyo

import (
	"github.com/developertom01/klaviyo-go/accounts"
	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/options"
	"github.com/developertom01/klaviyo-go/session"
)

type Auth interface{}

type KlaviyoApi struct {
	Accounts accounts.AccountsApi
}

func NewKlaviyoApi(apiKey string, retryOption *common.RetryOptions) *KlaviyoApi {
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := session.NewApiKeySession(opt, retryOption)

	return &KlaviyoApi{
		Accounts: accounts.NewAccountsApi(session, nil),
	}
}
