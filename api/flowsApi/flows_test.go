package flows

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

type FlowsApiTestSuite struct {
	suite.Suite
	api          FlowsApi
	mockedClient *common.MockHTTPClient
}

func (suit *FlowsApiTestSuite) SetupTest() {
	var apiKey = "test-key"
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())
	suit.mockedClient = common.NewMockHTTPClient()
	suit.api = NewFlowsApi(session, suit.mockedClient)
}

// ---- Test GetFlows
func (suit *FlowsApiTestSuite) TestGetFlowsMessageServerError() {
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	opt := &GetFlowsOptions{
		FlowActionFields: []models.FlowActionField{models.FlowActionFieldActionStatus},
		FlowFields:       []models.FlowField{models.FlowFieldArchived},
		TagFields:        []models.TagField{models.TagFieldName},
		Include:          []FlowsIncludeField{FlowsIncludeFieldTags, FlowsIncludeFieldFlowActions},
	}

	pageSize := 30
	var sortField FlowSortField = FlowSortFieldCreatedAtDESC
	paginationOpt := FlowPaginationOptions{
		PageSize: &pageSize,
		Sort:     &sortField,
	}

	_, err = suit.api.GetFlows(context.Background(), nil, opt, &paginationOpt)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *FlowsApiTestSuite) TestGetFlowsMessageBadRequest() {
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	opt := &GetFlowsOptions{
		FlowActionFields: []models.FlowActionField{models.FlowActionFieldActionStatus},
		FlowFields:       []models.FlowField{models.FlowFieldArchived},
		TagFields:        []models.TagField{models.TagFieldName},
		Include:          []FlowsIncludeField{FlowsIncludeFieldTags, FlowsIncludeFieldFlowActions},
	}

	pageSize := 30
	var sortField FlowSortField = FlowSortFieldCreatedAtDESC
	paginationOpt := FlowPaginationOptions{
		PageSize: &pageSize,
		Sort:     &sortField,
	}

	_, err = suit.api.GetFlows(context.Background(), nil, opt, &paginationOpt)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *FlowsApiTestSuite) TestGetFlowsMessageStatusOk() {
	mockedRespData := mockFlowsCollectionResource(3)

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	opt := &GetFlowsOptions{
		FlowActionFields: []models.FlowActionField{models.FlowActionFieldActionStatus},
		FlowFields:       []models.FlowField{models.FlowFieldArchived},
		TagFields:        []models.TagField{models.TagFieldName},
		Include:          []FlowsIncludeField{FlowsIncludeFieldTags, FlowsIncludeFieldFlowActions},
	}

	pageSize := 50
	var sortField FlowSortField = FlowSortFieldCreatedAtDESC
	paginationOpt := FlowPaginationOptions{
		PageSize: &pageSize,
		Sort:     &sortField,
	}

	res, err := suit.api.GetFlows(context.Background(), nil, opt, &paginationOpt)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data[0].ID, res.Data[0].ID)
}

// ---- Test GetFlow
func (suit *FlowsApiTestSuite) TestGetFlowMessageServerError() {
	var flowId = "test id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	opt := &GetFlowsOptions{
		FlowActionFields: []models.FlowActionField{models.FlowActionFieldActionStatus},
		FlowFields:       []models.FlowField{models.FlowFieldArchived},
		TagFields:        []models.TagField{models.TagFieldName},
		Include:          []FlowsIncludeField{FlowsIncludeFieldTags, FlowsIncludeFieldFlowActions},
	}

	_, err = suit.api.GetFlow(context.Background(), flowId, opt)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *FlowsApiTestSuite) TestGetFlowMessageBadRequest() {
	var flowId = "test id"
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	opt := &GetFlowsOptions{
		FlowActionFields: []models.FlowActionField{models.FlowActionFieldActionStatus},
		FlowFields:       []models.FlowField{models.FlowFieldArchived},
		TagFields:        []models.TagField{models.TagFieldName},
		Include:          []FlowsIncludeField{FlowsIncludeFieldTags, FlowsIncludeFieldFlowActions},
	}

	_, err = suit.api.GetFlow(context.Background(), flowId, opt)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

// ---- Test UpdateFlowStatus

func (suit *FlowsApiTestSuite) TestUpdateFlowStatusServerError() {
	var messageId = "message-id"
	reqData := mockUpdateFlowStatus()
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusInternalServerError, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.UpdateFlowStatus(context.Background(), messageId, reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *FlowsApiTestSuite) TestUpdateFlowStatusBadRequest() {
	var messageId = "message-id"
	reqData := mockUpdateFlowStatus()
	mockedRespData := common.MockedErrorResponse()

	err := common.PrepareMockResponse(http.StatusBadRequest, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	_, err = suit.api.UpdateFlowStatus(context.Background(), messageId, reqData)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *FlowsApiTestSuite) TestUpdateFlowStatusStatusOk() {
	var messageId = "message-id"
	reqData := mockUpdateFlowStatus()
	mockedRespData := mockFlowResource()

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.UpdateFlowStatus(context.Background(), messageId, reqData)

	suit.Nil(err)
	suit.Equal(mockedRespData.Data.ID, res.Data.ID)
}

func TestFlowsApiTestSuite(t *testing.T) {
	suite.Run(t, new(FlowsApiTestSuite))

}
