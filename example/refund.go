/*
 * File Created: Thursday, 21st September 2023 5:28:17 pm
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

func RefundCreate() {
	payload := durianpay.RefundPayload{
		RefID:         "order_ref_241",
		PaymentID:     "pay_y2yKEEWBYe1299",
		Amount:        "10000",
		UseRefundLink: false,
		Notes:         "rejected product",
	}

	res, err := c.Refund.Create(ctx, payload)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}

func RefundFetch() {
	options := durianpay.RefundFetchOption{
		From:  time.Now().Format("2006-01-02"),
		To:    time.Now().Add(60 * time.Second).Format("2006-01-02"),
		Skip:  10,
		Limit: 10,
	}

	res, err := c.Refund.FetchRefunds(ctx, options)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}
