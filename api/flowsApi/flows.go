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

		//Get a flow action from a flow with the given flow action ID.
		GetFlowAction(ctx context.Context, flowId string, opt *GetFlowActionOptions) (*models.FlowActionResource, error)

		//Get the flow message of a flow with the given message ID.
		GetFlowMessage(ctx context.Context, flowMessageID string, opt *GetFlowMessageOptions) (*models.FlowMessageResource, error)

		//Get all flow actions associated with the given flow ID.
		//Returns a maximum of 50 flows per request, which can be paginated with cursor-based pagination.
		GetFlowFlowActions(ctx context.Context, flowId string, opt *GetFlowActionOptions, paginationOpt *FlowActionPaginationOptions) (*models.FlowActionCollectionResource, error)

		//Return all tags associated with the given flow ID.
		GetFlowTags(ctx context.Context, flowId string, tagFields []models.TagField) (*models.FlowTagCollectionResource, error)

		//Get the flow associated with the given action ID.
		GetFlowForFlowAction(ctx context.Context, flowActionId string, flowsFields []models.FlowField) (*models.FlowActionResource, error)

		//Get all flow messages associated with the given action ID.
		GetFlowActionMessages(ctx context.Context, flowActionId string, filterStr *string, paginationOpt *FlowActionMessagePaginationOptions) (*models.FlowActionMessageCollectionResource, error)

		//Get the flow action for a flow message with the given message ID.
		GetFlowActionForMessage(ctx context.Context, actionMessageId string, flowActionFields []models.FlowActionField) (*models.FlowActionResource, error)
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

	if opt.FlowActionFields != nil {
		params = append(params, models.BuildFlowActionFieldsParam(opt.FlowActionFields))
	}

	if opt.FlowFields != nil {
		params = append(params, models.BuildFlowFieldsParam(opt.FlowFields))
	}

	if opt.TagFields != nil {
		params = append(params, models.BuildTagFieldParam(opt.TagFields))

	}

	if opt.Include != nil {
		params = append(params, buildIncludeFieldParam(opt.Include))
	}

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

type GetFlowActionOptions struct {
	FlowActionFields []models.FlowActionField
	FlowFields       []models.FlowField
	FlowMessageField []models.FlowMessageField
	Include          []FlowsActionIncludeField
}

func buildGetFlowActionOptionsQueryParams(opt *GetFlowActionOptions) string {
	if opt == nil {
		return ""
	}

	var params = make([]string, 0)

	if opt.FlowActionFields != nil {
		params = append(params, models.BuildFlowActionFieldsParam(opt.FlowActionFields))
	}

	if opt.FlowFields != nil {
		params = append(params, models.BuildFlowFieldsParam(opt.FlowFields))
	}

	if opt.FlowMessageField != nil {
		params = append(params, models.BuildFlowMessageFieldsParam(opt.FlowMessageField))

	}

	if opt.Include != nil {
		params = append(params, buildIncludeFieldParam(opt.Include))
	}

	return strings.Join(params, "&")
}

func (api flowsApi) GetFlowAction(ctx context.Context, flowId string, opt *GetFlowActionOptions) (*models.FlowActionResource, error) {
	params := buildGetFlowActionOptionsQueryParams(opt)
	url := fmt.Sprintf("%s/api/flows-actions/%s/?%s", api.baseApiUrl, flowId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var flowAction models.FlowActionResource
	err = json.Unmarshal(byteData, &flowAction)

	return &flowAction, err
}

type GetFlowMessageOptions struct {
	FlowActionFields  []models.FlowActionField
	FlowMessageFields []models.FlowMessageField
	TemplateFields    []models.TemplateField
	Include           []FlowMessageIncludeFieldParam
}

func buildGetFlowMessageOptionsQueryParams(opt *GetFlowMessageOptions) string {
	if opt == nil {
		return ""
	}

	var params = make([]string, 0)

	if opt.FlowActionFields != nil {
		params = append(params, models.BuildFlowActionFieldsParam(opt.FlowActionFields))
	}

	if opt.FlowMessageFields != nil {
		params = append(params, models.BuildTemplateFieldParam(opt.TemplateFields))

	}

	if opt.Include != nil {
		params = append(params, buildIncludeFieldParam(opt.Include))
	}

	return strings.Join(params, "&")
}

func (api *flowsApi) GetFlowMessage(ctx context.Context, flowMessageID string, opt *GetFlowMessageOptions) (*models.FlowMessageResource, error) {
	params := buildGetFlowMessageOptionsQueryParams(opt)
	url := fmt.Sprintf("%s/api/flow-messages/%s/?%s", api.baseApiUrl, flowMessageID, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var message models.FlowMessageResource
	err = json.Unmarshal(byteData, &message)

	return &message, err
}

func (api *flowsApi) GetFlowFlowActions(ctx context.Context, flowId string, opt *GetFlowActionOptions, paginationOpt *FlowActionPaginationOptions) (*models.FlowActionCollectionResource, error) {
	params := make([]string, 0)
	optParams := buildGetFlowActionOptionsQueryParams(opt)
	paginationOptParams := buildGetFlowActionsPaginationOptionsQueryParams(paginationOpt)
	if optParams != "" {
		params = append(params, optParams)
	}
	if paginationOptParams != "" {
		params = append(params, paginationOptParams)
	}

	url := fmt.Sprintf("%s/api/flows/%s/flow-actions/?%s", api.baseApiUrl, flowId, strings.Join(params, "&"))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var actions models.FlowActionCollectionResource
	err = json.Unmarshal(byteData, &actions)

	return &actions, err

}

func (api *flowsApi) GetFlowTags(ctx context.Context, flowId string, tagFields []models.TagField) (*models.FlowTagCollectionResource, error) {
	var params = ""

	tagsParam := models.BuildTagFieldParam(tagFields)
	if tagsParam != "" {
		params = fmt.Sprintf("?%s", tagsParam)
	}

	url := fmt.Sprintf("%s/api/flows/%s/%s", api.baseApiUrl, flowId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var tags models.FlowTagCollectionResource
	err = json.Unmarshal(byteData, &tags)

	return &tags, err
}

func (api *flowsApi) GetFlowForFlowAction(ctx context.Context, flowActionId string, flowsFields []models.FlowField) (*models.FlowActionResource, error) {
	var params = ""

	flowsFieldParam := models.BuildFlowFieldsParam(flowsFields)
	if flowsFieldParam != "" {
		params = fmt.Sprintf("?%s", flowsFieldParam)
	}

	url := fmt.Sprintf("%s/api/flow-actions/%s/%s", api.baseApiUrl, flowActionId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var action models.FlowActionResource
	err = json.Unmarshal(byteData, &action)

	return &action, err
}

func (api *flowsApi) GetFlowActionMessages(ctx context.Context, flowActionId string, filterStr *string, paginationOpt *FlowActionMessagePaginationOptions) (*models.FlowActionMessageCollectionResource, error) {
	params := make([]string, 0)
	paginationParams := buildFlowActionMessagePaginationOptionsQueryParams(paginationOpt)
	if filterStr != nil {
		params = append(params, *filterStr)
	}
	if paginationParams != "" {
		params = append(params, paginationParams)
	}

	url := fmt.Sprintf("%s/api/flow-actions/%s/flow-messages/%s", api.baseApiUrl, flowActionId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var messages models.FlowActionMessageCollectionResource
	err = json.Unmarshal(byteData, &messages)

	return &messages, err
}

func (api *flowsApi) GetFlowActionForMessage(ctx context.Context, actionMessageId string, flowActionFields []models.FlowActionField) (*models.FlowActionResource, error) {
	var params = ""

	flowsFieldParam := models.BuildFlowActionFieldsParam(flowActionFields)
	if flowsFieldParam != "" {
		params = fmt.Sprintf("?%s", flowsFieldParam)
	}

	url := fmt.Sprintf("%s/api/flow-messages/%s/%s", api.baseApiUrl, actionMessageId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(getFlowsApiCallError, err)
	}

	var action models.FlowActionResource
	err = json.Unmarshal(byteData, &action)

	return &action, err
}
