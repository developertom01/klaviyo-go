package models

type Content struct {
	Subject      string `json:"subject"`
	PreviewText  string `json:"preview_text"`
	FromEmail    string `json:"from_email"`
	FromLabel    string `json:"from_label"`
	ReplyToEmail string `json:"reply_to_email"`
	CcEmail      string `json:"cc_email"`
	BccEmail     string `json:"bcc_email"`
}

type RenderOptions struct {
	ShortenLinks      bool `json:"shorten_links"`
	AddOrgPrefix      bool `json:"add_org_prefix"`
	AddInfoLink       bool `json:"add_info_link"`
	AddOptOutLanguage bool `json:"add_opt_out_language"`
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
	Included []string `json:"included"`
	Excluded []string `json:"excluded"`
}

type SendOptions struct {
	UseSmartSending bool `json:"use_smart_sending"`
}

type SendStrategy struct {
	Method           string           `json:"method"`
	OptionsStatic    OptionsStatic    `json:"options_static"`
	OptionsThrottled OptionsThrottled `json:"options_throttled"`
	OptionsSto       OptionsSto       `json:"options_sto"`
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
	IsAddUtm         bool       `json:"is_add_utm"`
	UtmParams        []UtmParam `json:"utm_params"`
	IsTrackingClicks bool       `json:"is_tracking_clicks"`
	IsTrackingOpens  bool       `json:"is_tracking_opens"`
}

type UtmParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type StreetAddress struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Zip      string `json:"zip"`
}
