package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/developertom01/klaviyo-go"
	"github.com/developertom01/klaviyo-go/models"
	"github.com/developertom01/klaviyo-go/options"
)

func main() {
	var apiKey = "test-key"

	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)

	klaviyoApi := klaviyo.NewKlaviyoApi(opt, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	accounts, err := klaviyoApi.Accounts.GetAccounts(ctx, []models.AccountsField{models.AccountsFieldContactInformation, models.AccountsFieldContactInformation_DefaultSenderName})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(accounts)
}
