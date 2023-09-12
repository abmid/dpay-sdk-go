/*
 * File Created: Tuesday, 5th September 2023 11:13:37 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package payment

import "time"

// ChargeVA use for response Payment Charge API (VA)
type ChargeVA struct {
	Type     string           `json:"type"`
	Response chargeResponseVA `json:""response`
}

// ChargeBNPL use for response Payment Charge API (Buy Now PayLater)
type ChargeBNPL struct {
	Type     string             `json:"type"`
	Response chargeResponseBNPL `type:"response"`
}

// ChargeEwallet use for response Payment Charge API (E-Wallet)
type ChargeEwallet struct {
	Type     string                `json:"type"`
	Response chargeResponseEwallet `json:"response"`
}

// ChargeRetailStore use for response Payment Charge API (Retail Store)
type ChargeRetailStore struct {
	Type     string                    `json:"type"`
	Response chargeResponseRetailStore `json:"response"`
}

// ChargeOnlineBank use for response Payment Charge API (Online Bank)
type ChargeOnlineBank struct {
	Type     string                   `json:"type"`
	Response chargeResponseOnlineBank `json:"response"`
}

// ChargeQRIS use for response Payment Charge API (QRIS)
type ChargeQRIS struct {
	Type     string             `json:"type"`
	Response chargeResponseQRIS `json:"response"`
}

// ChargeCard use for response Payment Charge API (Card)
type ChargeCard struct {
	Type     string             `json:"type"`
	Response chargeResponseCard `json:"response"`
}

// ChargeResponseVA represents response for Payment Charge use Virtual Account
type chargeResponseVA struct {
	PaymentID          string    `json:"payment_id"`
	OrderID            string    `json:"order_id"`
	AccountNumber      string    `json:"account_number"`
	PaymentRefID       string    `json:"payment_ref_id"`
	ExpirationTime     time.Time `json:"expiration_time"`
	PaidAmount         string    `json:"paid_amount"`
	PaymentInstruction struct {
		EN paymentInstruction `json:"en"`
		ID paymentInstruction `json:"ID"`
	} `json:"payment_instruction"`
}

// paymentInstruction is part of ChargeResponseVA for attribute PaymentInstruction
type paymentInstruction struct {
	Atm struct {
		Heading         string `json:"heading"`
		InstructionText string `json:"instruction_text"`
	} `json:"atm"`
	MobileApp struct {
		Heading         string `json:"heading"`
		AppStoreURL     string `json:"appstore_url"`
		PlayStoreURL    string `json:"playstore_url"`
		InstructionText string `json:"instruction_text"`
	} `json:"mobile_app"`
	InternetBanking struct {
		Heading         string `json:"heading"`
		InstructionText string `json:"instruction_text"`
	} `json:"internet_banking"`
}

// chargeResponseEwallet represents response for Payment Charge use E-Wallet
type chargeResponseEwallet struct {
	PaymentID      string    `json:"payment_id"`
	OrderID        string    `json:"order_id"`
	Mobile         string    `json:"mobile"`
	Status         string    `json:"status"`
	ExpirationTime time.Time `json:"expiration_time"`
	CheckoutURL    string    `json:"checkout_url"`
	WebURL         string    `json:"web_url"`
	UniqueID       string    `json:"unique_id"`
	PaidAmount     string    `json:"paid_amount"`
}

// chargeResponseRetailStore represents response for Payment Charge use Retail Store (Alfamart, Indomaret)
type chargeResponseRetailStore struct {
	PaymentID      string    `json:"payment_id"`
	OrderID        string    `json:"order_id"`
	AccountNumber  string    `json:"account_number"`
	PaymentRefID   string    `json:"payment_ref_id"`
	ExpirationTime time.Time `json:"expiration_time"`
	PaidAmount     string    `json:"paid_amount"`
}

// chargeResponseOnlineBank represents response for payment charge use Online Bank like JeniusPay.
type chargeResponseOnlineBank struct {
	PaymentID      string    `json:"payment_id"`
	OrderID        string    `json:"order_id"`
	Mobile         string    `json:"mobile"`
	Status         string    `json:"status"`
	ExpirationTime time.Time `json:"expiration_time"`
	WebURL         string    `json:"web_url"`
	UniqueID       string    `json:"unique_id"`
	PaidAmount     string    `json:"paid_amount"`
}

// chargeResponseQRIS represents response for payment char use QRIS
type chargeResponseQRIS struct {
	PaymentID      string    `json:"payment_id"`
	OrderID        string    `json:"order_id"`
	Status         string    `json:"status"`
	ExpirationTime time.Time `json:"expiration_time"`
	CreationTime   time.Time `json:"creation_time"`
	QRString       string    `json:"qr_string"`
	UniqueID       string    `json:"unique_id"`
	Metadata       Metadata  `json:"metadata"`
	Amount         string    `json:"amount"`
	QRCode         string    `json:"qr_code"`
}

// Metadata is part of ChargeResponseQRIS & ChargeResponseCard
type Metadata struct {
	MerchantName string `json:"merchant_name"`
	MerchantID   string `json:"merchant_id"`
}

// chargeResponseCard represents response for payment charge use card
type chargeResponseCard struct {
	PaymentID    string   `json:"payment_id"`
	OrderID      string   `json:"order_id"`
	PaymentRefID string   `json:"payment_ref_id"`
	TokenID      string   `json:"token_id"`
	Status       string   `json:"status"`
	PaidAmount   string   `json:"paid_amount"`
	Metadata     Metadata `json:"metadata"`
	CheckoutURL  string   `json:"checkout_url"`
}

// chargeResponseBNPL represents response for payment charge use Buy Now PayLater
type chargeResponseBNPL struct {
	PaymentID    string   `json:"payment_id"`
	OrderID      string   `json:"order_id"`
	PaymentRefID string   `json:"payment_ref_id"`
	RedirectURL  string   `json:"redirect_url"`
	PaidAmount   string   `json:"paid_amount"`
	Metadata     Metadata `json:"metadata"`
}
