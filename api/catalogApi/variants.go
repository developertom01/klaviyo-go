package catalog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type (
	CatalogVariantApi interface {
		//Get all variants in an account.
		//Variants can be sorted by the following fields, in ascending and descending order: created
		//Currently, the only supported integration type is $custom, and the only supported catalog type is $default.
		//Returns a maximum of 100 variants per request.
		GetCatalogVariants(ctx context.Context, options *CatalogVariantsApiOptions) (*models.CatalogVariantCollectionResource, error)
		//Create a new variant for a related catalog item.
		CreateCatalogVariant(ctx context.Context, payload CreateCatalogItemVariantPayload) (*models.CatalogVariantResource, error)
		//Get a catalog item variant with the given variant ID.
		//The catalog variant ID is a compound ID (string), with format: {integration}:::{catalog}:::{external_id}. Currently, the only supported integration type is $custom, and the only supported catalog is $default.
		GetCatalogVariant(ctx context.Context, catalogVariantId string, options *GetCatalogVariantApiOptions) (*models.CatalogVariantResource, error)
		//Update a catalog item variant with the given variant ID.
		UpdateCatalogVariant(ctx context.Context, catalogVariantId string, payload UpdateCatalogVariantPayload) (*models.CatalogVariantResource, error)
		//Delete a catalog item variant with the given variant ID.
		DeleteCatalogVariant(ctx context.Context, catalogVariantId string) error
	}
)

type CatalogVariantsApiOptions struct {
	CatalogVariantFields []models.CatalogVariantField    //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
	PageCursor           *string                         //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination
	SortField            *models.CatalogVariantSortField //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sorting
	//Allowed field(s)/operator(s):
	//integration.name: equals
	//integration.category: equals
	//For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering
	filterString *string
}

func buildCatalogVariantApiOptionsParams(options *CatalogVariantsApiOptions) string {
	if options == nil {
		return ""
	}

	var params = []string{}

	if options.filterString != nil {
		params = append(params, *options.filterString)
	}

	if options.CatalogVariantFields != nil {
		params = append(params, models.BuildCatalogVariantFieldParams(options.CatalogVariantFields))
	}

	if options.CatalogVariantFields != nil {
		params = append(params, models.BuildCatalogVariantFieldParams((options.CatalogVariantFields)))
	}

	if options.SortField != nil {
		params = append(params, fmt.Sprintf("sort=%s", *options.SortField))
	}

	if options.PageCursor != nil {
		params = append(params, fmt.Sprintf("page[cursor]=%s", *options.PageCursor))
	}

	return strings.Join(params, "&")
}

func (api *catalogApi) GetCatalogVariants(ctx context.Context, options *CatalogVariantsApiOptions) (*models.CatalogVariantCollectionResource, error) {
	queryParams := buildCatalogVariantApiOptionsParams(options)
	url := fmt.Sprintf("%s/api/catalog-variants/?%s", api.baseApiUrl, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var catalogVariantCollection models.CatalogVariantCollectionResource
	err = json.Unmarshal(byteData, &catalogVariantCollection)

	return &catalogVariantCollection, nil
}

func (api *catalogApi) CreateCatalogVariant(ctx context.Context, payload CreateCatalogItemVariantPayload) (*models.CatalogVariantResource, error) {
	url := fmt.Sprintf("%s/api/catalog-variants/", api.baseApiUrl)

	reqData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	reqDataBuffer := bytes.NewBuffer(reqData)

	req, err := http.NewRequest(http.MethodPost, url, reqDataBuffer)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var catalogVariantResource models.CatalogVariantResource
	err = json.Unmarshal(byteData, &catalogVariantResource)

	return &catalogVariantResource, err
}

type GetCatalogVariantApiOptions struct {
	CatalogVariantFields []models.CatalogVariantField //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
}

func buildGetCatalogVariantApiOptionsParams(options *GetCatalogVariantApiOptions) string {
	if options == nil {
		return ""
	}
	var params = []string{}

	if options.CatalogVariantFields != nil {
		params = append(params, models.BuildCatalogVariantFieldParams(options.CatalogVariantFields))
	}

	return strings.Join(params, "&")
}

func (api *catalogApi) GetCatalogVariant(ctx context.Context, catalogVariantId string, options *GetCatalogVariantApiOptions) (*models.CatalogVariantResource, error) {
	queryParams := buildGetCatalogVariantApiOptionsParams(options)
	url := fmt.Sprintf("%s/api/catalog-variants/%s/?%s", api.baseApiUrl, catalogVariantId, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var catalogVariant models.CatalogVariantResource
	err = json.Unmarshal(byteData, &catalogVariant)

	return &catalogVariant, nil
}

func (api *catalogApi) UpdateCatalogVariant(ctx context.Context, catalogVariantId string, payload UpdateCatalogVariantPayload) (*models.CatalogVariantResource, error) {
	url := fmt.Sprintf("%s/api/catalog-variants/%s/", api.baseApiUrl, catalogVariantId)

	reqData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	reqDataBuffer := bytes.NewBuffer(reqData)

	req, err := http.NewRequest(http.MethodPatch, url, reqDataBuffer)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var catalogVariantsResource models.CatalogVariantResource
	err = json.Unmarshal(byteData, &catalogVariantsResource)

	return &catalogVariantsResource, err
}

func (api *catalogApi) DeleteCatalogVariant(ctx context.Context, catalogVariantId string) error {
	url := fmt.Sprintf("%s/api/catalog-variants/%s/", api.baseApiUrl, catalogVariantId)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	_, err = common.RetrieveData(api.httpClient, req, api.session, api.revision)

	return err
}
