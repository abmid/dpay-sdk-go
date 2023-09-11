/*
 * File Created: Tuesday, 5th September 2023 11:13:19 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package payment

import (
	"context"
	"net/http"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/common"
)

type Client struct {
	ServerKey string
	Api       common.Api
}

const (
	PATH_PAYMENT        = durianpay.DURIANPAY_URL + "/v1/payments"
	PATH_PAYMENT_CHARGE = PATH_PAYMENT + "/charge"
)

// ChargeVA returns a response from Payment Charge API.
//
//	[Doc Payment Charge API VA]: https://durianpay.id/docs/api/payments/charge/
func (c *Client) ChargeVA(ctx context.Context, payload durianpay.PaymentChargeVAPayload) (*ChargeVA, *durianpay.Error) {
	reqPayload := chargePayload{
		Type:          "VA",
		Request:       payload,
		SandboxOption: payload.SandboxOption,
	}

	res := struct {
		Data ChargeVA `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_PAYMENT_CHARGE, nil, reqPayload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// ChargeBNPL returns a response from Payment Charge API for type Buy Now PayLater
//
//	[Doc Payment Charge API VA]: https://durianpay.id/docs/api/payments/charge/
func (c *Client) ChargeBNPL(ctx context.Context, payload durianpay.PaymentChargeBNPLPayload) (*ChargeBNPL, *durianpay.Error) {
	reqPayload := chargePayload{
		Type:          "BNPL",
		Request:       payload,
		SandboxOption: payload.SandboxOption,
	}

	res := struct {
		Data ChargeBNPL `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_PAYMENT_CHARGE, nil, reqPayload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
