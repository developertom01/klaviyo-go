package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type (
	CatalogItemApi interface {
		//Get all catalog items in an account.
		//Catalog items can be sorted by the following fields, in ascending and descending order: created
		//Currently, the only supported integration type is $custom, and the only supported catalog type is $default
		GetCatalogItems(ctx context.Context, filterString string, options *CatalogItemApiOptions) (*models.CatalogItemCollectionResource, error)
	}
)

type CatalogItemApiOptions struct {
	CatalogItemFields    []models.CatalogItemField         //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
	CatalogVariantFields []models.CatalogVariantField      //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
	PageCursor           *string                           //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination
	SortField            *models.CatalogItemSortField      //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting
	Include              []models.CatalogItemIncludedField //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships
}

func buildCatalogItemApiOptionsParams(filterString string, options *CatalogItemApiOptions) string {
	if options == nil {
		return filterString
	}
	var params = []string{filterString}

	if options.CatalogItemFields != nil {
		params = append(params, models.BuildCatalogItemFieldParams(options.CatalogItemFields))
	}

	if options.CatalogVariantFields != nil {
		params = append(params, models.BuildCatalogVariantFieldParams((options.CatalogVariantFields)))
	}

	if options.PageCursor != nil {
		var includedStr = make([]string, 0)

		for _, inc := range options.Include {
			includedStr = append(includedStr, string(inc))
		}
		params = append(params, fmt.Sprintf("include=%s", strings.Join(includedStr, ",")))
	}

	return strings.Join(params, "&")

}

func (api catalogApi) GetCatalogItems(ctx context.Context, filterString string, options *CatalogItemApiOptions) (*models.CatalogItemCollectionResource, error) {
	queryParams := buildCatalogItemApiOptionsParams(filterString, options)
	url := fmt.Sprintf("%s/api/catalog-items/?%s", api.baseApiUrl, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var catalogItemsCollection models.CatalogItemCollectionResource
	err = json.Unmarshal(byteData, &catalogItemsCollection)
	return &catalogItemsCollection, nil
}
