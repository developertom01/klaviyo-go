package campaigns

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type CampaignRelationshipsAPi interface {
	//Returns the campaign ID of the related campaign. `type`: campaign, `id`: campaign ID
	GetCampaignMessageRelationshipsCampaign(ctx context.Context, messageId string) (*models.RelationshipData, error)

	//Returns the template ID of the related template.  `type`: template, `id`: template ID
	GetCampaignMessageRelationshipsTemplate(ctx context.Context, messageId string) (*models.RelationshipData, error)

	//Returns the IDs of all tags associated with the given campaign. [`type`: tag, `id`: tag ID]
	GetCampaignRelationshipsTags(ctx context.Context, campaignId string) (*models.RelationshipDataCollectionResponse, error)

	//Returns the IDs of all messages associated with the given campaign. [`type`: campaign-messages, `id`: message ID]
	GetCampaignRelationshipsCampaignMessages(ctx context.Context, campaignId string) (*models.RelationshipDataCollectionResponse, error)
}

func (api *campaignsApi) GetCampaignMessageRelationshipsCampaign(ctx context.Context, messageId string) (*models.RelationshipData, error) {
	url := fmt.Sprintf("%s/api/campaign-messages/%s/relationships/campaign/", api.baseApiUrl, messageId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var res models.RelationshipData
	err = json.Unmarshal(byteData, &res)

	return &res, err
}

func (api *campaignsApi) GetCampaignMessageRelationshipsTemplate(ctx context.Context, messageId string) (*models.RelationshipData, error) {
	url := fmt.Sprintf("%s/api/campaign-messages/%s/relationships/template/", api.baseApiUrl, messageId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var res models.RelationshipData
	err = json.Unmarshal(byteData, &res)

	return &res, err
}

func (api *campaignsApi) GetCampaignRelationshipsTags(ctx context.Context, campaignId string) (*models.RelationshipDataCollectionResponse, error) {
	url := fmt.Sprintf("%s/api/campaigns/%s/relationships/tags/", api.baseApiUrl, campaignId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var res models.RelationshipDataCollectionResponse
	err = json.Unmarshal(byteData, &res)

	return &res, err
}

func (api *campaignsApi) GetCampaignRelationshipsCampaignMessages(ctx context.Context, campaignId string) (*models.RelationshipDataCollectionResponse, error) {
	url := fmt.Sprintf("%s/api/campaigns/%s/relationships/campaign-messages/", api.baseApiUrl, campaignId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, err
	}

	var res models.RelationshipDataCollectionResponse
	err = json.Unmarshal(byteData, &res)

	return &res, err
}
