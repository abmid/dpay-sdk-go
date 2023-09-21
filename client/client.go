/*
 * File Created: Sunday, 30th July 2023 12:22:26 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package client

import (
	"github.com/abmid/dpay-sdk-go/common"
	"github.com/abmid/dpay-sdk-go/disbursement"
	"github.com/abmid/dpay-sdk-go/ewalletaccount"
	"github.com/abmid/dpay-sdk-go/invoice"
	"github.com/abmid/dpay-sdk-go/order"
	"github.com/abmid/dpay-sdk-go/payment"
	"github.com/abmid/dpay-sdk-go/promo"
	"github.com/abmid/dpay-sdk-go/refund"
	"github.com/abmid/dpay-sdk-go/settlement"
	"github.com/abmid/dpay-sdk-go/virtualaccount"
)

// Client represents return for NewClient.
type Client struct {
	Opts           Options
	Order          *order.Client
	Payment        *payment.Client
	Promo          *promo.Client
	Disbursement   *disbursement.Client
	Settlement     *settlement.Client
	Refund         *refund.Client
	EWalletAccount *ewalletaccount.Client
	VA             *virtualaccount.Client
	Invoice        *invoice.Client
}

// Options represents of parameter option for NewClient.
type Options struct {
	ServerKey string
}

func (c *Client) Init() {
	api := common.NewAPI(c.Opts.ServerKey)
	c.Order = &order.Client{ServerKey: c.Opts.ServerKey, Api: api}
	c.Payment = &payment.Client{ServerKey: c.Opts.ServerKey, Api: api}
	c.Promo = &promo.Client{ServerKey: c.Opts.ServerKey, Api: api}
	c.Disbursement = &disbursement.Client{ServerKey: c.Opts.ServerKey, Api: api}
	c.Settlement = &settlement.Client{ServerKey: c.Opts.ServerKey, Api: api}
	c.Refund = &refund.Client{ServerKey: c.Opts.ServerKey, Api: api}
	c.EWalletAccount = &ewalletaccount.Client{ServerKey: c.Opts.ServerKey, Api: api}
	c.VA = &virtualaccount.Client{ServerKey: c.Opts.ServerKey, Api: api}
	c.Invoice = &invoice.Client{ServerKey: c.Opts.ServerKey, Api: api}
}

// NewClient represents the creation of new client with options to access all the different resources.
func NewClient(opts Options) *Client {
	client := Client{
		Opts: opts,
	}

	client.Init()

	return &client
}
