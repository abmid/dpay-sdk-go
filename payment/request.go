/*
 * File Created: Saturday, 9th September 2023 1:17:58 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package payment

import durianpay "github.com/abmid/dpay-sdk-go"

// chargePayload use for multiple type Payment Charge payload
type chargePayload struct {
	Type          string                          `json:"type"`
	Request       interface{}                     `json:"request"`
	SandboxOption *durianpay.PaymentSandboxOption `json:"sandbox_options"`
}
