package models

import (
	"encoding/json"
	"time"

	"github.com/jaswdr/faker"
)

type (
	//MessageContent can be either MessageSmsContent | MessageSmsContent
	// Default type of `MessageContent` is map[string]any but had
	// iSMessageContentSms method which returns (MessageSmsContent, bool)
	// and isMessageContentEmail which returns (MessageEmailContent, bool)
	MessageContent map[string]any

	MessageEmailContent struct {
		Subject      *string `json:"subject,omitempty"`        //The subject of the message
		PreviewText  *string `json:"preview_text,omitempty"`   //Preview text associated with the message
		FromEmail    *string `json:"from_email,omitempty"`     //The email the message should be sent from
		FromLabel    *string `json:"from_label,omitempty"`     //The label associated with the from_email
		ReplyToEmail *string `json:"reply_to_email,omitempty"` //Optional Reply-To email address
		CcEmail      *string `json:"cc_email,omitempty"`       //Optional CC email address
		BccEmail     *string `json:"bcc_email,omitempty"`      //Optional BCC email address
	}

	MessageSmsContent struct {
		Body     *string `json:"body,omitempty"`      //The message body
		MediaUrl *string `json:"media_url,omitempty"` //URL for included media
	}

	//Additional options for rendering the message
	MessageRenderOptions struct {
		ShortenLinks      *bool `json:"shorten_links,omitempty"`
		AddOrgPrefix      *bool `json:"add_org_prefix,omitempty"`
		AddInfoLink       *bool `json:"add_info_link,omitempty"`
		AddOptOutLanguage *bool `json:"add_opt_out_language,omitempty"`
	}

	SendTime struct {
		Datetime string `json:"datetime"`
		IsLocal  bool   `json:"is_local"`
	}

	DataLinks struct {
		Self string `json:"self"`
	}

	Relationships struct {
		Data  *RelationshipData  `json:"data,omitempty"`
		Links *RelationshipLinks `json:"links,omitempty"`
	}

	RelationshipsRequestPayload struct {
		Data RelationshipData `json:"data,omitempty"`
	}

	RelationshipsCollectionRequestPayload struct {
		Data []RelationshipData `json:"data"`
	}

	RelationshipData struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}

	RelationshipDataCollectionResponse struct {
		Data []RelationshipData `json:"data"`
	}

	RelationshipLinks struct {
		Self    string `json:"self"`
		Related string `json:"related"`
	}

	Audiences struct {
		Included []string `json:"included,omitempty"` //A list of included audiences
		Excluded []string `json:"excluded,omitempty"` //An optional list of excluded audiences
	}

	// Options to use when sending a campaign
	SendOptions struct {
		UseSmartSending bool `json:"use_smart_sending"` //Use smart sending. Defaults to True
	}

	SendStrategy struct {
		Method           SendStrategyMethod `json:"method"`                      //Describes the shape of the options object. Allowed values: ['static', 'throttled', 'immediate', 'smart_send_time']
		OptionsStatic    *OptionsStatic     `json:"options_static,omitempty"`    //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'static' method.
		OptionsThrottled *OptionsThrottled  `json:"options_throttled,omitempty"` //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'throttled' method.
		OptionsSto       *OptionsSto        `json:"options_sto"`                 //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'smart_send_time' method.
	}

	OptionsStatic struct {
		Datetime                      string `json:"datetime"`
		IsLocal                       bool   `json:"is_local"`
		SendPastRecipientsImmediately bool   `json:"send_past_recipients_immediately"`
	}

	OptionsSto struct {
		Date time.Time `json:"date"`
	}

	OptionsThrottled struct {
		Datetime           string `json:"datetime"`
		ThrottlePercentage int64  `json:"throttle_percentage"`
	}

	TrackingOptions struct {
		IsAddUtm         *bool      `json:"is_add_utm,omitempty"`         //Whether the campaign needs UTM parameters. If set to False, UTM params will not be used.
		UtmParams        []UtmParam `json:"utm_params,omitempty"`         //A list of UTM parameters. If an empty list is given and is_add_utm is True, uses company defaults.
		IsTrackingClicks *bool      `json:"is_tracking_clicks,omitempty"` //Whether the campaign is tracking click events. If not specified, uses company defaults.
		IsTrackingOpens  *bool      `json:"is_tracking_opens,omitempty"`  //Whether the campaign is tracking open events. If not specified, uses company defaults.
	}

	UtmParam struct {
		Name  string `json:"name"`  //Name of the UTM param
		Value string `json:"value"` //Value of the UTM param. Can be templated data.
	}

	StreetAddress struct {
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City     string `json:"city"`
		Region   string `json:"region"`
		Country  string `json:"country"`
		Zip      string `json:"zip"`
	}

	Links struct {
		Self     *string `json:"self,omitempty"`
		First    *string `json:"first,omitempty"`
		Last     *string `json:"last,omitempty"`
		Previous *string `json:"previous,omitempty"`
		Next     *string `json:"next,omitempty"`
	}
)

// Returns `*MessageEmailContent` if type is email content
func (mc MessageContent) IsMessageContentEmail() (*MessageEmailContent, bool) {
	byteContent, err := json.Marshal(mc)
	if err != nil {
		return nil, false
	}

	var emailContent MessageEmailContent
	err = json.Unmarshal(byteContent, &emailContent)
	if err != nil {
		return nil, false
	}

	return &emailContent, true
}

// Returns `*MessageSmsContent` if type is sms content
func (mc MessageContent) IsMessageContentSms() (*MessageSmsContent, bool) {
	byteContent, err := json.Marshal(mc)
	if err != nil {
		return nil, false
	}

	var smsContent MessageSmsContent
	err = json.Unmarshal(byteContent, &smsContent)
	if err != nil {
		return nil, false
	}

	return &smsContent, true
}

// Describes the shape of the options object. Allowed values: ['static', 'throttled', 'immediate', 'smart_send_time']
type SendStrategyMethod string

const (
	SendStrategyMethodStatic        SendStrategyMethod = "static"
	SendStrategyMethodThrottled     SendStrategyMethod = "throttled"
	SendStrategyMethodImmediate     SendStrategyMethod = "immediate"
	SendStrategyMethodSmartSendTime SendStrategyMethod = "smart_send_time"
)

func MockedLinkResponse() Links {
	fake := faker.New()

	self := fake.Internet().URL()
	first := fake.Internet().URL()
	last := fake.Internet().URL()
	previous := fake.Internet().URL()
	next := fake.Internet().URL()

	return Links{
		Self:     &self,
		First:    &first,
		Last:     &last,
		Previous: &previous,
		Next:     &next,
	}

}

type PaginationOptionsParam interface {
	GetPageSize() int
	GetCursor() *string
	GetSortField() *string
}
