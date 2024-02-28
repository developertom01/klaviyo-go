package campaigns

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

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
		GetCampaignRecipientEstimation(ctx context.Context, id string, fields []models.CampaignRecipientEstimationField) (*models.CampaignRecipientCountResponse, error)

		//Clones an existing campaign, returning a new campaign based on the original with a new ID and name.
		CreateCampaignClone(ctx context.Context, data CreateCampaignCloneRequestData) (*models.CampaignResponse, error)

		//Return the related campaign
		GetCampaignMessageCampaign(ctx context.Context, messageId string, campaignFields []models.CampaignsField) (*models.CampaignResponse, error)

		//Return the related template for `messageId`
		GetCampaignMessageTemplate(ctx context.Context, messageId string, templateFields []models.TemplateField) (*models.TemplateResponse, error)
		//Return all tags that belong to the given campaign.
		GetCampaignTags(ctx context.Context, campaignId string, tagFields []models.TagField) (*models.TagsCollectionResponse, error)

		//Return all messages that belong to the given campaign.
		GetCampaignMessages(ctx context.Context, campaignId string, options *GetCampaignMessagesOptions) (*models.CampaignMessageCollectionResponse, error)

		//Campaign Messages API
		MessagesApi

		//Campaign Send Jobs API
		CampaignJobsApi

		//Relationships API
		CampaignRelationshipsAPi
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
	CampaignFields        []models.CampaignsField
	CampaignMessageFields []models.CampaignMessageField
	TagFields             []models.TagField
	PageCursor            *string
	Sort                  *models.CampaignSortField
	Include               []models.CampaignIncludeField
}

func buildGetCampaignsParams(filter string, opt *GetCampaignsOptions) string {
	if opt == nil {
		return ""
	}

	var params = []string{filter}

	if opt.CampaignFields != nil {
		params = append(params, models.BuildCampaignFieldsParam(opt.CampaignFields))
	}

	if opt.CampaignMessageFields != nil {
		params = append(params, models.BuildCampaignMessageFieldsParam(opt.CampaignMessageFields))
	}

	if opt.TagFields != nil {
		params = append(params, models.BuildTagFieldParam(opt.TagFields))
	}

	if opt.Include != nil {
		params = append(params, models.BuildCampaignIncludeFieldParam(opt.Include))
	}

	if opt.PageCursor != nil {
		params = append(params, fmt.Sprintf("page[cursor]=%s", *opt.PageCursor))
	}

	if opt.Sort != nil {
		params = append(params, fmt.Sprintf("sort=%s", *opt.Sort))
	}

	return strings.Join(params, "&")
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

func (api *campaignsApi) GetCampaignRecipientEstimation(ctx context.Context, id string, fields []models.CampaignRecipientEstimationField) (*models.CampaignRecipientCountResponse, error) {
	var fieldsParam = models.BuildCampaignRecipientEstimateFieldParam(fields)
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

func (api *campaignsApi) GetCampaignMessageCampaign(ctx context.Context, messageId string, campaignFields []models.CampaignsField) (*models.CampaignResponse, error) {

	var params = models.BuildCampaignFieldsParam(campaignFields)
	url := fmt.Sprintf("%s/api/campaign-messages/%s/campaign/?%s", api.baseApiUrl, messageId, params)

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

func (api *campaignsApi) GetCampaignMessageTemplate(ctx context.Context, messageId string, templateFields []models.TemplateField) (*models.TemplateResponse, error) {
	var params = models.BuildTemplateFieldParam(templateFields)
	url := fmt.Sprintf("%s/api/campaign-messages/%s/template/?%s", api.baseApiUrl, messageId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var template models.TemplateResponse
	err = json.Unmarshal(byteData, &template)

	return &template, err
}

// -- Untested
func (api *campaignsApi) GetCampaignTags(ctx context.Context, campaignId string, tagFields []models.TagField) (*models.TagsCollectionResponse, error) {
	var params = models.BuildTagFieldParam(tagFields)

	url := fmt.Sprintf("%s/api/campaigns/%s/tags/?%s", api.baseApiUrl, campaignId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var tags models.TagsCollectionResponse
	err = json.Unmarshal(byteData, &tags)

	return &tags, err
}

// Query parameters for GetCampaignMessages
type GetCampaignMessagesOptions struct {
	campaignMessageField []models.CampaignMessageField
	campaignFields       []models.CampaignsField
	templateFields       []models.TemplateField
	Include              []models.CampaignMessageIncludeField
}

func buildGetCampaignMessagesParams(opt *GetCampaignMessagesOptions) string {

	if opt == nil {
		return ""
	}

	var params = make([]string, 0)

	if opt.campaignMessageField != nil {
		params = append(params, models.BuildCampaignMessageFieldsParam(opt.campaignMessageField))
	}
	if opt.campaignFields != nil {
		params = append(params, models.BuildCampaignFieldsParam(opt.campaignFields))
	}
	if opt.campaignFields != nil {
		params = append(params, models.BuildTemplateFieldParam(opt.templateFields))
	}
	if opt.campaignFields != nil {
		params = append(params, models.BuildCampaignMessageIncludeFieldParam(opt.Include))
	}

	return strings.Join(params, "&")
}

func (api *campaignsApi) GetCampaignMessages(ctx context.Context, campaignId string, options *GetCampaignMessagesOptions) (*models.CampaignMessageCollectionResponse, error) {
	var params = buildGetCampaignMessagesParams(options)

	url := fmt.Sprintf("%s/api/campaigns/%s/campaign-messages/?%s", api.baseApiUrl, campaignId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var msg models.CampaignMessageCollectionResponse
	err = json.Unmarshal(byteData, &msg)

	return &msg, err
}
