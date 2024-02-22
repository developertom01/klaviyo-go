package options

// Latest revision date
const DEFAULT_REVISION string = "2024-02-15"

type (
	Options interface {
		//Set Klaviyo API revision date
		WithRevision(revision string) Options

		//Set api key - Needed for
		WithApiKey(apiKey string) Options

		//Set Company ID - Needed when calling client APIs
		WithCompanyId(companyId string) Options

		//Returns revision
		Revision() string

		//Returns API key
		ApiKey() *string

		//Returns company ID
		CompanyId() *string
	}

	options struct {
		revision  string
		apiKey    *string
		companyId *string
	}
)

func NewOptions() Options {
	return &options{}
}

func NewOptionsWithDefaultValues() Options {
	return &options{
		revision: DEFAULT_REVISION,
	}
}

func (opt *options) WithRevision(revision string) Options {
	opt.revision = revision
	return opt
}

func (opt *options) WithApiKey(apiKey string) Options {
	opt.apiKey = &apiKey
	return opt
}

func (opt *options) WithCompanyId(companyId string) Options {
	opt.companyId = &companyId
	return opt
}

func (opt *options) Revision() string {
	return opt.revision
}

func (opt *options) ApiKey() *string {
	return opt.apiKey
}

func (opt *options) CompanyId() *string {
	return opt.companyId
}
