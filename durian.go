/*
 * File Created: Friday, 28th July 2023 5:26:55 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

const (
	DURIANPAY_URL = "https://api.durianpay.id"
)

type Customer struct {
	CustomerRefID string          `json:"customer_ref_id"`
	GivenName     string          `json:"given_name"`
	Email         string          `json:"email"`
	Mobile        string          `json:"mobile"`
	Address       CustomerAddress `json:"address"`
}

// CustomerAddress is part of Customer for attribute Address
type CustomerAddress struct {
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
