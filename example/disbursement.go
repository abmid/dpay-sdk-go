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
	"github.com/abmid/dpay-sdk-go/internal/tests"
)

func DisbursementValidate() {
	c := client.NewClient(client.Options{
		ServerKey: "XXX-XXX",
	})

	payload := durianpay.DisbursementValidatePayload{
		XIdempotencyKey: "1",
		AccountNumber:   "12345678",
		BankCode:        "bca",
	}

	res, err := c.Disbursement.Validate(context.TODO(), payload)
	if err != nil {
		// Handle error
	}

	// Will be return response from DurianPay
	fmt.Println(res)
}

func DisbursementSubmit() {
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

	// Without params
	res, err := c.Disbursement.Submit(context.TODO(), payload, nil)
	if err != nil {
		// Handle error
	}

	// With params
	opt := durianpay.DisbursementOption{
		ForceDisburse:  tests.ToPtr(true),
		SkipValidation: tests.ToPtr(false),
	}
	res, err = c.Disbursement.Submit(context.TODO(), payload, &opt)
	if err != nil {
		// Handle error
	}

	// Will be return response from DurianPay
	fmt.Println(res)
}

func DisbursementApprove() {
	c := client.NewClient(client.Options{
		ServerKey: "xxx-xxx",
	})

	payload := durianpay.DisbursementApprovePayload{
		XIdempotencyKey: "1",
		ID:              "dis_xxx",
	}

	// Without params
	res, err := c.Disbursement.Approve(context.TODO(), payload, nil)
	if err != nil {
		// Handle error
	}

	// With params
	opt := durianpay.DisbursementApproveOption{
		IgnoreInvalid: tests.ToPtr(false),
	}
	res, err = c.Disbursement.Approve(context.TODO(), payload, &opt)
	if err != nil {
		// Handle error
	}

	// Will be return response from DurianPay
	fmt.Println(res)
}

func DisbursementFetchItemsByID() {
	c := client.NewClient(client.Options{
		ServerKey: "xxx-xxx",
	})

	// With params
	opt := durianpay.DisbursementFetchItemsOption{
		Skip:  10,
		Limit: 10,
	}
	res, err := c.Disbursement.FetchItemsByID(context.TODO(), "dis_xxx", &opt)
	if err != nil {
		// Handle error
	}

	// Without params
	// res, err := c.Disbursement.FetchDisbursementItemsByID(context.TODO(), "dis_xxx", nil)

	// Will be return response from DurianPay
	fmt.Println(res)
}

func DisbursementFetchByID() {
	c := client.NewClient(client.Options{
		ServerKey: "xxx-xxx",
	})

	res, err := c.Disbursement.FetchByID(context.TODO(), "dis_xxx")
	if err != nil {
		// Handle error
	}

	// Will be return response from DurianPay
	fmt.Println(res)
}

func DisbursementDelete() {
	c := client.NewClient(client.Options{
		ServerKey: "xxx-xxx",
	})

	res, err := c.Disbursement.Delete(context.TODO(), "dis_xxx")
	if err != nil {
		// Handle error
	}

	// Will be return response from DurianPay
	fmt.Println(res)
}
