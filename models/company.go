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
		string        `json:"type"`
		ID            string             `json:"id"`
		Attributes    CampaignAttributes `json:"attributes"`
		Links         DataLinks          `json:"links"`
		Relationships Relationships      `json:"relationships"`
	}

	CampaignAttributes struct {
		Name            string          `json:"name"`
		Status          string          `json:"status"`
		Archived        bool            `json:"archived"`
		Audiences       Audiences       `json:"audiences"`
		SendOptions     SendOptions     `json:"send_options"`
		TrackingOptions TrackingOptions `json:"tracking_options"`
		SendStrategy    SendStrategy    `json:"send_strategy"`
		CreatedAt       string          `json:"created_at"`
		ScheduledAt     string          `json:"scheduled_at"`
		UpdatedAt       string          `json:"updated_at"`
		SendTime        string          `json:"send_time"`
	}

	CampaignIncluded struct {
		Type       string                     `json:"type"`
		ID         string                     `json:"id"`
		Attributes CampaignIncludedAttributes `json:"attributes"`
		Links      DataLinks                  `json:"links"`
	}

	CampaignIncludedAttributes struct {
		Label         *string        `json:"label,omitempty"`
		Channel       *string        `json:"channel,omitempty"`
		Content       *Content       `json:"content,omitempty"`
		SendTimes     []SendTime     `json:"send_times,omitempty"`
		RenderOptions *RenderOptions `json:"render_options,omitempty"`
		CreatedAt     *string        `json:"created_at,omitempty"`
		UpdatedAt     *string        `json:"updated_at,omitempty"`
		Name          *string        `json:"name,omitempty"`
	}
)
