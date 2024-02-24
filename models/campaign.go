package models

import (
	"github.com/developertom01/klaviyo-go/common"
)

type (
	CampaignsCollectionResponse struct {
		Data     []Campaign         `json:"data"`
		Links    common.Links       `json:"links"`
		Included []CampaignIncluded `json:"included"`
	}
	CampaignsResponse struct {
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
		RenderOptions *MessageRenderOptions `json:"render_options,omitempty"` ////Additional options for rendering the message
		CreatedAt     *string               `json:"created_at,omitempty"`
		UpdatedAt     *string               `json:"updated_at,omitempty"`
		Name          *string               `json:"name,omitempty"`
	}
)
