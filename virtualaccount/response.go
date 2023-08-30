/*
 * File Created: Tuesday, 29th August 2023 10:51:29 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package virtualaccount

import (
	"time"
)

// VirtualAccount is struct for response Fetch VAs API
type VirtualAccount struct {
	ID                      string    `json:"id"`
	BankCode                string    `json:"bank_code"`
	AccountNumber           string    `json:"account_number"`
	Name                    string    `json:"name"`
	IsClosed                bool      `json:"is_closed"`
	Amount                  uint32    `json:"amount"`
	Currency                string    `json:"currency"`
	CustomerID              string    `json:"customer_id"`
	IsSandbox               bool      `json:"is_sandbox"`
	CreatedAt               time.Time `json:"created_at"`
	ExpiryAt                time.Time `json:"expiry_at"`
	IsDisabled              bool      `json:"is_disabled"`
	IsPaid                  bool      `json:"is_paid"`
	IsReusable              bool      `json:"is_reusable"`
	MinAmount               *uint32   `json:"min_amount"`
	MaxAmount               *uint32   `json:"max_amount"`
	VaRefID                 string    `json:"va_ref_id"`
	AutoDisableAfterPayment bool      `json:"auto_disable_after_payment"`
}

// Create is struct for response Virtual Account Create API
type Create struct {
	CustomerID     string         `json:"customer_id"`
	VirtualAccount VirtualAccount `json:"virtual_account"`
}

// Detail is struct for response Fetch By ID & Patch By ID API
type Detail struct {
	VirtualAccount       VirtualAccount         `json:"virtual_account"`
	VirtualAccountStatus string                 `json:"virtual_account_status"`
	Customer             VirtualAccountCustomer `json:"customer"`
}

// VirtualAccountCustomer is part of Detail
type VirtualAccountCustomer struct {
	ID            string `json:"id"`
	CustomerRefID string `json:"customer_ref_id"`
	GivenName     string `json:"given_name"`
	MiddleName    string `json:"middle_name"`
	SurName       string `json:"sur_name"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
}
