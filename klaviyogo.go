package klaviyo

import (
	accounts "github.com/developertom01/klaviyo-go/api/accountsApi"
	campaigns "github.com/developertom01/klaviyo-go/api/campaignsApi"
	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/options"
)

type KlaviyoApi struct {
	Accounts  accounts.AccountsApi
	Campaigns campaigns.CampaignsApi
}

func NewKlaviyoApi(apiKey string, retryOption *common.RetryOptions) *KlaviyoApi {
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, retryOption)

	return &KlaviyoApi{
		Accounts:  accounts.NewAccountsApi(session, nil),
		Campaigns: campaigns.NewCampaignsApi(session, nil),
	}
}
