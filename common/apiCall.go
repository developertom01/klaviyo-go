package common

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/developertom01/klaviyo-go/exceptions"
)

type (
	MultipartOptions struct {
		File          io.Reader
		FileFieldName string
		FileName      string

		Meta map[string]string
	}

	MultipartRequestOption struct {
		HttpClient HTTPClient
		Session    Session
		Url        string
		Revision   string
	}
)

func executeRequest(httpClient HTTPClient, req *http.Request, session Session, revision string) (*http.Response, error) {
	req.Header.Add("revision", revision)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	session.ApplyToRequest(session.GetOptions(), req)

	execFn := func() (*http.Response, error) {
		return httpClient.Do(req)
	}
	return Retry(execFn, session.GetRetryOptions())
}

func RetrieveData(httpClient HTTPClient, req *http.Request, session Session, revision string) ([]byte, error) {
	res, err := executeRequest(httpClient, req, session, revision)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if !exceptions.IsHttpCodeOk(res.StatusCode) {
		var errorRes exceptions.ApiErrorResponse
		err := json.NewDecoder(res.Body).Decode(&errorRes)
		if err != nil {
			return nil, err
		}
		return nil, exceptions.NewResponseError(errorRes)
	}

	return io.ReadAll(res.Body)
}

func MakeMultipartRequest(ctx context.Context, requestOptions MultipartRequestOption, multipartOptions MultipartOptions) ([]byte, error) {
	var requestWriter bytes.Buffer

	//Create multipart writer
	writer := multipart.NewWriter(&requestWriter)
	defer writer.Close()

	//Create file field
	multipartFileField, err := writer.CreateFormFile(multipartOptions.FileFieldName, multipartOptions.FileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(multipartFileField, multipartOptions.File)
	if err != nil {
		return nil, err
	}

	for key, value := range multipartOptions.Meta {
		err := writer.WriteField(key, value)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestOptions.Url, &requestWriter)
	req.Header.Add("revision", requestOptions.Revision)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	requestOptions.Session.ApplyToRequest(requestOptions.Session.GetOptions(), req)

	execFn := func() (*http.Response, error) {
		return requestOptions.HttpClient.Do(req)
	}

	res, err := Retry(execFn, requestOptions.Session.GetRetryOptions())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if !exceptions.IsHttpCodeOk(res.StatusCode) {
		var errorRes exceptions.ApiErrorResponse
		err := json.NewDecoder(res.Body).Decode(&errorRes)
		if err != nil {
			return nil, err
		}
		return nil, exceptions.NewResponseError(errorRes)
	}

	return io.ReadAll(res.Body)
}
