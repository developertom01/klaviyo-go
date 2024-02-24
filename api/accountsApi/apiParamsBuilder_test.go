package accounts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountsFieldParamsBuilder(t *testing.T) {
	var expectedString = "fields[account]=[context_information,context_information.default_sender_email,context_information.street_address.address1]"

	builder := NewAccountsFieldParamsBuilder()
	builder.Add(AccountsFieldContactInformation).Add(AccountsFieldContactInformation_DefaultSenderEmail).Add(AccountsFieldContactInformation_StreetAddress_Address1)

	assert.Equal(t, expectedString, builder.Build())
}
