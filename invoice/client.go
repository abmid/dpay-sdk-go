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