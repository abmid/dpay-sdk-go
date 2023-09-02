/*
 * File Created: Saturday, 2nd September 2023 2:00:00 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package ewalletaccount

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
	PATH_EWALLET_ACCOUNT        = durianpay.DURIANPAY_URL + "/v1/ewallet/account"
	PATH_EWALLET_ACCOUNT_BIND   = PATH_EWALLET_ACCOUNT + "/bind"
	PATH_EWALLET_ACCOUNT_UNBIND = PATH_EWALLET_ACCOUNT + "/:id/unbind"
)

// Link return a response from Link E-Wallet Account API.
//
//	[Doc Link E-Wallet Account API]: https://durianpay.id/docs/api/ewallet/link/
func (c *Client) Link(ctx context.Context, payload durianpay.EwalletAccountLinkPayload) (*Link, *durianpay.Error) {
	headers := map[string]string{"Is-live": "true"}
	res := struct {
		Data Link `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_EWALLET_ACCOUNT_BIND, nil, payload, headers, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Unlink return a response from Unlink E-Wallet Account API.
//
//	[Doc Unlink E-Wallet Account API]: https://durianpay.id/docs/api/ewallet/unlink/
func (c *Client) Unlink(ctx context.Context, ID string) (*Unlink, *durianpay.Error) {
	url := strings.ReplaceAll(PATH_EWALLET_ACCOUNT_UNBIND, ":id", ID)
	headers := map[string]string{"Is-live": "true"}
	res := struct {
		Data Unlink `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPut, url, nil, nil, headers, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
