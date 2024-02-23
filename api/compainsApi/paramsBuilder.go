package campaigns

import "fmt"

type CampaignsField string

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

type CampaignsFieldParamBuilder struct {
	params []CampaignsField
}

func NewCampaignsFieldParamBuilder() *CampaignsFieldParamBuilder {
	return &CampaignsFieldParamBuilder{}
}

func (builder *CampaignsFieldParamBuilder) Add(field CampaignsField) *CampaignsFieldParamBuilder {
	builder.params = append(builder.params, field)
	return builder
}

func (builder *CampaignsFieldParamBuilder) Build() string {
	return fmt.Sprintf("fields[campaign]=%v", builder.params)
}
