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
	"strings"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/common"
)

type Client struct {
	ServerKey string
	Api       common.Api
}

const (
	PATH_ORDER              = "/v1/orders"
	PATCH_FETCH_ORDER_BY_ID = PATH_ORDER + "/:id"
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

// FetchOrders returns a response from Orders Fetch API.
//
//	[Doc Orders Fetch API]: https://durianpay.id/docs/api/orders/fetch/
func (c *Client) FetchOrders(ctx context.Context, opt durianpay.OrderFetchOption) (*FetchOrders, *durianpay.Error) {

	res := struct {
		Data FetchOrders `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, durianpay.DURIANPAY_URL+PATH_ORDER, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchOrderByID returns a response from Order Fetch By ID API.
//
//	[Doc Order Fetch By ID API]: https://durianpay.id/docs/api/orders/fetch-one/
func (c *Client) FetchOrderByID(ctx context.Context, ID string, opt durianpay.OrderFetchByIDOption) (*FetchOrder, *durianpay.Error) {
	url := durianpay.DURIANPAY_URL + PATCH_FETCH_ORDER_BY_ID
	url = strings.ReplaceAll(url, ":id", ID)

	res := struct {
		Data FetchOrder `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// CreatePaymentLink returns a response from Create Payment Link API.
//
//	[Doc Create Payment Link API]: https://durianpay.id/docs/api/orders/create-link/
func (c *Client) CreatePaymentLink(ctx context.Context, payload durianpay.OrderPaymentLinkPayload) (*OrderCreate, *durianpay.Error) {
	res := struct {
		Data OrderCreate `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, durianpay.DURIANPAY_URL+PATH_ORDER, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
