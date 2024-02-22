package common

import (
	"fmt"
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

func Retry(options RetryOptions, fn func() error) error {
	ticker := time.Tick(options.Interval)
	var err error
	for retries := 0; retries < options.MaxRetries; retries++ {
		if err = fn(); err == nil {
			return nil
		}
		fmt.Printf("Error occurred: %v. Retrying...\n", err)
		<-ticker // Wait for the next tick
	}
	return err
}

var retryStatusCode = []int{500, 400, 300}
