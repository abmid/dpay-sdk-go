/*
 * File Created: Friday, 28th July 2023 5:33:23 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

import "time"

// ValidateDisbursementData is data object part of ValidateDisbursement
type ValidateDisbursementData struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	AccountHolder string `json:"account_holder"`
	Status        string `json:"status"`
}

// ValidateDisbursement is struct for response validate disbursement API
type ValidateDisbursement struct {
	Message string                   `json:"message"`
	Data    ValidateDisbursementData `json:"data"`
}

// ValidateDisbursementPayload is payload for request validate disbursement API
type ValidateDisbursementPayload struct {
	IdempotenKey  string `json:"-"`
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
}

// DisbursementPayload is payload for request disbursement API
type DisbursementPayload struct {
	AccountOwnerName string `json:"account_owner_name" validate:"required"`
	BankCode         string `json:"bank_code" validate:"required"`
	Amount           string `json:"amount" validate:"required"`
	AccountNumber    string `json:"accont_number" validate:"required"`
	EmailRecipient   string `json:"email_recipient"`
	PhoneNumber      string `json:"phone_number"`
	Notes            string `json:"notes"`
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

// Disbursement is response from disbursement API
type Disbursement struct {
	Message string           `json:"message"`
	Data    DisbursementData `json:"data"`
}

// DisbursementItem is response from disbursement items API
type DisbursementItem struct {
	SubmissionStatus       string                  `json:"submission_status"`
	Count                  uint16                  `json:"count"`
	DisbursementBatchItems []DisbursementBatchItem `json:"disbursement_batch_items"`
}

// DisbursementBatchItemInvalidField is part of DisbursementBatchItem
type DisbursementBatchItemInvalidField struct {
	Key     string `json:"key"`
	Message string `json:"message"`
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
