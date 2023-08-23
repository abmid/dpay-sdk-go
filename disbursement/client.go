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
	PATH_DISBURSEMENT                   = "/v1/disbursements"
	PATH_DISBURSEMENT_VALIDATE          = PATH_DISBURSEMENT + "/validate"
	PATH_DISBURSEMENT_SUBMIT            = PATH_DISBURSEMENT + "/submit"
	PATH_DISBURSEMENT_APPROVE           = PATH_DISBURSEMENT + "/:id/approve"
	PATH_DISBURSEMENT_FETCH_ITEMS_BY_ID = PATH_DISBURSEMENT + "/:id/items"
	PATH_DISBURSEMENT_FETCH_BY_ID       = PATH_DISBURSEMENT + "/:id"
	PATH_DISBURSEMENT_DELETE            = PATH_DISBURSEMENT + "/:id"
	PATH_DISBURSEMENT_FETCH_BANKS       = PATH_DISBURSEMENT + "/banks"
)

// Validate returns a response from Validate Disbursement API.
// Validate disbursement can be used to fetch the bank account and account number validation
//
//	[Doc Validate Disbursement API]: https://durianpay.id/docs/api/disbursements/validate/
func (c *Client) Validate(ctx context.Context, payload durianpay.DisbursementValidatePayload) (*durianpay.DisbursementValidate, *durianpay.Error) {
	res := &durianpay.DisbursementValidate{}
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, "")

	err := c.Api.Req(ctx, http.MethodPost, durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_VALIDATE, nil, payload, headers, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Submit returns a response from Submit Disbursement API.
// Options about skip_validation & force_disburse you can input in durianpay.DisbursementOption
//
//	[Doc Submit Disbursement API]: https://durianpay.id/docs/api/disbursements/submit/
func (c *Client) Submit(ctx context.Context, payload durianpay.DisbursementPayload, opt *durianpay.DisbursementOption) (*durianpay.Disbursement, *durianpay.Error) {
	res := &durianpay.Disbursement{}
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, payload.IdempotencyKey)

	err := c.Api.Req(ctx, http.MethodPost, durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_SUBMIT, opt, payload, headers, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Approve returns a response from Approve Disbursement API.
// Options about ignore_invalid you can input in durianpay.DisbursementApproveOption
//
//	[Doc Approve Disbursement API]: https://durianpay.id/docs/api/disbursements/approve/
func (c *Client) Approve(ctx context.Context, payload durianpay.DisbursementApprovePayload, opt *durianpay.DisbursementApproveOption) (*durianpay.Disbursement, *durianpay.Error) {
	res := &durianpay.Disbursement{}
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, "")

	url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_APPROVE
	url = strings.ReplaceAll(url, ":id", payload.ID)

	err := c.Api.Req(ctx, http.MethodPost, url, opt, payload, headers, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// FetchItemsByID returns a response from Fetch Disbursement Items By ID API.
// Options about skip & limit pagination can be fill in durianpay.DisbursementFetchItemsOption
//
//	[Doc Fetch Disbursement Items by ID]: https://durianpay.id/docs/api/disbursements/fetch-items/
func (c *Client) FetchItemsByID(ctx context.Context, ID string, opt *durianpay.DisbursementFetchItemsOption) (*durianpay.DisbursementItem, *durianpay.Error) {
	url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_FETCH_ITEMS_BY_ID
	url = strings.ReplaceAll(url, ":id", ID)

	tempRes := struct {
		Data durianpay.DisbursementItem `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, opt, nil, nil, &tempRes)
	if err != nil {
		return nil, err
	}

	return &tempRes.Data, nil
}

// FetchByID returns a response from Fetch Disbursement by ID API.
//
//	[Docs Fetch Disbursement]: https://durianpay.id/docs/api/disbursements/fetch-one/
func (c *Client) FetchByID(ctx context.Context, ID string) (*durianpay.DisbursementData, *durianpay.Error) {
	url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_FETCH_BY_ID
	url = strings.ReplaceAll(url, ":id", ID)

	tempRes := struct {
		Data durianpay.DisbursementData `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &tempRes)
	if err != nil {
		return nil, err
	}

	return &tempRes.Data, nil
}

// Delete returns a response from Delete Disbursement by ID API
//
//	[Docs Delete Disbursement]: https://durianpay.id/docs/api/disbursements/delete/
func (c *Client) Delete(ctx context.Context, ID string) (string, *durianpay.Error) {
	url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_DELETE
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
//	[Docs Delete Disbursement]: https://durianpay.id/docs/api/disbursements/fetch-banks/
func (c *Client) FetchBanks(ctx context.Context) ([]durianpay.DisbursementBank, *durianpay.Error) {
	tempRes := struct {
		Data []durianpay.DisbursementBank `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_FETCH_BANKS, nil, nil, nil, &tempRes)
	if err != nil {
		return tempRes.Data, err
	}

	return tempRes.Data, nil
}
