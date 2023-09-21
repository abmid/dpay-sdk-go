/*
 * File Created: Wednesday, 23rd August 2023 9:53:20 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package disbursement

import "time"

// DisbursementValidate is struct for response validate disbursement API
type DisbursementValidate struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	AccountHolder string `json:"account_holder"`
	Status        string `json:"status"`
}

// Disbursement is response from disbursement API
type Disbursement struct {
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

// DisbursementItem is response from fetch disbursement items API
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
	AccountNumber           string                              `json:"account_number"`
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

// DisbursementBank is response from fetch banks API
type DisbursementBank struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DisbursementTopup is response from Topup Amount API
type DisbursementTopup struct {
	SenderBank  string                      `json:"sender_bank"`
	TotalAmount string                      `json:"total_amount"`
	Status      string                      `json:"status"`
	ExpiryDate  time.Time                   `json:"expiry_date"`
	TransferTo  DisbursementTopupTransferTo `json:"transfer_to"`
}

// DisbursementTopupTransferTo is part of DisbursementTopup
type DisbursementTopupTransferTo struct {
	BankCode          string `json:"bank_code"`
	BankName          string `json:"bank_name"`
	AtmBersamaCode    string `json:"atm_bersama_code"`
	BankAccountNumber string `json:"bank_account_number"`
	AccountHolderName string `json:"account_holder_name"`
	UniqueCode        int    `json:"unique_code"`
}
