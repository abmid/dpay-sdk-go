/*
 * File Created: Saturday, 29th July 2023 10:39:48 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package common

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	durianpay "github.com/abmid/dpay-sdk-go"
)

type Api interface {
	Req(ctx context.Context, method string, url string, body any, idempotenKey string) (res []byte, durianErr *durianpay.Error)
}

type ApiImplement struct {
	ServerKey string
}

func NewAPI(serverKey string) *ApiImplement {
	return &ApiImplement{
		ServerKey: serverKey,
	}
}

func (c *ApiImplement) Req(ctx context.Context, method string, url string, body any, idempotenKey string) (res []byte, durianErr *durianpay.Error) {
	parseBody, err := json.Marshal(body)
	if err != nil {
		return nil, durianpay.FromSDKError(err)
	}

	base64SecretKey := base64.RawStdEncoding.EncodeToString([]byte(c.ServerKey + ":"))
	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(parseBody))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64SecretKey))
	httpReq.Header.Add("X-Idempotency-Key", idempotenKey)

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, durianpay.FromSDKError(err)
	}
	defer httpRes.Body.Close()

	resBody, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, durianpay.FromSDKError(err)
	}

	if httpRes.StatusCode != 200 {
		return nil, durianpay.FromAPI(httpRes.StatusCode, resBody)
	}

	return resBody, nil
}
