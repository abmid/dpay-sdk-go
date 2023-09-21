/*
 * File Created: Wednesday, 23rd August 2023 11:54:14 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

import "time"

/*
 Payloads
*/

// OrderPayload is payload for requests Create Orders API
type OrderPayload struct {
	Amount        string         `json:"amount"`
	PaymentOption string         `json:"payment_option"`
	Currency      string         `json:"currency"`
	OrderRefID    string         `json:"order_ref_id"`
	Customer      Customer       `json:"customer"`
	Items         []OrderItem    `json:"items"`
	Metadata      map[string]any `json:"metadata"`
	ExpiryDate    time.Time      `json:"expiry_date"`
}

// OrderItem is part of CreatePayload for attribute Items
type OrderItem struct {
	Name  string `json:"name"`
	Qty   uint16 `json:"qty"`
	Price string `json:"price"`
	Logo  string `json:"logo"`
}

// OrderPaymentLinkPayload is payload for requests Create Payment Link API
type OrderPaymentLinkPayload struct {
	Amount        string                   `json:"amount"`
	Currency      string                   `json:"currency"`
	OrderRefID    string                   `json:"order_ref_id"`
	IsPaymentLink bool                     `json:"is_payment_link"`
	Customer      OrderPaymentLinkCustomer `json:"customer"`
}

// OrderPaymentLinkCustomer is part of OrderPaymentLinkPayload
type OrderPaymentLinkCustomer struct {
	Email string `json:"email"`
}

/*
 Options
*/

// OrderFetchOption is parameter for requests Orders Fetch API
type OrderFetchOption struct {
	From  string `url:"form"`
	To    string `url:"to"`
	Skip  uint16 `url:"skip"`
	Limit uint16 `url:"limit"`
}

// OrderFetchByIDOption is parameter for requests Order Fetch By ID API
type OrderFetchByIDOption struct {
	Expand string `url:"expand"`
}
