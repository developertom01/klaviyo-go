package campaigns

import (
	"bytes"
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
		//Returns some or all campaigns based on filters.
		GetCampaigns(ctx context.Context, filter string, options *GetCampaignsOptions) (*models.CampaignsCollectionResponse, error)

		//Creates a campaign given a set of parameters, then returns it.
		CreateCampaigns(ctx context.Context, data CreateCampaignRequestData) (*models.CampaignsResponse, error)

		//Returns a specific campaign based on a required id.
		GetCampaign(ctx context.Context, id string, filter string, options *GetCampaignsOptions) (*models.CampaignsResponse, error)

		//Update a campaign with the given campaign ID.
		UpdateCampaigns(ctx context.Context, id string, data CreateCampaignRequestData) (*models.CampaignsResponse, error)

		//Delete a campaign with the given campaign ID.
		DeleteCampaigns(ctx context.Context, id string) error
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

func (api *campaignsApi) CreateCampaigns(ctx context.Context, data CreateCampaignRequestData) (*models.CampaignsResponse, error) {
	reqData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reqDataBuffer := bytes.NewBuffer(reqData)

	req, err := http.NewRequest(http.MethodPost, api.baseApiUrl, reqDataBuffer)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var resp models.CampaignsResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

func (api *campaignsApi) GetCampaign(ctx context.Context, id string, filter string, options *GetCampaignsOptions) (*models.CampaignsResponse, error) {
	queryParams := buildGetCampaignsParams(filter, options)
	url := fmt.Sprintf("%s/%s/?%s", api.baseApiUrl, id, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var campaignResp models.CampaignsResponse
	err = json.Unmarshal(byteData, &campaignResp)

	return &campaignResp, err
}

func (api *campaignsApi) UpdateCampaigns(ctx context.Context, id string, data CreateCampaignRequestData) (*models.CampaignsResponse, error) {
	url := fmt.Sprintf("%s/%s/", api.baseApiUrl, id)

	reqData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reqDataBuffer := bytes.NewBuffer(reqData)

	req, err := http.NewRequest(http.MethodPatch, url, reqDataBuffer)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var resp models.CampaignsResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

func (api *campaignsApi) DeleteCampaigns(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/%s/", api.baseApiUrl, id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = common.RetrieveData(api.httpClient, req, api.session, api.revision)
	return err
}
