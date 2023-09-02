/*
 * File Created: Saturday, 2nd September 2023 3:34:42 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package settlement

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
	PATH_SETTLEMENT             = "/v1/settlements"
	PATH_SETTLEMENT_DETAIL      = PATH_SETTLEMENT + "/details"
	PATH_SETTLEMENT_FETCH_BY_ID = PATH_SETTLEMENT + "/:id"
)

// FetchSettlements return a response from Settlements Fetch API.
//
//	[Doc Settlements Fetch API]: https://durianpay.id/docs/api/settlements/settlements-fetch-list/
func (c *Client) FetchSettlements(ctx context.Context, opt durianpay.SettlementOption) (*FetchSettlements, *durianpay.Error) {
	res := FetchSettlements{}

	err := c.Api.Req(ctx, http.MethodGet, PATH_SETTLEMENT, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FetchDetails return a response from Settlements Details Fetch API.
//
//	[Doc Settlements Details Fetch API]: https://durianpay.id/docs/api/settlements/settlements-fetch-details/
func (c *Client) FetchDetails(ctx context.Context, opt durianpay.SettlementOption) (*FetchDetails, *durianpay.Error) {
	res := FetchDetails{}

	err := c.Api.Req(ctx, http.MethodGet, PATH_SETTLEMENT_DETAIL, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// StatusByPaymentID return a response from Settlements Status By Payment ID API.
//
//	[Doc Settlements Status By Payment ID API]: https://durianpay.id/docs/api/settlements/settlements-fetch-by-payment-id/
func (c *Client) StatusByPaymentID(ctx context.Context, paymentID string) (*SettlementDetail, *durianpay.Error) {
	res := struct {
		Data SettlementDetail `json:"data"`
	}{}

	params := struct {
		PaymentID string `url:"payment_id"`
	}{PaymentID: paymentID}

	err := c.Api.Req(ctx, http.MethodGet, PATH_SETTLEMENT_DETAIL, params, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchSettlementByID return a response from Settlements By ID API.
//
//	[Doc Settlements By ID API]: https://durianpay.id/docs/api/settlements/settlements-fetch-by-id/
func (c *Client) FetchSettlementByID(ctx context.Context, ID string) (*Settlement, *durianpay.Error) {
	url := strings.ReplaceAll(PATH_SETTLEMENT_FETCH_BY_ID, ":id", ID)
	res := struct {
		Data Settlement `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
