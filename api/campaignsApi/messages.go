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
	MessagesApi interface {
		//Returns a specific message based on a required id.
		GetCampaignMessage(ctx context.Context, messageId string, options *GetCampaignMessageOptions) (*models.CampaignMessageResponse, error)

		//Update a campaign message
		UpdateCampaignMessage(ctx context.Context, messageId string, payload UpdateCampaignMessagePayload) (*models.CampaignMessageResponse, error)
	}
)

type GetCampaignMessageOptions struct {
	CampaignFields        []models.CampaignsField
	CampaignMessageFields []models.CampaignMessageField
	TemplateFields        []models.TemplateField
	Include               []models.CampaignIncludeField
}

func buildGetCampaignMessageParams(opt *GetCampaignMessageOptions) string {
	var params = ""

	if opt == nil {
		return params
	}

	params = fmt.Sprintf("%s&%s", params, models.BuildCampaignFieldsParam(opt.CampaignFields))
	params = fmt.Sprintf("%s&%s", params, models.BuildCampaignMessageFieldsParam(opt.CampaignMessageFields))
	params = fmt.Sprintf("%s&%s", params, models.BuildTemplateFieldParam(opt.TemplateFields))
	params = fmt.Sprintf("%s&%s", params, models.BuildCampaignIncludeFieldParam(opt.Include))

	return params
}

func (api *campaignsApi) GetCampaignMessage(ctx context.Context, messageId string, options *GetCampaignMessageOptions) (*models.CampaignMessageResponse, error) {
	queryParams := buildGetCampaignMessageParams(options)
	url := fmt.Sprintf("%s/api/campaign-messages/%s/?%s", api.baseApiUrl, messageId, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var messageResponse models.CampaignMessageResponse
	err = json.Unmarshal(byteData, &messageResponse)

	return &messageResponse, err
}

func (api *campaignsApi) UpdateCampaignMessagePayload(ctx context.Context, data CreateCampaignRequestData) (*models.CampaignResponse, error) {
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

func (api *campaignsApi) UpdateCampaignMessage(ctx context.Context, messageId string, payload UpdateCampaignMessagePayload) (*models.CampaignMessageResponse, error) {
	url := fmt.Sprintf("%s/api/campaign-messages/%s/", api.baseApiUrl, messageId)

	reqData, err := json.Marshal(payload)
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

	var resp models.CampaignMessageResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

func (api *campaignsApi) AssignCampaignMessageTemplate(ctx context.Context, payload AssignCampaignMessageTemplatePayload) (*models.CampaignMessageResponse, error) {
	url := fmt.Sprintf("%s/api/campaign-message-assign-template/", api.baseApiUrl)

	reqData, err := json.Marshal(payload)
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

	var resp models.CampaignMessageResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}
