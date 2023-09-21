/*
 * File Created: Thursday, 21st September 2023 5:12:53 pm
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

func PaymentCharge() {
	// Example case use Virtual Account
	vaPayload := durianpay.PaymentChargeVAPayload{
		OrderID:      "ord_WkJWY1ysZ57194",
		BankCode:     "MANDIRI",
		Name:         "Name Appear in ATM",
		Amount:       "20000",
		PaymentRefID: "pay_ref_123",
	}

	res, err := c.Payment.ChargeVA(ctx, vaPayload)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}

func PaymentFetchPayments() {
	options := durianpay.PaymentFetchOption{
		From:  time.Now().Format("2006-01-02"),
		To:    time.Now().Add(60 * time.Second).Format("2006-01-02"),
		Skip:  10,
		Limit: 10,
	}

	res, err := c.Payment.FetchPayments(ctx, options)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}
