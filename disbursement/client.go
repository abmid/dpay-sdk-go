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

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/common"
)

type Client struct {
	ServerKey string
	Api       common.Api
}

const (
	PATH = "/v1/disbursements/validate"
)

func (c *Client) ValidateDisbursement(ctx context.Context, payload durianpay.ValidateDisbursementPayload) (res *durianpay.ValidateDisbursement, err *durianpay.Error) {
	apiRes, err := c.Api.Req(ctx, http.MethodPost, durianpay.DURIANPAY_URL+PATH, payload, payload.IdempotenKey)
	if err != nil {
		return nil, err
	}

	jsonErr := json.Unmarshal(apiRes, &res)
	if jsonErr != nil {
		return nil, durianpay.FromSDKError(jsonErr)
	}

	return res, err
}
