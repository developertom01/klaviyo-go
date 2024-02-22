package accounts

type AccountsField string

const (
	ContactInformation                      AccountsField = "context_information"
	ContactInformationDefaultSenderName     AccountsField = "context_information.default_sender_name"
	ContactInformationDefaultSenderEmail    AccountsField = "context_information.default_sender_email"
	ContactInformationDefaultWebsiteUrl     AccountsField = "context_information.default_website_url"
	ContactInformationOrganizationName      AccountsField = "context_information.organization_name"
	ContactInformationStreetAddress         AccountsField = "context_information.street_address"
	ContactInformationStreetAddressAddress1 AccountsField = "context_information.street_address.address1"
	ContactInformationStreetAddressAddress2 AccountsField = "context_information.street_address.address2"
	ContactInformationStreetAddressCity     AccountsField = "context_information.street_address.city"
	ContactInformationStreetAddressRegion   AccountsField = "context_information.street_address.region"
	ContactInformationStreetAddressCountry  AccountsField = "context_information.street_address.country"
	ContactInformationStreetAddressZip      AccountsField = "context_information.street_address.zip"
	Industry                                AccountsField = "industry"
	timezone                                AccountsField = "timezone"
	preferred_currency                      AccountsField = "preferred_currency"
	public_api_key                          AccountsField = "public_api_key"
)

func accountsFieldsToStrings(fields []AccountsField) []string {
	convertedStr := make([]string, 0)

	for _, field := range fields {
		convertedStr = append(convertedStr, string(field))
	}

	return convertedStr
}
