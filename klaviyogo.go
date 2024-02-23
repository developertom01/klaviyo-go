package klaviyo

import (
	accounts "github.com/developertom01/klaviyo-go/api/accountsApi"
	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/options"
)

type KlaviyoApi struct {
	Accounts accounts.AccountsApi
}

func NewKlaviyoApi(apiKey string, retryOption *common.RetryOptions) *KlaviyoApi {
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, retryOption)

	return &KlaviyoApi{
		Accounts: accounts.NewAccountsApi(session, nil),
	}
}
