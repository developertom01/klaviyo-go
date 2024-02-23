package campaigns

import (
	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/session"
)

type (
	CampaignsCollectionResponse struct {
		Data     []Campaign   `json:"data"`
		Links    common.Links `json:"links"`
		Included []Included   `json:"included"`
	}

	Campaign struct {
		string        `json:"type"`
		ID            string             `json:"id"`
		Attributes    CampaignAttributes `json:"attributes"`
		Links         DatumLinks         `json:"links"`
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

	Audiences struct {
		Included []string `json:"included"`
		Excluded []string `json:"excluded"`
	}

	SendOptions struct {
		UseSmartSending bool `json:"use_smart_sending"`
	}

	SendStrategy struct {
		Method           string           `json:"method"`
		OptionsStatic    OptionsStatic    `json:"options_static"`
		OptionsThrottled OptionsThrottled `json:"options_throttled"`
		OptionsSto       OptionsSto       `json:"options_sto"`
	}

	OptionsStatic struct {
		Datetime                      string `json:"datetime"`
		IsLocal                       bool   `json:"is_local"`
		SendPastRecipientsImmediately bool   `json:"send_past_recipients_immediately"`
	}

	OptionsSto struct {
		Date string `json:"date"`
	}

	OptionsThrottled struct {
		Datetime           string `json:"datetime"`
		ThrottlePercentage int64  `json:"throttle_percentage"`
	}

	TrackingOptions struct {
		IsAddUtm         bool       `json:"is_add_utm"`
		UtmParams        []UtmParam `json:"utm_params"`
		IsTrackingClicks bool       `json:"is_tracking_clicks"`
		IsTrackingOpens  bool       `json:"is_tracking_opens"`
	}

	UtmParam struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	DatumLinks struct {
		Self string `json:"self"`
	}

	Relationships struct {
		CampaignMessages CampaignMessages `json:"campaign-messages"`
		Tags             CampaignMessages `json:"tags"`
	}

	CampaignMessages struct {
		Data  []CampaignMessagesDatum `json:"data"`
		Links CampaignMessagesLinks   `json:"links"`
	}

	CampaignMessagesDatum struct {
		string `json:"type"`
		ID     string `json:"id"`
	}

	CampaignMessagesLinks struct {
		Self    string `json:"self"`
		Related string `json:"related"`
	}

	Included struct {
		string     `json:"type"`
		ID         string             `json:"id"`
		Attributes IncludedAttributes `json:"attributes"`
		Links      DatumLinks         `json:"links"`
	}

	IncludedAttributes struct {
		Label         *string        `json:"label,omitempty"`
		Channel       *string        `json:"channel,omitempty"`
		Content       *Content       `json:"content,omitempty"`
		SendTimes     []SendTime     `json:"send_times,omitempty"`
		RenderOptions *RenderOptions `json:"render_options,omitempty"`
		CreatedAt     *string        `json:"created_at,omitempty"`
		UpdatedAt     *string        `json:"updated_at,omitempty"`
		Name          *string        `json:"name,omitempty"`
	}

	Content struct {
		Subject      string `json:"subject"`
		PreviewText  string `json:"preview_text"`
		FromEmail    string `json:"from_email"`
		FromLabel    string `json:"from_label"`
		ReplyToEmail string `json:"reply_to_email"`
		CcEmail      string `json:"cc_email"`
		BccEmail     string `json:"bcc_email"`
	}

	RenderOptions struct {
		ShortenLinks      bool `json:"shorten_links"`
		AddOrgPrefix      bool `json:"add_org_prefix"`
		AddInfoLink       bool `json:"add_info_link"`
		AddOptOutLanguage bool `json:"add_opt_out_language"`
	}

	SendTime struct {
		Datetime string `json:"datetime"`
		IsLocal  bool   `json:"is_local"`
	}

	CampaignsApi interface {
	}
	campaignsApi struct {
		session    session.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}
)

func NewCampaignsApi(session session.Session, httpClient common.HTTPClient) CampaignsApi {
	return &campaignsApi{}
}
