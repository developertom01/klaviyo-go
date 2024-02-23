package common

type Links struct {
	Self     *string `json:"self,omitempty"`
	First    *string `json:"first,omitempty"`
	Last     *string `json:"last,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Next     *string `json:"next,omitempty"`
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
