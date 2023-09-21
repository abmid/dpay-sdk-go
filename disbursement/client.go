/*
 * File Created: Friday, 28th July 2023 6:42:39 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package disbursement

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
	pathDisbursement   = "/v1/disbursements"
	pathValidate       = pathDisbursement + "/validate"
	pathSubmit         = pathDisbursement + "/submit"
	pathApprove        = pathDisbursement + "/:id/approve"
	pathFetchItemsByID = pathDisbursement + "/:id/items"
	pathFetchByID      = pathDisbursement + "/:id"
	pathDelete         = pathDisbursement + "/:id"
	pathFetchBanks     = pathDisbursement + "/banks"
	pathTopupAmount    = pathDisbursement + "/topup"
	pathFetchBalance   = pathDisbursement + "/topup/balance"
)

// Validate returns a response from Validate Disbursement API.
// Validate disbursement can be used to fetch the bank account and account number validation
//
//	[Doc Validate Disbursement API]: https://durianpay.id/docs/api/disbursements/validate/
func (c *Client) Validate(ctx context.Context, payload durianpay.DisbursementValidatePayload) (*DisbursementValidate, *durianpay.Error) {
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, "")

	res := struct {
		Data DisbursementValidate `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, durianpay.DurianpayURL+pathValidate, nil, payload, headers, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Submit returns a response from Submit Disbursement API.
// Options about skip_validation & force_disburse you can input in durianpay.DisbursementOption
//
//	[Doc Submit Disbursement API]: https://durianpay.id/docs/api/disbursements/submit/
func (c *Client) Submit(ctx context.Context, payload durianpay.DisbursementPayload, opt *durianpay.DisbursementOption) (*Disbursement, *durianpay.Error) {
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, payload.IdempotencyKey)

	res := struct {
		Data Disbursement `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, durianpay.DurianpayURL+pathSubmit, opt, payload, headers, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Approve returns a response from Approve Disbursement API.
// Options about ignore_invalid you can input in durianpay.DisbursementApproveOption
//
//	[Doc Approve Disbursement API]: https://durianpay.id/docs/api/disbursements/approve/
func (c *Client) Approve(ctx context.Context, payload durianpay.DisbursementApprovePayload, opt *durianpay.DisbursementApproveOption) (*Disbursement, *durianpay.Error) {
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, "")

	res := struct {
		Data Disbursement `json:"data"`
	}{}

	url := durianpay.DurianpayURL + pathApprove
	url = strings.ReplaceAll(url, ":id", payload.ID)

	err := c.Api.Req(ctx, http.MethodPost, url, opt, payload, headers, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchItemsByID returns a response from Fetch Disbursement Items By ID API.
// Options about skip & limit pagination can be fill in durianpay.DisbursementFetchItemsOption
//
//	[Doc Fetch Disbursement Items by ID]: https://durianpay.id/docs/api/disbursements/fetch-items/
func (c *Client) FetchItemsByID(ctx context.Context, ID string, opt *durianpay.DisbursementFetchItemsOption) (*DisbursementItem, *durianpay.Error) {
	url := durianpay.DurianpayURL + pathFetchItemsByID
	url = strings.ReplaceAll(url, ":id", ID)

	res := struct {
		Data DisbursementItem `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchByID returns a response from Fetch Disbursement by ID API.
//
//	[Docs Fetch Disbursement]: https://durianpay.id/docs/api/disbursements/fetch-one/
func (c *Client) FetchByID(ctx context.Context, ID string) (*Disbursement, *durianpay.Error) {
	url := durianpay.DurianpayURL + pathFetchByID
	url = strings.ReplaceAll(url, ":id", ID)

	res := struct {
		Data Disbursement `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Delete returns a response from Delete Disbursement by ID API
//
//	[Docs Delete Disbursement]: https://durianpay.id/docs/api/disbursements/delete/
func (c *Client) Delete(ctx context.Context, ID string) (string, *durianpay.Error) {
	url := durianpay.DurianpayURL + pathDelete
	url = strings.ReplaceAll(url, ":id", ID)

	tempRes := struct {
		Data string `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodDelete, url, nil, nil, nil, &tempRes)
	if err != nil {
		return "", err
	}

	return tempRes.Data, nil
}

// Delete returns a response from Fetch Bank List API
//
//	[Docs Fetch Banks]: https://durianpay.id/docs/api/disbursements/fetch-banks/
func (c *Client) FetchBanks(ctx context.Context) ([]DisbursementBank, *durianpay.Error) {
	tempRes := struct {
		Data []DisbursementBank `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, durianpay.DurianpayURL+pathFetchBanks, nil, nil, nil, &tempRes)
	if err != nil {
		return tempRes.Data, err
	}

	return tempRes.Data, nil
}

// TopupAmount returns a response from Topup Amount API
//
//	[Docs Topup Amount]: https://durianpay.id/docs/api/disbursements/topup/
func (c *Client) TopupAmount(ctx context.Context, payload durianpay.DisbursementTopupPayload) (*DisbursementTopup, *durianpay.Error) {
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, "")

	tempRes := struct {
		Data DisbursementTopup `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, durianpay.DurianpayURL+pathTopupAmount, nil, payload, headers, &tempRes)
	if err != nil {
		return nil, err
	}

	return &tempRes.Data, nil
}

// FetchBalance returns a response from Fetch Durianpay Balance API
//
//	[Docs Fetch Durianpay Balance]: https://durianpay.id/docs/api/disbursements/balance/
func (c *Client) FetchBalance(ctx context.Context) (*int, *durianpay.Error) {
	tempRes := struct {
		Data struct {
			Balance int `json:"balance"`
		} `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, durianpay.DurianpayURL+pathFetchBalance, nil, nil, nil, &tempRes)
	if err != nil {
		return nil, err
	}

	return &tempRes.Data.Balance, nil
}
