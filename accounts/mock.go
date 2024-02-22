package accounts

import (
	"github.com/developertom01/klaviyo-go/common"
	"github.com/jaswdr/faker"
)

func mockAccountsData() Account {
	fake := faker.New()

	return Account{
		Type:  fake.Lorem().Word(),
		ID:    fake.UUID().V4(),
		Links: common.MockedLinkResponse(),
		Attributes: Attributes{
			Industry:          fake.Lorem().Word(),
			Timezone:          fake.Time().Timezone(),
			PublicAPIKey:      fake.UUID().V4(),
			PreferredCurrency: fake.Currency().Currency(),
			ContactInformation: AttributesContactInformation{
				DefaultSenderName:  fake.Person().FirstName(),
				DefaultSenderEmail: fake.Person().Contact().Email,
				WebsiteURL:         fake.Internet().URL(),
				OrganizationName:   fake.Company().Name(),
				StreetAddress: StreetAddress{
					Address1: fake.Address().StreetAddress(),
					Address2: fake.Address().StreetAddress(),
					City:     fake.Address().City(),
					Region:   fake.Address().State(),
					Country:  fake.Address().Country(),
					Zip:      fake.Address().PostCode(),
				},
			},
		},
	}
}

func mockedAccountsCollectionResponse(n int) AccountsCollectionResponse {

	accounts := make([]Account, 0)

	for i := 0; i < n; i++ {
		accounts = append(accounts, mockAccountsData())
	}

	return AccountsCollectionResponse{
		Links: common.MockedLinkResponse(),
		Data:  accounts,
	}
}

func mockedAccountResponse() AccountResponse {

	return AccountResponse{
		Data: mockAccountsData(),
	}
}
