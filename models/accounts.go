package models

import (
	"fmt"
	"strings"
)

type (
	Account struct {
		Type       string            `json:"type"`
		ID         string            `json:"id"`
		Attributes AccountAttributes `json:"attributes"`
		Links      Links             `json:"links"`
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
		Data  []Account `json:"data"`
		Links Links     `json:"links"`
	}

	AccountResponse struct {
		Data Account `json:"data"`
	}
)

type AccountsField string

const (
	AccountsFieldContactInformation                        AccountsField = "contact_information.information"
	AccountsFieldContactInformation_DefaultSenderName      AccountsField = "contact_information.information.default_sender_name"
	AccountsFieldContactInformation_DefaultSenderEmail     AccountsField = "contact_information.information.default_sender_email"
	AccountsFieldContactInformation_DefaultWebsiteUrl      AccountsField = "contact_information.information.default_website_url"
	AccountsFieldContactInformation_OrganizationName       AccountsField = "contact_information.information.organization_name"
	AccountsFieldContactInformation_StreetAddress          AccountsField = "contact_information.information.street_address"
	AccountsFieldContactInformation_StreetAddress_Address1 AccountsField = "contact_information.information.street_address.address1"
	AccountsFieldContactInformation_StreetAddress_Address2 AccountsField = "contact_information.information.street_address.address2"
	AccountsFieldContactInformation_StreetAddress_City     AccountsField = "contact_information.information.street_address.city"
	AccountsFieldContactInformation_StreetAddress_Region   AccountsField = "contact_information.information.street_address.region"
	AccountsFieldContactInformation_StreetAddress_Country  AccountsField = "contact_information.information.street_address.country"
	AccountsFieldContactInformation_StreetAddress_Zip      AccountsField = "contact_information.information.street_address.zip"
	AccountsFieldIndustry                                  AccountsField = "industry"
	AccountsFieldTimezone                                  AccountsField = "timezone"
	AccountsFieldPreferredCurrency                         AccountsField = "preferred_currency"
	AccountsFieldPublicApiKey                              AccountsField = "public_api_key"
)

func BuildAccountFieldsParam(fields []AccountsField) string {

	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[account]=%v", strings.Join(formattedFields, ","))

}
