package common

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	"github.com/developertom01/klaviyo-go/exceptions"
	"github.com/jaswdr/faker/v2"
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

func PrepareMockResponse(respObj any) (io.ReadCloser, error) {
	responseByte, err := json.Marshal(respObj)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewBuffer(responseByte)), nil

}
