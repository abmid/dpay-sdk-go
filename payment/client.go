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
	"strings"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/common"
)

type Client struct {
	ServerKey string
	Api       common.Api
}

const (
	PATH_PAYMENT              = durianpay.DURIANPAY_URL + "/v1/payments"
	PATH_PAYMENT_CHARGE       = PATH_PAYMENT + "/charge"
	PATH_PAYMENT_FETCH_BY_ID  = PATH_PAYMENT + "/:id"
	PATH_PAYMENT_CHECK_STATUS = PATH_PAYMENT + "/:id/status"
	PATH_PAYMENT_VERIFY       = PATH_PAYMENT + "/:id/verify"
	PATH_PAYMENT_CAPTURE      = PATH_PAYMENT + "/:id/capture"
	PATH_PAYMENT_CANCEL       = PATH_PAYMENT + "/:id/cancel"
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

// ChargeQRIS returns a response from Payment Charge API for QRIS type
//
//	[Doc Payment Charge API Online Bank]: https://durianpay.id/docs/api/payments/charge/
func (c *Client) ChargeQRIS(ctx context.Context, payload durianpay.PaymentChargeQRISPayload) (*ChargeQRIS, *durianpay.Error) {
	reqPayload := chargePayload{
		Type:    "QRIS",
		Request: payload,
	}

	res := struct {
		Data ChargeQRIS `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_PAYMENT_CHARGE, nil, reqPayload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// ChargeCard returns a response from Payment Charge API for CARD type
//
//	[Doc Payment Charge API Online Bank]: https://durianpay.id/docs/api/payments/charge/
func (c *Client) ChargeCard(ctx context.Context, payload durianpay.PaymentChargeCardPayload) (*ChargeCard, *durianpay.Error) {
	reqPayload := chargePayload{
		Type:    "CARD",
		Request: payload,
	}

	res := struct {
		Data ChargeCard `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_PAYMENT_CHARGE, nil, reqPayload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchPayments returns a response from Payment Fetch API
//
//	[Doc Payment Fetch API]: https://durianpay.id/docs/api/payments/fetch/
func (c *Client) FetchPayments(ctx context.Context, opt durianpay.PaymentFetchOption) (*FetchPayments, *durianpay.Error) {
	res := struct {
		Data FetchPayments `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, PATH_PAYMENT_CHARGE, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchPaymentByID returns a response from Payment Fetch by ID API.
//
//	[Doc Payment Fetch by ID API]: https://durianpay.id/docs/api/payments/fetch-one/
func (c *Client) FetchPaymentByID(ctx context.Context, ID string, opt durianpay.PaymentFetchByIDOption) (*Payment, *durianpay.Error) {
	url := strings.ReplaceAll(PATH_PAYMENT_FETCH_BY_ID, ":id", ID)

	res := struct {
		Data Payment `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// CheckPaymentStatus returns a response from Check Payments Status API.
//
//	[Doc Check Payments Status API]: https://durianpay.id/docs/api/payments/status/
func (c *Client) CheckPaymentStatus(ctx context.Context, ID string) (*CheckPaymentStatus, *durianpay.Error) {
	url := strings.ReplaceAll(PATH_PAYMENT_CHECK_STATUS, ":id", ID)

	res := struct {
		Data CheckPaymentStatus `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Verify returns a response from Verify Payments Status API.
//
//	[Doc Verify Payments Status API]: https://durianpay.id/docs/api/payments/verify/
func (c *Client) Verify(ctx context.Context, ID string, payload durianpay.PaymentVerifyPayload) (bool, *durianpay.Error) {
	url := strings.ReplaceAll(PATH_PAYMENT_VERIFY, ":id", ID)

	res := struct {
		Data bool `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, url, nil, payload, nil, &res)
	if err != nil {
		return false, err
	}

	return res.Data, nil
}

// Capture returns a response from Payment Capture API
//
//	[Doc Payment Capture API]: https://durianpay.id/docs/api/payments/capture/
func (c *Client) Capture(ctx context.Context, ID string, payload durianpay.PaymentCapturePayload) (*Capture, *durianpay.Error) {
	url := strings.ReplaceAll(PATH_PAYMENT_CAPTURE, ":id", ID)

	res := struct {
		Data Capture `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, url, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Cancel returns a response from Cancel Payment API
//
//	[Doc Cancel Payment API]: https://durianpay.id/docs/api/payments/cancel/
func (c *Client) Cancel(ctx context.Context, ID string) (*Cancel, *durianpay.Error) {
	url := strings.ReplaceAll(PATH_PAYMENT_CANCEL, ":id", ID)

	res := struct {
		Data Cancel `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPut, url, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
