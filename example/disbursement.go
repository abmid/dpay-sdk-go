/*
 * File Created: Sunday, 30th July 2023 5:19:51 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package example

import (
	"context"
	"fmt"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/client"
)

func ExampleValidateDisbursement() {
	c := client.NewClient(client.Options{
		ServerKey: "XXX-XXX",
	})

	payload := durianpay.ValidateDisbursementPayload{
		XIdempotencyKey: "1",
		AccountNumber:   "12345678",
		BankCode:        "bca",
	}

	res, err := c.Disbursement.ValidateDisbursement(context.TODO(), payload)
	if err != nil {
		// Handle error
	}

	// Will be return response from DurianPay
	fmt.Println(res)
}

func ExampleSubmitDisbursement() {
	c := client.NewClient(client.Options{
		ServerKey: "xxx-xxxx",
	})

	payload := durianpay.DisbursementPayload{
		XIdempotencyKey: "1",
		IdempotencyKey:  "1",
		Name:            "Test",
		Description:     "Desc test",
		Items: []durianpay.DisbursementItemPayload{
			{
				AccountOwnerName: "Jane Doe",
				BankCode:         "bca",
				Amount:           "10000",
				AccountNumber:    "222444",
				EmailRecipient:   "jane_doe@nomail.com",
				PhoneNumber:      "081234567890",
				Notes:            "Salary",
			},
		},
	}

	res, err := c.Disbursement.SubmitDisbursement(context.TODO(), payload, nil)
	if err != nil {
		// Handle error
	}

	// Will be return response from DurianPay
	fmt.Println(res)
}
