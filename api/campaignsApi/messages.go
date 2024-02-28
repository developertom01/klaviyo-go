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
	MessagesApi interface {
		//Returns a specific message based on a required id.
		GetCampaignMessage(ctx context.Context, messageId string, options *GetCampaignMessageOptions) (*models.CampaignMessageResponse, error)

		//Update a campaign message
		UpdateCampaignMessage(ctx context.Context, messageId string, payload UpdateCampaignMessagePayload) (*models.CampaignMessageResponse, error)

		//Creates a non-reusable version of the template and assigns it to the message.
		AssignCampaignMessageTemplate(ctx context.Context, payload AssignCampaignMessageTemplatePayload) (*models.CampaignMessageResponse, error)
	}
)

type GetCampaignMessageOptions struct {
	CampaignFields        []models.CampaignsField
	CampaignMessageFields []models.CampaignMessageField
	TemplateFields        []models.TemplateField
	Include               []models.CampaignIncludeField
}

func buildGetCampaignMessageParams(opt *GetCampaignMessageOptions) string {
	if opt == nil {
		return ""
	}
	var params = make([]string, 0)

	if opt.CampaignFields != nil {
		params = append(params, models.BuildCampaignFieldsParam(opt.CampaignFields))
	}

	if opt.CampaignMessageFields != nil {
		params = append(params, models.BuildCampaignMessageFieldsParam(opt.CampaignMessageFields))
	}

	if opt.TemplateFields != nil {
		params = append(params, models.BuildTemplateFieldParam(opt.TemplateFields))
	}

	if opt.Include != nil {
		params = append(params, models.BuildCampaignIncludeFieldParam(opt.Include))
	}

	return strings.Join(params, "&")
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

func (api *campaignsApi) UpdateCampaignMessage(ctx context.Context, messageId string, payload UpdateCampaignMessagePayload) (*models.CampaignMessageResponse, error) {
	url := fmt.Sprintf("%s/api/campaign-messages/%s/", api.baseApiUrl, messageId)

	reqData, err := json.Marshal(payload)
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
