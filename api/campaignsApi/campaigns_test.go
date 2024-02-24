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

func (suit *CampaignsApiTestSuite) TestDeleteCampaignsServerError() {
	var campaignId = "test id2"

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

	err = suit.api.DeleteCampaigns(context.Background(), campaignId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestDeleteCampaignsBadRequest() {
	var campaignId = "test id"

	mockedRespData := common.MockedErrorResponse()
	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "400 bad Gateway",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	err = suit.api.DeleteCampaigns(context.Background(), campaignId)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestDeleteCampaignsOkRequest() {
	var campaignId = "test id"

	mockedRespData := common.MockedErrorResponse()
	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "204 No content",
		StatusCode: http.StatusNoContent,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	err = suit.api.DeleteCampaigns(context.Background(), campaignId)

	suit.Nil(err)
}

func (suit *CampaignsApiTestSuite) TestCreateCampaignServerError() {
	reqData := mockCreateCampaignRequestData()
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

	_, err = suit.api.CreateCampaign(context.Background(), reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestCreateCampaignBadRequest() {
	reqData := mockCreateCampaignRequestData()
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "400 bad Request",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.CreateCampaign(context.Background(), reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestCreateCampaignOKRequest() {
	reqData := mockCreateCampaignRequestData()
	mockedRespData := mockCampaignResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "201 Created",
		StatusCode: http.StatusCreated,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	res, err := suit.api.CreateCampaign(context.Background(), reqData)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func (suit *CampaignsApiTestSuite) TestUpdateCampaignsServerError() {
	var campaignId = "123232"
	reqData := mockCreateCampaignRequestData()
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

	_, err = suit.api.UpdateCampaigns(context.Background(), campaignId, reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestUpdateCampaignsBadRequest() {
	var campaignId = "123232"
	reqData := mockCreateCampaignRequestData()
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

	_, err = suit.api.UpdateCampaigns(context.Background(), campaignId, reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestUpdateCampaignsOKRequest() {
	var campaignId = "123232"
	reqData := mockCreateCampaignRequestData()
	mockedRespData := mockCampaignResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "200 Created",
		StatusCode: http.StatusOK,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	res, err := suit.api.UpdateCampaigns(context.Background(), campaignId, reqData)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func (suit *CampaignsApiTestSuite) TestGetCampaignRecipientEstimationServerError() {
	var campaignId = "123232"
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

	_, err = suit.api.GetCampaignRecipientEstimation(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestGetCampaignRecipientEstimationBadRequest() {
	var campaignId = "123232"
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "400 bad Gateway",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.GetCampaignRecipientEstimation(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *CampaignsApiTestSuite) TestGetCampaignRecipientEstimationOkStatus() {
	var campaignId = "test-campaign-1"
	mockedRespData := mockCampaignResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "200 bad Gateway",
		StatusCode: http.StatusOK,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	res, err := suit.api.GetCampaignRecipientEstimation(context.Background(), campaignId, nil)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func TestCampaignsApiTestSuite(t *testing.T) {
	suite.Run(t, new(CampaignsApiTestSuite))

}
