/*
 * File Created: Monday, 4th September 2023 4:44:39 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

/*
Payloads
*/

// PaymentChargeVAPayload is requests payload for Payment Charge API.
// This request for type `VA`
type PaymentChargeVAPayload struct {
	OrderID       string                `json:"order_id"`
	BankCode      string                `json:"bank_code"`
	Name          string                `json:"name"`
	Amount        string                `json:"amount"`
	PaymentRefID  string                `json:"payment_ref_id"`
	SandboxOption *PaymentSandboxOption // If you want send request as Sandbox use this option
}

// PaymentChargeEwalletPayload is requests payload for Payment Charge API.
// This request for type E-WALLET
type PaymentChargeEwalletPayload struct {
	OrderID       string                `json:"order_id"`
	Amount        string                `json:"amount"`
	Mobile        string                `json:"mobile"`
	WalletType    string                `json:"wallet_type"`
	SandboxOption *PaymentSandboxOption `json:"-"` // If you want send request as Sandbox use this option
}

// PaymentChargeRetailStorePayload is requests payload for Payment Charge API.
// This request for type `Retail Store`
type PaymentChargeRetailStorePayload struct {
	OrderID       string                `json:"order_id"`
	BankCode      string                `json:"bank_code"`
	Name          string                `json:"name"`
	Amount        string                `json:"amount"`
	PaymentRefID  string                `json:"payment_ref_id"`
	SandboxOption *PaymentSandboxOption `json:"-"` // If you want send request as Sandbox use this option
}

// PaymentChargeOnlineBankingPayload is requests payload for Payment Charge API.
// This requests for type `Online Banking`
type PaymentChargeOnlineBankingPayload struct {
	OrderID      string              `json:"order_id"`
	Type         string              `json:"type"`
	Name         string              `json:"name"`
	Amount       string              `json:"amount"`
	CustomerInfo PaymentCustomerInfo `json:"customer_info"`
	Mobile       string              `json:"mobile"`
}

// PaymentCustomerInfo is part of PaymentRequestOnlineBanking for attribute Customer Info
type PaymentCustomerInfo struct {
	Email     string `json:"email"`
	GivenName string `json:"given_name"`
	ID        string `json:"id"`
}

// PaymentChargeQRISPayload is requests payload for Payment Charge API.
// This requests for type `QRIS`
type PaymentChargeQRISPayload struct {
	OrderID string `json:"order_id"`
	Type    string `json:"type"`
	Amount  string `json:"amount"`
	Name    string `json:"name"`
}

// PaymentChargeCardPayload is requests payload for Payment Charge API.
// This requests for type `CARD`
type PaymentChargeCardPayload struct {
	OrderID      string              `json:"order_id"`
	Amount       string              `json:"amount"`
	PaymentRefID string              `json:"payment_ref_id"`
	CustomerInfo PaymentCustomerInfo `json:"customer_info"`
}

// PaymentChargeBNPLPayload is requests payload for Payment Charge API.
// This requests for `BNPL`
type PaymentChargeBNPLPayload struct {
	OrderID               string                `json:"order_id"`
	Amount                string                `json:"amount"`
	PaymentRefID          string                `json:"payment_ref_id"`
	PaymentMethodUniqueID string                `json:"payment_method_unique_id"`
	CustomerInfo          PaymentCustomerInfo   `json:"customer_info"`
	SandboxOption         *PaymentSandboxOption `json:"-"` // If you want send request as Sandbox use this option
}

// PaymentSandboxOption is option for request payment charge as Sanbox Mode.
type PaymentSandboxOption struct {
	ForceFail bool `json:"force_fail"`
	DelayMS   int  `json:"delay_ms"`
}

// PaymentVerifyPayload is payload for Verify Payments API
type PaymentVerifyPayload struct {
	VerificationSignature string `json:"verification_signature"`
}

// PaymentCapturePayload is payload for Payment Capture API
type PaymentCapturePayload struct {
	Amount string `json:"amount"`
}

/*
Options
*/

// PaymentFetchOption is parameter for Payment Fetch API.
type PaymentFetchOption struct {
	From  string `url:"from"`
	To    string `url:"to"`
	Skip  uint16 `url:"skip"`
	Limit uint16 `url:"limit"`
}

// PaymentFetchByIDOption is parameter for Payment Fetch by ID API.
type PaymentFetchByIDOption struct {
	Expand string `url:"expand"` // customer or order
}

// PaymentMDRFeesOption is parameter for MDR Fees Calculation API.
type PaymentMDRFeesOption struct {
	Amount        string `url:"amount"`
	PaymentMethod string `url:"payment_method"`
}
