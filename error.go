/*
 * File Created: Friday, 28th July 2023 6:23:33 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

import "encoding/json"

const (
	SDKErrorCode = "SDK_ERROR"
)

// Error is commons response error DurianPay
type Error struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"Message"`
}

func FromSDKError(err error) *Error {
	return &Error{
		Error:     err.Error(),
		ErrorCode: SDKErrorCode,
		Message:   err.Error(),
	}
}

func FromAPI(statusCode int, responseBody []byte) *Error {
	tempErr := Error{}

	err := json.Unmarshal(responseBody, &tempErr)
	if err != nil {
		return FromSDKError(err)
	}

	return &tempErr
}
