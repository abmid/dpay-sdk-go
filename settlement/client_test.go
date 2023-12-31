/*
 * File Created: Saturday, 2nd September 2023 3:34:48 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package settlement

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
	pathResponse           = "../internal/tests/response/"
	pathResponseSettlement = pathResponse + "/settlement/"
	pathPayloadSettlement  = "../internal/tests/payload/settlement/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_FetchSettlements(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		opt durianpay.SettlementOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *FetchSettlements
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				opt: durianpay.SettlementOption{
					From: time.Now().Format("2006-01-02"),
					To:   time.Now().Format("2006-01-02"),
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", pathSettlement, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseSettlement+"fetch_settlements_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchSettlements{
				TotalCount: 1,
				SettlementDetail: []Settlement{
					{
						ID:                     "set_WDizQUoyWy1234",
						SettlementAmount:       "20000.00",
						Status:                 "settled",
						Fee:                    "200.00",
						TotalTransactionAmount: "20200.00",
						CreatedAt:              tests.StringToTime("2021-05-17T08:30:56.73529Z"),
						SettledAt:              tests.StringToTime("2021-05-17T08:32:00.628182Z"),
						Currency:               "IDR",
					},
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
					Req(gomock.Any(), "GET", pathSettlement, args.opt, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.FetchSettlements(tt.args.ctx, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchSettlements() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchSettlements() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchDetails(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		opt durianpay.SettlementOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *FetchDetails
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				opt: durianpay.SettlementOption{
					From: time.Now().Format("2006-01-02"),
					To:   time.Now().Format("2006-01-02"),
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", pathDetail, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseSettlement+"details_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchDetails{
				SettlementCount:  1,
				TransactionCount: 1,
				SettlementDetail: []SettlementDetail{
					{
						SettlementID:       "set_WDizQUoyWy1234",
						PaymentID:          "pay_O7laHvW2BI1234",
						OrderID:            "ord_KlQRqodMqq0540",
						OrderReference:     "order_ref_001",
						Status:             "settled",
						Currency:           "IDR",
						SettlementAmount:   "19635.00",
						TotalSettlementFee: "365.00",
						SettledAt:          tests.StringToTime("2021-05-17T08:32:00.628182Z"),
						PaymentAmount:      "20000.00",
						PaymentDate:        tests.StringToTime("2021-05-17T08:26:43.990125Z"),
						TransactionAmount:  "20000.00",
						PaymentDetailsType: "ewallet_details",
						PaymentMethodID:    "SHOPEEPAY",
					},
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
					Req(gomock.Any(), "GET", pathDetail, args.opt, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.FetchDetails(parseArgs.ctx, parseArgs.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchDetails() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchDetails() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_StatusByPaymentID(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx       context.Context
		paymentID string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *SettlementDetail
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx:       context.Background(),
				paymentID: "pay_O7laHvW2BI5527",
			},
			prepare: func(m mocks, args args) {
				params := struct {
					PaymentID string `url:"payment_id"`
				}{PaymentID: args.paymentID}

				m.api.EXPECT().
					Req(gomock.Any(), "GET", pathDetail, params, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseSettlement+"status_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &SettlementDetail{
				SettlementID:       "set_WDizQUoyWy8680",
				PaymentID:          "pay_O7laHvW2BI5527",
				OrderID:            "ord_KlQRqodMqq0540",
				OrderReference:     "order_ref_001",
				Status:             "settled",
				Currency:           "IDR",
				SettlementAmount:   "19635.00",
				TotalSettlementFee: "365.00",
				SettledAt:          tests.StringToTime("2021-05-17T08:32:00.628182Z"),
				Group:              "A",
				PaymentAmount:      "20000.00",
				PaymentDate:        tests.StringToTime("2021-05-17T08:26:43.990125Z"),
				TransactionAmount:  "20000.00",
				PaymentChannel:     "ewallet_details",
				PaymentSubchannel:  "SHOPEEPAY",
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", pathDetail, gomock.Any(), nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.StatusByPaymentID(tt.args.ctx, tt.args.paymentID)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.StatusByPaymentID() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.StatusByPaymentID() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchSettlementByID(t *testing.T) {
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
		wantRes *Settlement
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				ID:  "set_WDizQUoyWy8680",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(pathFetchByID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseSettlement+"fetch_settlement_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Settlement{
				ID:               "set_WDizQUoyWy8680",
				SettlementAmount: "50000",
				Status:           "settled",
				Fee:              "10000",
				CreatedAt:        tests.StringToTime("2021-05-17T08:32:00.628182Z"),
				SettledAt:        tests.StringToTime("2021-05-27T08:32:00.628182Z"),
				PromoAmount:      "10000",
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
				ID:  "set_WDizQUoyWy8680",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(pathFetchByID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, gomock.Any(), nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.FetchSettlementByID(tt.args.ctx, tt.args.ID)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchSettlementByID() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchSettlementByID() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
