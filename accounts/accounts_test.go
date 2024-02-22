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
	"github.com/developertom01/klaviyo-go/session"
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
	session := session.NewApiKeySession(apiKey, common.NewRetryOptionsWithDefaultValues())
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
	_, err = suit.api.GetAccounts(context.Background())

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
	_, err = suit.api.GetAccounts(context.Background())

	suit.ErrorAs(err, &exceptions.ErrorResponse{}, nil)
}

func (suit *AccountsApiTestSuite) TestGetAccountsOkResponse() {
	//Mock response of API call to return 200 error
	mockedAccount := mockedAccountResponse(1)
	responseByte, err := json.Marshal(mockedAccount)
	if err != nil {
		suit.T().Fatal(err)
	}

	buff := io.NopCloser(bytes.NewBuffer(responseByte))
	response := http.Response{
		Status:     "200 ok",
		StatusCode: http.StatusOK,
		Body:       buff,
	}

	suit.mockedClient.On("Do", mock.Anything).Return(&response, nil)
	accountResp, err := suit.api.GetAccounts(context.Background())

	suit.Nil(err)
	suit.Equal(mockedAccount.Data[0].ID, accountResp.Data[0].ID)
}

func TestAccountsApiTestSuite(t *testing.T) {
	suite.Run(t, new(AccountsApiTestSuite))
}
