package campaigns

import "github.com/developertom01/klaviyo-go/models"

type CreateCampaignRequestData struct {
	Data CreateCampaignData `json:"data"`
}

type UpdateCampaignRequestData struct {
	Data UpdateCampaignData `json:"data"`
}

type CreateCampaignData struct {
	Type       string                           `json:"type"`
	Attributes CreateCampaignDataDataAttributes `json:"attributes"` //The audiences to be included and/or excluded from the campaign
}

type UpdateCampaignData struct {
	ID         string                           `json:"id"` //The campaign ID to be retrieved
	Type       string                           `json:"type"`
	Attributes UpdateCampaignDataDataAttributes `json:"attributes"` //The audiences to be included and/or excluded from the campaign
}

type UpdateCampaignDataDataAttributes struct {
	Name             *string                            `json:"name,omitempty"`             //The campaign name
	Audiences        *CampaignDataAttributesAudiences   `json:"audiences,omitempty"`        //The audiences to be included and/or excluded from the campaign
	SendStrategy     *CampaignDataAttributeSendStrategy `json:"send_strategy,omitempty"`    //The send strategy the campaign will send with. Defaults to 'Immediate' send strategy.
	SendOptions      *SendOptions                       `json:"send_options,omitempty"`     //Options to use when sending a campaign
	TrackingOptions  *TrackingOptions                   `json:"tracking_options,omitempty"` //The tracking options associated with the campaign
	CampaignMessages CampaignAttributesMessages         `json:"campaign-messages"`          //The message(s) associated with the campaign
}

type CreateCampaignDataDataAttributes struct {
	Name             string                             `json:"name"`                       //The campaign name
	Audiences        CampaignDataAttributesAudiences    `json:"audiences"`                  //The audiences to be included and/or excluded from the campaign
	SendStrategy     *CampaignDataAttributeSendStrategy `json:"send_strategy,omitempty"`    //The send strategy the campaign will send with. Defaults to 'Immediate' send strategy.
	SendOptions      *SendOptions                       `json:"send_options,omitempty"`     //Options to use when sending a campaign
	TrackingOptions  *TrackingOptions                   `json:"tracking_options,omitempty"` //The tracking options associated with the campaign
	CampaignMessages CampaignAttributesMessages         `json:"campaign-messages"`          //The message(s) associated with the campaign
}

type CampaignDataAttributesAudiences struct {
	Included []string `json:"included"`
	Excluded []string `json:"excluded"`
}

type CampaignAttributesMessages struct {
	Data []CampaignAttributesMessagesData `json:"data"`
}

// The tracking options associated with the campaign
type TrackingOptions struct {
	IsAddUtm         *bool      `json:"is_add_utm,omitempty"`         //Whether the campaign needs UTM parameters. If set to False, UTM params will not be used.
	UtmParams        []UtmParam `json:"utm_params"`                   //A list of UTM parameters. If an empty list is given and is_add_utm is True, uses company defaults.
	IsTrackingClicks *bool      `json:"is_tracking_clicks,omitempty"` //Whether the campaign is tracking click events. If not specified, uses company defaults.
	IsTrackingOpens  *bool      `json:"is_tracking_opens,omitempty"`  //Whether the campaign is tracking open events. If not specified, uses company defaults.
}

type CampaignAttributesMessagesData struct {
	Type       string                                   `json:"type"`
	Attributes CampaignAttributesMessagesDataAttributes `json:"attributes"`
}

type CampaignAttributesMessagesDataAttributes struct {
	Channel       string                `json:"channel"`                  //The channel the message is to be sent on (email or sms, for example)
	Label         *string               `json:"label,omitempty"`          //The label or name on the message
	Content       *MessageContent       `json:"content,omitempty"`        //Additional attributes relating to the content of the message
	RenderOptions *MessageRenderOptions `json:"render_options,omitempty"` //Additional options for rendering the message
}

type MessageContent struct {
	Subject      *string `json:"subject,omitempty"`        //The subject of the message
	PreviewText  *string `json:"preview_text,omitempty"`   //Preview text associated with the message
	FromEmail    *string `json:"from_email,omitempty"`     //The email the message should be sent from
	FromLabel    *string `json:"from_label,omitempty"`     //The label associated with the from_email
	ReplyToEmail *string `json:"reply_to_email,omitempty"` //Optional Reply-To email address
	CcEmail      *string `json:"cc_email,omitempty"`       //Optional CC email address
	BccEmail     *string `json:"bcc_email,omitempty"`      //Optional BCC email address
}

type MessageRenderOptions struct {
	ShortenLinks      *bool `json:"shorten_links,omitempty"`
	AddOrgPrefix      *bool `json:"add_org_prefix,omitempty"`
	AddInfoLink       *bool `json:"add_info_link,omitempty"`
	AddOptOutLanguage *bool `json:"add_opt_out_language,omitempty"`
}

// Options to use when sending a campaign
type SendOptions struct {
	UseSmartSending *bool `json:"use_smart_sending"` //Use smart sending. Defaults to True
}

type CampaignDataAttributeSendStrategy struct {
	Method           models.SendStrategyMethod `json:"method"`                      //Describes the shape of the options object. Allowed values: ['static', 'throttled', 'immediate', 'smart_send_time']
	OptionsStatic    *OptionsStatic            `json:"options_static,omitempty"`    //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'static' method.
	OptionsThrottled *OptionsThrottled         `json:"options_throttled,omitempty"` //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'throttled' method.
	OptionsSto       OptionsSto                `json:"options_sto"`
}

type OptionsStatic struct {
	Datetime                      string `json:"datetime"`                                   //The time to send at
	IsLocal                       *bool  `json:"is_local,omitempty"`                         //If the campaign should be sent with local recipient timezone send (requires UTC time) or statically sent at the given time. Defaults to False.
	SendPastRecipientsImmediately *bool  `json:"send_past_recipients_immediately,omitempty"` //Determines if we should send to local recipient timezone if the given time has passed. Only applicable to local sends. Defaults to False.
}

// The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'static' method.
type OptionsSto struct {
	Date string `json:"date"` //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'smart_send_time' method.
}

// The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'throttled' method.
type OptionsThrottled struct {
	Datetime           string `json:"datetime"`                      //The time to send at . eg 2022-11-08T00:00:00
	ThrottlePercentage *int64 `json:"throttle_percentage,omitempty"` //The percentage of recipients per hour to send to. Allowed values: [10, 11, 13, 14, 17, 20, 25, 33, 50]
}

type UtmParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
