/*
 * File Created: Saturday, 2nd September 2023 2:00:07 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package ewalletaccount

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
	path_response_ewalletaccount = "../internal/tests/response/ewalletaccount/"
	path_response                = "../internal/tests/response/"
	path_payload_ewalletaccount  = "../internal/tests/payload/ewalletaccount/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_Link(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.EwalletAccountLinkPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *Link
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.EwalletAccountLinkPayload{
					Mobile:      "8888888888",
					WalletType:  "GOPAY",
					RedirectURL: "https://redirect_url.com/",
				},
			},
			prepare: func(m mocks, args args) {
				headers := map[string]string{
					"Is-live": "true",
				}
				m.api.EXPECT().Req(gomock.Any(), "POST", PATH_EWALLET_ACCOUNT_BIND, nil, args.payload, headers, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, headers map[string]string, response any) *durianpay.Error {

						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_ewalletaccount+"link_200.json"), response)
						if err != nil {
							panic(err)
						}
						return nil
					})
			},
			wantRes: &Link{
				WalletType:     "GOPAY",
				Mobile:         "8888888888",
				RefID:          "7f125e70-095e-481d-8db8-241df9d5b86d",
				Status:         "pending",
				AppRedirectURL: "https://simulator.sandbox.midtrans.com/gopay/partner/web/otp?id=14c95e30-0586-4270-961e-f3b0b3d3d2b0",
				Message:        "use redirection url to bind the account",
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_EWALLET_ACCOUNT_BIND, nil, args.payload, gomock.Any(), gomock.Any()).
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
				payload := durianpay.EwalletAccountLinkPayload{}

				if !featureWrap.DeepEqualPayload(path_payload_ewalletaccount+"link.json", &payload, &tt.args.payload) {
					t.Errorf("Client.Link() gotPayload = %v, wantPayload %v", payload, tt.args.payload)
				}
			}

			gotRes, gotErr := c.Link(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Link() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Link() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
