package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type (
	CampaignMessageResponse struct {
		Data     CampaignMessage                    `json:"data"`
		Included []CampaignMessageIncludedUnionType `json:"included,omitempty"` //This can either be array of Template or Campaign object.
	}

	CampaignMessageCollectionResponse struct {
		Data     []CampaignMessage                  `json:"data"`
		Links    Links                              `json:"links"`
		Included []CampaignMessageIncludedUnionType `json:"included,omitempty"` //This can either be array of Template or Campaign object.
	}

	CampaignMessage struct {
		Type          string                    `json:"type"`
		ID            string                    `json:"id"`
		Attributes    CampaignMessageAttributes `json:"attributes"`
		Links         DataLinks                 `json:"links"`
		Relationships Relationships             `json:"relationships"`
	}

	CampaignMessageAttributes struct {
		Label         string               `json:"label"`
		Channel       string               `json:"channel"`
		Content       MessageContent       `json:"content"`
		SendTimes     []SendTime           `json:"send_times"`
		RenderOptions MessageRenderOptions `json:"render_options"`
		CreatedAt     time.Time            `json:"created_at"`
		UpdatedAt     time.Time            `json:"updated_at"`
	}

	CampaignRelationships struct {
		Campaign Relationships `json:"campaign"`
		Template Relationships `json:"template"`
	}
)

type CampaignMessageField string

const (
	CampaignMessageFieldLabel                           CampaignMessageField = "label"
	CampaignMessageFieldChannel                         CampaignMessageField = "channel"
	CampaignMessageFieldContent                         CampaignMessageField = "content"
	CampaignMessageFieldSendTimes                       CampaignMessageField = "send_times"
	CampaignMessageFieldRenderOptions                   CampaignMessageField = "render_options"
	CampaignMessageFieldRenderOptions_ShortenLinks      CampaignMessageField = "render_options.shorten_links"
	CampaignMessageFieldRenderOptions_AddOrgPrefix      CampaignMessageField = "render_options.add_org_prefix"
	CampaignMessageFieldRenderOptions_AddOrgLink        CampaignMessageField = "render_options.add_org_link"
	CampaignMessageFieldRenderOptions_AddOptOutLanguage CampaignMessageField = "render_options.add_opt_out_language"
	CampaignMessageFieldCreatedAt                       CampaignMessageField = "created_at"
	CampaignMessageFieldUpdatedAt                       CampaignMessageField = "updated_at"
)

// Build query param string. eg. fields[campaign-message]=[name,contact_information]
func BuildCampaignMessageFieldsParam(fields []CampaignMessageField) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[campaign-message]=%s", strings.Join(formattedFields, ","))
}

type CampaignMessageIncludeField string

const (
	CampaignMessageIncludeFieldCampaign CampaignMessageIncludeField = "campaign"
	CampaignMessageIncludeFieldTemplate CampaignMessageIncludeField = "template"
)

func BuildCampaignMessageIncludeFieldParam(fields []CampaignMessageIncludeField) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("includes=%s", strings.Join(formattedFields, ","))

}

type CampaignMessageIncludedUnionType map[string]any

func (cms *CampaignMessageIncludedUnionType) GetCampaign() (*Campaign, bool) {
	data, err := json.Marshal(cms)
	if err != nil {
		return nil, false
	}

	var campaign Campaign
	if err = json.Unmarshal(data, &campaign); err != nil {
		return nil, false
	}

	return &campaign, true
}

func (cms *CampaignMessageIncludedUnionType) GetTemplate() (*Template, bool) {
	data, err := json.Marshal(cms)
	if err != nil {
		return nil, false
	}

	var template Template
	if err = json.Unmarshal(data, &template); err != nil {
		return nil, false
	}

	return &template, true
}
