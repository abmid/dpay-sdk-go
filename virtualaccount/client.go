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

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/common"
)

type Client struct {
	ServerKey string
	Api       common.Api
}

const (
	PATH_VA = "/v1/virtual-accounts"
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
