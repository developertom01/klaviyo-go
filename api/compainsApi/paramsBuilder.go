package campaigns

import (
	"fmt"
	"strings"
)

type CampaignsField string

type CampaignMessageField string

type CampaignSortField string

const (
	CampaignsFieldName                                                    CampaignsField = "name"
	CampaignsFieldStatus                                                  CampaignsField = "status"
	CampaignsFieldArchived                                                CampaignsField = "archived"
	CampaignsFieldAudience                                                CampaignsField = "audience"
	CampaignsFieldAudience_Include                                        CampaignsField = "audience.include"
	CampaignsFieldAudience_Exclude                                        CampaignsField = "audience.exclude"
	CampaignsFieldSendOptions                                             CampaignsField = "send_options"
	CampaignsFieldTrackingOptions                                         CampaignsField = "tracking_options"
	CampaignsFieldSendStrategy                                            CampaignsField = "send_strategy"
	CampaignsFieldSendStrategy_Method                                     CampaignsField = "send_strategy.method"
	CampaignsFieldSendStrategy_OptionStatic                               CampaignsField = "send_strategy.option_static"
	CampaignsFieldSendStrategy_OptionStatic_DateTime                      CampaignsField = "send_strategy.option_static.datetime"
	CampaignsFieldSendStrategy_OptionStatic_SendPastRecipientsImmediately CampaignsField = "send_strategy.option_static.send_past_recipients_immediately"
	CampaignsFieldSendStrategy_OptionsThrottled                           CampaignsField = "send_strategy.options_throttled"
	CampaignsFieldSendStrategy_OptionsThrottled_ThrottledPercentage       CampaignsField = "send_strategy.options_throttled.throttle_percentage"
	CampaignsFieldSendStrategy_OptionsSto                                 CampaignsField = "send_strategy.options_sto"
	CampaignsFieldSendStrategy_OptionsSto_Date                            CampaignsField = "send_strategy.options_sto.date"
	CampaignsCreatedAt                                                    CampaignsField = "created_at"
	CampaignsUpdatedAt                                                    CampaignsField = "updated_at"
	CampaignsSendTime                                                     CampaignsField = "send_time"
)

const (
	CampaignMessageFieldLabel                           CampaignMessageField = "label"
	CampaignMessageFieldChannel                         CampaignMessageField = "channel"
	CampaignMessageFieldContent                         CampaignMessageField = "content"
	CampaignMessageFieldSendTimes                       CampaignMessageField = "send_times"
	CampaignMessageFieldRenderOptions                   CampaignMessageField = "render_options"
	CampaignMessageFieldRenderOptions_ShortenLinks      CampaignMessageField = "render_options.shorten_links"
	CampaignMessageFieldRenderOptions_AddOrgPrefix      CampaignMessageField = "render_options.add_org_prefix"
	CampaignMessageFieldRenderOptions_AddOrgLink        CampaignMessageField = "render_options.add_org_link"
	CampaignMessageFieldRenderOptions_AddOptOutLanguage CampaignMessageField = "render_options.add_opt_out_language"
	CampaignMessageFieldCreatedAt                       CampaignMessageField = "created_at"
	CampaignMessageFieldUpdatedAt                       CampaignMessageField = "updated_at"
)

const (
	CampaignSortFieldNameAsc  CampaignSortField = "name"
	CampaignSortFieldNameDesc CampaignSortField = "-name"

	CampaignSortFieldIdAsc  CampaignSortField = "id"
	CampaignSortFieldIdDesc CampaignSortField = "-id"

	CampaignSortFieldScheduledAtAsc  CampaignSortField = "scheduled_at"
	CampaignSortFieldScheduledAtDesc CampaignSortField = "-scheduled_at"

	CampaignSortFieldCreatedAtAsc  CampaignSortField = "created_at"
	CampaignSortFieldCreatedAtDesc CampaignSortField = "-created_at"

	CampaignSortFieldUpdatedAtAsc  CampaignSortField = "updated_at"
	CampaignSortFieldUpdatedAtDesc CampaignSortField = "-updated_at"
)

type CampaignsFieldParamBuilder struct {
	params []CampaignsField
}

type CampaignMessageFieldParamBuilder struct {
	params []CampaignMessageField
}

func NewCampaignsFieldParamBuilder() *CampaignsFieldParamBuilder {
	return &CampaignsFieldParamBuilder{}
}

func NewCampaignMessageFieldParamBuilder() *CampaignMessageFieldParamBuilder {
	return &CampaignMessageFieldParamBuilder{}
}

func (builder *CampaignsFieldParamBuilder) Add(field CampaignsField) *CampaignsFieldParamBuilder {
	builder.params = append(builder.params, field)

	return builder
}

func (builder *CampaignsFieldParamBuilder) Build() string {
	return strings.ReplaceAll(fmt.Sprintf("fields[campaign]=%v", builder.params), " ", ",")
}

func (builder *CampaignMessageFieldParamBuilder) Add(field CampaignMessageField) *CampaignMessageFieldParamBuilder {
	builder.params = append(builder.params, field)

	return builder
}

func (builder *CampaignMessageFieldParamBuilder) Build() string {
	return strings.ReplaceAll(fmt.Sprintf("fields[campaign-message]=%v", builder.params), " ", ",")
}
