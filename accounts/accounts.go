package accounts

import (
	"context"
	"encoding/json"
	"errors"
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

	ContactInformation struct {
		DefaultSenderName  string        `json:"default_sender_name"`
		DefaultSenderEmail string        `json:"default_sender_email"`
		WebsiteURL         string        `json:"website_url"`
		OrganizationName   string        `json:"organization_name"`
		StreetAddress      StreetAddress `json:"street_address"`
	}

	Attributes struct {
		ContactInformation ContactInformation `json:"contact_information"`
		Industry           string             `json:"industry"`
		Timezone           string             `json:"timezone"`
		PreferredCurrency  string             `json:"preferred_currency"`
		PublicAPIKey       string             `json:"public_api_key"`
	}

	Account struct {
		Type       string       `json:"type"`
		ID         string       `json:"id"`
		Attributes Attributes   `json:"attributes"`
		Links      common.Links `json:"links"`
	}
	AccountResponse struct {
		Data  []Account    `json:"data"`
		Links common.Links `json:"links"`
	}

	AccountsApi interface {
		//Retrieve the account(s) associated with a given private API key.
		GetAccount(ctx context.Context) (*AccountResponse, error)
	}

	accountApi struct {
		session    session.Session
		baseApiUrl string
		revision   string
	}
)

func NewAccountsApi(session session.Session) AccountsApi {

	return &accountApi{
		session:    session,
		baseApiUrl: "https://a.klaviyo.com/api/accounts",
		revision:   "2024-02-15",
	}
}

func (api *accountApi) GetAccount(ctx context.Context) (*AccountResponse, error) {
	var res *http.Response

	reqFn := func() error {
		req, err := http.NewRequestWithContext(ctx, "Get", api.baseApiUrl, nil)
		if err != nil {
			return err
		}

		req.Header.Add("revision", api.revision)
		req.Header.Add("accept", "application/json")

		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		return nil
	}

	err := common.Retry(api.session.GetRetryOptions(), reqFn)
	if err != nil {
		return nil, errors.Join(getAccountApiCallError, err)
	}

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Join(serializationError, err)
	}

	if !exceptions.IsHttpCodeOk(res.StatusCode) {
		var errorRes exceptions.ApiErrorResponse

		err = json.Unmarshal(byteData, &errorRes)
		if err != nil {
			return nil, errors.Join(serializationError, err)
		}

		return nil, exceptions.NewResponseError(errorRes)
	}

	var accountResp AccountResponse
	err = json.Unmarshal(byteData, &accountResp)

	return &accountResp, err
}
