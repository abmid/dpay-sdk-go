/*
 * File Created: Monday, 18th September 2023 11:33:54 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package invoice

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
	urlInvoice             = durianpay.DurianpayURL + "/v1/invoices"
	urlGenerateCheckoutURL = urlInvoice + "/generate_checkout_url/:customer_id"
	urlFetchByID           = urlInvoice + "/:id"
	urlUpdateByID          = urlInvoice + "/:id"
	urlPay                 = urlInvoice + "/pay"
	urlManualPayment       = urlInvoice + "/manual_transaction"
	urlDeleteByID          = urlInvoice + "/:id"
)

// Create returns a response from Create Invoice API.
//
//	[Doc Create Invoice API]: https://durianpay.id/docs/api/invoices/create/
func (c *Client) Create(ctx context.Context, payload durianpay.InvoiceCreatePayload) (*Create, *durianpay.Error) {
	res := struct {
		Data Create `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, urlInvoice, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// GenerateCheckoutURL returns a response from Generate Checkout URL API.
//
//	[Doc Generate Checkout URL API]: https://durianpay.id/docs/api/invoices/generate-checkout-url/
func (c *Client) GenerateCheckoutURL(ctx context.Context, customerID string) (*GenerateCheckoutURL, *durianpay.Error) {
	url := strings.ReplaceAll(urlGenerateCheckoutURL, ":customer_id", customerID)

	res := struct {
		Data GenerateCheckoutURL `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, url, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchInvoiceByID returns a response from Invoice Fetch by ID API.
//
//	[Doc Invoice Fetch by ID API]: https://durianpay.id/docs/api/invoices/fetch-one/
func (c *Client) FetchInvoiceByID(ctx context.Context, ID string) (*FetchInvoiceByID, *durianpay.Error) {
	url := strings.ReplaceAll(urlFetchByID, ":id", ID)

	res := struct {
		Data FetchInvoiceByID `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// FetchInvoices returns a response from List Invoices API
//
//	[Doc List Invoices API]: https://durianpay.id/docs/api/invoices/fetch/
func (c *Client) FetchInvoices(ctx context.Context, opt durianpay.InvoiceFetchOption) (*FetchInvoices, *durianpay.Error) {
	res := struct {
		Data FetchInvoices `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodGet, urlInvoice, opt, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Update returns a response from Update Invoice API
//
//	[Doc Update Invoice API]: https://durianpay.id/docs/api/invoices/update/
func (c *Client) Update(ctx context.Context, ID string, payload durianpay.InvoiceUpdatePayload) (*Update, *durianpay.Error) {
	url := strings.ReplaceAll(urlUpdateByID, ":id", ID)

	res := struct {
		Data Update `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPut, url, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Pay returns a response from Pay Invoice API
//
//	[Doc Pay Invoice API]: https://durianpay.id/docs/api/invoices/pay/
func (c *Client) Pay(ctx context.Context, payload durianpay.InvoicePayPayload) (*Pay, *durianpay.Error) {
	res := struct {
		Data Pay `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, urlPay, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
