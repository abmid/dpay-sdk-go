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
	Amount        string        `json:"amount"`
	PaymentOption string        `json:"payment_option"`
	Currency      string        `json:"currency"`
	OrderRefID    string        `json:"order_ref_id"`
	Customer      OrderCustomer `json:"customer"`
	Items         []OrderItem   `json:"items"`
	Metadata      OrderMetadata `json:"metadata"`
	ExpiryDate    time.Time     `json:"expiry_date"`
}

// OrderCustomer is part of CreatePayload for attribute Customer
type OrderCustomer struct {
	CustomerRefID string               `json:"customer_ref_id"`
	GivenName     string               `json:"given_name"`
	Email         string               `json:"email"`
	Mobile        string               `json:"mobile"`
	Address       OrderCustomerAddress `json:"address"`
}

// OrderCustomerAddress is part of CreateCustomer for attribute Address
type OrderCustomerAddress struct {
	ReceiverName  string `json:"receiver_name"`
	ReceiverPhone string `json:"receiver_phone"`
	Label         string `json:"label"`
	AddressLine1  string `json:"address_line_1"`
	AddressLine2  string `json:"address_line_2"`
	City          string `json:"city"`
	Region        string `json:"Region"`
	Country       string `json:"country"`
	PostalCode    string `json:"postal_code"`
	Landmark      string `json:"landmark"`
}

// OrderItem is part of CreatePayload for attribute Items
type OrderItem struct {
	Name  string `json:"name"`
	Qty   uint16 `json:"qty"`
	Price string `json:"price"`
	Logo  string `json:"logo"`
}

// OrderMetadata is part of CreatePayload for attribute Metadata
type OrderMetadata struct {
	MyMetaKey       string `json:"my-meta-key"`
	SettlementGroup string `json:"SettlementGroup"`
}

// OrderPaymentLinkPayload is payload for requests Create Payment Link API
type OrderPaymentLinkPayload struct {
	Amount        string                   `json:"amount"`
	Currency      string                   `json:"id"`
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

// OrderFetchIDOption is parameter for requests Order Fetch By ID API
type OrderFetchIDOption struct {
	Expand string `url:"expand"`
}
