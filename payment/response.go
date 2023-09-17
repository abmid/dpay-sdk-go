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
	PaymentID      string            `json:"payment_id"`
	OrderID        string            `json:"order_id"`
	Status         string            `json:"status"`
	ExpirationTime time.Time         `json:"expiration_time"`
	CreationTime   time.Time         `json:"creation_time"`
	QRString       string            `json:"qr_string"`
	UniqueID       string            `json:"unique_id"`
	Metadata       map[string]string `json:"metadata"`
	Amount         string            `json:"amount"`
	QRCode         string            `json:"qr_code"`
}

// chargeResponseCard represents response for payment charge use card
type chargeResponseCard struct {
	PaymentID    string            `json:"payment_id"`
	OrderID      string            `json:"order_id"`
	PaymentRefID string            `json:"payment_ref_id"`
	TokenID      string            `json:"token_id"`
	Status       string            `json:"status"`
	PaidAmount   string            `json:"paid_amount"`
	Metadata     map[string]string `json:"metadata"`
	CheckoutURL  string            `json:"checkout_url"`
}

// chargeResponseBNPL represents response for payment charge use Buy Now PayLater
type chargeResponseBNPL struct {
	PaymentID    string            `json:"payment_id"`
	OrderID      string            `json:"order_id"`
	PaymentRefID string            `json:"payment_ref_id"`
	RedirectURL  string            `json:"redirect_url"`
	PaidAmount   string            `json:"paid_amount"`
	Metadata     map[string]string `json:"metadata"`
}

// PaymentCustomer is part of Payment for attribute Customer
type PaymentCustomer struct {
	ID            string    `json:"id"`
	CustomerRefID string    `json:"customer_ref_id"`
	Email         string    `json:"email"`
	Mobile        string    `json:"mobile"`
	GivenName     string    `json:"given_name"`
	MiddleName    string    `json:"middle_name"`
	SurName       string    `json:"sur_name"`
	AddressLine1  string    `json:"address_line_1"`
	AddressLine2  string    `json:"address_line_2"`
	City          string    `json:"city"`
	Region        string    `json:"region"`
	Country       string    `json:"country"`
	PostalCode    string    `json:"postal_code"`
	CreatedAt     time.Time `json:"created_at"`
}

// PaymentOrder is part of Payment for attribute Order
type PaymentOrder struct {
	ID           string    `json:"id"`
	MerchantID   string    `json:"merchant_id"`
	CustomerID   string    `json:"customer_id"`
	OrderRefID   string    `json:"order_ref_id"`
	OrderDsRefID string    `json:"order_ds_ref_id"`
	Amount       string    `json:"amount"`
	Currency     string    `json:"currency"`
	Status       string    `json:"status"`
	IsLive       bool      `json:"is_live"`
	CreatedAt    time.Time `json:"created_at"`
}

// Payments is part of FetchPayments for attribute payments.
type Payments struct {
	ID                 string            `json:"id"`
	OrderID            string            `json:"order_id"`
	PaymentRefID       string            `json:"payment_ref_id"`
	SettlementID       string            `json:"settlement_id"`
	PaymentDsRefID     string            `json:"payment_ds_ref_id"`
	Amount             string            `json:"amount"`
	Status             string            `json:"status"`
	IsLive             bool              `json:"is_live"`
	ExpirationDate     time.Time         `json:"expiration_date"`
	PaymentDetailsType string            `json:"payment_details_type"`
	MethodID           string            `json:"method_id"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
	Metadata           map[string]string `json:"metadata"`
	RetryCount         uint16            `json:"retry_count"`
	Discount           string            `json:"discount"`
	PaidAmount         string            `json:"paid_amount"`
	ProviderID         string            `json:"provider_id"`
	TotalFee           string            `json:"total_fee"`
	PromoID            string            `json:"promo_id"`
	ShippingFee        string            `json:"shipping_fee"`
	DsErrorMetadata    map[string]string `json:"ds_error_metadata"`
	CustomerID         string            `json:"customer_id"`
	GivenName          string            `json:"given_name"`
	Email              string            `json:"email"`
	OrderRefID         string            `json:"order_ref_id"`
	Currency           string            `json:"currency"`
	FailureReason      map[string]string `json:"failure_reason"`
	SettlementStatus   string            `json:"settlement_status"`
}

// Payment represents for response Payment Fetch By ID API.
type Payment struct {
	ID                 string            `json:"id"`
	OrderID            string            `json:"order_id"`
	MerchantID         string            `json:"merchant_id"`
	PaymentRefID       string            `json:"payment_ref_id"`
	PaymentDsRefID     string            `json:"payment_ds_ref_id"`
	SettlementID       string            `json:"settlement_id"`
	Amount             string            `json:"amount"`
	Status             string            `json:"status"`
	IsLive             bool              `json:"is_live"`
	ExpirationDate     time.Time         `json:"expiration_date"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
	Customer           PaymentCustomer   `json:"customer"`
	Order              PaymentOrder      `json:"order"`
	Metadata           map[string]string `json:"metadata"`
	PaymentDetailsType string            `json:"payment_details_type"`
	MethodID           string            `json:"method_id"`
	Discount           string            `json:"discount"`
	PromoID            string            `json:"promo_id"`
	PaidAmount         string            `json:"paid_amount"`
	ShippingFee        string            `json:"shipping_fee"`
	FailureReason      map[string]string `json:"failure_reason"`
}

// FetchPayments represents for response Payment Fetch API.
type FetchPayments struct {
	Payments []Payments `json:"payments"`
	Total    int        `json:"total"`
}

// CheckPaymentStatus represents for response Check Payments Status API.
type CheckPaymentStatus struct {
	Status      string `json:"status"`
	IsCompleted bool   `json:"is_completed"`
	Signature   string `json:"signature"`
	ErrorCode   string `json:"error_code"`
}

// Capture represents for response Payment Capture API.
type Capture struct {
	PaymentID           string            `json:"payment_id"`
	OrderID             string            `json:"order_id"`
	PreauthorizedAmount string            `json:"preauthorized_amount"`
	AccountID           string            `json:"account_id"`
	PaidAmount          string            `json:"paid_amount"`
	Status              string            `json:"status"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	Metadata            map[string]string `json:"metadata"`
}

// Cancel represents for response Cancel Payment API.
type Cancel struct {
	ID                 string            `json:"id"`
	OrderID            string            `json:"order_id"`
	PaymentRefID       string            `json:"payment_ref_id"`
	SettlementID       string            `json:"settlement_id"`
	PaymentDsRefID     string            `json:"payment_ds_ref_id"`
	Amount             string            `json:"amount"`
	Status             string            `json:"status"`
	IsLive             bool              `json:"is_live"`
	ExpirationDate     time.Time         `json:"expiration_date"`
	PaymentDetailsType string            `json:"payment_details_type"`
	MethodID           string            `json:"method_id"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
	Metadata           map[string]string `json:"metadata"`
	RetryCount         uint16            `json:"retry_count"`
	Discount           string            `json:"discount"`
	PaidAmount         string            `json:"paid_amount"`
	ProviderID         string            `json:"provider_id"`
	TotalFee           string            `json:"total_fee"`
	PromoID            string            `json:"promo_id"`
	ShippingFee        string            `json:"shipping_fee"`
	DsErrorMetadata    map[string]string `json:"ds_error_metadata"`
	FailureReason      map[string]string `json:"failure_reason"`
	SettlementStatus   string            `json:"settlement_status"`
}

// MDRFee is part of MDRFeesCalculation
type MDRFee struct {
	ActualAmount float32 `json:"actual_amount"`
	Fees         float32 `json:"fees"`
	TotalAmount  float32 `json:"total_amount"`
}

// MDRFeesCalculation represents for response MDR Fees Calculation API
type MDRFeesCalculation struct {
	GOPAY     MDRFee `json:"GOPAY"`
	DD_CIMB   MDRFee `json:"DD_CIMB"`
	QRIS      MDRFee `json:"QRIS"`
	SHOPEEPAY MDRFee `json:"SHOPEEPAY"`
	OTHERS    MDRFee `json:"OTHERS"`
	INDOMARET MDRFee `json:"INDOMARET"`
	BRI       MDRFee `json:"BRI"`
	ALFAMART  MDRFee `json:"ALFAMART"`
	JENIUSPAY MDRFee `json:"JENIUSPAY"`
	BCA       MDRFee `json:"BCA"`
	DD_BRI    MDRFee `json:"DD_BRI"`
	CARD      MDRFee `json:"CARD"`
	BNI       MDRFee `json:"BNI"`
	OVO       MDRFee `json:"OVO"`
	DANA      MDRFee `json:"DANA"`
	MANDIRI   MDRFee `json:"MANDIRI"`
	PERMATA   MDRFee `json:"PERMATA"`
	LINKAJA   MDRFee `json:"LINKAJA"`
	DANAMON   MDRFee `json:"DANAMON"`
	CIMB      MDRFee `json:"CIMB"`
	SYARIAH   MDRFee `json:"SYARIAH"`
}
