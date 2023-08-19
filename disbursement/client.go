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
	PATH_DISBURSEMENT          = "/v1/disbursements"
	PATH_DISBURSEMENT_VALIDATE = PATH_DISBURSEMENT + "/validate"
	PATH_DISBURSEMENT_SUBMIT   = PATH_DISBURSEMENT + "/submit"
	PATH_DISBURSEMENT_APPROVE  = PATH_DISBURSEMENT + "/:id/approve"
)

// ValidateDisbursement returns a response from validate disbursement API.
// Validate disbursement can be used to fetch the bank account and account number validation
//
//	[Doc Validate Disbursement API]: https://durianpay.id/docs/api/disbursements/validate/
func (c *Client) ValidateDisbursement(ctx context.Context, payload durianpay.ValidateDisbursementPayload) (res *durianpay.ValidateDisbursement, err *durianpay.Error) {
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

// SubmitDisbursement returns a response from submit disbursement API.
// Options about skip_validation & force_disburse you can input in durianpay.DisbursementOption
//
//	[Doc Submit Disbursement API]: https://durianpay.id/docs/api/disbursements/submit/
func (c *Client) SubmitDisbursement(ctx context.Context, payload durianpay.DisbursementPayload, opt *durianpay.DisbursementOption) (res *durianpay.Disbursement, err *durianpay.Error) {
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

// ApproveDisbursement returns a response from approve disbursement API.
// Options about ignore_invalid you can input in durianpay.ApproveDisbursementOption
//
//	[Doc Approve Disbursement API]: https://durianpay.id/docs/api/disbursements/approve/
func (c *Client) ApproveDisbursement(ctx context.Context, payload durianpay.ApproveDisbursementPayload, opt *durianpay.ApproveDisbursementOption) (res *durianpay.Disbursement, err *durianpay.Error) {
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
