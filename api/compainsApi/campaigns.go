package campaigns

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type (
	CampaignsApi interface {
		GetCampaigns(ctx context.Context, filter string, options *GetCampaignsOptions) (*models.CampaignsCollectionResponse, error)
	}

	campaignsApi struct {
		session    common.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}
)

func NewCampaignsApi(session common.Session, httpClient common.HTTPClient) CampaignsApi {
	return &campaignsApi{
		session:    session,
		httpClient: httpClient,
		baseApiUrl: fmt.Sprintf("%s/api/campaigns", common.BASE_URL),
		revision:   common.API_REVISION,
	}
}

type GetCampaignsOptions struct {
	CampaignFields        *string
	CampaignMessageFields *string
	TagFields             *string
	PageCursor            *string
	Sort                  *CampaignSortField
	Include               *string
}

func buildGetCampaignsParams(filter string, opt *GetCampaignsOptions) string {
	var params = filter

	if opt != nil && opt.CampaignFields != nil {
		params = fmt.Sprintf("%s&%s", params, *opt.CampaignFields)
	}
	if opt != nil && opt.CampaignMessageFields != nil {
		params = fmt.Sprintf("%s&%s", params, *opt.CampaignMessageFields)
	}
	if opt != nil && opt.TagFields != nil {
		params = fmt.Sprintf("%s&%s", params, *opt.TagFields)
	}
	if opt != nil && opt.PageCursor != nil {
		params = fmt.Sprintf("%s&page[cursor]=%s", params, *opt.PageCursor)
	}
	if opt != nil && opt.Sort != nil {
		params = fmt.Sprintf("%s&sort=%s", params, *opt.Sort)
	}

	return params
}

func (api *campaignsApi) GetCampaigns(ctx context.Context, filter string, options *GetCampaignsOptions) (*models.CampaignsCollectionResponse, error) {
	queryParams := buildGetCampaignsParams(filter, options)
	url := fmt.Sprintf("%s/?%s", api.baseApiUrl, queryParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var campaignResp models.CampaignsCollectionResponse
	err = json.Unmarshal(byteData, &campaignResp)
	
	return &campaignResp, err
}
