package accounts

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/developertom01/klaviyo-go/models"
	"github.com/developertom01/klaviyo-go/options"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AccountsApiTestSuite struct {
	suite.Suite
	api          AccountsApi
	mockedClient *common.MockHTTPClient
}

func (suit *AccountsApiTestSuite) SetupTest() {
	var apiKey = "test-key"
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())
	suit.mockedClient = common.NewMockHTTPClient()
	suit.api = NewAccountsApi(session, suit.mockedClient)
}

func (suit *AccountsApiTestSuite) TestGetAccountsBadRequest() {
	//Mock response of API call to return 400 error
	errBody := common.MockedErrorResponse()
	errByte, err := json.Marshal(errBody)
	if err != nil {
		suit.T().Fatal(err)
	}
	buff := io.NopCloser(bytes.NewBuffer(errByte))
	response := http.Response{
		Status:     "400 Bad Request",
		StatusCode: http.StatusBadRequest,
		Body:       buff,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)
	_, err = suit.api.GetAccounts(context.Background(), nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *AccountsApiTestSuite) TestGetAccountsServerError() {
	//Mock response of API call to return 500 error
	errBody := common.MockedErrorResponse()
	errByte, err := json.Marshal(errBody)
	if err != nil {
		suit.T().Fatal(err)
	}

	buff := io.NopCloser(bytes.NewBuffer(errByte))
	response := http.Response{
		Status:     "500 server error",
		StatusCode: http.StatusInternalServerError,
		Body:       buff,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)
	_, err = suit.api.GetAccounts(context.Background(), nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *AccountsApiTestSuite) TestGetAccountsOkResponse() {
	//Mock response of API call to return 200 error
	mockedAccount := mockedAccountsCollectionResponse(1)
	resp, err := common.PrepareMockResponse(mockedAccount)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "200 ok",
		StatusCode: http.StatusOK,
		Body:       resp,
	}

	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)
	accountResp, err := suit.api.GetAccounts(context.Background(), []models.AccountsField{models.AccountsFieldContactInformation, models.AccountsFieldContactInformation_DefaultSenderName})

	suit.Nil(err)
	suit.Equal(mockedAccount.Data[0].ID, accountResp.Data[0].ID)
}

func (suit *AccountsApiTestSuite) TestGetAccountBadRequest() {
	//Mock response of API call to return 400 error
	var companyId = "12345"
	errBody := common.MockedErrorResponse()
	bodyResp, err := common.PrepareMockResponse(errBody)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "400 Bad Request",
		StatusCode: http.StatusBadRequest,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)
	_, err = suit.api.GetAccount(context.Background(), companyId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *AccountsApiTestSuite) TestGetAccountServerError() {
	//Mock response of API call to return 400 error
	var companyId = "12345"
	errBody := common.MockedErrorResponse()
	bodyResp, err := common.PrepareMockResponse(errBody)

	if err != nil {
		suit.T().Fatal(err)
	}
	response := http.Response{
		Status:     "500 Bad Request",
		StatusCode: http.StatusInternalServerError,
		Body:       bodyResp,
	}
	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)
	_, err = suit.api.GetAccount(context.Background(), companyId, nil)

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *AccountsApiTestSuite) TestGetAccountOkResponse() {
	var companyId = "12345"

	//Mock response of API call to return 200 error
	mockedAccount := mockedAccountResponse()
	bodyResp, err := common.PrepareMockResponse(mockedAccount)
	if err != nil {
		suit.T().Fatal(err)
	}

	response := http.Response{
		Status:     "200 ok",
		StatusCode: http.StatusOK,
		Body:       bodyResp,
	}

	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)
	accountResp, err := suit.api.GetAccount(context.Background(), companyId, []models.AccountsField{models.AccountsFieldContactInformation, models.AccountsFieldContactInformation_DefaultSenderName})

	suit.Nil(err)
	suit.Equal(mockedAccount.Data.ID, accountResp.Data.ID)
}

func TestAccountsApiTestSuite(t *testing.T) {
	suite.Run(t, new(AccountsApiTestSuite))
}
