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
	"io"
	"net/http"

	durianpay "github.com/abmid/dpay-sdk-go"
	goquery "github.com/google/go-querystring/query"
)

// Api interface build for dependecy injection and for testing purposes.
type Api interface {
	Req(ctx context.Context, method string, url string, param any, body any, headers map[string]string) (res []byte, durianErr *durianpay.Error)
}

type ApiImplement struct {
	ServerKey string
}

func NewAPI(serverKey string) *ApiImplement {
	return &ApiImplement{
		ServerKey: serverKey,
	}
}

// Req is an http request made specifically to hit the durian pay endpoint.
// If the HTTP status code returned is not 2xx then an error will be returned
func (c *ApiImplement) Req(ctx context.Context, method string, url string, param any, body any, headers map[string]string) (res []byte, durianErr *durianpay.Error) {
	parseBody, err := json.Marshal(body)
	if err != nil {
		return nil, durianpay.FromSDKError(err)
	}

	base64SecretKey := base64.StdEncoding.EncodeToString([]byte(c.ServerKey + ":"))
	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(parseBody))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64SecretKey))

	if headers != nil {
		for key, value := range headers {
			httpReq.Header.Add(key, value)
		}
	}

	if param != nil {
		parseParam, err := goquery.Values(param)
		if err != nil {
			return nil, durianpay.FromSDKError(err)
		}

		httpReq.URL.RawQuery = parseParam.Encode()
	}

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, durianpay.FromSDKError(err)
	}
	defer httpRes.Body.Close()

	resBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, durianpay.FromSDKError(err)
	}

	isStatusCodeSuccess := (httpRes.StatusCode >= 200) && (httpRes.StatusCode < 300)

	if !isStatusCodeSuccess {
		return nil, durianpay.FromAPI(httpRes.StatusCode, resBody)
	}

	return resBody, nil
}

// HeaderIdempotencyKey returns X-Idempotency-Key & idempotency-key values for DurianPay idempotency purposes.
// Difference about X-Idempotency-Key and idempotency-key you can read on
// [Docs Idempotent] https://durianpay.id/docs/integration/disbursements/idempotent/
func HeaderIdempotencyKey(xIdempotencyKey, idempotencyKey string) map[string]string {
	headers := map[string]string{}

	if xIdempotencyKey != "" {
		headers["X-Idempotency-Key"] = xIdempotencyKey
	}

	if idempotencyKey != "" {
		headers["idempotency_key"] = idempotencyKey
	}

	return headers
}
