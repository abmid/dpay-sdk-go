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

// ChargeVA returns a response from Payment Charge API for Virtual Account type.
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

// ChargeBNPL returns a response from Payment Charge API for Buy Now PayLater type
//
//	[Doc Payment Charge API BNPL]: https://durianpay.id/docs/api/payments/charge/
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

// ChargeEwallet returns a response from Payment Charge API for E-Wallet type
//
//	[Doc Payment Charge API E-Wallet]: https://durianpay.id/docs/api/payments/charge/
func (c *Client) ChargeEwallet(ctx context.Context, payload durianpay.PaymentChargeEwalletPayload) (*ChargeEwallet, *durianpay.Error) {
	reqPayload := chargePayload{
		Type:    "EWALLET",
		Request: payload,
	}

	res := struct {
		Data ChargeEwallet `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_PAYMENT_CHARGE, nil, reqPayload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// ChargeRetailStore returns a response from Payment Charge API for Retail Store type (ex: Indomaret / Alfamaret)
//
//	[Doc Payment Charge API Retail Store]: https://durianpay.id/docs/api/payments/charge/
func (c *Client) ChargeRetailStore(ctx context.Context, payload durianpay.PaymentChargeRetailStorePayload) (*ChargeRetailStore, *durianpay.Error) {
	reqPayload := chargePayload{
		Type:    "RETAILSTORE",
		Request: payload,
	}

	res := struct {
		Data ChargeRetailStore `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_PAYMENT_CHARGE, nil, reqPayload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// ChargeOnlineBank returns a response from Payment Charge API for Online Banking type (ex: JeniusPay)
//
//	[Doc Payment Charge API Online Bank]: https://durianpay.id/docs/api/payments/charge/
func (c *Client) ChargeOnlineBank(ctx context.Context, payload durianpay.PaymentChargeOnlineBankingPayload) (*ChargeOnlineBank, *durianpay.Error) {
	reqPayload := chargePayload{
		Type:    "ONLINE_BANKING",
		Request: payload,
	}

	res := struct {
		Data ChargeOnlineBank `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_PAYMENT_CHARGE, nil, reqPayload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
