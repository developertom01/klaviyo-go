package klaviyo

import (
	"github.com/developertom01/klaviyo-go/accounts"
	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/session"
)

type Auth interface{}

type KlaviyoApi struct {
	Accounts accounts.AccountsApi
}

func NewKlaviyoApi(apiKey string, retryOption *common.RetryOptions) *KlaviyoApi {
	session := session.NewApiKeySession(apiKey, retryOption)

	return &KlaviyoApi{
		Accounts: accounts.NewAccountsApi(session, nil),
	}
}
