/*
 * File Created: Friday, 1st September 2023 11:29:13 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package refund

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/internal/tests"
	mock_common "github.com/abmid/dpay-sdk-go/internal/tests/mock"
	"github.com/golang/mock/gomock"
)

const (
	path_response_refund = "../internal/tests/response/refund/"
	path_response        = "../internal/tests/response/"
	path_payload_refund  = "../internal/tests/payload/refund/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_Create(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.RefundPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *Refund
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.RefundPayload{
					RefID:         "order_ref_241",
					PaymentID:     "pay_y2yKEEWBYe1299",
					Amount:        "10000",
					UseRefundLink: false,
					Notes:         "rejected product",
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().Req(args.ctx, "POST", PATH_REFUND, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_refund+"create_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Refund{
				ID:            "rfn_iLNvzkCakx0330",
				RefID:         "order_ref_241",
				Amount:        "10000.00",
				RefundType:    "full",
				Status:        "done",
				CreatedAt:     tests.StringToTime("2023-08-30T16:54:39.593123Z"),
				UpdatedAt:     tests.StringToTime("2023-08-30T16:54:39.593123Z"),
				Source:        "api",
				CustomerID:    "cus_IwDIb0MDY20938",
				CustomerName:  "Abdul",
				CustomerEmail: "abdul.surel@gmail.com",
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_REFUND, nil, args.payload, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response+"internal_server_error_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response+"internal_server_error_500.json")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := mock_common.NewMockApi(featureWrap.Ctrl)
			parseArgs := tt.args

			c := &Client{
				ServerKey: featureWrap.ServerKey,
				Api:       apiMock,
			}

			tt.prepare(mocks{api: apiMock}, parseArgs)

			if tt.wantRes != nil {
				payload := durianpay.RefundPayload{}

				if !featureWrap.DeepEqualPayload(path_payload_refund+"create.json", &payload, &tt.args.payload) {
					t.Errorf("Client.Create() gotPayload = %v, wantPayload %v", payload, tt.args.payload)
				}
			}

			gotRes, gotErr := c.Create(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Create() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Create() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
