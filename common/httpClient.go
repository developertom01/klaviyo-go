package common

import (
	"net/http"

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
