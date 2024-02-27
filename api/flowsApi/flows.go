package flows

import (
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
)

type (
	FlowsApi interface {
	}

	flowsApi struct {
		session    common.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}
)

func NewFlowsApi(session common.Session, httpClient common.HTTPClient) FlowsApi {
	var client common.HTTPClient
	if httpClient == nil {
		client = http.DefaultClient
	} else {
		client = httpClient
	}

	return &flowsApi{
		session:    session,
		baseApiUrl: common.BASE_URL,
		revision:   common.API_REVISION,
		httpClient: client,
	}
}
