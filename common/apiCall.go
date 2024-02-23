package common

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/developertom01/klaviyo-go/exceptions"
)

func executeRequest(httpClient HTTPClient, req *http.Request, session Session, revision string) (*http.Response, error) {
	req.Header.Add("revision", revision)
	req.Header.Add("accept", "application/json")
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
