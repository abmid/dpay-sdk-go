/*
 * File Created: Monday, 18th September 2023 11:34:44 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

import "time"

/*
Payloads
*/

// InvoiceCreate represents payload for Create Invoice API.
type InvoiceCreatePayload struct {
	Amount                   string         `json:"amount"`
	RemainingAmount          string         `json:"remaining_amount"`
	Title                    string         `json:"title"`
	InvoiceRefID             string         `json:"invoice_ref_id"`
	Customer                 Customer       `json:"customer"`
	EnablePartialTransaction bool           `json:"enable_partial_transaction"`
	PartialTransactionConfig map[string]any `json:"partial_transaction_config"` // Key-Value pair that can be used to store configuration about partial transactions like minimum acceptable amount for a partial transaction
	StartDate                time.Time      `json:"start_date"`
	DueDate                  time.Time      `json:"due_date"`
}

// InvoiceUpdate represents payload for Update Invoice API.
type InvoiceUpdatePayload struct {
	InvoiceRefID             string         `json:"invoice_ref_id"`
	Title                    string         `json:"title"`
	StartDate                time.Time      `json:"start_date"`
	DueDate                  time.Time      `json:"due_date"`
	InvoiceURL               string         `json:"invoice_url"`
	EnablePartialTransaction bool           `json:"enable_partial_transaction"`
	PartialTransactionConfig map[string]any `json:"partial_transaction_config"` // Key-Value pair that can be used to store configuration about partial transactions like minimum acceptable amount for a partial transaction
	IsBlocked                bool           `json:"is_blocked"`
	RemainingAmount          string         `json:"remaining_amount"`
	Metadata                 map[string]any `json:"metadata"`
}

// InvoicePay represents payload for Pay Invoice API.
type InvoicePayPayload struct {
	BankCode   string    `json:"bank_code"`
	Invoices   []Invoice `json:"invoices"`
	CustomerID string    `json:"customer"`
}

// Invoice is part of InvoicePay for attribute Invoices.
type Invoice struct {
	ID                string `json:"id"`
	TransactionAmount string `json:"transaction_amount"`
}

// InvoiceManualPay represents payload for Manual Payment for Invoice API.
type InvoiceManualPayPayload struct {
	ID     string `json:"id"`
	Amount string `json:"amount"`
}

/*
Options
*/

// InvoiceFetchOption represents paramater for List Invoices API.
type InvoiceFetchOption struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Skip   uint32 `json:"skip"`
	Limit  uint16 `json:"limit"`
	Status string `json:"status"`
}
