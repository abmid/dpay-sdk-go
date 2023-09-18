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
	"strings"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/common"
)

type Client struct {
	ServerKey string
	Api       common.Api
}

const (
	pathPromo      = durianpay.DurianpayURL + "/v1/merchants/promos"
	pathFetchByID  = pathPromo + "/:id"
	pathDeleteByID = pathPromo + "/:id"
	pathUpdateByID = pathPromo + "/:id"
)

// Create return a response from Create Promos API.
//
//	[Doc Create Promos API]: https://durianpay.id/docs/api/promos/create/
func (c *Client) Create(ctx context.Context, payload durianpay.PromoPayload) (*Promo, *durianpay.Error) {
	res := struct {
		Data Promo `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, pathPromo, nil, payload, nil, &res)
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

	err := c.Api.Req(ctx, http.MethodGet, pathPromo, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

// FetchPromoByID return a response from Promos Fetch By ID API.
//
//	[Doc Promos Fetch By ID API]: https://durianpay.id/docs/api/promos/fetch-one/
func (c *Client) FetchPromoByID(ctx context.Context, ID string) (*Promo, *durianpay.Error) {
	url := strings.ReplaceAll(pathFetchByID, ":id", ID)
	res := struct {
		Data Promo `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Delete return a response from Delete Promo API.
//
//	[Doc Delete Promo API]: https://durianpay.id/docs/api/promos/delete/
func (c *Client) Delete(ctx context.Context, ID string) (string, *durianpay.Error) {
	url := strings.ReplaceAll(pathDeleteByID, ":id", ID)
	res := struct {
		Data string `json:"message"`
	}{}

	err := c.Api.Req(ctx, http.MethodDelete, url, nil, nil, nil, &res)
	if err != nil {
		return "", err
	}

	return res.Data, nil
}

// Update return a response from Update Promos API.
//
//	[Doc Update Promos API]: https://durianpay.id/docs/api/promos/update/
func (c *Client) Update(ctx context.Context, ID string, payload durianpay.PromoPayload) (*Promo, *durianpay.Error) {
	url := strings.ReplaceAll(pathFetchByID, ":id", ID)
	res := struct {
		Data Promo `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPatch, url, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
