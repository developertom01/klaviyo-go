package campaigns

import (
	"context"
	"net/http"
	"testing"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/developertom01/klaviyo-go/models"
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

// Test for GetCampaignRecipientEstimation if it returns 5xx code
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

// Test for GetCampaignRecipientEstimation if it returns 4xx code
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

// Test for GetCampaignRecipientEstimation if it returns 2xx code
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

// Test for CreateCampaignClone if it returns 5xx code
func (suit *CampaignsApiTestSuite) TestCreateCampaignCloneServerError() {
	reqData := mockCreateCampaignCloneRequestDataRequestData()
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "500 Server error",
		StatusCode: http.StatusInternalServerError,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.CreateCampaignClone(context.Background(), reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

// Test for CreateCampaignClone if it returns 4xx code
func (suit *CampaignsApiTestSuite) TestCreateCampaignCloneBadRequest() {
	reqData := mockCreateCampaignCloneRequestDataRequestData()
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "400 Server error",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.CreateCampaignClone(context.Background(), reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

// Test for CreateCampaignClone if it returns 201
func (suit *CampaignsApiTestSuite) TestCreateCampaignCloneOKRequest() {
	reqData := mockCreateCampaignCloneRequestDataRequestData()
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

	res, err := suit.api.CreateCampaignClone(context.Background(), reqData)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

// Test for CreateCampaignClone if it returns 4xx code
func (suit *CampaignsApiTestSuite) TestGetCampaignMessageCampaignBadRequest() {
	var messageId = "message-id"
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "400 Bad request",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.GetCampaignMessageCampaign(context.Background(), messageId, []models.CampaignsField{models.CampaignsFieldArchived, models.CampaignsFieldAudience_Exclude})

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

// Test for CreateCampaignClone if it returns 5xx code
func (suit *CampaignsApiTestSuite) TestGetCampaignMessageCampaignServerError() {
	var messageId = "message-id"
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "500 Internal Server Error",
		StatusCode: http.StatusInternalServerError,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.GetCampaignMessageCampaign(context.Background(), messageId, []models.CampaignsField{models.CampaignsFieldArchived, models.CampaignsFieldAudience})

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

// Test for CreateCampaignClone if it returns 2xx code
func (suit *CampaignsApiTestSuite) TestGetCampaignMessageCampaignStatusOk() {
	var messageId = "message-id"
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

	res, err := suit.api.GetCampaignMessageCampaign(context.Background(), messageId, []models.CampaignsField{models.CampaignsFieldAudience, models.CampaignsFieldAudience_Exclude})

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

// Test when GetCampaignMessages returns 4xx response
func (suit *CampaignsApiTestSuite) TestGetCampaignMessagesBadRequest() {
	var campaignId = "campaign-id1"
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "400 Created",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.GetCampaignMessages(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

// Test when GetCampaignMessages returns 5xx response
func (suit *CampaignsApiTestSuite) TestGetCampaignMessagesServerError() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "500 Created",
		StatusCode: http.StatusInternalServerError,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.GetCampaignMessages(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

// Test when GetCampaignMessages returns 200 response
func (suit *CampaignsApiTestSuite) TestGetCampaignMessagesStatusOk() {
	var campaignId = "campaign-id"
	mockedRespData := mockCampaignMessageCollectionResponse(3)

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

	res, err := suit.api.GetCampaignMessages(context.Background(), campaignId, nil)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data[0].ID, res.Data[0].ID)
}

// ---- Test GetCampaignMessageTemplate

// Test when GetCampaignMessageTemplate returns 5xx response
func (suit *CampaignsApiTestSuite) TestGetCampaignMessageTemplateServerError() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "500 Created",
		StatusCode: http.StatusInternalServerError,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.GetCampaignMessageTemplate(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

// Test when GetCampaignMessageTemplate returns 4xx response
func (suit *CampaignsApiTestSuite) TestGetCampaignMessageTemplateBadRequest() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "4000 Bad Request",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.GetCampaignMessageTemplate(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

// Test when GetCampaignMessages returns 200 response
func (suit *CampaignsApiTestSuite) TestGetCampaignMessageTemplateStatusOk() {
	var campaignId = "campaign-id"
	mockedRespData := models.MockTemplateResponse()

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
	res, err := suit.api.GetCampaignMessageTemplate(context.Background(), campaignId, []models.TemplateField{models.TemplateFieldCreatedAt, models.TemplateFieldCreatedAt})

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func (suit *CampaignsApiTestSuite) TestGetCampaignTagsServerError() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "500 Created",
		StatusCode: http.StatusInternalServerError,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.GetCampaignTags(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

func (suit *CampaignsApiTestSuite) TestGetCampaignTagsBadRequest() {
	var campaignId = "campaign-id"
	mockedRespData := common.MockedErrorResponse()

	bodyResp, err := common.PrepareMockResponse(mockedRespData)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "400 Bad Request",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)

	_, err = suit.api.GetCampaignTags(context.Background(), campaignId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)

}

// Test when GetCampaignMessages returns 200 response
func (suit *CampaignsApiTestSuite) TestGetCampaignTagsStatusOk() {
	var campaignId = "campaign-id"
	mockedRespData := models.MockTagsCollectionResponse(3)

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
	res, err := suit.api.GetCampaignTags(context.Background(), campaignId, []models.TagField{models.TagFieldName})

	suit.Nil(err)
	suit.Equal(mockedRespData.Data[0].ID, res.Data[0].ID)
}

func TestCampaignsApiTestSuite(t *testing.T) {
	suite.Run(t, new(CampaignsApiTestSuite))

}
