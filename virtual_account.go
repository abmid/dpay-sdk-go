/*
 * File Created: Tuesday, 29th August 2023 2:47:43 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

/*
Payloads
*/

// VirtualAccountPayload is payload for Virtual Account Create API.
type VirtualAccountPayload struct {
	BankCode                string                 `json:"bank_code"`
	Name                    string                 `json:"name"`
	IsClosed                bool                   `json:"is_closed"`
	Amount                  string                 `json:"amount"`
	Customer                VirtualAccountCustomer `json:"customer"`
	ExpiryMinutes           uint32                 `json:"expiry_minutes"`
	AccountSuffix           string                 `json:"account_suffix"`
	IsReusable              bool                   `json:"is_reusable"`
	VaRefID                 string                 `json:"va_ref_id"`
	MinAmount               uint32                 `json:"min_amount"`
	MaxAmount               uint32                 `json:"max_amount"`
	AutoDisableAfterPayment bool                   `json:"auto_disable_after_payment"`
}

// VirtualAccountCustomer is part of VirtualAccountPayload for attribute Customer
type VirtualAccountCustomer struct {
	GivenName string `json:"given_name"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
}

// VirtualAccountPatchPayload is payload for Virtual Account Patch By ID API
type VirtualAccountPatchPayload struct {
	ExpiryMinutes uint32 `json:"expiry_minutes"`
	MinAmount     uint32 `json:"min_amount"`
	MaxAmount     uint32 `json:"max_amount"`
	Amount        uint32 `json:"amount"`
	IsDisabled    bool   `json:"is_disabled"`
	VaRefID       string `json:"va_ref_id"`
}
