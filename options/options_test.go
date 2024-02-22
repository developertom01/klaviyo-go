package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testApiKey = "test_api_key"

func TestNewOptions(t *testing.T) {
	options := NewOptions()
	assert.Implements(t, (*Options)(nil), options)
}

func TestNewOptionsWithDefaultValues(t *testing.T) {
	options := NewOptionsWithDefaultValues()
	assert.Equal(t, DEFAULT_REVISION, options.Revision())
}

func TestWithRevisionAndRevision(t *testing.T) {
	var revision = "2023-02-15"
	options := NewOptions()

	options.WithRevision(revision)
	assert.Equal(t, revision, options.Revision())
}

func TestWithApiKeyAndApiKey(t *testing.T) {
	options := NewOptions()
	options.WithApiKey(testApiKey)
	assert.Equal(t, testApiKey, *options.ApiKey())
}

func TestWithCompanyIdAndCompanyId(t *testing.T) {
	var companyId = "test_company_id"
	options := NewOptions()
	options.WithCompanyId(companyId)
	assert.Equal(t, companyId, *options.CompanyId())
}
