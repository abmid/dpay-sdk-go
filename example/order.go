/*
 * File Created: Thursday, 21st September 2023 4:36:47 pm
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

func OrderCreate() {
	payload := durianpay.OrderPayload{
		Amount:        "1000",
		PaymentOption: "full_payment",
		Currency:      "IDR",
		OrderRefID:    "order_ref_001",
		Customer: durianpay.Customer{
			CustomerRefID: "cust_001",
			GivenName:     "Jane Doe",
			Email:         "jane_doe@gmail.com",
			Mobile:        "85722173217",
			Address: durianpay.CustomerAddress{
				ReceiverName:  "Jude Casper",
				ReceiverPhone: "8987654321",
				Label:         "Home Address",
				AddressLine1:  "Jl. HR. Rasuna Said",
				AddressLine2:  "Apartment #786",
				City:          "Jakarta Selatan",
				Region:        "Jakarta",
				Country:       "Indonesia",
				PostalCode:    "560008",
				Landmark:      "Kota Jakarta Selatan",
			},
		},
		Items: []durianpay.OrderItem{
			{
				Name:  "LED Television",
				Qty:   1,
				Price: "10001.00",
				Logo:  "https://merchant.com/tv_image.jpg",
			},
		},
		Metadata: map[string]any{
			"my-meta-key":     "my-meta-value",
			"SettlementGroup": "BranchName",
		},
		ExpiryDate: time.Now().Add(60 * time.Minute),
	}

	res, err := c.Order.Create(ctx, payload)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}

func OrderFetchOrders() {
	options := durianpay.OrderFetchOption{
		From:  time.Now().Format("2006-01-02"),
		To:    time.Now().Add(60 * time.Second).Format("2006-01-02"),
		Skip:  10,
		Limit: 10,
	}
	res, err := c.Order.FetchOrders(ctx, options)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}
