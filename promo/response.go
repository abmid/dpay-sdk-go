/*
 * File Created: Sunday, 3rd September 2023 10:52:49 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package promo

import "time"

// Promo use for response Create, Update, Fetch Promos & Fetch By ID API
type Promo struct {
	Currency           string       `json:"currency"`
	Label              string       `json:"label"`
	Description        string       `json:"description"`
	MinOrderAmount     string       `json:"min_order_amount"`
	MaxDiscountAmount  string       `json:"max_discount_amount"`
	StartsAt           time.Time    `json:"starts_at"`
	EndsAt             time.Time    `json:"ends_at"`
	Discount           string       `json:"discount"`
	DiscountType       string       `json:"discount_type"`
	Type               string       `json:"type"`
	PromoDetails       PromoDetails `json:"promo_details"`
	SubType            string       `json:"sub_type"`
	LimitType          string       `json:"limit_type"`
	LimitValue         string       `json:"limit_value"`
	PriceDeductionType string       `json:"price_deduction_type"`
	Status             string       `json:"status"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
	IsLive             bool         `json:"is_live"`
	PromoUsage         string       `json:"promo_usage"`
	ID                 string       `json:"id"`
}

// PromoDetails is part of Promo
type PromoDetails struct {
	PromoID   string   `json:"promo_id"`
	BinList   []int    `json:"bin_list"`
	BankCodes []string `json:"bank_codes"`
}
