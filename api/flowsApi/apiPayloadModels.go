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

func buildIncludeFieldParam[T FlowsIncludeField | FlowsActionIncludeField](fields []T) string {
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

func buildFlowSortFieldParam(sField *FlowSortField) string {
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
