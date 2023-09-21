/*
 * File Created: Thursday, 21st September 2023 5:38:28 pm
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

func InvoiceCreate() {
	payload := durianpay.InvoiceCreatePayload{
		Amount:          "20000.67",
		RemainingAmount: "5000.67",
		Title:           "sample",
		InvoiceRefID:    "inv_ref_001",
		Customer: durianpay.Customer{
			CustomerRefID: "cust_001",
			GivenName:     "Jane Doe",
			Email:         "jane_doe@nomail.com",
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
		EnablePartialTransaction: true,
		PartialTransactionConfig: map[string]any{
			"min_acceptable_amount": 1000,
		},
		StartDate: time.Now(),
		DueDate:   time.Now(),
	}

	res, err := c.Invoice.Create(ctx, payload)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}

func InvoiceFetch() {
	options := durianpay.InvoiceFetchOption{
		From:  time.Now().Format("2006-01-02"),
		To:    time.Now().Add(60 * time.Second).Format("2006-01-02"),
		Skip:  10,
		Limit: 10,
	}

	res, err := c.Invoice.FetchInvoices(ctx, options)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}
