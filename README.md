# DurianPay SDK for Go #

[![GoDoc](https://godoc.org/github.com/abmid/dpay-sdk-go?status.svg)](https://godoc.org/github.com/abmid/dpay-sdk-go)
![Test Status](https://github.com/abmid/dpay-sdk-go/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/abmid/dpay-sdk-go)](https://goreportcard.com/report/github.com/abmid/dpay-sdk-go)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Documentation](#documentation)
- [API Supports](#api-supports)
- [Contributing](#contributing)
- [License](#license)

## Overview

🚧 *The SDK is currently undergoing heavy development with frequent changes, because of this the current major version is zero (v0.x.x)* 🚧

Durianpay is a payments platform and aggregator which helps business to connect with different payment service providers (PSPs) and gateways.

Durianpay provides SDKs in several programming languages but not Go. Because of this, this SDK was created.

For more information, visit the [DurianPay API Official documentation](https://durianpay.id/docs/api/).


## Installation

Make sure you are using go version `1.18` or later

```bash
go get github.com/abmid/dpay-sdk-go
```

## Documentation

```go
package main

import (
	"context"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/client"
)

func main() {
	// Init client to access all difference resources
	c := client.NewClient(client.Options{
		ServerKey: "XXX-XXX",
	})

	//----------------------------------------------
	// Example Validate Disbursement
	//----------------------------------------------
	payload := durianpay.DisbursementValidatePayload{
		XIdempotencyKey: "c128fc41-46e7-42fd-93ef-cb147a8f96c8",
		AccountNumber:   "12345678",
		BankCode:        "bca",
	}

	res, err := c.Disbursement.Validate(context.Background(), payload)
	if err != nil {
		// Handle error
	}
}
```

Below is an example of implementing Idempotent Request (X-Idempotency-Key and idempotency_key) in this SDK, for the differences between `X-Idempotency-Key` and `idempotency_key` you can read in the [Idempotent Request Implementation Guide](https://durianpay.id/docs/integration/disbursements/idempotent/)

```go
  payload := durianpay.DisbursementPayload{
    XIdempotencyKey: "c128fc41-46e7-42fd-93ef-cb147a8f96c8",
    IdempotencyKey:  "1",
    Name:            "Test Name",
    Description:     "Test Desc",
    Items: []durianpay.DisbursementItemPayload{
      {
        AccountOwnerName: "Goodman",
        BankCode:         "bca",
        Amount:           "10000",
        AccountNumber:    "222444",
        EmailRecipient:   "goodman@domain.com",
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
```

For more examples, please check directory [example](https://github.com/abmid/dpay-sdk-go/tree/master/example) and [Godoc](https://godoc.org/github.com/abmid/dpay-sdk-go)

## API Supports

- ORDERS
  - [x] Create Order
  - [x] Fetch Orders
  - [x] Fetch By ID
  - [x] Create Payment Link
- PAYMENTS
  - [x] Charge Payment
  - [x] Fetch Payments
  - [x] Fetch Payment By ID
  - [x] Check Payment Status
  - [x] Verify Payment
  - [x] Cancel Payment
  - [x] MDR Fees Calculation
- PROMOS
  - [x] Create Promo
  - [x] Fetch Promos
  - [x] Fetch Promo By ID
  - [x] Delete Promo
  - [x] Update Promo
- DISBURSEMENTS
  - [x] Submit Disbursement
  - [x] Approve Disbursment
  - [x] Validate Disbursement (Tested)
  - [x] Fetch Bank List
  - [x] Topup Amount
  - [x] Fetch Topup Detail By ID
  - [x] Fetch Balance
  - [x] Fetch Disbursement Items By ID
  - [x] Fetch Disbursement By ID
  - [x] Delete Disbusement By ID
- SETTLEMENTS
  - [x] Fetch Settlements
  - [x] Detail Settlement By ID
  - [x] Status Settlement By ID
  - [x] Fetch Settlement By ID
- REFUNDS
  - [x] Create Refund
  - [x] Fetch Refunds
  - [x] Fetch Refund By ID
- E-WALLET Account
  - [x] Link E-Wallet Account
  - [x] Unlink E-Wallet Account
  - [x] Detail E-Wallet Account
- VIRTUAL ACCOUNTS
  - [x] Create VA
  - [x] Fetch VAs
  - [x] Fetch VA By ID
  - [x] Patch VA By ID
  - [x] Simulate VA Payment
- INVOINCES
  - [x] Create Invoince
  - [x] Generate Checkout URL
  - [x] Fetch Invoice By ID
  - [x] Fetch Invoices / List Invoices
  - [x] Update Invoice
  - [x] Pay Invoice
  - [x] Manual Payment Invoice
  - [x] Delete Invoice

## Contributing

We are open  and grateful for any contribution. If you want to contribute please do PR and follow the code guide.

## License

Copyright (c) 2023-present [Abdul Hamid](https://github.com/abmid) and [Contributors](https://github.com/abmid/dpay-sdk-go/graphs/contributors). This SDK is free and open-source software licensed under the [MIT License](https://github.com/abmid/dpay-sdk-go/tree/master/LICENSE).
