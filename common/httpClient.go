package common

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/stretchr/testify/mock"
)

// HTTPClient interface for making HTTP requests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// MockHTTPClient is a mock implementation of HTTPClient for testing
type MockHTTPClient struct {
	mock.Mock
}

// Do mocks the Do method of the HTTP client
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func NewMockHTTPClient() *MockHTTPClient {
	return &MockHTTPClient{}
}

func BuildURLWithQueryParams(baseUrl string, queryParams map[string][]string) (string, error) {
	params := url.Values{}
	for key, val := range queryParams {
		if val != nil && len(val) > 0 {
			fieldsStr := strings.Join(val, ",")
			params.Set(key, fieldsStr)
		}

	}

	apiUrl, err := url.Parse(fmt.Sprintf("%s/?%s", baseUrl, params.Encode()))

	if err != nil {
		return "", err
	}
	fmt.Println(apiUrl)

	return apiUrl.String(), nil
}
