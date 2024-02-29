package flows

import (
	"context"
	"net/http"
	"testing"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
	"github.com/developertom01/klaviyo-go/options"
	"github.com/stretchr/testify/suite"
)

type FlowsRelationshipsApiTestSuite struct {
	suite.Suite
	api          FlowsApi
	mockedClient *common.MockHTTPClient
}

func (suit *FlowsRelationshipsApiTestSuite) SetupTest() {
	var apiKey = "test-key"
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())
	suit.mockedClient = common.NewMockHTTPClient()
	suit.api = NewFlowsApi(session, suit.mockedClient)
}

func (suit *FlowsRelationshipsApiTestSuite) TestGetFlowRelationshipsFlowActions() {
	var flowId = "flow-1"
	mockedRespData := models.MockRelationshipDataCollectionResponse("flow", 3)

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	filter := common.NewFilterBuilder().Contains("name", "flow-1")
	filterPram := filter.Build()
	res, err := suit.api.GetFlowRelationshipsFlowActions(context.Background(), flowId, &filterPram, nil)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

func (suit *FlowsRelationshipsApiTestSuite) TestGetFlowRelationshipsTags() {
	var flowId = "flow-1"
	mockedRespData := models.MockRelationshipDataCollectionResponse("tags", 3)

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetFlowRelationshipsTags(context.Background(), flowId)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

func (suit *FlowsRelationshipsApiTestSuite) TestGetGetFlowActionRelationshipsFlow() {
	var flowActionId = "flow-action-1"
	mockedRespData := models.MockRelationshipData("flow")

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetFlowActionRelationshipsFlow(context.Background(), flowActionId)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

func (suit *FlowsRelationshipsApiTestSuite) TestGetFlowActionRelationshipsMessages() {
	var flowId = "flow-1"
	mockedRespData := models.MockRelationshipDataCollectionResponse("flow-action-message", 3)

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	filter := common.NewFilterBuilder().Contains("name", "flow-1")
	filterPram := filter.Build()
	res, err := suit.api.GetFlowActionRelationshipsMessages(context.Background(), flowId, &filterPram, nil)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

func (suit *FlowsRelationshipsApiTestSuite) TestGetFlowMessageRelationshipsAction() {
	var flowMessageId = "flow-message-1"
	mockedRespData := models.MockRelationshipData("flow-action")

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetFlowMessageRelationshipsAction(context.Background(), flowMessageId)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

func (suit *FlowsRelationshipsApiTestSuite) TestGetFlowMessageRelationshipsTemplate() {
	var flowMessageId = "flow-message-1"
	mockedRespData := models.MockTemplateResponse()

	err := common.PrepareMockResponse(http.StatusOK, mockedRespData, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	res, err := suit.api.GetFlowMessageRelationshipsTemplate(context.Background(), flowMessageId, nil)

	suit.Nil(err)
	suit.Equal(mockedRespData, *res)
}

func TestFlowsRelationshipsApiTestSuite(t *testing.T) {
	suite.Run(t, new(FlowsRelationshipsApiTestSuite))

}
