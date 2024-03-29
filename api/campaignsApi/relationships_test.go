package campaigns

import (
	"context"
	"net/http"
	"testing"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/developertom01/klaviyo-go/models"
	"github.com/developertom01/klaviyo-go/options"
	"github.com/stretchr/testify/suite"
)

type CampaignsRelationshipsApiTestSuite struct {
	suite.Suite
	api          CampaignsApi
	mockedClient *common.MockHTTPClient
}

func (suit *CampaignsRelationshipsApiTestSuite) SetupTest() {
	var apiKey = "test-key"
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())
	suit.mockedClient = common.NewMockHTTPClient()
	suit.api = NewCampaignsApi(session, suit.mockedClient)
}

//----- Test GetCampaignMessageRelationshipsCampaign

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignMessageRelationshipsCampaignBadRequest() {
	var messageId = "message-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignMessageRelationshipsCampaign(context.Background(), messageId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignMessageRelationshipsCampaignServerError() {
	var messageId = "message-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignMessageRelationshipsCampaign(context.Background(), messageId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignMessageRelationshipsCampaignStatusOK() {
	var messageId = "message-id"
	mockedRespData := models.MockRelationshipData("campaign")

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetCampaignMessageRelationshipsCampaign(context.Background(), messageId)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

//----- Test GetCampaignMessageRelationshipsTemplate

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignMessageRelationshipsTemplateBadRequest() {
	var messageId = "message-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignMessageRelationshipsTemplate(context.Background(), messageId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignMessageRelationshipsTemplateServerError() {
	var messageId = "message-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignMessageRelationshipsTemplate(context.Background(), messageId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignMessageRelationshipsTemplateCampaignStatusOK() {
	var messageId = "message-id"
	mockedRespData := models.MockRelationshipData("templates")

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetCampaignMessageRelationshipsTemplate(context.Background(), messageId)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

//----- Test GetCampaignRelationshipsTags

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignRelationshipsTagsBadRequest() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignRelationshipsTags(context.Background(), campaignId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignRelationshipsTagsServerError() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadGateway, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignRelationshipsTags(context.Background(), campaignId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignRelationshipsTagsCampaignStatusOK() {
	var campaignId = "campaign-id"
	mockedRespData := models.MockRelationshipDataCollectionResponse("tags", 3)

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetCampaignRelationshipsTags(context.Background(), campaignId)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

//----- Test GetCampaignRelationshipsCampaignMessages

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignRelationshipsCampaignMessagesBadRequest() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignRelationshipsCampaignMessages(context.Background(), campaignId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignRelationshipsCampaignMessagesServerError() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadGateway, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignRelationshipsCampaignMessages(context.Background(), campaignId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsJobsApiTestSuite) TestGetCampaignRelationshipsCampaignMessagesStatusOK() {
	var campaignId = "campaign-id"
	mockedRespData := models.MockRelationshipDataCollectionResponse("campaign-messages", 3)

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetCampaignRelationshipsCampaignMessages(context.Background(), campaignId)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

func TestCampaignsRelationshipsApiTestSuite(t *testing.T) {
	suite.Run(t, new(CampaignsRelationshipsApiTestSuite))
}
