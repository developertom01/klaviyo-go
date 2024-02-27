package campaigns

import (
	"context"
	"net/http"
	"testing"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/developertom01/klaviyo-go/options"
	"github.com/stretchr/testify/suite"
)

type CampaignsJobsApiTestSuite struct {
	suite.Suite
	api          CampaignsApi
	mockedClient *common.MockHTTPClient
}

func (suit *CampaignsJobsApiTestSuite) SetupTest() {
	var apiKey = "test-key"
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())
	suit.mockedClient = common.NewMockHTTPClient()
	suit.api = NewCampaignsApi(session, suit.mockedClient)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignSendJobServerError() {
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignSendJob(context.Background(), nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignSendJobBadRequest() {
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignSendJob(context.Background(), nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignSendJobStatusOk() {
	mockedRespData := mockCampaignJobResponse()

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetCampaignSendJob(context.Background(), nil)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func (suit *CampaignsJobsApiTestSuite) TestUpdateCampaignSendJobServerError() {
	var jobId = "job-id"
	reqData := mockUpdateCampaignSendJobPayload()
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.UpdateCampaignSendJob(context.Background(), jobId, reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestUpdateCampaignSendJobBadRequest() {
	var jobId = "job-id"
	reqData := mockUpdateCampaignSendJobPayload()
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.UpdateCampaignSendJob(context.Background(), jobId, reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestUpdateCampaignSendJobStatusOk() {
	var jobId = "job-id"
	reqData := mockUpdateCampaignSendJobPayload()
	mockedRespData := mockCampaignJobResponse()

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.UpdateCampaignSendJob(context.Background(), jobId, reqData)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignRecipientEstimationJobBadRequest() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignRecipientEstimationJob(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignRecipientEstimationJobServerError() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadGateway, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignRecipientEstimationJob(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignRecipientEstimationJobStatusOk() {
	var campaignId = "campaign-id"
	mockedRespData := mockCampaignJobResponse()

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetCampaignRecipientEstimationJob(context.Background(), campaignId, nil)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func (suit *CampaignsJobsApiTestSuite) TestCreateCampaignSendJobServerError() {
	reqData := mockCampaignSendCreationJobPayload()
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.CreateCampaignSendJob(context.Background(), reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestCreateCampaignSendJobBadRequest() {
	reqData := mockCampaignSendCreationJobPayload()
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.CreateCampaignSendJob(context.Background(), reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestCreateCampaignSendJobStatusOk() {
	reqData := mockCampaignSendCreationJobPayload()
	mockedRespData := mockCampaignJobResponse()

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.CreateCampaignSendJob(context.Background(), reqData)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

// ------ Test CreateCampaignRecipientEstimationJob
func (suit *CampaignsJobsApiTestSuite) TestCreateCampaignRecipientEstimationJobServerError() {
	reqData := mockCampaignSendCreationJobPayload()
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.CreateCampaignRecipientEstimationJob(context.Background(), reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestCreateCampaignRecipientEstimationJobBadRequest() {
	reqData := mockCampaignSendCreationJobPayload()
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.CreateCampaignRecipientEstimationJob(context.Background(), reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestCreateCampaignRecipientEstimationJobStatusOk() {
	reqData := mockCampaignSendCreationJobPayload()
	mockedRespData := mockCampaignJobResponse()

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.CreateCampaignRecipientEstimationJob(context.Background(), reqData)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func TestCampaignsJobsApiTestSuite(t *testing.T) {
	suite.Run(t, new(CampaignsJobsApiTestSuite))

}
