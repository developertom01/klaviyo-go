package klaviyo

import (
	accounts "github.com/developertom01/klaviyo-go/api/accountsApi"
	campaigns "github.com/developertom01/klaviyo-go/api/campaignsApi"
	flows "github.com/developertom01/klaviyo-go/api/flowsApi"
	images "github.com/developertom01/klaviyo-go/api/imagesApi"
	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/options"
)

type KlaviyoApi struct {
	Accounts  accounts.AccountsApi   //Accounts API
	Campaigns campaigns.CampaignsApi //Campaigns API
	Flows     flows.FlowsApi         //Flows API
	Images    images.ImagesApi
}

func NewKlaviyoApi(options options.Options, retryOption *common.RetryOptions) *KlaviyoApi {
	session := common.NewApiKeySession(options, retryOption)

	return &KlaviyoApi{
		Accounts:  accounts.NewAccountsApi(session, nil),
		Campaigns: campaigns.NewCampaignsApi(session, nil),
		Flows:     flows.NewFlowsApi(session, nil),
		Images:    images.NewImagesApi(session, nil),
	}
}
