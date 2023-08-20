/*
 * File Created: Friday, 28th July 2023 6:42:39 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package disbursement

import (
	"context"
	"encoding/json"
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
)

// Validate returns a response from Validate Disbursement API.
// Validate disbursement can be used to fetch the bank account and account number validation
//
//	[Doc Validate Disbursement API]: https://durianpay.id/docs/api/disbursements/validate/
func (c *Client) Validate(ctx context.Context, payload durianpay.DisbursementValidatePayload) (res *durianpay.DisbursementValidate, err *durianpay.Error) {
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, "")

	apiRes, err := c.Api.Req(ctx, http.MethodPost, durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_VALIDATE, nil, payload, headers)
	if err != nil {
		return nil, err
	}

	jsonErr := json.Unmarshal(apiRes, &res)
	if jsonErr != nil {
		return nil, durianpay.FromSDKError(jsonErr)
	}

	return res, err
}

// Submit returns a response from Submit Disbursement API.
// Options about skip_validation & force_disburse you can input in durianpay.DisbursementOption
//
//	[Doc Submit Disbursement API]: https://durianpay.id/docs/api/disbursements/submit/
func (c *Client) Submit(ctx context.Context, payload durianpay.DisbursementPayload, opt *durianpay.DisbursementOption) (res *durianpay.Disbursement, err *durianpay.Error) {
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, payload.IdempotencyKey)

	apiRes, err := c.Api.Req(ctx, http.MethodPost, durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_SUBMIT, opt, payload, headers)
	if err != nil {
		return nil, err
	}

	jsonErr := json.Unmarshal(apiRes, &res)
	if jsonErr != nil {
		return nil, durianpay.FromSDKError(jsonErr)
	}

	return res, err
}

// Approve returns a response from Approve Disbursement API.
// Options about ignore_invalid you can input in durianpay.DisbursementApproveOption
//
//	[Doc Approve Disbursement API]: https://durianpay.id/docs/api/disbursements/approve/
func (c *Client) Approve(ctx context.Context, payload durianpay.DisbursementApprovePayload, opt *durianpay.DisbursementApproveOption) (res *durianpay.Disbursement, err *durianpay.Error) {
	headers := common.HeaderIdempotencyKey(payload.XIdempotencyKey, "")

	url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_APPROVE
	url = strings.ReplaceAll(url, ":id", payload.ID)

	apiRes, err := c.Api.Req(ctx, http.MethodPost, url, opt, payload, headers)
	if err != nil {
		return nil, err
	}

	jsonErr := json.Unmarshal(apiRes, &res)
	if jsonErr != nil {
		return nil, durianpay.FromSDKError(jsonErr)
	}

	return res, err
}

// FetchItemsByID returns a response from Fetch Disbursement Items By ID API.
// Options about skip & limit pagination can be fill in durianpay.DisbursementFetchItemsOption
//
//	[Doc Fetch Disbursement Items by ID]: https://durianpay.id/docs/api/disbursements/fetch-items/
func (c *Client) FetchItemsByID(ctx context.Context, ID string, opt *durianpay.DisbursementFetchItemsOption) (res *durianpay.DisbursementItem, err *durianpay.Error) {
	url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_FETCH_ITEMS_BY_ID
	url = strings.ReplaceAll(url, ":id", ID)

	apiRes, err := c.Api.Req(ctx, http.MethodGet, url, opt, nil, nil)
	if err != nil {
		return nil, err
	}

	tempRes := struct {
		Data durianpay.DisbursementItem `json:"data"`
	}{}

	jsonErr := json.Unmarshal(apiRes, &tempRes)
	if jsonErr != nil {
		return nil, durianpay.FromSDKError(jsonErr)
	}

	res = &tempRes.Data

	return res, err
}

// FetchDisbursementByID returns a response from Fetch Disbursement API.
//
//	[Docs Fetch Disbursement]: https://durianpay.id/docs/api/disbursements/fetch-one/
func (c *Client) FetchDisbursementByID(ctx context.Context, ID string) (res *durianpay.DisbursementData, err *durianpay.Error) {
	url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_FETCH_BY_ID
	url = strings.ReplaceAll(url, ":id", ID)

	apiRes, err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	tempRes := struct {
		Data durianpay.DisbursementData `json:"data"`
	}{}

	jsonErr := json.Unmarshal(apiRes, &tempRes)
	if jsonErr != nil {
		return nil, durianpay.FromSDKError(jsonErr)
	}

	res = &tempRes.Data

	return res, err
}

// DeleteDisbursement returns a response from Delete Disbursement API
//
//	[Docs Delete Disbursement]: https://durianpay.id/docs/api/disbursements/delete/
func (c *Client) DeleteDisbursement(ctx context.Context, ID string) (res string, err *durianpay.Error) {
	url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_DELETE
	url = strings.ReplaceAll(url, ":id", ID)

	apiRes, err := c.Api.Req(ctx, http.MethodDelete, url, nil, nil, nil)
	if err != nil {
		return "", err
	}

	parseRes := struct {
		Data string `json:"data"`
	}{}

	jsonErr := json.Unmarshal(apiRes, &parseRes)
	if jsonErr != nil {
		return "", durianpay.FromSDKError(jsonErr)
	}

	return parseRes.Data, err
}
