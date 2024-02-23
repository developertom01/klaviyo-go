package models

type CampaignMessageResponse struct {
	Data     CampaignMessage           `json:"data"`
	Included []CampaignMessageIncluded `json:"included"`
}

type CampaignMessageIncluded struct {
	Type       string                            `json:"type"`
	ID         string                            `json:"id"`
	Attributes CampaignMessageIncludedAttributes `json:"attributes"`
	Links      DataLinks                         `json:"links"`
}

type CampaignMessageIncludedAttributes struct {
	Name            string           `json:"name"`
	Status          *string          `json:"status,omitempty"`
	Archived        *bool            `json:"archived,omitempty"`
	Audiences       *Audiences       `json:"audiences,omitempty"`
	SendOptions     *SendOptions     `json:"send_options,omitempty"`
	TrackingOptions *TrackingOptions `json:"tracking_options,omitempty"`
	SendStrategy    *SendStrategy    `json:"send_strategy,omitempty"`
	CreatedAt       *string          `json:"created_at,omitempty"`
	ScheduledAt     *string          `json:"scheduled_at,omitempty"`
	UpdatedAt       *string          `json:"updated_at,omitempty"`
	SendTime        *string          `json:"send_time,omitempty"`
	EditorType      *string          `json:"editor_type,omitempty"`
	HTML            *string          `json:"html,omitempty"`
	Text            *string          `json:"text,omitempty"`
	Created         *string          `json:"created,omitempty"`
	Updated         *string          `json:"updated,omitempty"`
}

type CampaignMessage struct {
	Type          string                    `json:"type"`
	ID            string                    `json:"id"`
	Attributes    CampaignMessageAttributes `json:"attributes"`
	Links         DataLinks                 `json:"links"`
	Relationships Relationships             `json:"relationships"`
}

type CampaignMessageAttributes struct {
	Label         string        `json:"label"`
	Channel       string        `json:"channel"`
	Content       Content       `json:"content"`
	SendTimes     []SendTime    `json:"send_times"`
	RenderOptions RenderOptions `json:"render_options"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}

type CampaignRelationships struct {
	Campaign Relationships `json:"campaign"`
	Template Relationships `json:"template"`
}

type MessageIncludedAttributes struct {
	Name            string           `json:"name"`
	Status          *string          `json:"status,omitempty"`
	Archived        *bool            `json:"archived,omitempty"`
	Audiences       *Audiences       `json:"audiences,omitempty"`
	SendOptions     *SendOptions     `json:"send_options,omitempty"`
	TrackingOptions *TrackingOptions `json:"tracking_options,omitempty"`
	SendStrategy    *SendStrategy    `json:"send_strategy,omitempty"`
	CreatedAt       *string          `json:"created_at,omitempty"`
	ScheduledAt     *string          `json:"scheduled_at,omitempty"`
	UpdatedAt       *string          `json:"updated_at,omitempty"`
	SendTime        *string          `json:"send_time,omitempty"`
	EditorType      *string          `json:"editor_type,omitempty"`
	HTML            *string          `json:"html,omitempty"`
	Text            *string          `json:"text,omitempty"`
	Created         *string          `json:"created,omitempty"`
	Updated         *string          `json:"updated,omitempty"`
}
