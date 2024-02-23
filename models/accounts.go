package models

import "github.com/developertom01/klaviyo-go/common"

type (
	Account struct {
		Type       string            `json:"type"`
		ID         string            `json:"id"`
		Attributes AccountAttributes `json:"attributes"`
		Links      common.Links      `json:"links"`
	}

	AccountAttributes struct {
		ContactInformation AccountAttributesContactInformation `json:"contact_information"`
		Industry           string                              `json:"industry"`
		Timezone           string                              `json:"timezone"`
		PreferredCurrency  string                              `json:"preferred_currency"`
		PublicAPIKey       string                              `json:"public_api_key"`
	}

	AccountAttributesContactInformation struct {
		DefaultSenderName  string        `json:"default_sender_name"`
		DefaultSenderEmail string        `json:"default_sender_email"`
		WebsiteURL         string        `json:"website_url"`
		OrganizationName   string        `json:"organization_name"`
		StreetAddress      StreetAddress `json:"street_address"`
	}

	AccountsCollectionResponse struct {
		Data  []Account    `json:"data"`
		Links common.Links `json:"links"`
	}

	AccountResponse struct {
		Data Account `json:"data"`
	}
)
