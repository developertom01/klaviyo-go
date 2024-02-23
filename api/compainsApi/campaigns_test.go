package campaigns

import (
	"context"
	"net/http"
	"testing"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/developertom01/klaviyo-go/options"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CampaignsApiTestSuite struct {
	suite.Suite
	api          CampaignsApi
	mockedClient *common.MockHTTPClient
}

func (suit *CampaignsApiTestSuite) SetupTest() {
	var apiKey = "test-key"
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())
	suit.mockedClient = common.NewMockHTTPClient()
	suit.api = NewCampaignsApi(session, suit.mockedClient)
}

func (suit *CampaignsApiTestSuite) TestGetCampaignsBadRequest() {
	mockedRespData := common.MockedErrorResponse()
	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "400 bad request",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	filter := common.NewFilterBuilder().Equal("name", "sam")
	_, err = suit.api.GetCampaigns(context.Background(), filter.Build(), nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestGetCampaignsServerError() {
	mockedRespData := common.MockedErrorResponse()
	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "500 bad Gateway",
		StatusCode: http.StatusBadGateway,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	filter := common.NewFilterBuilder().Equal("name", "sam")
	_, err = suit.api.GetCampaigns(context.Background(), filter.Build(), nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestGetCampaignsOkResponse() {
	mockedRespData := mockCampaignCollectionResponse(3)
	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "200 ok",
		StatusCode: http.StatusOK,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	filter := common.NewFilterBuilder().Equal("name", "sam")
	resp, err := suit.api.GetCampaigns(context.Background(), filter.Build(), nil)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data[0].ID, resp.Data[0].ID)
}

func TestCampaignsApiTestSuite(t *testing.T) {
	suite.Run(t, new(CampaignsApiTestSuite))

}
