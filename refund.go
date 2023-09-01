/*
 * File Created: Wednesday, 30th August 2023 11:47:06 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

/*
Payloads
*/

// RefundPayload is struct for payload Create Refund API
type RefundPayload struct {
	RefID         string `json:"ref_id"`
	PaymentID     string `json:"payment_id"`
	Amount        string `json:"amount"`
	UseRefundLink bool   `json:"use_refund_link"`
	Notes         string `json:"notes"`
}

/*
Options
*/

// RefundFetchOption is parameter for Refund Fetch API
type RefundFetchOption struct {
	From  string `url:"from"`
	To    string `url:"to"`
	Skip  string `url:"skip"`
	Limit string `url:"limit"`
}
