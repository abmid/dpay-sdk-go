/*
 * File Created: Friday, 28th July 2023 5:33:23 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

/*
Payloads
*/

// DisbursementValidatePayload is payload for request validate disbursement API
type DisbursementValidatePayload struct {
	XIdempotencyKey string `json:"-"`
	AccountNumber   string `json:"account_number"`
	BankCode        string `json:"bank_code"`
}

// DisbursementPayload is payload for request disbursement API
type DisbursementPayload struct {
	XIdempotencyKey string                    `json:"-" validate:"required"`
	IdempotencyKey  string                    `json:"-" validate:"required"`
	Name            string                    `json:"name"`
	Description     string                    `json:"description"`
	Items           []DisbursementItemPayload `json:"items"`
}

// DisbursementItemPayload is part of DisbursementPayload for attribute items
type DisbursementItemPayload struct {
	AccountOwnerName string `json:"account_owner_name" validate:"required"`
	BankCode         string `json:"bank_code" validate:"required"`
	Amount           string `json:"amount" validate:"required"`
	AccountNumber    string `json:"account_number" validate:"required"`
	EmailRecipient   string `json:"email_recipient"`
	PhoneNumber      string `json:"phone_number"`
	Notes            string `json:"notes"`
}

// DisbursementApprovePayload is payload for request approve disbursement API
type DisbursementApprovePayload struct {
	XIdempotencyKey string `json:"-"`
	ID              string `json:"id"` //Disbursement ID
}

// DisbursementTopupPayload is payload for request Topup Amount API
type DisbursementTopupPayload struct {
	XIdempotencyKey string `json:"-"`
	BankID          uint16 `json:"bank_id"`
	Amount          string `json:"amount"`
}

/*
Options
*/

// DisbursementOption is parameter for submit disbursement API
type DisbursementOption struct {
	ForceDisburse  *bool `url:"force_disburse"`
	SkipValidation *bool `url:"skip_validation"`
}

// DisbursementApproveOption is paramaeter for approve disbursement API
type DisbursementApproveOption struct {
	IgnoreInvalid *bool `url:"ignore_invalid"`
}

// DisbursementFetchItemsOption is parameter for Fetch Disbursement Items API
type DisbursementFetchItemsOption struct {
	Skip  uint16 `json:"skip"`
	Limit uint16 `json:"limit"`
}
