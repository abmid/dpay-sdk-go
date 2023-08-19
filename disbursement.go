/*
 * File Created: Friday, 28th July 2023 5:33:23 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

import "time"

/*
Structs for Response API
*/

// ValidateDisbursement is struct for response validate disbursement API
type ValidateDisbursement struct {
	Message string                   `json:"message"`
	Data    ValidateDisbursementData `json:"data"`
}

// ValidateDisbursementData is data object part of ValidateDisbursement
type ValidateDisbursementData struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	AccountHolder string `json:"account_holder"`
	Status        string `json:"status"`
}

// Disbursement is response from disbursement API
type Disbursement struct {
	Message string           `json:"message"`
	Data    DisbursementData `json:"data"`
}

// DisbursementData is data object part of Disbursement
type DisbursementData struct {
	ID                 string    `json:"id"`
	IdempotencyKey     string    `json:"idempotency_key"`
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	Status             string    `json:"status"`
	TotalAmount        string    `json:"total_amount"`
	TotalDisbursements uint16    `json:"total_disbursements"`
	Description        string    `json:"description"`
	Fees               uint32    `json:"fees"`
	CreatedAt          time.Time `json:"created_at"`
}

// DisbursementItem is response from disbursement items API
type DisbursementItem struct {
	SubmissionStatus       string                  `json:"submission_status"`
	Count                  uint16                  `json:"count"`
	DisbursementBatchItems []DisbursementBatchItem `json:"disbursement_batch_items"`
}

// DisburementBatchItem is part of DisbursementItem
type DisbursementBatchItem struct {
	ID                      string                              `json:"id"`
	DisbursementBatchID     string                              `json:"disbursement_batch_id"`
	AccountOwnerName        string                              `json:"account_owner_name"`
	RealName                string                              `json:"real_name"`
	BankCode                string                              `json:"bank_code"`
	Amount                  string                              `json:"amount"`
	AccountNumber           string                              `json:"account_numer"`
	EmailRecipient          string                              `json:"email_recipient"`
	PhoneNumber             string                              `json:"phone_number"`
	InvalidFields           []DisbursementBatchItemInvalidField `json:"invalid_fields"`
	Status                  string                              `json:"status"`
	Notes                   string                              `json:"notes"`
	ApproverNotes           string                              `json:"approver_notes"`
	IsDeleted               bool                                `json:"is_deleted"`
	CreatedBy               string                              `json:"created_by"`
	UpdatedBy               string                              `json:"updated_by"`
	CreatedAt               time.Time                           `json:"created_at"`
	UpdatedAt               time.Time                           `json:"updated_at"`
	AllowRetrigger          bool                                `json:"allow_retrigger"`
	SplitID                 string                              `json:"split_id"`
	Receipt                 string                              `json:"receipt"`
	Fee                     string                              `json:"fee"`
	DisbursementStatusSetAt time.Time                           `json:"disbursement_status_set_at"`
	FailureReson            string                              `json:"failure_reason"`
}

// DisbursementBatchItemInvalidField is part of DisbursementBatchItem
type DisbursementBatchItemInvalidField struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

/*
Structs for Payload
*/

// ValidateDisbursementPayload is payload for request validate disbursement API
type ValidateDisbursementPayload struct {
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

// DisbursementOption is parameter for submit disbursement API
type DisbursementOption struct {
	ForceDisburse  *bool `url:"force_disburse"`
	SkipValidation *bool `url:"skip_validation"`
}
