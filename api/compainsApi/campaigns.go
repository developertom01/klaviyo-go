package campaigns

import (
	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/session"
)

type (
	CampaignsApi interface {
		// GetCampaigns()
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

// func
