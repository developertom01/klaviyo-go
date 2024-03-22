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
	CatalogItemApi interface {
		//Get all catalog items in an account.
		//Catalog items can be sorted by the following fields, in ascending and descending order: created
		//Currently, the only supported integration type is $custom, and the only supported catalog type is $default
		GetCatalogItems(ctx context.Context, filterString string, options *CatalogItemApiOptions) (*models.CatalogItemCollectionResource, error)
		//Create a new catalog item.
		CreateCatalogItem(ctx context.Context, payload CreateCatalogItemPayload) (*models.CatalogItemResource, error)
		//Get a specific catalog item with the given item ID.
		//CatalogItemId: The catalog item ID is a compound ID (string), with format: {integration}:::{catalog}:::{external_id}. Currently, the only supported integration type is $custom, and the only supported catalog is $default.
		GetCatalogItem(ctx context.Context, catalogItemId string, options *GetCatalogItemApiOptions) (*models.CatalogItemResource, error)
		//Update a catalog item with the given item ID.
		UpdateCatalogItem(ctx context.Context, catalogItemId string, payload UpdateCatalogItemPayload) (*models.CatalogItemResource, error)
		//Delete a catalog item with the given item ID.
		//The catalog item ID is a compound ID (string), with format: {integration}:::{catalog}:::{external_id}. Currently, the only supported integration type is $custom, and the only supported catalog is $default.
		DeleteCatalogItem(ctx context.Context, catalogItemId string) error
		//Get all catalog item bulk create jobs.
		//Returns a maximum of 100 jobs per request.
		GetCreateItemsJobs(ctx context.Context, options *GetBulkItemsJobsOptions) (*models.CatalogItemBulkJobCollectionResource, error)
		//Create a catalog item bulk create job to create a batch of catalog items
		//Accepts up to 100 catalog items per request. The maximum allowed payload size is 5MB.
		//The maximum number of jobs in progress at one time is 500.
		SpawnCreateItemsJob(ctx context.Context, payload SpawnCreateItemsJobPayload) (*models.CatalogItemBulkJobResource, error)
		//Get a catalog item bulk create job with the given job ID.
		//An include parameter can be provided to get the following related resource data: items.
		GetCreateItemsJob(ctx context.Context, buildCreateJobId string, options *GetBulkItemsJobOptions) (*models.CatalogItemBulkJobResource, error)
		//Get all catalog item bulk update jobs.
		GetUpdateItemsJobs(ctx context.Context, options *GetBulkItemsJobsOptions) (*models.CatalogItemBulkJobCollectionResource, error)
		//Create a catalog item bulk update job to update a batch of catalog items.
		//Accepts up to 100 catalog items per request. The maximum allowed payload size is 5MB.
		//The maximum number of jobs in progress at one time is 500.
		SpawnUpdateItemsJob(ctx context.Context, payload SpawnUpdateItemsJobPayload) (*models.CatalogItemBulkJobResource, error)
		//Get a catalog item bulk update job with the given job ID.
		//An include parameter can be provided to get the following related resource data: items.
		GetUpdateItemsJob(ctx context.Context, buildUpdateJobId string, options *GetBulkItemsJobOptions) (*models.CatalogItemBulkJobResource, error)
		//Get all catalog item bulk delete jobs.
		//Returns a maximum of 100 jobs per request.
		GetDeleteItemsJobs(ctx context.Context, options *GetBulkItemsJobsOptions) (*models.CatalogItemBulkJobCollectionResource, error)
		//Create a catalog item bulk delete job to delete a batch of catalog items.
		//Accepts up to 100 catalog items per request. The maximum allowed payload size is 5MB.
		//The maximum number of jobs in progress at one time is 500.
		SpawnDeleteItemsJob(ctx context.Context, payload SpawnDeleteItemsJobPayload) (*models.CatalogItemBulkJobResource, error)
		//Get a catalog item bulk Delete job with the given job ID.
		//An include parameter can be provided to get the following related resource data: items.
		GetDeleteItemsJob(ctx context.Context, buildDeleteJobId string, options *GetBulkItemsJobOptions) (*models.CatalogItemBulkJobResource, error)
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

	if options.Include != nil {
		var includedStr = make([]string, 0)

		for _, inc := range options.Include {
			includedStr = append(includedStr, string(inc))
		}
		params = append(params, fmt.Sprintf("include=%s", strings.Join(includedStr, ",")))
	}

	if options.SortField != nil {
		params = append(params, fmt.Sprintf("sort=%s", *options.SortField))
	}

	if options.PageCursor != nil {
		params = append(params, fmt.Sprintf("page[cursor]=%s", *options.PageCursor))

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

func (api *catalogApi) CreateCatalogItem(ctx context.Context, payload CreateCatalogItemPayload) (*models.CatalogItemResource, error) {
	url := fmt.Sprintf("%s/api/catalog-items/", api.baseApiUrl)

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

	var catalogItemResource models.CatalogItemResource
	err = json.Unmarshal(byteData, &catalogItemResource)

	return &catalogItemResource, err
}

type GetCatalogItemApiOptions struct {
	CatalogItemFields    []models.CatalogItemField         //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
	CatalogVariantFields []models.CatalogVariantField      //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
	Include              []models.CatalogItemIncludedField //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships
}

func buildGetCatalogItemApiOptionsParams(options *GetCatalogItemApiOptions) string {
	if options == nil {
		return ""
	}
	var params = []string{}

	if options.CatalogItemFields != nil {
		params = append(params, models.BuildCatalogItemFieldParams(options.CatalogItemFields))
	}

	if options.CatalogVariantFields != nil {
		params = append(params, models.BuildCatalogVariantFieldParams((options.CatalogVariantFields)))
	}

	if options.Include != nil {
		var includedStr = make([]string, 0)

		for _, inc := range options.Include {
			includedStr = append(includedStr, string(inc))
		}
		params = append(params, fmt.Sprintf("include=%s", strings.Join(includedStr, ",")))
	}

	return strings.Join(params, "&")
}

func (api *catalogApi) GetCatalogItem(ctx context.Context, catalogItemId string, options *GetCatalogItemApiOptions) (*models.CatalogItemResource, error) {
	queryParams := buildGetCatalogItemApiOptionsParams(options)
	url := fmt.Sprintf("%s/api/catalog-items/%s/?%s", api.baseApiUrl, catalogItemId, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var catalogItem models.CatalogItemResource
	err = json.Unmarshal(byteData, &catalogItem)

	return &catalogItem, nil
}

func (api catalogApi) UpdateCatalogItem(ctx context.Context, catalogItemId string, payload UpdateCatalogItemPayload) (*models.CatalogItemResource, error) {
	url := fmt.Sprintf("%s/api/catalog-items/%s/", api.baseApiUrl, catalogItemId)

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

	var catalogItemResource models.CatalogItemResource
	err = json.Unmarshal(byteData, &catalogItemResource)

	return &catalogItemResource, err
}

func (api *catalogApi) DeleteCatalogItem(ctx context.Context, catalogItemId string) error {
	url := fmt.Sprintf("%s/api/catalog-items/%s/", api.baseApiUrl, catalogItemId)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	_, err = common.RetrieveData(api.httpClient, req, api.session, api.revision)

	return err
}

type GetBulkItemsJobsOptions struct {
	Filter         *string                          //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#filtering Allowed field(s)/operator(s):status: equals
	PageCursor     *string                          //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination
	ItemJobsFields []models.CatalogItemBulkJobField //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
}

// func
func buildGetBulkItemsJobsOptionsParams(options *GetBulkItemsJobsOptions) string {
	if options == nil {
		return ""
	}
	var params = []string{}

	if options.Filter != nil {
		params = append(params, *options.Filter)
	}

	if options.PageCursor != nil {
		params = append(params, fmt.Sprintf("page[cursor]=%s", *options.PageCursor))
	}

	if options.ItemJobsFields != nil {
		params = append(params, models.BuildCatalogItemBulkJobFieldParams(options.ItemJobsFields))
	}

	return strings.Join(params, "&")
}

func (api catalogApi) GetCreateItemsJobs(ctx context.Context, options *GetBulkItemsJobsOptions) (*models.CatalogItemBulkJobCollectionResource, error) {
	queryParams := buildGetBulkItemsJobsOptionsParams(options)
	url := fmt.Sprintf("%s/api/catalog-item-bulk-create-jobs/?%s", api.baseApiUrl, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var itemJobs models.CatalogItemBulkJobCollectionResource
	err = json.Unmarshal(byteData, &itemJobs)

	return &itemJobs, nil
}

func (api *catalogApi) SpawnCreateItemsJob(ctx context.Context, payload SpawnCreateItemsJobPayload) (*models.CatalogItemBulkJobResource, error) {
	url := fmt.Sprintf("%s/api/catalog-item-bulk-create-jobs/", api.baseApiUrl)

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

	var catalogItemResource models.CatalogItemBulkJobResource
	err = json.Unmarshal(byteData, &catalogItemResource)

	return &catalogItemResource, err
}

type GetBulkItemsJobOptions struct {
	ItemJobsFields     []models.CatalogItemBulkJobField        //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
	CatalogItemsFields []models.CatalogItemField               //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#sparse-fieldsets
	Include            []models.CatalogItemBulkJobIncludeField //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#relationships

}

func buildGetBulkItemsJobOptionsParams(options *GetBulkItemsJobOptions) string {
	if options == nil {
		return ""
	}
	var params = []string{}

	if options.ItemJobsFields != nil {
		params = append(params, models.BuildCatalogItemBulkJobFieldParams(options.ItemJobsFields))
	}

	if options.CatalogItemsFields != nil {
		params = append(params, models.BuildCatalogItemFieldParams((options.CatalogItemsFields)))
	}

	if options.Include != nil {
		var includedStr = make([]string, 0)

		for _, inc := range options.Include {
			includedStr = append(includedStr, string(inc))
		}
		params = append(params, fmt.Sprintf("include=%s", strings.Join(includedStr, ",")))
	}

	return strings.Join(params, "&")
}

func (api *catalogApi) GetCreateItemsJob(ctx context.Context, createBuildItemJobId string, options *GetBulkItemsJobOptions) (*models.CatalogItemBulkJobResource, error) {
	queryParams := buildGetBulkItemsJobOptionsParams(options)
	url := fmt.Sprintf("%s/api/catalog-item-bulk-create-jobs/%s/?%s", api.baseApiUrl, createBuildItemJobId, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var createBuildItemJob models.CatalogItemBulkJobResource
	err = json.Unmarshal(byteData, &createBuildItemJob)

	return &createBuildItemJob, nil
}

func (api *catalogApi) GetUpdateItemsJobs(ctx context.Context, options *GetBulkItemsJobsOptions) (*models.CatalogItemBulkJobCollectionResource, error) {
	queryParams := buildGetBulkItemsJobsOptionsParams(options)
	url := fmt.Sprintf("%s/api/catalog-item-bulk-update-jobs/?%s", api.baseApiUrl, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var itemJobs models.CatalogItemBulkJobCollectionResource
	err = json.Unmarshal(byteData, &itemJobs)

	return &itemJobs, nil
}

func (api *catalogApi) SpawnUpdateItemsJob(ctx context.Context, payload SpawnUpdateItemsJobPayload) (*models.CatalogItemBulkJobResource, error) {
	url := fmt.Sprintf("%s/api/catalog-item-bulk-update-jobs/", api.baseApiUrl)

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

	var catalogItemResource models.CatalogItemBulkJobResource
	err = json.Unmarshal(byteData, &catalogItemResource)

	return &catalogItemResource, err
}

func (api *catalogApi) GetUpdateItemsJob(ctx context.Context, buildUpdateJobId string, options *GetBulkItemsJobOptions) (*models.CatalogItemBulkJobResource, error) {
	queryParams := buildGetBulkItemsJobOptionsParams(options)
	url := fmt.Sprintf("%s/api/catalog-item-bulk-update-jobs/%s/?%s", api.baseApiUrl, buildUpdateJobId, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var createBuildItemJob models.CatalogItemBulkJobResource
	err = json.Unmarshal(byteData, &createBuildItemJob)

	return &createBuildItemJob, nil
}

func (api *catalogApi) GetDeleteItemsJobs(ctx context.Context, options *GetBulkItemsJobsOptions) (*models.CatalogItemBulkJobCollectionResource, error) {
	queryParams := buildGetBulkItemsJobsOptionsParams(options)
	url := fmt.Sprintf("%s/api/catalog-item-bulk-delete-jobs/?%s", api.baseApiUrl, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var itemJobs models.CatalogItemBulkJobCollectionResource
	err = json.Unmarshal(byteData, &itemJobs)

	return &itemJobs, nil
}

func (api *catalogApi) SpawnDeleteItemsJob(ctx context.Context, payload SpawnDeleteItemsJobPayload) (*models.CatalogItemBulkJobResource, error) {
	url := fmt.Sprintf("%s/api/catalog-item-bulk-delete-jobs/", api.baseApiUrl)

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

	var catalogItemResource models.CatalogItemBulkJobResource
	err = json.Unmarshal(byteData, &catalogItemResource)

	return &catalogItemResource, err
}

func (api *catalogApi) GetDeleteItemsJob(ctx context.Context, buildDeleteJobId string, options *GetBulkItemsJobOptions) (*models.CatalogItemBulkJobResource, error) {
	queryParams := buildGetBulkItemsJobOptionsParams(options)
	url := fmt.Sprintf("%s/api/catalog-item-bulk-delete-jobs/%s/?%s", api.baseApiUrl, buildDeleteJobId, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var createBuildItemJob models.CatalogItemBulkJobResource
	err = json.Unmarshal(byteData, &createBuildItemJob)

	return &createBuildItemJob, nil
}
