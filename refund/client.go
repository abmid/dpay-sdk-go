/*
 * File Created: Friday, 1st September 2023 11:29:06 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package refund

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
	PATH_REFUND = durianpay.DURIANPAY_URL + "/v1/refunds"
)

// Create return a response from Create Refund API.
//
//	[Doc Create Refund API]: https://durianpay.id/docs/api/refunds/create/
func (c *Client) Create(ctx context.Context, payload durianpay.RefundPayload) (*Refund, *durianpay.Error) {
	res := struct {
		Data Refund `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_REFUND, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
