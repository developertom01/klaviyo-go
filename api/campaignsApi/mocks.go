package campaigns

import (
	"encoding/json"
	"time"

	"github.com/developertom01/klaviyo-go/models"
	"github.com/jaswdr/faker"
)

const campaignMessageType string = "campaign-message"

func mockCampaignsData() models.Campaign {
	fake := faker.New()

	boolData := fake.Bool()
	return models.Campaign{
		ID: fake.UUID().V4(),
		Relationships: &models.CampaignRelationship{
			CampaignMessage: &models.Relationships{
				Data: &models.RelationshipData{
					Type: campaignMessageType,
					ID:   fake.UUID().V4(),
				},
				Links: &models.RelationshipLinks{
					Self:    fake.Internet().URL(),
					Related: fake.Internet().URL(),
				},
			},
			Tags: &models.Relationships{
				Data: &models.RelationshipData{
					Type: "tags",
					ID:   fake.UUID().V4(),
				},
				Links: &models.RelationshipLinks{
					Self:    fake.Internet().URL(),
					Related: fake.Internet().URL(),
				},
			},
		},
		Attributes: models.CampaignAttributes{
			Name:     fake.Company().Name(),
			Status:   fake.Lorem().Text(6),
			Archived: fake.Bool(),
			Audiences: models.Audiences{
				Included: []string{fake.Lorem().Word()},
				Excluded: []string{},
			},
			SendOptions: models.SendOptions{
				UseSmartSending: fake.Bool(),
			},
			TrackingOptions: models.TrackingOptions{
				IsAddUtm:         &boolData,
				IsTrackingClicks: &boolData,
				IsTrackingOpens:  &boolData,
				UtmParams: []models.UtmParam{
					{
						Name:  fake.Lorem().Word(),
						Value: fake.Lorem().Word(),
					},
					{
						Name:  fake.Lorem().Word(),
						Value: fake.Lorem().Word(),
					},
				},
			},
			SendStrategy: models.SendStrategy{
				Method: models.SendStrategyMethodImmediate,
				OptionsSto: &models.OptionsSto{
					Date: time.Now().Format(time.RFC3339),
				},
			},
			CreatedAt:   time.Now().Format(time.RFC3339),
			ScheduledAt: time.Now().Format(time.RFC3339),
			UpdatedAt:   time.Now().Format(time.RFC3339),
			SendTime:    time.Now().Format(time.RFC3339),
		},
		Links: models.DataLinks{
			Self: fake.Internet().URL(),
		},
	}
}

func mockCampaignCollectionResponse(n int) models.CampaignsCollectionResponse {
	campaigns := make([]models.Campaign, 0)
	for i := 0; i < n; i++ {
		campaigns = append(campaigns, mockCampaignsData())
	}
	return models.CampaignsCollectionResponse{
		Data:  campaigns,
		Links: models.MockedLinkResponse(),
	}
}

func mockCampaignResponse() models.CampaignResponse {
	return models.CampaignResponse{
		Data: mockCampaignsData(),
	}
}

func mockCreateCampaignRequestData() CreateCampaignRequestData {
	fake := faker.New()

	return CreateCampaignRequestData{
		Data: CreateCampaignData{
			Type: "Campaign",
			Attributes: CreateCampaignDataDataAttributes{
				Name: fake.Company().Name(),
				SendStrategy: &CampaignDataAttributeSendStrategy{
					Method: models.SendStrategyMethodSmartSendTime,
					OptionsSto: OptionsSto{
						Date: time.Now().Format(time.RFC822),
					},
				},
				Audiences: CampaignDataAttributesAudiences{
					Included: []string{},
					Excluded: []string{},
				},
				CampaignMessages: CampaignAttributesMessages{
					Data: []CampaignAttributesMessagesData{{
						Type: campaignMessageType,
						Attributes: CreateCampaignMessagesDataAttributes{
							Channel: "email",
						},
					}},
				},
			},
		},
	}
}

func mockUpdateCampaignRequestData() UpdateCampaignRequestData {
	fake := faker.New()

	name := fake.Company().Name()
	return UpdateCampaignRequestData{
		Data: UpdateCampaignData{
			Type: "Campaign",
			Attributes: UpdateCampaignDataDataAttributes{
				Name: &name,
				SendStrategy: &CampaignDataAttributeSendStrategy{
					Method: models.SendStrategyMethodSmartSendTime,
					OptionsSto: OptionsSto{
						Date: time.Now().Format(time.RFC822),
					},
				},
				Audiences: &CampaignDataAttributesAudiences{
					Included: []string{},
					Excluded: []string{},
				},
			},
		},
	}
}

func mockCreateCampaignCloneRequestDataRequestData() CreateCampaignCloneRequestData {
	fake := faker.New()

	return CreateCampaignCloneRequestData{
		Data: CreateCampaignCloneData{
			Type: "campaign",
			ID:   fake.UUID().V4(),
			Attributes: CreateCampaignCloneRequestAttributes{
				NewName: fake.Company().Name(),
			},
		},
	}
}

func mockCampaignMessage() models.CampaignMessage {
	fake := faker.New()

	var emailSubject = fake.Lorem().Sentence(20)
	var contentEmail = fake.Lorem().Sentence(20)
	emailContent := models.MessageEmailContent{
		Subject:   &emailSubject,
		FromEmail: &contentEmail,
	}

	emailContentByte, _ := json.Marshal(emailContent)
	var messageContent models.MessageContent
	json.Unmarshal(emailContentByte, &messageContent)

	return models.CampaignMessage{
		Type: "campaign-message",
		ID:   fake.UUID().V4(),
		Attributes: models.CampaignMessageAttributes{
			Label:         fake.Lorem().Word(),
			Channel:       "email",
			Content:       messageContent,
			RenderOptions: models.MessageRenderOptions{},
			SendTimes:     []models.SendTime{},
			CreatedAt:     time.Now().Format(time.RFC1123),
			UpdatedAt:     time.Now().Format(time.RFC1123),
		},
	}
}

func mockCampaignMessageResponse() models.CampaignMessageResponse {
	return models.CampaignMessageResponse{
		Data: mockCampaignMessage(),
	}
}

func mockCampaignMessageCollectionResponse(n int) models.CampaignMessageCollectionResponse {
	data := make([]models.CampaignMessage, 0)
	for i := 0; i < n; i++ {
		data = append(data, mockCampaignMessage())
	}

	return models.CampaignMessageCollectionResponse{
		Data:  data,
		Links: models.MockedLinkResponse(),
	}

}

func mockUpdateCampaignMessagePayload() UpdateCampaignMessagePayload {
	fake := faker.New()

	smsBody := "text message"
	smsContent := MessageSMSContent{
		Body: &smsBody,
	}
	msgContent, _ := smsContent.ToMessageContent()

	return UpdateCampaignMessagePayload{
		Type: "campaign-message",
		ID:   fake.UUID().V4(),
		Attributes: UpdateCampaignMessageAttributes{
			Content: msgContent,
		},
	}
}

func mockAssignCampaignMessageTemplatePayload() AssignCampaignMessageTemplatePayload {
	fake := faker.New()

	return AssignCampaignMessageTemplatePayload{
		Type: "campaign-message",
		ID:   fake.UUID().V4(),
		Relationships: AssignCampaignMessageTemplateRelationships{
			Template: TemplateRelationshipPayload{
				Data: models.RelationshipData{
					Type: "template",
					ID:   fake.UUID().V4(),
				},
			},
		},
	}
}
