package catalog

import (
	"testing"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/options"
	"github.com/stretchr/testify/suite"
)

type CatalogItemTestSuit struct {
	suite.Suite
	api          CatalogApi
	mockedClient *common.MockHTTPClient
}

func (suit *CatalogItemTestSuit) SetupTest() {
	var apiKey = "test-key"

	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())
	suit.mockedClient = common.NewMockHTTPClient()
	suit.api = NewCatalogApi(session, suit.mockedClient)
}

// func (suit *CatalogItemTestSuit) TestGetCatalogItems() {
// 	filterBuilder := common.NewFilterBuilder()
// 	filterBuilder.Any("id", []string{"id1", "id2"})
// 	filterStr := filterBuilder.Build()

// 	res, err := suit.api.GetCatalogItems(context.Background(), filterStr, nil)
// 	suit.Nil(err)
// 	suit.NotNil(res)

// }

func TestCatalogItemTestSuit(t *testing.T) {
	suite.Run(t, new(CatalogItemTestSuit))

}
