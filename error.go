/*
 * File Created: Friday, 28th July 2023 6:23:33 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

import "encoding/json"

const (
	ErrorCodeSDK                    = "SDK_ERROR"
	ErrorCodeDPAYInternalError      = "DPAY_INTERNAL_ERROR"
	ErrorCodeDPAYUnauthorizedAccess = "DPAY_UNAUTHORIZED_ACCESS"
	ErrorCodeDPAYInvalidRequest     = "DPAY_INVALID_REQUEST"
)

// Error is commons response error DurianPay
type Error struct {
	StatusCode int      // Response from http status code
	Error      string   `json:"error"`
	ErrorCode  string   `json:"error_code"`
	Errors     []Errors `json:"errors"`
	Message    string   `json:"message"`
}

type Errors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FromSDKError(err error) *Error {
	return &Error{
		Error:     err.Error(),
		ErrorCode: ErrorCodeSDK,
		Message:   err.Error(),
	}
}

func FromAPI(statusCode int, responseBody []byte) *Error {
	tempErr := Error{StatusCode: statusCode}

	err := json.Unmarshal(responseBody, &tempErr)
	if err != nil {
		return FromSDKError(err)
	}

	return &tempErr
}
