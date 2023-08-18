/*
 * File Created: Sunday, 30th July 2023 12:28:21 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package common

import (
	"context"
	"encoding/json"
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
					SkipValidation: tests.BoolPtr(true),
					AnotherParam:   "test-param",
				},
				body: durianpay.ValidateDisbursementPayload{
					IdempotenKey:  "1",
					AccountNumber: "123737383830",
					BankCode:      "bca",
				},
				durianRes: durianpay.ValidateDisbursement{},
				headers:   HeaderIdempotencyKey("x-123", ""),
			},
			prepare: func(args args) {
				query := map[string]string{}
				query["skip_validation"] = "true"
				query["another_param"] = "test-param"

				httpmock.RegisterMatcherResponderWithQuery(args.method, args.url, query, httpmock.Matcher{},
					tests.HttpMockResJSON(200, "../internal/tests/response/validate_disbursement_200.json", args.headers))
			},
			wantRes: durianpay.ValidateDisbursement{
				Data: durianpay.ValidateDisbursementData{
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
					tests.HttpMockResJSON(400, "../internal/tests/response/validate_disbursement_400.json", args.headers))
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

			res := durianpay.ValidateDisbursement{}
			gotRes, gotDurianErr := c.Req(tt.args.ctx, tt.args.method, tt.args.url, tt.args.param, tt.args.body, tt.args.headers)

			if tt.wantRes != nil {
				err := json.Unmarshal(gotRes, &res)
				if err != nil {
					t.Fatal(err)
				}

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
