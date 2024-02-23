package campaigns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCampaignsFieldParamBuilderAdd(t *testing.T) {
	var expectedString = "fields[campaign]=[created_at,archived,name]"

	builder := NewCampaignsFieldParamBuilder()
	builder.Add(CampaignsCreatedAt).Add(CampaignsFieldArchived).Add(CampaignsFieldName)

	assert.Equal(t, expectedString, builder.Build())
}

func TestCampaignMessageFieldParamBuilder(t *testing.T) {
	var expectedString = "fields[campaign-message]=[channel,label,content]"

	builder := NewCampaignMessageFieldParamBuilder()
	builder.Add(CampaignMessageFieldChannel).Add(CampaignMessageFieldLabel).Add(CampaignMessageFieldContent)

	assert.Equal(t, expectedString, builder.Build())
}
