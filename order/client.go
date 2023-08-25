/*
 * File Created: Thursday, 24th August 2023 6:36:16 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package order

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
	PATH_ORDER = "/v1/orders"
)

// Create returns a response from Create Order API.
//
//	[Doc Create Order API]: https://durianpay.id/docs/api/orders/create/
func (c *Client) Create(ctx context.Context, payload durianpay.OrderPayload) (*OrderCreate, *durianpay.Error) {
	res := struct {
		Data OrderCreate `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, durianpay.DURIANPAY_URL+PATH_ORDER, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
