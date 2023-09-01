/*
 * File Created: Friday, 1st September 2023 11:29:23 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package refund

import "time"

// Refund is general response from Refund API.
// Currently this response use for Create and Fetch By ID API.
type Refund struct {
	ID            string    `json:"id"`
	RefID         string    `json:"ref_id"`
	Amount        string    `json:"amount"`
	RefundType    string    `json:"refund_type"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ApprovedAt    time.Time `json:"approved_at"`
	Source        string    `json:"source"`
	CustomerID    string    `json:"customer_id"`
	CustomerName  string    `json:"customer_name"`
	CustomerEmail string    `json:"customer_email"`
	CustomerPhone string    `json:"customer_phone"`
	FailureReason string    `json:"failure_reason"`
}

// FetchRefunds is response for Refund Fetch API
type FetchRefunds struct {
	Refunds   []Refunds `json:"refund"`
	TotalData uint32    `json:"total_data"`
}

// Refunds is part of FetchRefunds for attribute Refunds
type Refunds struct {
	ID                 string    `json:"id"`
	MerchantID         string    `json:"merchant_id"`
	Status             string    `json:"status"`
	DisbursementID     string    `json:"disbursement_id"`
	TotalAmount        string    `json:"total_amount"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	ApprovedAt         time.Time `json:"approved_at"`
	PaymentID          string    `json:"payment_id"`
	RefundRefID        string    `json:"json:"refund_ref_id""`
	IsLive             bool      `json:"is_live"`
	Type               string    `json:"type"`
	OrderID            string    `json:"order_id"`
	CustomerID         string    `json:"customer_id"`
	RefundPartial      string    `json:"refund_partial"`
	RefundNotes        string    `json:"refund_notes"`
	PaymentPaidAmount  string    `json:"payment_paid_amount"`
	PaymentDetailsType string    `json:"payment_details_type"`
	PaymentMethodID    string    `json:"payment_method_id"`
	CustomerName       string    `json:"customer_name"`
	CustomerEmail      string    `json:"customer_email"`
	CustomerPhone      string    `json:"customer_phone"`
	Destination        string    `json:"destination"`
	EwalletName        string    `json:"ewallet_name"`
	AccountName        string    `json:"account_name"`
	AccountNumber      string    `json:"account_number"`
	CreatedBy          uint16    `json:"created_by"`
	CreatedByName      string    `json:"created_by_name"`
	UpdatedBy          uint16    `json:"updated_by"`
	UpdatedByName      string    `json:"updated_by_name"`
	History            string    `json:"history"`
	Source             string    `json:"source"`
	AllowRetrigger     bool      `json:"allow_retrigger"`
	FailureReason      string    `json:"failure_reason"`
}
