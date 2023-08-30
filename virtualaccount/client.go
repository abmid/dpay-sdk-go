/*
 * File Created: Tuesday, 29th August 2023 10:51:07 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package virtualaccount

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
	PATH_VA          = "/v1/virtual-accounts"
	PATH_FETCH_BY_ID = PATH_VA + "/:id"
	PATH_PATCH_BY_ID = PATH_VA + "/:id"
)

// Create returns a response from Virtual Account Create API.
//
//	[Doc Virtual Account Create API]: https://durianpay.id/docs/api/virtual-accounts/create/
func (c *Client) Create(ctx context.Context, payload durianpay.VirtualAccountPayload) (*Create, *durianpay.Error) {
	res := struct {
		Data Create `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, durianpay.DURIANPAY_URL+PATH_VA, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchVirtualAccounts returns a response from Virtual Accounts Fetch API
//
//	[Doc Virtual Accounts Fetch API]: https://durianpay.id/docs/api/virtual-accounts/fetch/
func (c *Client) FetchVirtualAccounts(ctx context.Context, opt durianpay.VirtualAccountFetchOption) (*FetchVirtualAccounts, *durianpay.Error) {
	res := struct {
		Data FetchVirtualAccounts `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, durianpay.DURIANPAY_URL+PATH_VA, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchVirtualAccountByID returns a response from Virtual Accounts Fetch By ID API.
//
//	[Doc Virtual Accounts Fetch By ID API]: https://durianpay.id/docs/api/virtual-accounts/fetch-one/
func (c *Client) FetchVirtualAccountByID(ctx context.Context, ID string) (*FetchVirtualAccount, *durianpay.Error) {
	url := durianpay.DURIANPAY_URL + PATH_FETCH_BY_ID
	url = strings.ReplaceAll(url, ":id", ID)
	res := struct {
		Data FetchVirtualAccount `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// PatchByID returns a response from Virtual Accounts Patch By ID API.
//
//	[Doc Virtual Accounts Patch By ID API]: https://durianpay.id/docs/api/virtual-accounts/patch-one/
func (c *Client) PatchByID(ctx context.Context, ID string, payload durianpay.VirtualAccountPatchPayload) (*FetchVirtualAccount, *durianpay.Error) {
	url := durianpay.DURIANPAY_URL + PATH_PATCH_BY_ID
	url = strings.ReplaceAll(url, ":id", ID)
	res := struct {
		Data FetchVirtualAccount `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPatch, url, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
