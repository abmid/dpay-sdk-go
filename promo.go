/*
 * File Created: Sunday, 3rd September 2023 10:43:37 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

import "time"

/*
Payloads
*/

// PromoPayload use for Create & Update Promo API
type PromoPayload struct {
	Type               string       `json:"type"`
	Label              string       `json:"label"`
	Currency           string       `json:"currency"`
	PromoDetails       PromoDetails `json:"promo_details"`
	DiscountType       string       `json:"discount_type"`
	Discount           string       `json:"discount"`
	MinOrderAmount     string       `json:"min_order_amount"`
	MaxDiscountAmount  string       `json:"max_discount_amount"`
	StartsAt           time.Time    `json:"starts_at"`
	EndsAt             time.Time    `json:"ends_at"`
	PromoType          string       `json:"promo_type"`
	Description        string       `json:"description"`
	SubType            string       `json:"sub_type"`
	LimitType          string       `json:"limit_type"`
	LimitValue         string       `json:"limit_value"`
	PriceDeductionType string       `json:"price_deduction_type"`
	Code               string       `json:"code"`
}

// PromoDetails is part of PromoPayload
type PromoDetails struct {
	BinList   []int    `json:"bin_list"`
	BankCodes []string `json:"bank_codes"`
}

/*
Options
*/
