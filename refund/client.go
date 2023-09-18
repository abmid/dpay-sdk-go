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
	"strings"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/common"
)

type Client struct {
	ServerKey string
	Api       common.Api
}

const (
	pathRefund    = durianpay.DurianpayURL + "/v1/refunds"
	pathFetchByID = pathRefund + "/:id"
)

// Create return a response from Create Refund API.
//
//	[Doc Create Refund API]: https://durianpay.id/docs/api/refunds/create/
func (c *Client) Create(ctx context.Context, payload durianpay.RefundPayload) (*Refund, *durianpay.Error) {
	res := struct {
		Data Refund `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, pathRefund, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchRefunds return a response from Refund Fetch API.
//
//	[Doc Refund Fetch API]: https://durianpay.id/docs/api/refunds/fetch/
func (c *Client) FetchRefunds(ctx context.Context, opt durianpay.RefundFetchOption) (*FetchRefunds, *durianpay.Error) {
	res := struct {
		Data FetchRefunds `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, pathRefund, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchRefundByID return a response from Refund Fetch By ID API.
//
//	[Doc Refund Fetch By ID API]: https://durianpay.id/docs/api/refunds/fetch-one/
func (c *Client) FetchRefundByID(ctx context.Context, ID string) (*Refund, *durianpay.Error) {
	url := strings.ReplaceAll(pathFetchByID, ":id", ID)

	res := struct {
		Data Refund `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
