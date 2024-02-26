package campaigns

import (
	"encoding/json"

	"github.com/developertom01/klaviyo-go/models"
)

// Create and Update campaign API request models
type (
	CreateCampaignRequestData struct {
		Data CreateCampaignData `json:"data"`
	}

	UpdateCampaignRequestData struct {
		Data UpdateCampaignData `json:"data"`
	}

	CreateCampaignData struct {
		Type       string                           `json:"type"`
		Attributes CreateCampaignDataDataAttributes `json:"attributes"` //The audiences to be included and/or excluded from the campaign
	}

	UpdateCampaignData struct {
		ID         string                           `json:"id"` //The campaign ID to be retrieved
		Type       string                           `json:"type"`
		Attributes UpdateCampaignDataDataAttributes `json:"attributes"` //The audiences to be included and/or excluded from the campaign
	}

	UpdateCampaignDataDataAttributes struct {
		Name            *string                            `json:"name,omitempty"`             //The campaign name
		Audiences       *CampaignDataAttributesAudiences   `json:"audiences,omitempty"`        //The audiences to be included and/or excluded from the campaign
		SendStrategy    *CampaignDataAttributeSendStrategy `json:"send_strategy,omitempty"`    //The send strategy the campaign will send with. Defaults to 'Immediate' send strategy.
		SendOptions     *SendOptions                       `json:"send_options,omitempty"`     //Options to use when sending a campaign
		TrackingOptions *TrackingOptions                   `json:"tracking_options,omitempty"` //The tracking options associated with the campaign
	}

	CreateCampaignDataDataAttributes struct {
		Name             string                             `json:"name"`                       //The campaign name
		Audiences        CampaignDataAttributesAudiences    `json:"audiences"`                  //The audiences to be included and/or excluded from the campaign
		SendStrategy     *CampaignDataAttributeSendStrategy `json:"send_strategy,omitempty"`    //The send strategy the campaign will send with. Defaults to 'Immediate' send strategy.
		SendOptions      *SendOptions                       `json:"send_options,omitempty"`     //Options to use when sending a campaign
		TrackingOptions  *TrackingOptions                   `json:"tracking_options,omitempty"` //The tracking options associated with the campaign
		CampaignMessages CampaignAttributesMessages         `json:"campaign-messages"`          //The message(s) associated with the campaign
	}

	CampaignDataAttributesAudiences struct {
		Included []string `json:"included"`
		Excluded []string `json:"excluded"`
	}

	CampaignAttributesMessages struct {
		Data []CampaignAttributesMessagesData `json:"data"`
	}

	// The tracking options associated with the campaign
	TrackingOptions struct {
		IsAddUtm         *bool      `json:"is_add_utm,omitempty"`         //Whether the campaign needs UTM parameters. If set to False, UTM params will not be used.
		UtmParams        []UtmParam `json:"utm_params"`                   //A list of UTM parameters. If an empty list is given and is_add_utm is True, uses company defaults.
		IsTrackingClicks *bool      `json:"is_tracking_clicks,omitempty"` //Whether the campaign is tracking click events. If not specified, uses company defaults.
		IsTrackingOpens  *bool      `json:"is_tracking_opens,omitempty"`  //Whether the campaign is tracking open events. If not specified, uses company defaults.
	}

	CampaignAttributesMessagesData struct {
		Type       string                               `json:"type"`
		Attributes CreateCampaignMessagesDataAttributes `json:"attributes"`
	}

	CreateCampaignMessagesDataAttributes struct {
		Channel       string                `json:"channel"`                  //The channel the message is to be sent on (email or sms, for example)
		Label         *string               `json:"label,omitempty"`          //The label or name on the message
		Content       *MessageContent       `json:"content,omitempty"`        //Additional attributes relating to the content of the message
		RenderOptions *MessageRenderOptions `json:"render_options,omitempty"` //Additional options for rendering the message
	}

	MessageRenderOptions struct {
		ShortenLinks      *bool `json:"shorten_links,omitempty"`
		AddOrgPrefix      *bool `json:"add_org_prefix,omitempty"`
		AddInfoLink       *bool `json:"add_info_link,omitempty"`
		AddOptOutLanguage *bool `json:"add_opt_out_language,omitempty"`
	}

	// Options to use when sending a campaign
	SendOptions struct {
		UseSmartSending *bool `json:"use_smart_sending"` //Use smart sending. Defaults to True
	}

	CampaignDataAttributeSendStrategy struct {
		Method           models.SendStrategyMethod `json:"method"`                      //Describes the shape of the options object. Allowed values: ['static', 'throttled', 'immediate', 'smart_send_time']
		OptionsStatic    *OptionsStatic            `json:"options_static,omitempty"`    //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'static' method.
		OptionsThrottled *OptionsThrottled         `json:"options_throttled,omitempty"` //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'throttled' method.
		OptionsSto       OptionsSto                `json:"options_sto"`
	}

	OptionsStatic struct {
		Datetime                      string `json:"datetime"`                                   //The time to send at
		IsLocal                       *bool  `json:"is_local,omitempty"`                         //If the campaign should be sent with local recipient timezone send (requires UTC time) or statically sent at the given time. Defaults to False.
		SendPastRecipientsImmediately *bool  `json:"send_past_recipients_immediately,omitempty"` //Determines if we should send to local recipient timezone if the given time has passed. Only applicable to local sends. Defaults to False.
	}

	// The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'static' method.
	OptionsSto struct {
		Date string `json:"date"` //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'smart_send_time' method.
	}

	// The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'throttled' method.
	OptionsThrottled struct {
		Datetime           string `json:"datetime"`                      //The time to send at . eg 2022-11-08T00:00:00
		ThrottlePercentage *int64 `json:"throttle_percentage,omitempty"` //The percentage of recipients per hour to send to. Allowed values: [10, 11, 13, 14, 17, 20, 25, 33, 50]
	}

	UtmParam struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
)

type (
	//Clones a campaign from an existing campaign
	CreateCampaignCloneRequestData struct {
		Data CreateCampaignCloneData `json:"data"`
	}

	CreateCampaignCloneData struct {
		Type       string                               `json:"type"` //Campaign
		Attributes CreateCampaignCloneRequestAttributes `json:"attributes"`
		ID         string                               `json:"id"` //The campaign ID to be cloned
	}

	CreateCampaignCloneRequestAttributes struct {
		NewName string `json:"new_name"` //The name for the new cloned campaign
	}
)

// --- Update Campaign Message
type (
	UpdateCampaignMessagePayload struct {
		Type       string                          `json:"type"` // campaign-message
		ID         string                          `json:"id"`   //The message ID to be retrieved
		Attributes UpdateCampaignMessageAttributes `json:"attributes"`
	}

	UpdateCampaignMessageAttributes struct {
		Label         *string               `json:"label,omitempty"`          //The label or name on the message
		Content       *MessageContent       `json:"content,omitempty"`        //Additional attributes relating to the content of the message
		RenderOptions *MessageRenderOptions `json:"render_options,omitempty"` //Additional options for rendering the message
	}
)

// -------- Message content UNION type

type (
	//Generic map object that exports IsSMSContent and IsEmailContent
	MessageContent map[string]any

	MessageEmailContent struct {
		Subject      *string `json:"subject,omitempty"`        //The subject of the message
		PreviewText  *string `json:"preview_text,omitempty"`   //Preview text associated with the message
		FromEmail    *string `json:"from_email,omitempty"`     //The email the message should be sent from
		FromLabel    *string `json:"from_label,omitempty"`     //The label associated with the from_email
		ReplyToEmail *string `json:"reply_to_email,omitempty"` //Optional Reply-To email address
		CcEmail      *string `json:"cc_email,omitempty"`       //Optional CC email address

	}

	MessageSMSContent struct {
		Body *string `json:"body,omitempty"` //The message body
	}
)

// Converts data from `MessageEmailContent` or `MessageSMSContent` to `MessageContent`
func toMessageContent[T MessageEmailContent | MessageSMSContent](data T) (*MessageContent, error) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var messageContent MessageContent
	err = json.Unmarshal(dataByte, &messageContent)

	return &messageContent, err
}

// Converts to `MessageContent` type
func (mec MessageEmailContent) ToMessageContent() (*MessageContent, error) {
	return toMessageContent(mec)
}

// Converts to `MessageContent` type
func (msc MessageSMSContent) ToMessageContent() (*MessageContent, error) {
	return toMessageContent(msc)
}

// Returns MessageEmailContent and true if MessageContent is email
func (mc MessageContent) IsEmailContent() (*MessageEmailContent, bool) {
	byteData, err := json.Marshal(mc)
	if err != nil {
		return nil, false
	}

	var messageEmailContent MessageEmailContent
	err = json.Unmarshal(byteData, &messageEmailContent)
	if err != nil {
		return nil, false
	}

	return &messageEmailContent, true
}

// Returns MessageSMSContent and true if MessageContent is sms
func (mc MessageContent) IsSMSContent() (*MessageSMSContent, bool) {
	byteData, err := json.Marshal(mc)
	if err != nil {
		return nil, false
	}

	var msgSMSContent MessageSMSContent
	err = json.Unmarshal(byteData, &msgSMSContent)
	if err != nil {
		return nil, false
	}

	return &msgSMSContent, true
}

// ----- AssignCampaignMessageTemplate payloads

type (
	AssignCampaignMessageTemplatePayload struct {
		Type          string                                     `json:"type"` // campaign-message
		ID            string                                     `json:"id"`   //The message ID to be assigned to
		Relationships AssignCampaignMessageTemplateRelationships `json:"relationships"`
	}

	AssignCampaignMessageTemplateRelationships struct {
		Template TemplateRelationshipPayload `json:"template"`
	}

	TemplateRelationshipPayload struct {
		Data models.RelationshipData // Type should be `template` and ID is the template ID to assign
	}
)
