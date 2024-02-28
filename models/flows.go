package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ---- Flow

type (
	FlowCollectionResource struct {
		Data  []Flow `json:"data"`
		Links Links  `json:"links"`
		// Array of type Tag | FlowAction. FlowIncludesUnionType exports isTag and isFlowAction methods
		Included []FlowIncludesUnionType `json:"included,omitempty"`
	}

	FlowResource struct {
		Data     Flow                    `json:"data"`
		Included []FlowIncludesUnionType `json:"included,omitempty"`
	}

	Flow struct {
		Type       string         `json:"type"` //flow
		ID         string         `json:"id"`
		Attributes FlowAttributes `json:"attributes"`
	}

	FlowAttributes struct {
		Name          *string                      `json:"name,omitempty"`
		Status        *FlowsStatus                 `json:"status,omitempty"`
		Archived      *bool                        `json:"archived,omitempty"`
		CreatedAt     *time.Time                   `json:"created_at,omitempty"`
		UpdatedAt     *time.Time                   `json:"updated_at,omitempty"`
		TriggerType   *FlowTriggerType             `json:"trigger_type,omitempty"` //Corresponds to the object which triggered the flow. [`Added to List` `Date Based` `Low Inventory` `Metric` `Price Drop` `Unconfigured`]
		Links         DataLinks                    `json:"links"`
		RelationShips *FlowAttributesRelationShips `json:"relationships,omitempty"`
	}

	FlowTriggerType string //[`Added to List` `Date Based` `Low Inventory` `Metric` `Price Drop` `Unconfigured`]

	// ['draft', 'manual', or 'live']
	FlowsStatus string

	FlowAttributesRelationShips struct {
		FlowAction Relationships `json:"flow-actions"`
		Tags       Relationships `json:"tags"`
	}
)

const (
	FlowsStatusDraft  FlowsStatus = "draft"
	FlowsStatusManual FlowsStatus = "manual"
	FlowsStatusLive   FlowsStatus = "live"
)

const (
	FlowTriggerTypeAddToList    FlowTriggerType = "Added to List"
	FlowTriggerTypeDateBased    FlowTriggerType = "Date Based"
	FlowTriggerTypeLowInventory FlowTriggerType = "Low Inventory"
	FlowTriggerTypeMetric       FlowTriggerType = "Metric"
	FlowTriggerPriceDrop        FlowTriggerType = "Price Drop"
	FlowTriggerUnconfigured     FlowTriggerType = "Unconfigured"
)

// ---- FlowIncludesUnionType

type FlowIncludesUnionType map[string]any

// Return Tag type and True if FlowIncludesUnionType is a Tag
func (fi FlowIncludesUnionType) IsTag() (*Tag, bool) {
	byteData, err := json.Marshal(fi)
	if err != nil {
		return nil, false
	}

	var tag Tag
	err = json.Unmarshal(byteData, &tag)

	if err != nil {
		return nil, false
	}

	return &tag, true
}

func (fi FlowIncludesUnionType) IsFlowAction() (*FlowAction, bool) {
	byteData, err := json.Marshal(fi)
	if err != nil {
		return nil, false
	}

	var flowAction FlowAction
	err = json.Unmarshal(byteData, &flowAction)

	if err != nil {
		return nil, false
	}

	return &flowAction, true
}

// ---- FlowAction

type (
	FlowAction struct {
		Type          string                   `json:"type"` //flow-action
		ID            string                   `json:"id"`
		Attributes    FlowActionAttribute      `json:"attribute"`
		Links         DataLinks                `json:"links"`
		Relationships *FlowActionRelationships `json:"relationships,omitempty"`
		// Array of type Flow | FlowMessage. FlowIncludesUnionType exports IsFlow and IsFlowMessage methods
		Includes []FlowActionIncludesUnionType `json:"includes"`
	}

	FlowActionAttribute struct {
		ActionType      *string               `json:"action_type,omitempty"`
		Status          *string               `json:"status,omitempty"`
		CreatedAt       *time.Time            `json:"created_at,omitempty"`
		UpdatedAt       *time.Time            `json:"updated_at,omitempty"`
		Settings        *FlowActionSettings   `json:"settings,omitempty"`
		TrackingOptions *TrackingOptions      `json:"tracking_options,omitempty"`
		SendOptions     *SendOptions          `json:"send_options,omitempty"`
		RenderOptions   *MessageRenderOptions `json:"render_options,omitempty"`
	}

	FlowActionRelationships struct {
		Flow        Relationships `json:"flow"`
		FlowMessage Relationships `json:"flow-message"`
	}

	FlowActionSettings interface{}
)

// ---- FlowActionIncludesUnionType

type FlowActionIncludesUnionType map[string]any

func (actionIncludes FlowActionIncludesUnionType) IsFlow() (*Flow, bool) {
	byteData, err := json.Marshal(actionIncludes)
	if err != nil {
		return nil, false
	}

	var flow Flow
	err = json.Unmarshal(byteData, &flow)

	if err != nil {
		return nil, false
	}

	return &flow, true
}

func (actionIncludes FlowActionIncludesUnionType) IsFlowMessage() (*FlowMessage, bool) {
	byteData, err := json.Marshal(actionIncludes)
	if err != nil {
		return nil, false
	}

	var message FlowMessage
	err = json.Unmarshal(byteData, &message)

	if err != nil {
		return nil, false
	}

	return &message, true
}

// ---- FlowMessage

type (
	FlowMessage struct {
		Type       string                        `json:"type"` //flow-message
		ID         string                        `json:"id"`
		Attributes FlowMessageAttributes         `json:"attributes"`
		Includes   []FlowMessageIncludeUnionType `json:"includes,omitempty"`
	}

	FlowMessageAttributes struct {
		Name      string         `json:"name"`
		Channel   string         `json:"channel"`
		Content   MessageContent `json:"content"`
		CreatedAt *time.Time     `json:"created_at,omitempty"`
		UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	}
)

// ---- FlowMessageIncludeUnionType

type FlowMessageIncludeUnionType map[string]any

// Returns pointer to FlowAction and true if instance is FlowActon
func (un FlowMessageIncludeUnionType) IsFlowAction() (*FlowAction, bool) {
	byteData, err := json.Marshal(un)
	if err != nil {
		return nil, false
	}

	var action FlowAction
	err = json.Unmarshal(byteData, &action)

	if err != nil {
		return nil, false
	}

	return &action, true
}

// Returns pointer to Template and true if instance is Template
func (un FlowMessageIncludeUnionType) IsTemplate() (*Template, bool) {
	byteData, err := json.Marshal(un)
	if err != nil {
		return nil, false
	}

	var template Template
	err = json.Unmarshal(byteData, &template)

	if err != nil {
		return nil, false
	}

	return &template, true
}

// --- FlowField

type FlowField string

const (
	FlowFieldName        FlowField = "name"
	FlowFieldStatus      FlowField = "status"
	FlowFieldArchived    FlowField = "archived"
	FlowFieldCreatedAt   FlowField = "created_at"
	FlowFieldUpdatedAt   FlowField = "updated_at"
	FlowFieldTriggerType FlowField = "trigger_type"
)

func BuildFlowFieldsParam(fields []FlowField) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[flow]=%s", strings.Join(formattedFields, ","))
}

// ---- FlowActionField

type FlowActionField string

const (
	FlowActionFieldActionType                      FlowActionField = "action_type"
	FlowActionFieldActionStatus                    FlowActionField = "status"
	FlowActionFieldCreated                         FlowActionField = "created"
	FlowActionFieldUpdated                         FlowActionField = "updated"
	FlowActionFieldSettings                        FlowActionField = "settings"
	FlowActionFieldTrackingOption                  FlowActionField = "tracking_option"
	FlowActionFieldSendOption                      FlowActionField = "send_option"
	FlowActionFieldSendOption_UseSmartSending      FlowActionField = "send_option.use_smart_Sending"
	FlowActionFieldSendOption_IsTransactional      FlowActionField = "send_option.is_transactional"
	FlowActionFieldRenderOptions                   FlowActionField = "render_options"
	FlowActionFieldRenderOptions_ShortLinks        FlowActionField = "render_options.short_links"
	FlowActionFieldRenderOptions_AddInfoLink       FlowActionField = "render_options.add_info_link"
	FlowActionFieldRenderOptions_AddOptOutLanguage FlowActionField = "render_options.add_opt_out_language"
)

func BuildFlowActionFieldsParam(fields []FlowActionField) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[flow-action]=%s", strings.Join(formattedFields, ","))
}
