package accounts

import (
	"github.com/developertom01/klaviyo-go/common"
	"github.com/jaswdr/faker"
)

func mockedAccountResponse(n int) AccountResponse {
	fake := faker.New()

	accounts := make([]Account, 0)

	for i := 0; i < n; i++ {
		accounts = append(accounts, Account{
			Type:  fake.Lorem().Word(),
			ID:    fake.UUID().V4(),
			Links: common.MockedLinkResponse(),
			Attributes: Attributes{
				Industry:          fake.Lorem().Word(),
				Timezone:          fake.Time().Timezone(),
				PublicAPIKey:      fake.UUID().V4(),
				PreferredCurrency: fake.Currency().Currency(),
				ContactInformation: ContactInformation{
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
		})
	}

	return AccountResponse{
		Links: common.MockedLinkResponse(),
		Data:  accounts,
	}
}
