package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/mock"
)

func MockedErrorResponse() exceptions.ApiErrorResponse {
	fake := faker.New()

	error := exceptions.ApiError{
		Id:     fake.UUID().V4(),
		Code:   strconv.FormatInt(int64(fake.Internet().StatusCode()), 10),
		Title:  fake.Lorem().Sentence(5),
		Detail: fake.Lorem().Sentence(10),
	}
	return exceptions.ApiErrorResponse{
		Errors: []exceptions.ApiError{error},
	}

}

func serializeMockResponse(respObj any) (io.ReadCloser, error) {
	responseByte, err := json.Marshal(respObj)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewBuffer(responseByte)), nil

}

func PrepareMockResponse(statusCode int, mockedRespData any, mockedHttpClient *MockHTTPClient) error {
	bodyResp, err := serializeMockResponse(mockedRespData)
	if err != nil {
		return err
	}

	response := http.Response{
		Status:     http.StatusText(statusCode),
		StatusCode: statusCode,
		Body:       bodyResp,
	}
	mockedHttpClient.On("Do", mock.Anything).Return(&response, nil)
	return nil
}
