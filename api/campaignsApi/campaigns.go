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
		CreateCampaign(ctx context.Context, data CreateCampaignRequestData) (*models.CampaignResponse, error)

		//Returns a specific campaign based on a required id.
		GetCampaign(ctx context.Context, id string, filter string, options *GetCampaignsOptions) (*models.CampaignResponse, error)

		//Update a campaign with the given campaign ID.
		UpdateCampaigns(ctx context.Context, id string, data CreateCampaignRequestData) (*models.CampaignResponse, error)

		//Delete a campaign with the given campaign ID.
		DeleteCampaigns(ctx context.Context, id string) error

		//Get the estimated recipient count for a campaign with the provided campaign ID.
		//You can refresh this count by using the Create Campaign Recipient Estimation Job endpoint.
		GetCampaignRecipientEstimation(ctx context.Context, id string, fieldsStr *string) (*models.CampaignRecipientCountResponse, error)

		//Clones an existing campaign, returning a new campaign based on the original with a new ID and name.
		CreateCampaignClone(ctx context.Context, data CreateCampaignCloneRequestData) (*models.CampaignResponse, error)
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
		baseApiUrl: common.BASE_URL,
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
	url := fmt.Sprintf("%s/api/campaigns/?%s", api.baseApiUrl, queryParams)

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

func (api *campaignsApi) CreateCampaign(ctx context.Context, data CreateCampaignRequestData) (*models.CampaignResponse, error) {
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

	var resp models.CampaignResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

func (api *campaignsApi) GetCampaign(ctx context.Context, id string, filter string, options *GetCampaignsOptions) (*models.CampaignResponse, error) {
	queryParams := buildGetCampaignsParams(filter, options)
	url := fmt.Sprintf("%s/api/campaigns/%s/?%s", api.baseApiUrl, id, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var campaignResp models.CampaignResponse
	err = json.Unmarshal(byteData, &campaignResp)

	return &campaignResp, err
}

func (api *campaignsApi) UpdateCampaigns(ctx context.Context, id string, data CreateCampaignRequestData) (*models.CampaignResponse, error) {
	url := fmt.Sprintf("%s/api/campaigns/%s/", api.baseApiUrl, id)

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

	var resp models.CampaignResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

func (api *campaignsApi) DeleteCampaigns(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/api/campaigns/%s/", api.baseApiUrl, id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = common.RetrieveData(api.httpClient, req, api.session, api.revision)
	return err
}

func (api *campaignsApi) GetCampaignRecipientEstimation(ctx context.Context, id string, fieldsStr *string) (*models.CampaignRecipientCountResponse, error) {
	var fieldsParam = ""
	if fieldsStr != nil {
		fieldsParam = *fieldsStr
	}
	url := fmt.Sprintf("%s/api/campaign-recipient-estimations/%s/?%s", api.baseApiUrl, id, fieldsParam)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var resp models.CampaignRecipientCountResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

// --Untested
func (api *campaignsApi) CreateCampaignClone(ctx context.Context, data CreateCampaignCloneRequestData) (*models.CampaignResponse, error) {
	url := fmt.Sprintf("%s/api/campaign-clone/", api.baseApiUrl)

	reqData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reqDataBuffer := bytes.NewBuffer(reqData)

	req, err := http.NewRequest(http.MethodPost, url, reqDataBuffer)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var resp models.CampaignResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

func (api *campaignsApi) GetCampaignMessageCampaign(ctx context.Context, messageId string, fieldsParam string) (*models.CampaignResponse, error) {
	url := fmt.Sprintf("%s/api/campaign-messages/%s/campaign/?%s", api.baseApiUrl, messageId, fieldsParam)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var campaignResp models.CampaignResponse
	err = json.Unmarshal(byteData, &campaignResp)

	return &campaignResp, err
}
