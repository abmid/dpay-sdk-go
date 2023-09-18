/*
 * File Created: Thursday, 24th August 2023 6:36:34 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package order

import (
	"time"

	durianpay "github.com/abmid/dpay-sdk-go"
)

// Create is struct for response Create Order or Create Payment Link API.
// For case Create Payment Link API, attribute PaymentLinkUrl will be filled
type Create struct {
	ID             string                  `json:"id"`
	CustomerID     string                  `json:"customer_id"`
	OrderRefID     string                  `json:"order_ref_id"`
	OrderDsRefID   string                  `json:"order_ds_ref_id"`
	Amount         string                  `json:"amount"`
	PaymentOption  string                  `json:"payment_option"`
	PendingAmount  string                  `json:"pending_amount"`
	Currency       string                  `json:"currency"`
	Status         string                  `json:"status"`
	IsLive         bool                    `json:"is_live"`
	CreatedAt      time.Time               `json:"created_at"`
	UpdatedAt      time.Time               `json:"updated_at"`
	MetaData       durianpay.OrderMetadata `json:"metadata"`
	Items          []durianpay.OrderItem   `json:"items"`
	AccessToken    string                  `json:"access_token"`
	ExpireTime     time.Time               `json:"expire_time"`
	ExpiryDate     time.Time               `json:"expiry_date"`
	PaymentLinkUrl string                  `json:"payment_link_url"`
	AddressID      uint32                  `json:"address_id"`
	Fees           string                  `json:"fees"`
	ShippingFee    string                  `json:"shipping_fee"`
	AdminFeeMethod string                  `json:"admin_fee_method"`
}

// FetchOrders is struct for response Fetch Orders API
type FetchOrders struct {
	Orders []Orders `json:"orders"`
	Count  uint     `json:"count"`
}

// Orders is part of FetchOrders for attribute Orders
type Orders struct {
	ID                    string    `json:"id"`
	CustomerID            string    `json:"customer_id"`
	OrderRefID            string    `json:"order_ref_id"`
	OrderDsRefID          string    `json:"order_ds_ref_id"`
	Amount                string    `json:"amount"`
	Currency              string    `json:"currency"`
	Status                string    `json:"status"`
	IsLive                bool      `json:"is_live"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	ExpiryDate            time.Time `json:"expiry_date"`
	GivenName             string    `json:"given_name"`
	SurName               string    `json:"sur_name"`
	Email                 string    `json:"email"`
	Mobile                string    `json:"mobile"`
	PaymentOption         string    `json:"payment_option"`
	PaymentID             string    `json:"payment_id"`
	PaymentDetailsType    string    `json:"payment_details_type"`
	PaymentStatus         string    `json:"payment_status"`
	PaymentDate           time.Time `json:"payment_date"`
	Description           string    `json:"description"`
	PaymentLinkUrl        string    `json:"payment_link_url"`
	IsNotificationEnabled bool      `json:"is_notification_enabled"`
	EmailSubject          string    `json:"email_subject"`
	EmailContent          string    `json:"email_content"`
	PaymentMethodID       string    `json:"payment_method_id"`
}

// FetchOrder is struct for response Fetch Order API
type FetchOrder struct {
	ID                    string                  `json:"id"`
	CustomerID            string                  `json:"customer_id"`
	OrderRefID            string                  `json:"order_ref_id"`
	OrderDsRefID          string                  `json:"order_ds_ref_id"`
	Amount                string                  `json:"amount"`
	PaymentOption         string                  `json:"payment_option"`
	PendingAmount         string                  `json:"pending_amount"`
	Currency              string                  `json:"currency"`
	Status                string                  `json:"status"`
	IsLive                bool                    `json:"is_live"`
	CreatedAt             time.Time               `json:"created_at"`
	UpdatedAt             time.Time               `json:"updated_at"`
	Metadata              durianpay.OrderMetadata `json:"metadata"`
	Items                 []durianpay.OrderItem   `json:"items"`
	ExpiryDate            time.Time               `json:"expiry_date"`
	Description           string                  `json:"description"`
	PaymentLinkUrl        string                  `json:"payment_link_url"`
	IsNotificationEnabled bool                    `json:"is_notification_enabled"`
	EmailSubject          string                  `json:"email_subject"`
	EmailContent          string                  `json:"email_content"`
	Fees                  string                  `json:"fees"`
	ShippingFee           string                  `json:"shipping_fee"`
	AdminFeeMethod        string                  `json:"admin_fee_method"`
	Customer              durianpay.Customer      `json:"customer"` // Will be filled if use query expand=customer
	Payments              []Payment               `json:"payments"` // Will be filled if use query expand=payments
}

// Payment is part of FetchOrder for attribute Payments
type Payment struct {
	ID                 string                  `json:"id"`
	OrderID            string                  `json:"order_id"`
	PaymentRefID       string                  `json:"payment_ref_id"`
	SettlementID       string                  `json:"settlement_id"`
	PaymentDsRefID     string                  `json:"payment_ds_ref_id"`
	Amount             string                  `json:"amount"`
	Status             string                  `json:"status"`
	IsLive             bool                    `json:"is_live"`
	ExpirationDate     time.Time               `json:"expiration_date"`
	PaymentDetailsType string                  `json:"payment_details_type"`
	MethodID           string                  `json:"method_id"`
	CreatedAt          time.Time               `json:"created_at"`
	UpdatedAt          time.Time               `json:"updated_at"`
	Metadata           durianpay.OrderMetadata `json:"metadata"`
	RetryCount         uint16                  `json:"retry_count"`
	Discount           string                  `json:"discount"`
	PaidAmount         string                  `json:"paid_amount"`
	ProvideID          string                  `json:"provider_id"`
	TotalFee           string                  `json:"total_fee"`
	PromoID            string                  `json:"promo_id"`
	ShippingFee        string                  `json:"shipping_fee"`
	SettlementStatus   string                  `json:"settlement_status"`
	// TODO: ds_error_metadata, failure_reson
}
