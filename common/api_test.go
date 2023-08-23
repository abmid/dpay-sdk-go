/*
 * File Created: Sunday, 30th July 2023 12:28:21 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package common

import (
	"context"
	"reflect"
	"testing"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/internal/tests"
	"github.com/jarcoal/httpmock"
)

func TestApiImplement_Req(t *testing.T) {
	type param struct {
		SkipValidation *bool  `url:"skip_validation"`
		AnotherParam   string `url:"another_param"`
	}

	type responseData struct {
		AccountNumber string `json:"account_number"`
		BankCode      string `json:"bank_code"`
		AccountHolder string `json:"account_holder"`
		Status        string `json:"status"`
	}

	type response struct {
		Message string       `json:"message"`
		Data    responseData `json:"data"`
	}

	type fields struct {
		ServerKey string
	}

	type args struct {
		ctx       context.Context
		method    string
		url       string
		param     any
		body      any
		headers   map[string]string
		durianRes any
	}

	tests := []struct {
		name          string
		fields        fields
		args          args
		prepare       func(args args)
		wantRes       any
		wantDurianErr *durianpay.Error
	}{
		{
			name: "Success http response 200",
			fields: fields{
				ServerKey: "dpay_test_xxx",
			},
			args: args{
				ctx:    context.Background(),
				method: "POST",
				url:    durianpay.DURIANPAY_URL,
				param: param{
					SkipValidation: tests.ToPtr(true),
					AnotherParam:   "test-param",
				},
				body: durianpay.DisbursementValidatePayload{
					XIdempotencyKey: "1",
					AccountNumber:   "123737383830",
					BankCode:        "bca",
				},
				durianRes: response{},
				headers:   HeaderIdempotencyKey("x-123", ""),
			},
			prepare: func(args args) {
				query := map[string]string{}
				query["skip_validation"] = "true"
				query["another_param"] = "test-param"

				httpmock.RegisterMatcherResponderWithQuery(args.method, args.url, query, httpmock.Matcher{},
					tests.HttpMockResJSON(200, "../internal/tests/response/disbursement/validate_disbursement_200.json", args.headers))
			},
			wantRes: response{
				Data: responseData{
					AccountNumber: "123737383830",
					BankCode:      "bca",
					Status:        "processing",
				},
			},
			wantDurianErr: nil,
		},
		{
			name: "Failed http response 4xx",
			fields: fields{
				ServerKey: "dpay_test_xxx",
			},
			args: args{
				ctx:    context.Background(),
				method: "POST",
				url:    durianpay.DURIANPAY_URL,
			},
			prepare: func(args args) {
				httpmock.RegisterResponder(args.method, args.url,
					tests.HttpMockResJSON(400, "../internal/tests/response/disbursement/validate_disbursement_400.json", args.headers))
			},
			wantRes: nil,
			wantDurianErr: &durianpay.Error{
				StatusCode: 400,
				Error:      "error reading request body",
				ErrorCode:  "DPAY_INTERNAL_ERROR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			c := &ApiImplement{
				ServerKey: tt.fields.ServerKey,
			}

			tt.prepare(tt.args)

			res := response{}
			gotDurianErr := c.Req(tt.args.ctx, tt.args.method, tt.args.url, tt.args.param, tt.args.body, tt.args.headers, &res)

			if tt.wantRes != nil {
				if !reflect.DeepEqual(res, tt.wantRes) {
					t.Errorf("ApiImplement.Req() gotRes = %v, want %v", res, tt.wantRes)
				}
			}

			if !reflect.DeepEqual(gotDurianErr, tt.wantDurianErr) {
				t.Errorf("ApiImplement.Req() gotDurianErr = %v, want %v", gotDurianErr, tt.wantDurianErr)
			}
		})
	}
}
