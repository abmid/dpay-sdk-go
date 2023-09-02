/*
 * File Created: Saturday, 2nd September 2023 3:34:54 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package settlement

import "time"

// Settlementuse for response Settlements By ID API and part of Fetch Settlements
type Settlement struct {
	ID                     string    `json:"id"`
	SettlementAmount       string    `json:"settlement_amount"`
	Status                 string    `json:"status"`
	Fee                    string    `json:"fee"`
	CreatedAt              time.Time `json:"created_at"`
	SettledAt              time.Time `json:"settled_at"`
	PromoAmount            string    `json:"promo_amount"`
	TotalTransactionAmount string    `json:"total_transaction_amount"` // Special case for Settlements Fetch API
	Currency               string    `json:"currency"`                 // Special case for Settlements Fetch API
}

// SettlementDetail use for response Status By Payment ID API and Settlements Details Fetch API
type SettlementDetail struct {
	SettlementID       string    `json:"settlement_id"`
	PaymentID          string    `json:"payment_id"`
	PaymentReference   string    `json:"payment_reference"`
	OrderID            string    `json:"order_id"`
	OrderReference     string    `json:"order_reference"`
	Status             string    `json:"status"`
	Currency           string    `json:"currency"`
	SettlementAmount   string    `json:"settlement_amount"`
	TotalSettlementFee string    `json:"total_settlement_fee"`
	PaymentDiscount    string    `json:"payment_discount"`
	SettledAt          time.Time `json:"settled_at"`
	Group              string    `json:"group"`
	PaymentAmount      string    `json:"payment_amount"`
	PaymentDate        time.Time `json:"payment_date"`
	TransactionAmount  string    `json:"transaction_amount"`
	PaymentDetailsType string    `json:"payment_details_type"` // Special case for Settlement Details API
	PaymentMethodID    string    `json:"payment_method_id"`    // Special case for Settlement Details API
	PaymentChannel     string    `json:"payment_channel"`      // Special case for Status By Payment ID API
	PaymentSubchannel  string    `json:"payment_subchannel"`   // Special case for Status By Payment ID API
}

// FetchSettlements response for Settlements Fetch API
type FetchSettlements struct {
	TotalCount       uint32       `json:"total_count"`
	SettlementDetail []Settlement `json:"settlement_detail"`
}

// FetchDetails response for Settlements Details Fetch API
type FetchDetails struct {
	SettlementCount  uint32             `json:"settlement_count"`
	TransactionCount uint32             `json:"transaction_count"`
	SettlementDetail []SettlementDetail `json:"settlement_detail"`
}
