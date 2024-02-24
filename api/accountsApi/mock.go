package accounts

import (
	"github.com/developertom01/klaviyo-go/models"
	"github.com/jaswdr/faker"
)

func mockAccountsData() models.Account {
	fake := faker.New()

	return models.Account{
		Type:  fake.Lorem().Word(),
		ID:    fake.UUID().V4(),
		Links: models.MockedLinkResponse(),
		Attributes: models.AccountAttributes{
			Industry:          fake.Lorem().Word(),
			Timezone:          fake.Time().Timezone(),
			PublicAPIKey:      fake.UUID().V4(),
			PreferredCurrency: fake.Currency().Currency(),
			ContactInformation: models.AccountAttributesContactInformation{
				DefaultSenderName:  fake.Person().FirstName(),
				DefaultSenderEmail: fake.Person().Contact().Email,
				WebsiteURL:         fake.Internet().URL(),
				OrganizationName:   fake.Company().Name(),
				StreetAddress: models.StreetAddress{
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

func mockedAccountsCollectionResponse(n int) models.AccountsCollectionResponse {

	accounts := make([]models.Account, 0)

	for i := 0; i < n; i++ {
		accounts = append(accounts, mockAccountsData())
	}

	return models.AccountsCollectionResponse{
		Links: models.MockedLinkResponse(),
		Data:  accounts,
	}
}

func mockedAccountResponse() models.AccountResponse {

	return models.AccountResponse{
		Data: mockAccountsData(),
	}
}
