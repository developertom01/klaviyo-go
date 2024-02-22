package accounts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/developertom01/klaviyo-go/session"
)

type (
	StreetAddress struct {
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City     string `json:"city"`
		Region   string `json:"region"`
		Country  string `json:"country"`
		Zip      string `json:"zip"`
	}

	AttributesContactInformation struct {
		DefaultSenderName  string        `json:"default_sender_name"`
		DefaultSenderEmail string        `json:"default_sender_email"`
		WebsiteURL         string        `json:"website_url"`
		OrganizationName   string        `json:"organization_name"`
		StreetAddress      StreetAddress `json:"street_address"`
	}

	Attributes struct {
		ContactInformation AttributesContactInformation `json:"contact_information"`
		Industry           string                       `json:"industry"`
		Timezone           string                       `json:"timezone"`
		PreferredCurrency  string                       `json:"preferred_currency"`
		PublicAPIKey       string                       `json:"public_api_key"`
	}

	Account struct {
		Type       string       `json:"type"`
		ID         string       `json:"id"`
		Attributes Attributes   `json:"attributes"`
		Links      common.Links `json:"links"`
	}

	AccountsCollectionResponse struct {
		Data  []Account    `json:"data"`
		Links common.Links `json:"links"`
	}

	AccountResponse struct {
		Data Account `json:"data"`
	}

	AccountsApi interface {
		GetAccounts(ctx context.Context, fields []AccountsField) (*AccountsCollectionResponse, error)
		GetAccount(ctx context.Context, id string, fields []AccountsField) (*AccountResponse, error)
	}

	accountApi struct {
		session    session.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}
)

func NewAccountsApi(session session.Session, httpClient common.HTTPClient) AccountsApi {
	var client common.HTTPClient
	if httpClient == nil {
		client = http.DefaultClient
	} else {
		client = httpClient
	}
	return &accountApi{
		session:    session,
		baseApiUrl: "https://a.klaviyo.com/api/accounts",
		revision:   "2024-02-15",
		httpClient: client,
	}
}

func (api *accountApi) getAccountsInternal(ctx context.Context, fields []AccountsField) (*AccountsCollectionResponse, error) {
	queryParamMaps := map[string][]string{
		"fields[account]": accountsFieldsToStrings(fields),
	}
	url, err := common.BuildURLWithQueryParams(fmt.Sprintf("%s", api.baseApiUrl), queryParamMaps)
	if err != nil {
		return nil, errors.Join(urlSerializationError, err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := api.retrieveData(req)
	if err != nil {
		return nil, errors.Join(getAccountApiCallError, err)
	}

	var accountResp AccountsCollectionResponse
	err = json.Unmarshal(byteData, &accountResp)
	if err != nil {
		return nil, errors.Join(serializationError, err)
	}
	return &accountResp, nil
}

func (api *accountApi) GetAccounts(ctx context.Context, fields []AccountsField) (*AccountsCollectionResponse, error) {
	return api.getAccountsInternal(ctx, fields)
}

func (api *accountApi) GetAccount(ctx context.Context, id string, fields []AccountsField) (*AccountResponse, error) {
	queryParamMaps := map[string][]string{
		"fields[account]": accountsFieldsToStrings(fields),
	}
	url, err := common.BuildURLWithQueryParams(fmt.Sprintf("%s/%s", api.baseApiUrl, id), queryParamMaps)
	if err != nil {
		return nil, errors.Join(urlSerializationError, err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := api.retrieveData(req)
	if err != nil {
		return nil, errors.Join(getAccountApiCallError, err)
	}

	var accountResp AccountResponse
	err = json.Unmarshal(byteData, &accountResp)
	return &accountResp, err
}

func (api *accountApi) executeRequest(req *http.Request) (*http.Response, error) {
	req.Header.Add("revision", api.revision)
	req.Header.Add("accept", "application/json")
	api.session.ApplyToRequest(api.session.GetOptions(), req)

	return api.httpClient.Do(req)
}

func (api *accountApi) retrieveData(req *http.Request) ([]byte, error) {
	res, err := api.executeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if !exceptions.IsHttpCodeOk(res.StatusCode) {
		var errorRes exceptions.ApiErrorResponse
		err := json.NewDecoder(res.Body).Decode(&errorRes)
		if err != nil {
			return nil, err
		}
		return nil, exceptions.NewResponseError(errorRes)
	}

	return io.ReadAll(res.Body)
}
