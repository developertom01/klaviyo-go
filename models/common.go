package models

import "github.com/jaswdr/faker"

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

type SendTime struct {
	Datetime string `json:"datetime"`
	IsLocal  bool   `json:"is_local"`
}

type DataLinks struct {
	Self string `json:"self"`
}

type Relationships struct {
	Data  RelationshipData  `json:"data"`
	Links RelationshipLinks `json:"links"`
}

type RelationshipData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type RelationshipLinks struct {
	Self    string `json:"self"`
	Related string `json:"related"`
}

type Audiences struct {
	Included []string `json:"included,omitempty"` //A list of included audiences
	Excluded []string `json:"excluded,omitempty"` //An optional list of excluded audiences
}

// Options to use when sending a campaign
type SendOptions struct {
	UseSmartSending bool `json:"use_smart_sending"` //Use smart sending. Defaults to True
}

type SendStrategy struct {
	Method           SendStrategyMethod `json:"method"`                      //Describes the shape of the options object. Allowed values: ['static', 'throttled', 'immediate', 'smart_send_time']
	OptionsStatic    *OptionsStatic     `json:"options_static,omitempty"`    //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'static' method.
	OptionsThrottled *OptionsThrottled  `json:"options_throttled,omitempty"` //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'throttled' method.
	OptionsSto       *OptionsSto        `json:"options_sto"`                 //The send configuration options the campaign will send with. These define variables that alter the send strategy and must match the given method. Intended to be used with the 'smart_send_time' method.
}

type OptionsStatic struct {
	Datetime                      string `json:"datetime"`
	IsLocal                       bool   `json:"is_local"`
	SendPastRecipientsImmediately bool   `json:"send_past_recipients_immediately"`
}

type OptionsSto struct {
	Date string `json:"date"`
}

type OptionsThrottled struct {
	Datetime           string `json:"datetime"`
	ThrottlePercentage int64  `json:"throttle_percentage"`
}

type TrackingOptions struct {
	IsAddUtm         *bool      `json:"is_add_utm,omitempty"`         //Whether the campaign needs UTM parameters. If set to False, UTM params will not be used.
	UtmParams        []UtmParam `json:"utm_params,omitempty"`         //A list of UTM parameters. If an empty list is given and is_add_utm is True, uses company defaults.
	IsTrackingClicks *bool      `json:"is_tracking_clicks,omitempty"` //Whether the campaign is tracking click events. If not specified, uses company defaults.
	IsTrackingOpens  *bool      `json:"is_tracking_opens,omitempty"`  //Whether the campaign is tracking open events. If not specified, uses company defaults.
}

type UtmParam struct {
	Name  string `json:"name"`  //Name of the UTM param
	Value string `json:"value"` //Value of the UTM param. Can be templated data.
}

type StreetAddress struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Zip      string `json:"zip"`
}

// Describes the shape of the options object. Allowed values: ['static', 'throttled', 'immediate', 'smart_send_time']
type SendStrategyMethod string

const (
	SendStrategyMethodStatic        SendStrategyMethod = "static"
	SendStrategyMethodThrottled     SendStrategyMethod = "throttled"
	SendStrategyMethodImmediate     SendStrategyMethod = "immediate"
	SendStrategyMethodSmartSendTime SendStrategyMethod = "smart_send_time"
)

type Links struct {
	Self     *string `json:"self,omitempty"`
	First    *string `json:"first,omitempty"`
	Last     *string `json:"last,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Next     *string `json:"next,omitempty"`
}

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
