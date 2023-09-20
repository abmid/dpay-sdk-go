/*
 * File Created: Monday, 18th September 2023 11:34:33 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package invoice

import "time"

// Create represents a response from Create Invoice API.
type Create struct {
	ID                       string         `json:"id"`
	InvoiceRefID             string         `json:"invoice_ref_id"`
	Title                    string         `json:"title"`
	Status                   string         `json:"status"`
	Amount                   string         `json:"amount"`
	RemainingAmount          string         `json:"remaining_amount"`
	DueDate                  time.Time      `json:"due_date"`
	StartDate                time.Time      `json:"start_date"`
	CreatedAt                time.Time      `json:"created_at"`
	CustomerID               string         `json:"customer_id"`
	EnablePartialTransaction bool           `json:"enable_partial_transaction"`
	PartialTransactionConfig map[string]any `json:"partial_transaction_config"`
	CheckoutURL              string         `json:"checkout_url"`
	CheckoutURLExpiryAt      time.Time      `json:"checkout_url_expiry_at"`
}

// Transaction is part of Create for attribute Transactions
type Transaction struct {
	ID     string `json:"ID"`
	Amount string `json:"amount"`
	Status string `json:"status"`
}

// FetchInvoiceByID represents a response from Payment Fetch By ID API.
type FetchInvoiceByID struct {
	ID                          string         `json:"id"`
	InvoiceRefID                string         `json:"invoice_ref_id"`
	CustomerID                  string         `json:"customer_id"`
	IsLive                      bool           `json:"is_live"`
	Title                       string         `json:"title"`
	Status                      string         `json:"status"`
	Amount                      string         `json:"amount"`
	RemainingAmount             string         `json:"remaining_amount"`
	StartDate                   time.Time      `json:"start_date"`
	DueDate                     time.Time      `json:"due_date"`
	CreatedAt                   time.Time      `json:"created_at"`
	IsPartialTransactionEnabled bool           `json:"is_partial_transaction_enabled"`
	PartialTransactionConfig    map[string]any `json:"partial_transaction_config"`
	InvoiceURL                  string         `json:"invoice_url"`
	Metadata                    map[string]any `json:"metadata"`
	IsBlocked                   bool           `json:"is_blocked"`
	Transactions                []Transaction  `json:"transactions"`
}

// Invoices is part of FetchInvoice for attribute Invoices
type Invoices struct {
	ID                          string         `json:"id"`
	InvoiceRefID                string         `json:"invoice_ref_id"`
	CustomerID                  string         `json:"customer_id"`
	IsLive                      bool           `json:"is_live"`
	Title                       string         `json:"title"`
	Status                      string         `json:"status"`
	Amount                      string         `json:"amount"`
	RemainingAmount             string         `json:"remaining_amount"`
	StartDate                   time.Time      `json:"start_date"`
	DueDate                     time.Time      `json:"due_date"`
	CreatedAt                   time.Time      `json:"created_at"`
	IsPartialTransactionEnabled bool           `json:"is_partial_transaction_enabled"`
	PartialTransactionConfig    map[string]any `json:"partial_transaction_config"`
	InvoiceURL                  string         `json:"invoice_url"`
	IsBlocked                   bool           `json:"is_blocked"`
}

// FetchInvoice represents a response from List Invoices API.
type FetchInvoices struct {
	Invoices   []Invoices `json:"invoices"`
	TotalCount int        `json:"total_count"`
}

// GenerateCheckoutURL represents a response from Generate Checkout URL API.
type GenerateCheckoutURL struct {
	URL    string    `json:"url"`
	Expiry time.Time `json:"expiry"`
}

// Pay represents a response from Pay Invoice API.
type Pay struct {
	VANumber             string `json:"va_number"`
	Amount               string `json:"amount"`
	BankCode             string `json:"bank_code"`
	InvoiceTransactionID string `json:"invoice_transaction_id"`
}

// ManualPay represents a response from Manual Payment for Invoice API.
type ManualPay struct {
	ID string `json:"id"`
}

// Update represents a response from Update Invoice API.
type Update struct {
	ID                       string         `json:"id"`
	InvoiceRefID             string         `json:"invoice_ref_id"`
	CustomerID               string         `json:"customer_id"`
	Title                    string         `json:"title"`
	Status                   string         `json:"status"`
	Amount                   string         `json:"amount"`
	RemainingAmount          string         `json:"remaining_amount"`
	StartDate                time.Time      `json:"start_date"`
	DueDate                  time.Time      `json:"due_date"`
	CreatedAt                time.Time      `json:"created_at"`
	UpdatedAt                time.Time      `json:"updated_at"`
	EnablePartialTransaction bool           `json:"enable_partial_transaction"`
	PartialTransactionConfig map[string]any `json:"partial_transaction_config"`
	Metadata                 map[string]any `json:"metadata"`
	IsBlocked                bool           `json:"is_blocked"`
}
