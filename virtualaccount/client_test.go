/*
 * File Created: Wednesday, 30th August 2023 8:39:38 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package virtualaccount

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/internal/tests"
	mock_common "github.com/abmid/dpay-sdk-go/internal/tests/mock"
	"github.com/golang/mock/gomock"
)

const (
	pathResponse   = "../internal/tests/response/"
	pathResponseVA = pathResponse + "/virtualaccount/"
	pathPayloadVA  = "../internal/tests/payload/virtualaccount/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_Create(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.VirtualAccountPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *Create
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.VirtualAccountPayload{
					BankCode: "PERMATA",
					Name:     "Abdul Hamid",
					IsClosed: true,
					Amount:   "12333",
					Customer: durianpay.VirtualAccountCustomer{
						GivenName: "Abdul Hamid",
						Mobile:    "+6288888888",
						Email:     "abdul.surel@gmail.com",
					},
					ExpiryMinutes:           14400,
					AccountSuffix:           "123456",
					IsReusable:              true,
					VaRefID:                 "1234",
					MinAmount:               10000,
					MaxAmount:               15000,
					AutoDisableAfterPayment: true,
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().Req(args.ctx, "POST", pathVA, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseVA+"create_201.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Create{
				CustomerID: "cus_iODsnGTdCh3706",
				VirtualAccount: VirtualAccount{
					ID:                      "va_sample_Cre3I9gg962549",
					BankCode:                "BCA",
					AccountNumber:           "190061002123456",
					Name:                    "Abdul Hamid",
					IsClosed:                true,
					Amount:                  12333,
					Currency:                "IDR",
					CustomerID:              "cus_iODsnGTdCh3706",
					IsSandbox:               true,
					CreatedAt:               tests.StringToTime("2023-08-29T16:02:41.002599Z"),
					ExpiryAt:                tests.StringToTime("2023-09-08T16:02:41.009911Z"),
					IsDisabled:              false,
					IsPaid:                  false,
					IsReusable:              true,
					VaRefID:                 "1234",
					AutoDisableAfterPayment: true,
				},
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", pathVA, nil, args.payload, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")),
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
				payload := durianpay.VirtualAccountPayload{}

				if !featureWrap.DeepEqualPayload(pathPayloadVA+"create.json", &payload, &tt.args.payload) {
					t.Errorf("Client.Create() gotPayload = %v, wantPayload %v", tt.args.payload, payload)
				}
			}

			gotRes, gotErr := c.Create(parseArgs.ctx, parseArgs.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Create() gotRes = %v, want %v", gotRes, tt.wantErr)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Create() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchVirtualAccounts(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		opt durianpay.VirtualAccountFetchOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *FetchVirtualAccounts
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				opt: durianpay.VirtualAccountFetchOption{
					From:  time.Now().Format("2006-01-02"),
					To:    time.Now().Format("2006-01-02"),
					Skip:  10,
					Limit: 10,
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().Req(args.ctx, "GET", pathVA, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseVA+"fetch_virtualaccounts_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchVirtualAccounts{
				VirtualAccounts: []VirtualAccount{
					{
						ID:                      "va_sample_Cre3I9gg962549",
						BankCode:                "BCA",
						AccountNumber:           "190061002123456",
						Name:                    "Abdul Hamid",
						IsClosed:                true,
						Amount:                  12333,
						Currency:                "IDR",
						CustomerID:              "cus_iODsnGTdCh3706",
						IsSandbox:               true,
						CreatedAt:               tests.StringToTime("2023-08-29T16:02:41.002599Z"),
						ExpiryAt:                tests.StringToTime("2023-09-08T16:02:41.009911Z"),
						IsDisabled:              false,
						IsPaid:                  false,
						IsReusable:              true,
						VaRefID:                 "1234",
						AutoDisableAfterPayment: true,
					},
				},
				Total: 1,
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", pathVA, args.opt, nil, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")),
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

			gotRes, gotErr := c.FetchVirtualAccounts(parseArgs.ctx, parseArgs.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchVirtualAccounts() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchVirtualAccounts() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchVirtualAccountByID(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		ID  string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *FetchVirtualAccount
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				ID:  "va_sample_Cre3I9gg962549",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(pathFetchByID, ":id", args.ID)

				m.api.EXPECT().Req(args.ctx, "GET", url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseVA+"fetch_virtualaccount_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchVirtualAccount{
				VirtualAccount: VirtualAccount{
					ID:                      "va_sample_Cre3I9gg962549",
					BankCode:                "BCA",
					AccountNumber:           "190061002123456",
					Name:                    "Abdul Hamid",
					IsClosed:                true,
					Amount:                  12333,
					Currency:                "IDR",
					CustomerID:              "cus_iODsnGTdCh3706",
					IsSandbox:               true,
					CreatedAt:               tests.StringToTime("2023-08-29T16:02:41.002599Z"),
					ExpiryAt:                tests.StringToTime("2023-09-08T16:02:41.009911Z"),
					IsDisabled:              false,
					IsPaid:                  false,
					IsReusable:              true,
					VaRefID:                 "1234",
					AutoDisableAfterPayment: true,
				},
				VirtualAccountStatus: "VirtualAccountSuccess",
				Customer: VirtualAccountCustomer{
					ID:        "cus_iODsnGTdCh3706",
					GivenName: "Abdul Hamid",
					Email:     "abdul.surel@gmail.com",
					Mobile:    "+6288888888",
				},
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(pathFetchByID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, nil, nil, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")),
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

			gotRes, gotErr := c.FetchVirtualAccountByID(parseArgs.ctx, parseArgs.ID)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchVirtualAccountByID() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchVirtualAccountByID() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_PatchByID(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		ID      string
		payload durianpay.VirtualAccountPatchPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *FetchVirtualAccount
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.VirtualAccountPatchPayload{
					ExpiryMinutes: 1440,
					MinAmount:     11000,
					MaxAmount:     13000,
					Amount:        123456,
					IsDisabled:    true,
					VaRefID:       "1234412",
				},
				ID: "va_sample_Cre3I9gg962549",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(pathPatchByID, ":id", args.ID)

				m.api.EXPECT().Req(args.ctx, "PATCH", url, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseVA+"fetch_virtualaccount_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchVirtualAccount{
				VirtualAccount: VirtualAccount{
					ID:                      "va_sample_Cre3I9gg962549",
					BankCode:                "BCA",
					AccountNumber:           "190061002123456",
					Name:                    "Abdul Hamid",
					IsClosed:                true,
					Amount:                  12333,
					Currency:                "IDR",
					CustomerID:              "cus_iODsnGTdCh3706",
					IsSandbox:               true,
					CreatedAt:               tests.StringToTime("2023-08-29T16:02:41.002599Z"),
					ExpiryAt:                tests.StringToTime("2023-09-08T16:02:41.009911Z"),
					IsDisabled:              false,
					IsPaid:                  false,
					IsReusable:              true,
					VaRefID:                 "1234",
					AutoDisableAfterPayment: true,
				},
				VirtualAccountStatus: "VirtualAccountSuccess",
				Customer: VirtualAccountCustomer{
					ID:        "cus_iODsnGTdCh3706",
					GivenName: "Abdul Hamid",
					Email:     "abdul.surel@gmail.com",
					Mobile:    "+6288888888",
				},
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(pathPatchByID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "PATCH", url, nil, args.payload, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")),
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
				payload := durianpay.VirtualAccountPatchPayload{}

				if !featureWrap.DeepEqualPayload(pathPayloadVA+"patch.json", &payload, &parseArgs.payload) {
					t.Errorf("Client.PatchByID() gotPayload = %v, wantPayload %v", tt.args.payload, payload)
				}
			}

			gotRes, gotErr := c.PatchByID(parseArgs.ctx, parseArgs.ID, parseArgs.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.PatchByID() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.PatchByID() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_PaymentSimulate(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.VirtualAccountPaymentSimulatePayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes string
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.VirtualAccountPaymentSimulatePayload{
					Amount:        "12333",
					AccountNumber: "88565004532522",
					ForceFail:     false,
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().Req(args.ctx, "POST", pathPaymentSimulate, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseVA+"payment_simulate_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: "Payment Successful",
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", pathPaymentSimulate, nil, args.payload, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponse+"internal_server_error_500.json")),
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

			if tt.wantRes != "" {
				payload := durianpay.VirtualAccountPaymentSimulatePayload{}

				if !featureWrap.DeepEqualPayload(pathPayloadVA+"payment_simulate.json", &payload, &parseArgs.payload) {
					t.Errorf("Client.PaymentSimulate() gotPayload = %v, want %v", payload, parseArgs.payload)
				}
			}

			gotRes, gotErr := c.PaymentSimulate(tt.args.ctx, tt.args.payload)
			if gotRes != tt.wantRes {
				t.Errorf("Client.PaymentSimulate() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.PaymentSimulate() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
