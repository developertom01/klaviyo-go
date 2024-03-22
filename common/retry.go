package common

import (
	"log"
	"net/http"
	"slices"
	"time"
)

type RetryOptions struct {
	MaxRetries int
	Interval   time.Duration
}

func NewRetryOptionsWithDefaultValues() *RetryOptions {
	return &RetryOptions{
		MaxRetries: 5,
		Interval:   time.Millisecond * 500,
	}
}

type RetryableFunc func() (*http.Response, error)

func Retry(fn RetryableFunc, opt RetryOptions) (*http.Response, error) {
	var retryStatusCode = []int{http.StatusRequestTimeout, http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout}

	var resp *http.Response
	var err error

	for retries := 0; retries < opt.MaxRetries; retries++ {
		// Execute the provided function
		resp, err = fn()
		if err != nil {
			return nil, err
		}

		// If successful, break the loop
		if resp.StatusCode == http.StatusOK {
			break
		}

		if !slices.Contains(retryStatusCode, resp.StatusCode) {
			break
		}

		// Log the error
		log.Printf("Error making API call: %v\n", err)

		// If not the last retry, wait before retrying
		if retries < opt.MaxRetries-1 {
			log.Printf("Retrying in %v seconds...\n", opt.Interval.Seconds())
			time.Sleep(opt.Interval)
		}
	}

	return resp, err
}
