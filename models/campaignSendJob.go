package models

import (
	"fmt"
	"strings"
)

type (
	CampaignSendJobResponse struct {
		Data  CampaignSendJob `json:"data"`
		Links DataLinks       `json:"link"`
	}

	CampaignSendJob struct {
		Type       string                    `json:"type"` //campaign-send-job
		ID         string                    `json:"id"`   //The ID of the campaign to send
		Attributes CampaignSendJobAttributes `json:"attributes"`
	}

	CampaignSendJobAttributes struct {
		status CampaignSendJobStatus //The status of the send job. [`cancelled` `complete` `processing` `queued`]
	}

	//The status of the send job
	//[`cancelled` `complete` `processing` `queued`]
	CampaignSendJobStatus string
)

const (
	CampaignSendJobStatusCancelled  CampaignSendJobStatus = "canceled"
	CampaignSendJobStatusCompleted  CampaignSendJobStatus = "completed"
	CampaignSendJobStatusProcessing CampaignSendJobStatus = "processing"
	CampaignSendJobStatusQueued     CampaignSendJobStatus = "queued"
)

type CampaignSendJobField string

const (
	CampaignSendJobFieldStatus CampaignSendJobField = "status"
)

func BuildCampaignSendJobFieldsParam(fields []CampaignSendJobField) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[campaign-send-job]=%s", strings.Join(formattedFields, ","))
}
