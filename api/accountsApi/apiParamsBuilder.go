package accounts

import (
	"fmt"
	"slices"
	"strings"
)

type AccountsField string

const (
	AccountsFieldContactInformation                        AccountsField = "context_information"
	AccountsFieldContactInformation_DefaultSenderName      AccountsField = "context_information.default_sender_name"
	AccountsFieldContactInformation_DefaultSenderEmail     AccountsField = "context_information.default_sender_email"
	AccountsFieldContactInformation_DefaultWebsiteUrl      AccountsField = "context_information.default_website_url"
	AccountsFieldContactInformation_OrganizationName       AccountsField = "context_information.organization_name"
	AccountsFieldContactInformation_StreetAddress          AccountsField = "context_information.street_address"
	AccountsFieldContactInformation_StreetAddress_Address1 AccountsField = "context_information.street_address.address1"
	AccountsFieldContactInformation_StreetAddress_Address2 AccountsField = "context_information.street_address.address2"
	AccountsFieldContactInformation_StreetAddress_City     AccountsField = "context_information.street_address.city"
	AccountsFieldContactInformation_StreetAddress_Region   AccountsField = "context_information.street_address.region"
	AccountsFieldContactInformation_StreetAddress_Country  AccountsField = "context_information.street_address.country"
	AccountsFieldContactInformation_StreetAddress_Zip      AccountsField = "context_information.street_address.zip"
	AccountsFieldIndustry                                  AccountsField = "industry"
	AccountsFieldTimezone                                  AccountsField = "timezone"
	AccountsFieldPreferredCurrency                         AccountsField = "preferred_currency"
	AccountsFieldPublicApiKey                              AccountsField = "public_api_key"
)

func accountsFieldsToStrings(fields []AccountsField) []string {
	convertedStr := make([]string, 0)

	for _, field := range fields {
		convertedStr = append(convertedStr, string(field))
	}

	return convertedStr
}

type AccountsFieldParamsBuilder struct {
	params []AccountsField
}

func NewAccountsFieldParamsBuilder() *AccountsFieldParamsBuilder {
	return &AccountsFieldParamsBuilder{}
}

func (builder *AccountsFieldParamsBuilder) Add(field AccountsField) *AccountsFieldParamsBuilder {

	if !slices.Contains(builder.params, field) {
		builder.params = append(builder.params, field)
	}

	return builder
}

// Build query param string. eg. fields[account]=[name,contact_information]
func (builder *AccountsFieldParamsBuilder) Build() string {

	if len(builder.params) == 0 {
		return ""
	}

	return strings.ReplaceAll(fmt.Sprintf("fields[account]=%v", builder.params), " ", ",")

}
