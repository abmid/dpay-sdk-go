/*
 * File Created: Sunday, 3rd September 2023 10:52:35 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package promo

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
	PATH_PROMO = durianpay.DURIANPAY_URL + "/v1/merchants/promos"
)

// Create return a response from Create Promos API.
//
//	[Doc Create Promos API]: https://durianpay.id/docs/api/promos/create/
func (c *Client) Create(ctx context.Context, payload durianpay.PromoPayload) (*Promo, *durianpay.Error) {
	res := struct {
		Data Promo `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, PATH_PROMO, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchPromos return a response from Promos Fetch API.
//
//	[Doc Promos Fetch API]: https://durianpay.id/docs/api/promos/fetch/
func (c *Client) FetchPromos(ctx context.Context) ([]Promo, *durianpay.Error) {
	res := struct {
		Data []Promo `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, PATH_PROMO, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
