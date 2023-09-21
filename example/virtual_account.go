/*
 * File Created: Thursday, 21st September 2023 5:34:18 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package example

import (
	"fmt"
	"time"

	durianpay "github.com/abmid/dpay-sdk-go"
)

func VirtualAccountCreate() {
	payload := durianpay.VirtualAccountPayload{
		BankCode: "PERMATA",
		Name:     "Abdul Hamid",
		IsClosed: true,
		Amount:   "123000",
		Customer: durianpay.VirtualAccountCustomer{
			GivenName: "Abdul Hamid",
			Mobile:    "+6285555555555",
			Email:     "abdul.surel@gmail.com",
		},
		ExpiryMinutes:           120,
		AccountSuffix:           "123456",
		IsReusable:              true,
		VaRefID:                 "1234",
		MinAmount:               10000,
		MaxAmount:               20000,
		AutoDisableAfterPayment: true,
	}

	res, err := c.VA.Create(ctx, payload)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}

func VirtualAccountFetch() {
	options := durianpay.VirtualAccountFetchOption{
		From:  time.Now().Format("2006-01-02"),
		To:    time.Now().Add(60 * time.Second).Format("2006-01-02"),
		Skip:  10,
		Limit: 10,
	}

	res, err := c.VA.FetchVirtualAccounts(ctx, options)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}
