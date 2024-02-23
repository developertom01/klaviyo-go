package campaigns

import (
	"context"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
	"github.com/developertom01/klaviyo-go/session"
)

type (
	CampaignsApi interface {
		GetCampaigns(ctx context.Context) models.CampaignMessage
	}

	campaignsApi struct {
		session    session.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}
)

func NewCampaignsApi(session session.Session, httpClient common.HTTPClient) CampaignsApi {
	return &campaignsApi{}
}

func (*campaignsApi) GetCampaigns(ctx context.Context) models.CampaignMessage {
	panic("")
}

// func
