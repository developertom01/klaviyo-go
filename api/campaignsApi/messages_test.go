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

type CampaignsMessageApiTestSuite struct {
	suite.Suite
	api          CampaignsApi
	mockedClient *common.MockHTTPClient
}

func (suit *CampaignsMessageApiTestSuite) SetupTest() {
	var apiKey = "test-key"
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())
	suit.mockedClient = common.NewMockHTTPClient()
	suit.api = NewCampaignsApi(session, suit.mockedClient)
}

func (suit *CampaignsMessageApiTestSuite) TestGetGetCampaignMessageServerError() {
	var messageId = "message-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignMessage(context.Background(), messageId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

func (suit *CampaignsMessageApiTestSuite) TestGetGetCampaignMessageBadRequest() {
	var messageId = "message-id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.GetCampaignMessage(context.Background(), messageId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

// Test when GetCampaignMessage returns 200 response
func (suit *CampaignsMessageApiTestSuite) TestGetGetCampaignMessageStatusOk() {
	var messageId = "message-id"
	mockedRespData := mockCampaignMessageResponse()

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	opt := GetCampaignMessageOptions{
		CampaignFields:        []models.CampaignsField{models.CampaignsFieldArchived},
		CampaignMessageFields: []models.CampaignMessageField{models.CampaignMessageFieldContent},
		TemplateFields:        []models.TemplateField{models.TemplateFieldHtml},
		Include:               []models.CampaignIncludeField{models.CampaignIncludeFieldTags},
	}
	res, err := suit.api.GetCampaignMessage(context.Background(), messageId, &opt)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func TestCampaignsMessageApiTestSuite(t *testing.T) {
	suite.Run(t, new(CampaignsMessageApiTestSuite))

}
