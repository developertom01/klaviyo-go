package flows

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
	FlowRelationshipsApi interface {
		//Get all relationships for flow actions associated with the given flow ID.
		GetFlowRelationshipsFlowActions(ctx context.Context, flowId string, filterStr *string, paginationOption *FlowActionPaginationOptions) (*models.RelationshipDataCollection, error)

		//Return the tag IDs of all tags associated with the given flow.
		GetFlowRelationshipsTags(ctx context.Context, flowId string) (*models.RelationshipDataCollection, error)

		//Get the flow associated with the given action ID.
		GetFlowActionRelationshipsFlow(ctx context.Context, flowActionId string) (*models.RelationshipData, error)

		//Get all relationships for flow messages associated with the given flow action ID.
		GetFlowActionRelationshipsMessages(ctx context.Context, flowId string, filterStr *string, paginationOption *FlowActionMessagePaginationOptions) (*models.RelationshipDataCollection, error)

		//Get the relationship for a flow message's flow action, given the flow ID.
		GetFlowMessageRelationshipsAction(ctx context.Context, flowMessageId string) (*models.RelationshipData, error)

		//Returns the ID of the related template
		GetFlowMessageRelationshipsTemplate(ctx context.Context, flowMessageId string, templateFields []models.TemplateField) (*models.TemplateResponse, error)
	}
)

func (api *flowsApi) GetFlowRelationshipsFlowActions(ctx context.Context, flowId string, filterStr *string, paginationOption *FlowActionPaginationOptions) (*models.RelationshipDataCollection, error) {
	params := make([]string, 0)
	if filterStr != nil {
		params = append(params, *filterStr)
	}

	paginationPrams := buildGetFlowActionsPaginationOptionsQueryParams(paginationOption)
	if paginationPrams != "" {
		params = append(params, paginationPrams)
	}

	url := fmt.Sprintf("%s/api/flows/%s/relationships/flow-actions/%s", api.baseApiUrl, flowId, strings.Join(params, "&"))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var relationships models.RelationshipDataCollection
	err = json.Unmarshal(byteData, &relationships)

	return &relationships, err
}

func (api *flowsApi) GetFlowRelationshipsTags(ctx context.Context, flowId string) (*models.RelationshipDataCollection, error) {
	url := fmt.Sprintf("%s/api/flows/%s/relationships/tags/", api.baseApiUrl, flowId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var relationships models.RelationshipDataCollection
	err = json.Unmarshal(byteData, &relationships)

	return &relationships, err
}

func (api *flowsApi) GetFlowActionRelationshipsFlow(ctx context.Context, flowActionId string) (*models.RelationshipData, error) {
	url := fmt.Sprintf("%s/api/flow-actions/%s/relationships/flow/", api.baseApiUrl, flowActionId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var relationships models.RelationshipData
	err = json.Unmarshal(byteData, &relationships)

	return &relationships, err
}

func (api *flowsApi) GetFlowActionRelationshipsMessages(ctx context.Context, flowId string, filterStr *string, paginationOption *FlowActionMessagePaginationOptions) (*models.RelationshipDataCollection, error) {
	params := make([]string, 0)
	if filterStr != nil {
		params = append(params, *filterStr)
	}

	paginationPrams := buildFlowActionMessagePaginationOptionsQueryParams(paginationOption)
	if paginationPrams != "" {
		params = append(params, paginationPrams)
	}

	url := fmt.Sprintf("%s/api/flow-actions/%s/relationships/flow-messages/%s", api.baseApiUrl, flowId, strings.Join(params, "&"))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var relationships models.RelationshipDataCollection
	err = json.Unmarshal(byteData, &relationships)

	return &relationships, err
}

func (api *flowsApi) GetFlowMessageRelationshipsAction(ctx context.Context, flowMessageId string) (*models.RelationshipData, error) {
	url := fmt.Sprintf("%s/api/flow-actions/%s/relationships/flow/", api.baseApiUrl, flowMessageId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var relationships models.RelationshipData
	err = json.Unmarshal(byteData, &relationships)

	return &relationships, err
}

func (api *flowsApi) GetFlowMessageRelationshipsTemplate(ctx context.Context, flowMessageId string, templateFields []models.TemplateField) (*models.TemplateResponse, error) {
	var params = ""

	templateFieldsParam := models.BuildTemplateFieldParam(templateFields)
	if templateFieldsParam != "" {
		params = fmt.Sprintf("?%s", templateFieldsParam)
	}

	url := fmt.Sprintf("%s/api/flow-messages/%s/template/%s", api.baseApiUrl, flowMessageId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var template models.TemplateResponse
	err = json.Unmarshal(byteData, &template)

	return &template, err
}
