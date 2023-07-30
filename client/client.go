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
)

type Client struct {
	Opts         Options
	Disbursement *disbursement.Client
}

type Options struct {
	ServerKey string
}

func (c *Client) Init() {
	api := common.NewAPI(c.Opts.ServerKey)
	c.Disbursement = &disbursement.Client{ServerKey: c.Opts.ServerKey, Api: api}
}

func NewClient(opts Options) *Client {
	client := Client{
		Opts: opts,
	}

	return &client
}
