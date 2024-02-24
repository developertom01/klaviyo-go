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
		GetAccounts(ctx context.Context, fields []AccountsField) (*models.AccountsCollectionResponse, error)
		GetAccount(ctx context.Context, id string, fields []AccountsField) (*models.AccountResponse, error)
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
		baseApiUrl: fmt.Sprintf("%s/api/campaigns", common.BASE_URL),
		revision:   common.API_REVISION,
		httpClient: client,
	}
}

func (api *accountApi) getAccountsInternal(ctx context.Context, fields []AccountsField) (*models.AccountsCollectionResponse, error) {
	queryParamMaps := map[string][]string{
		"fields[account]": accountsFieldsToStrings(fields),
	}
	url, err := common.BuildURLWithQueryParams(fmt.Sprintf("%s", api.baseApiUrl), queryParamMaps)
	if err != nil {
		return nil, errors.Join(urlSerializationError, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getAccountApiCallError, err)
	}

	var accountResp models.AccountsCollectionResponse
	err = json.Unmarshal(byteData, &accountResp)
	if err != nil {
		return nil, errors.Join(serializationError, err)
	}
	return &accountResp, nil
}

func (api *accountApi) GetAccounts(ctx context.Context, fields []AccountsField) (*models.AccountsCollectionResponse, error) {
	return api.getAccountsInternal(ctx, fields)
}

func (api *accountApi) GetAccount(ctx context.Context, id string, fields []AccountsField) (*models.AccountResponse, error) {
	queryParamMaps := map[string][]string{
		"fields[account]": accountsFieldsToStrings(fields),
	}
	url, err := common.BuildURLWithQueryParams(fmt.Sprintf("%s/%s", api.baseApiUrl, id), queryParamMaps)
	if err != nil {
		return nil, errors.Join(urlSerializationError, err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getAccountApiCallError, err)
	}

	var accountResp models.AccountResponse
	err = json.Unmarshal(byteData, &accountResp)

	return &accountResp, err
}
