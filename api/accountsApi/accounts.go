package accounts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type (
	AccountsApi interface {
		//Retrieve the account(s) associated with a given private API key. This will return 1 account object within the array.
		GetAccounts(ctx context.Context, accountFields []models.AccountsField) (*models.AccountsCollectionResponse, error)
		//Retrieve a single account object by its account ID. You can only request the account by which the private API key was generated.
		GetAccount(ctx context.Context, id string, accountFields []models.AccountsField) (*models.AccountResponse, error)
	}

	accountApi struct {
		session    common.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}
)

func NewAccountsApi(session common.Session, httpClient common.HTTPClient) AccountsApi {
	var client common.HTTPClient
	if httpClient == nil {
		client = http.DefaultClient
	} else {
		client = httpClient
	}
	return &accountApi{
		session:    session,
		baseApiUrl: common.BASE_URL,
		revision:   common.API_REVISION,
		httpClient: client,
	}
}

func (api *accountApi) getAccountsInternal(ctx context.Context, accountFields []models.AccountsField) (*models.AccountsCollectionResponse, error) {

	var fieldsParam = models.BuildAccountFieldsParam(accountFields)
	url := fmt.Sprintf("%s/api/accounts/?%s", api.baseApiUrl, fieldsParam)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var accountResp models.AccountsCollectionResponse
	err = json.Unmarshal(byteData, &accountResp)
	if err != nil {
		return nil, errors.Join(serializationError, err)
	}
	return &accountResp, nil
}

func (api *accountApi) GetAccounts(ctx context.Context, accountFields []models.AccountsField) (*models.AccountsCollectionResponse, error) {
	return api.getAccountsInternal(ctx, accountFields)
}

func (api *accountApi) GetAccount(ctx context.Context, id string, accountFields []models.AccountsField) (*models.AccountResponse, error) {

	var fieldsParam = models.BuildAccountFieldsParam(accountFields)
	url := fmt.Sprintf("%s/api/accounts/%s/?%s", api.baseApiUrl, id, fieldsParam)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var accountResp models.AccountResponse
	err = json.Unmarshal(byteData, &accountResp)

	return &accountResp, err
}
