package catalog

import (
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
)

type (
	CatalogApi interface {
		//Catalog item API
		CatalogItemApi
	}

	catalogApi struct {
		session    common.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}
)

func NewCatalogApi(session common.Session, httpClient common.HTTPClient) CatalogApi {
	var client common.HTTPClient
	if httpClient == nil {
		client = http.DefaultClient
	} else {
		client = httpClient
	}

	return &catalogApi{
		session:    session,
		baseApiUrl: common.BASE_URL,
		revision:   common.API_REVISION,
		httpClient: client,
	}
}
