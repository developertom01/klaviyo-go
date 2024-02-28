package flows

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type (
	FlowsApi interface {
		//Get all flows in an account.
		//Returns a maximum of 50 flows per request, which can be paginated with cursor-based pagination.
		GetFlows(ctx context.Context, filterStr *string, options *GetFlowsOptions, paginationOpt *FlowPaginationOptions) (*models.FlowCollectionResource, error)

		//Get a flow with the given flow ID.
		GetFlow(ctx context.Context, flowId string, options *GetFlowsOptions) (*models.FlowResource, error)

		//Update the status of a flow with the given flow ID, and all actions in that flow.
		UpdateFlowStatus(ctx context.Context, flowId string, payload UpdateFlowStatusPayload) (*models.FlowResource, error)
	}

	flowsApi struct {
		session    common.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}
)

func NewFlowsApi(session common.Session, httpClient common.HTTPClient) FlowsApi {
	var client common.HTTPClient
	if httpClient == nil {
		client = http.DefaultClient
	} else {
		client = httpClient
	}

	return &flowsApi{
		session:    session,
		baseApiUrl: common.BASE_URL,
		revision:   common.API_REVISION,
		httpClient: client,
	}
}

type GetFlowsOptions struct {
	FlowActionFields []models.FlowActionField
	FlowFields       []models.FlowField
	TagFields        []models.TagField
	Include          []FlowsIncludeField
}

func buildGetFlowsOptionsQueryParams(filter *string, opt *GetFlowsOptions) string {
	if opt == nil {
		return ""
	}

	var params = make([]string, 0)

	if filter != nil {
		params = append(params, *filter)
	}
	params = append(params, models.BuildFlowActionFieldsParam(opt.FlowActionFields))
	params = append(params, models.BuildFlowFieldsParam(opt.FlowFields))
	params = append(params, models.BuildTagFieldParam(opt.TagFields))
	params = append(params, buildFlowsIncludeFieldParam(opt.Include))

	return strings.Join(params, "&")
}

func buildGetFlowsPaginationOptionsQueryParams(opt *FlowPaginationOptions) string {
	if opt == nil {
		return ""
	}
	var params = make([]string, 0)

	if opt.PageSize == nil {
		pageSize := 50
		opt.PageSize = &pageSize
	}

	params = append(params, fmt.Sprintf("page[size]=%d", *opt.PageSize))

	if opt.Sort != nil {
		params = append(params, buildFlowSortFieldParam(opt.Sort))
	}

	if opt.Cursor != nil {
		params = append(params, fmt.Sprintf("page[cursor]=%s", *opt.Cursor))
	}

	return strings.Join(params, "&")
}

func (api *flowsApi) GetFlows(ctx context.Context, filterStr *string, options *GetFlowsOptions, paginationOpt *FlowPaginationOptions) (*models.FlowCollectionResource, error) {
	url := fmt.Sprintf("%s/api/flows/?%s&%s", api.baseApiUrl, buildGetFlowsOptionsQueryParams(filterStr, options), buildGetFlowsPaginationOptionsQueryParams(paginationOpt))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var flowsCollectionResp models.FlowCollectionResource
	err = json.Unmarshal(byteData, &flowsCollectionResp)

	return &flowsCollectionResp, err
}

func (api flowsApi) GetFlow(ctx context.Context, flowId string, options *GetFlowsOptions) (*models.FlowResource, error) {
	url := fmt.Sprintf("%s/api/flows/%s/?%s", api.baseApiUrl, flowId, buildGetFlowsOptionsQueryParams(nil, options))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var flow models.FlowResource
	err = json.Unmarshal(byteData, &flow)

	return &flow, err

}

func (api *flowsApi) UpdateFlowStatus(ctx context.Context, flowId string, payload UpdateFlowStatusPayload) (*models.FlowResource, error) {
	url := fmt.Sprintf("%s/api/flows/%s", api.baseApiUrl, flowId)

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
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var flow models.FlowResource
	err = json.Unmarshal(byteData, &flow)

	return &flow, err
}
