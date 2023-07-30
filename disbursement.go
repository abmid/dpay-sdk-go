/*
 * File Created: Friday, 28th July 2023 5:33:23 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

// ValidateDisbursementData is data object from response validate disbursement API
type ValidateDisbursementData struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	AccountHolder string `json:"account_holder"`
	Status        string `json:"status"`
}

// ValidateDisbursement is response from validate disbursement API
type ValidateDisbursement struct {
	Message string                   `json:"message"`
	Data    ValidateDisbursementData `json:"data"`
}

// ValidateDisbursementPayload is payload for request disbursement API
type ValidateDisbursementPayload struct {
	IdempotenKey  string
	AccountNumber string
	BankCode      string
}
