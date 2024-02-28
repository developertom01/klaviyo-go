package flows

import (
	"fmt"
	"strings"

	"github.com/developertom01/klaviyo-go/models"
)

type FlowsIncludeField string

const (
	FlowsIncludeFieldFlowActions FlowsIncludeField = "flow-actions"
	FlowsIncludeFieldTags        FlowsIncludeField = "tags"
)

type FlowsActionIncludeField string

const (
	FlowsActionIncludeFieldFlow        FlowsActionIncludeField = "flow"
	FlowsActionIncludeFieldFlowMessage FlowsActionIncludeField = "flow-message"
)

func buildIncludeFieldParam[T FlowsIncludeField | FlowsActionIncludeField | FlowMessageIncludeFieldParam](fields []T) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("include=%s", strings.Join(formattedFields, ","))
}

type FlowPaginationOptions struct {
	PageSize *int    //Default: 50. Min: 1. Max: 50.
	Cursor   *string //For more information please visit
	Sort     *FlowSortField
}

// ---- FlowSortField

type FlowSortField string

const (
	FlowSortFieldCreatedAtASC  FlowSortField = "created_at"
	FlowSortFieldCreatedAtDESC FlowSortField = "-created_at"

	FlowSortFieldIdASC  FlowSortField = "id"
	FlowSortFieldIdDESC FlowSortField = "-id"

	FlowSortFieldNameASC  FlowSortField = "name"
	FlowSortFieldNameDESC FlowSortField = "-name"

	FlowSortFieldStatusASC  FlowSortField = "status"
	FlowSortFieldStatusDESC FlowSortField = "-status"

	FlowSortFieldUpdatedASC  FlowSortField = "updated"
	FlowSortFieldUpdatedDESC FlowSortField = "-updated"
)

func buildFlowSortFieldParam[T FlowSortField | FlowActionSortField | FlowActionMessageSortField](sField *T) string {
	if sField == nil {
		return ""
	}
	return fmt.Sprintf("sort=%s", *sField)
}

// ---- UpdateFlowStatusPayload

type (
	UpdateFlowStatusPayload struct {
		Type string `json:"type"` //flow
		ID   string `json:"id"`   //ID of the Flow to update. Ex: XVTP5Q

	}

	UpdateFlowStatusPayloadAttribute struct {
		Status models.FlowsStatus `json:"status"` //Status you want to update the flow to. ['draft', 'manual', or 'live']
	}
)

//---- FlowMessageIncludeFieldParam

type FlowMessageIncludeFieldParam string

const (
	FlowMessageIncludeFieldParamFlowAction FlowMessageIncludeFieldParam = "flow-action"
	FlowMessageIncludeFieldParamTemplate   FlowMessageIncludeFieldParam = "template"
)

// ---- FlowActionSortField

type FlowActionSortField string

const (
	FlowActionSortFieldActionTypeASC  FlowActionSortField = "action_type"
	FlowActionSortFieldActionTypeDESC FlowActionSortField = "-action_type"

	FlowActionSortFieldCreatedASC  FlowActionSortField = "created"
	FlowActionSortFieldCreatedDESC FlowActionSortField = "-created"

	FlowActionSortFieldCreatedIdASC  FlowActionSortField = "id"
	FlowActionSortFieldCreatedIdDESC FlowActionSortField = "-id"

	FlowActionSortFieldUpdatedASC  FlowActionSortField = "updated"
	FlowActionSortFieldUpdatedDESC FlowActionSortField = "-updated"

	FlowActionSortFieldStatusASC  FlowActionSortField = "status"
	FlowActionSortFieldStatusDESC FlowActionSortField = "-status"
)

// ---- FlowPaginationOptions
type FlowActionPaginationOptions struct {
	PageSize *int    //Default: 50. Min: 1. Max: 50.
	Cursor   *string //For more information please visit
	Sort     *FlowActionSortField
}

func buildGetFlowActionsPaginationOptionsQueryParams(opt *FlowActionPaginationOptions) string {
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

type FlowActionMessageSortField string

const (
	FlowActionMessageSortFieldIdASC  = "id"
	FlowActionMessageSortFieldIdDESC = "-id"

	FlowActionMessageSortFieldNameASC  = "name"
	FlowActionMessageSortFieldNameDESC = "-name"

	FlowActionMessageSortFieldCreatedASC  = "created"
	FlowActionMessageSortFieldCreatedDESC = "-created"

	FlowActionMessageSortFieldUpdatedASC  = "updated"
	FlowActionMessageSortFieldUpdatedDESC = "-updated"
)

// ---- FlowPaginationOptions
type FlowActionMessagePaginationOptions struct {
	PageSize *int    //Default: 50. Min: 1. Max: 50.
	Cursor   *string //For more information please visit
	Sort     *FlowActionMessageSortField
}

func buildFlowActionMessagePaginationOptionsQueryParams(opt *FlowActionMessagePaginationOptions) string {
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
