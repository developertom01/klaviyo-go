package models

import (
	"fmt"
	"strings"
)

type (
	CampaignsCollectionResponse struct {
		Data     []Campaign         `json:"data"`
		Links    Links              `json:"links"`
		Included []CampaignIncluded `json:"included"`
	}
	CampaignResponse struct {
		Data     Campaign           `json:"data"`
		Included []CampaignIncluded `json:"included"`
	}

	Campaign struct {
		Type          string                `json:"type"` //campaign
		ID            string                `json:"id"`   //The campaign ID
		Attributes    CampaignAttributes    `json:"attributes"`
		Links         DataLinks             `json:"links"`
		Relationships *CampaignRelationship `json:"relationships,omitempty"`
	}
	CampaignRelationship struct {
		CampaignMessage *Relationships `json:"campaign-messages,omitempty"`
		Tags            *Relationships `json:"tags,omitempty"`
	}

	CampaignAttributes struct {
		Name            string          `json:"name"`             //The campaign name
		Status          string          `json:"status"`           //The current status of the campaign
		Archived        bool            `json:"archived"`         //Whether the campaign has been archived or not
		Audiences       Audiences       `json:"audiences"`        //The audiences to be included and/or excluded from the campaign
		SendOptions     SendOptions     `json:"send_options"`     //Options to use when sending a campaign
		TrackingOptions TrackingOptions `json:"tracking_options"` //The tracking options associated with the campaign
		SendStrategy    SendStrategy    `json:"send_strategy"`    //The send strategy the campaign will send with
		CreatedAt       string          `json:"created_at"`       //The datetime when the campaign was created
		ScheduledAt     string          `json:"scheduled_at"`     // The datetime when the campaign was scheduled for future sending
		UpdatedAt       string          `json:"updated_at"`       //The datetime when the campaign was last updated by a user or the system
		SendTime        string          `json:"send_time"`        //The datetime when the campaign will be / was sent or None if not yet scheduled by a send_job.
	}

	CampaignIncluded struct {
		Type       string                     `json:"type"`
		ID         string                     `json:"id"`
		Attributes CampaignIncludedAttributes `json:"attributes"`
		Links      DataLinks                  `json:"links"`
	}

	CampaignIncludedAttributes struct {
		Label         *string               `json:"label,omitempty"`
		Channel       *string               `json:"channel,omitempty"`
		Content       *MessageContent       `json:"content,omitempty"` // //Additional attributes relating to the content of the message
		SendTimes     []SendTime            `json:"send_times,omitempty"`
		RenderOptions *MessageRenderOptions `json:"render_options,omitempty"` //Additional options for rendering the message
		CreatedAt     *string               `json:"created_at,omitempty"`
		UpdatedAt     *string               `json:"updated_at,omitempty"`
		Name          *string               `json:"name,omitempty"`
	}
)

type (
	CampaignRecipientCountResponse struct {
		Data CampaignRecipientCountData `json:"data"`
	}
	CampaignRecipientCountData struct {
		Type       string                           `json:"type"`       //campaign-recipient-estimation
		ID         string                           `json:"id"`         //The ID of the campaign for which to get the estimated number of recipients
		Attributes CampaignRecipientCountAttributes `json:"attributes"` //The estimated number of unique recipients the campaign will send to
		Links      DataLinks                        `json:"links"`
	}

	CampaignRecipientCountAttributes struct {
		EstimatedRecipientCount int64 `json:"estimated_recipient_count"` //The estimated number of unique recipients the campaign will send to
	}
)

type (
	CampaignsField string

	CampaignSortField string

	CampaignRecipientEstimationField string

	CampaignIncludeField string
)

const (
	CampaignsFieldName                                                    CampaignsField = "name"
	CampaignsFieldStatus                                                  CampaignsField = "status"
	CampaignsFieldArchived                                                CampaignsField = "archived"
	CampaignsFieldAudience                                                CampaignsField = "audience"
	CampaignsFieldAudience_Include                                        CampaignsField = "audience.include"
	CampaignsFieldAudience_Exclude                                        CampaignsField = "audience.exclude"
	CampaignsFieldSendOptions                                             CampaignsField = "send_options"
	CampaignsFieldTrackingOptions                                         CampaignsField = "tracking_options"
	CampaignsFieldSendStrategy                                            CampaignsField = "send_strategy"
	CampaignsFieldSendStrategy_Method                                     CampaignsField = "send_strategy.method"
	CampaignsFieldSendStrategy_OptionStatic                               CampaignsField = "send_strategy.option_static"
	CampaignsFieldSendStrategy_OptionStatic_DateTime                      CampaignsField = "send_strategy.option_static.datetime"
	CampaignsFieldSendStrategy_OptionStatic_SendPastRecipientsImmediately CampaignsField = "send_strategy.option_static.send_past_recipients_immediately"
	CampaignsFieldSendStrategy_OptionsThrottled                           CampaignsField = "send_strategy.options_throttled"
	CampaignsFieldSendStrategy_OptionsThrottled_ThrottledPercentage       CampaignsField = "send_strategy.options_throttled.throttle_percentage"
	CampaignsFieldSendStrategy_OptionsSto                                 CampaignsField = "send_strategy.options_sto"
	CampaignsFieldSendStrategy_OptionsSto_Date                            CampaignsField = "send_strategy.options_sto.date"
	CampaignsFieldSCreatedAt                                              CampaignsField = "created_at"
	CampaignsFieldSUpdatedAt                                              CampaignsField = "updated_at"
	CampaignsFieldSendTime                                                CampaignsField = "send_time"
)

const (
	CampaignSortFieldNameAsc  CampaignSortField = "name"
	CampaignSortFieldNameDesc CampaignSortField = "-name"

	CampaignSortFieldIdAsc  CampaignSortField = "id"
	CampaignSortFieldIdDesc CampaignSortField = "-id"

	CampaignSortFieldScheduledAtAsc  CampaignSortField = "scheduled_at"
	CampaignSortFieldScheduledAtDesc CampaignSortField = "-scheduled_at"

	CampaignSortFieldCreatedAtAsc  CampaignSortField = "created_at"
	CampaignSortFieldCreatedAtDesc CampaignSortField = "-created_at"

	CampaignSortFieldUpdatedAtAsc  CampaignSortField = "updated_at"
	CampaignSortFieldUpdatedAtDesc CampaignSortField = "-updated_at"
)

const (
	CampaignIncludeFieldCampaignMessage CampaignIncludeField = "campaign-message"
	CampaignIncludeFieldTags            CampaignIncludeField = "tags"
)

const (
	CampaignRecipientEstimationFieldEstimatedRecipientCount CampaignRecipientEstimationField = "estimated_recipient_count"
)

func BuildCampaignFieldsParam(fields []CampaignsField) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[campaign]=%s", strings.Join(formattedFields, ","))
}

// Build query param string. eg. fields[campaign-recipient-estimation]=[name,contact_information]
func BuildCampaignRecipientEstimateFieldParam(fields []CampaignRecipientEstimationField) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[campaign-recipient-estimation]=%s", strings.Join(formattedFields, ","))
}

func BuildCampaignIncludeFieldParam(fields []CampaignIncludeField) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("include=%s", strings.Join(formattedFields, ","))
}
