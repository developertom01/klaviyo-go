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

type CampaignJobsApi interface {
	//Get a campaign send job
	GetCampaignSendJob(ctx context.Context, jobFields []models.CampaignSendJobField) (*models.CampaignSendJobResponse, error)

	//Permanently cancel the campaign, setting the status to CANCELED or revert the campaign, setting the status back to DRAFT
	UpdateCampaignSendJob(ctx context.Context, jobId string, payload UpdateCampaignSendJobPayload) (*models.CampaignSendJobResponse, error)

	//Retrieve the status of a recipient estimation job triggered with the Create Campaign Recipient Estimation Job endpoint.
	//`campaignId`is the ID of the campaign to get recipient estimation status
	GetCampaignRecipientEstimationJob(ctx context.Context, campaignId string, jobFields []models.CampaignSendJobField) (*models.CampaignSendJobResponse, error)

	//Permanently cancel the campaign, setting the status to CANCELED or revert the campaign, setting the status back to DRAFT
	CreateCampaignSendJob(ctx context.Context, payload CreateCampaignSendJobPayload) (*models.CampaignSendJobResponse, error)

	//Trigger an asynchronous job to update the estimated number of recipients
	//for the given campaign ID. Use the `GetCampaignRecipientEstimationJob` method or Get Campaign Recipient Estimation Job endpoint to retrieve the status of this estimation job. Use the
	//Get Campaign Recipient Estimation endpoint to retrieve the estimated
	//recipient count for a given campaign.
	CreateCampaignRecipientEstimationJob(ctx context.Context, payload CreateCampaignRecipientEstimationJobPayload) (*models.CampaignSendJobResponse, error)
}

func (api *campaignsApi) GetCampaignSendJob(ctx context.Context, jobFields []models.CampaignSendJobField) (*models.CampaignSendJobResponse, error) {
	fieldsParam := models.BuildCampaignSendJobFieldsParam(jobFields)
	url := fmt.Sprintf("%s/api/campaigns/?%s", api.baseApiUrl, fieldsParam)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var jobResp models.CampaignSendJobResponse
	err = json.Unmarshal(byteData, &jobResp)
	return &jobResp, err
}

func (api *campaignsApi) UpdateCampaignSendJob(ctx context.Context, jobId string, payload UpdateCampaignSendJobPayload) (*models.CampaignSendJobResponse, error) {
	url := fmt.Sprintf("%s/api/campaign-send-jobs/%s", api.baseApiUrl, jobId)

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

	var resp models.CampaignSendJobResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

func (api *campaignsApi) GetCampaignRecipientEstimationJob(ctx context.Context, campaignId string, jobFields []models.CampaignSendJobField) (*models.CampaignSendJobResponse, error) {
	fieldsParam := models.BuildCampaignSendJobFieldsParam(jobFields)
	url := fmt.Sprintf("%s/api/campaign-recipient-estimation-jobs/%s/?%s", api.baseApiUrl, campaignId, fieldsParam)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getCampaignsApiCallError, err)
	}

	var jobResp models.CampaignSendJobResponse
	err = json.Unmarshal(byteData, &jobResp)
	return &jobResp, err
}

func (api campaignsApi) CreateCampaignSendJob(ctx context.Context, payload CreateCampaignSendJobPayload) (*models.CampaignSendJobResponse, error) {
	url := fmt.Sprintf("%s/api/campaign-send-jobs/", api.baseApiUrl)

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

	var resp models.CampaignSendJobResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}

func (api campaignsApi) CreateCampaignRecipientEstimationJob(ctx context.Context, payload CreateCampaignRecipientEstimationJobPayload) (*models.CampaignSendJobResponse, error) {
	url := fmt.Sprintf("%s/api/campaign-recipient-estimation-jobs/", api.baseApiUrl)

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

	var resp models.CampaignSendJobResponse
	err = json.Unmarshal(byteData, &resp)

	return &resp, err
}
