package campaigns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	var expectedString = "fields[campaign]=[created_at,archived,name]"

	builder := NewCampaignsFieldParamBuilder()
	builder.Add(CampaignsCreatedAt).Add(CampaignsFieldArchived).Add(CampaignsFieldName)

	assert.Equal(t, expectedString, builder.Build())
}
