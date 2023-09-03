# DurianPay SDK for Go #

![Test Status](https://github.com/abmid/dpay-sdk-go/actions/workflows/test.yml/badge.svg)

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Quickstart](#quickstart)
- [API Supports](#api-supports)
- [Contributing](#contributing)
- [License](#license)

## Overview

🚧 *The SDK is currently undergoing heavy development with frequent breaking changes* 🚧

For more information, visit the [DurianPay API Official documentation](https://durianpay.id/docs/api/).


## Installation

Make sure you are using go version `1.18` or later

```bash
go get github.com/abmid/dpay-sdk-go
```

## Quickstart

TODO

For more examples, please check directory [example](https://github.com/abmid/dpay-sdk-go/example).

## API Supports

- ORDERS
  - [x] Create Order
  - [x] Fetch Orders
  - [x] Fetch By ID
  - [x] Create Payment Link
- PAYMENTS
  - [ ] Charge Payment
  - [ ] Fetch Payments
  - [ ] Fetch Payment By ID
  - [ ] Check Payment Status
  - [ ] Verify Payment
  - [ ] Cancel Payment
  - [ ] MDR Fees Calculation
- PROMOS
  - [x] Create Promo
  - [x] Fetch Promos
  - [x] Fetch Promo By ID
  - [x] Delete Promo
  - [x] Update Promo
- DISBURSEMENTS
  - [x] Submit Disbursement
  - [x] Approve Disbursment
  - [x] Validate Disbursement
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
  - [ ] Create Invoince
  - [ ] Generate Checkout URL
  - [ ] Fetch Invoice By ID
  - [ ] Fetch Invoices / List Invoices
  - [ ] Update Invoice
  - [ ] Pay Invoice
  - [ ] Manual Payment Invoice
  - [ ] Delete Invoice

## Contributing

We are open to, and grateful for, any contribution. If you want to contribute please do PR and follow the code guide.

## License

Copyright (c) 2023-present [Abdul Hamid](https://github.com/abmid) and [Contributors](https://github.com/abmid/dpay-sdk-go/graphs/contributors).
