/*
 * File Created: Sunday, 3rd September 2023 10:52:42 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package promo

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
	path_response_promo = "../internal/tests/response/promo/"
	path_response       = "../internal/tests/response/"
	path_payload_promo  = "../internal/tests/payload/promo/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_Create(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.PromoPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *Promo
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PromoPayload{
					Type:     "card_promos",
					Label:    "SALE502022",
					Currency: "IDR",
					PromoDetails: durianpay.PromoDetails{
						BinList:   []int{424242},
						BankCodes: []string{},
					},
					DiscountType:       "percentage",
					Discount:           "10",
					StartsAt:           tests.StringToTime("2022-02-24T18:30:00.000Z"),
					EndsAt:             tests.StringToTime("2022-02-27T18:30:00.000Z"),
					SubType:            "direct_discount",
					LimitType:          "quota",
					LimitValue:         "100",
					PriceDeductionType: "total_price",
					Code:               "SALE2022",
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PROMO, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_promo+"create_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Promo{
				Currency:     "IDR",
				Label:        "SALE502022",
				StartsAt:     tests.StringToTime("2022-02-24T18:30:00Z"),
				EndsAt:       tests.StringToTime("2022-02-27T18:30:00Z"),
				Discount:     "10",
				DiscountType: "percentage",
				Type:         "card_promos",
				PromoDetails: PromoDetails{
					PromoID:   "prm_3eTlttAEF84045",
					BinList:   []int{424242},
					BankCodes: []string{},
				},
				SubType:            "direct_discount",
				LimitType:          "quota",
				LimitValue:         "100",
				PriceDeductionType: "total_price",
				Status:             "expired",
				CreatedAt:          tests.StringToTime("2023-09-03T03:27:06.882206Z"),
				UpdatedAt:          tests.StringToTime("2023-09-03T03:27:06.882206Z"),
				IsLive:             false,
				ID:                 "prm_3eTlttAEF84045",
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PROMO, nil, args.payload, nil, gomock.Any()).
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
				payload := durianpay.PromoPayload{}

				if !featureWrap.DeepEqualPayload(path_payload_promo+"create.json", &payload, &parseArgs.payload) {
					t.Errorf("Client.Create() gotPayload = %v, wantPayload %v", payload, parseArgs.payload)
				}
			}

			gotRes, gotErr := c.Create(parseArgs.ctx, parseArgs.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Create() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Create() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchPromos(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes []Promo
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", PATH_PROMO, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_promo+"fetch_promos_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: []Promo{
				{
					Currency:     "IDR",
					Label:        "SALE502022",
					StartsAt:     tests.StringToTime("2022-02-24T18:30:00Z"),
					EndsAt:       tests.StringToTime("2022-02-27T18:30:00Z"),
					Discount:     "10",
					DiscountType: "percentage",
					Type:         "card_promos",
					PromoDetails: PromoDetails{
						PromoID:   "prm_3eTlttAEF84045",
						BinList:   []int{424242},
						BankCodes: []string{},
					},
					SubType:            "direct_discount",
					LimitType:          "quota",
					LimitValue:         "100",
					PriceDeductionType: "total_price",
					Status:             "expired",
					CreatedAt:          tests.StringToTime("2023-09-03T03:27:06.882206Z"),
					UpdatedAt:          tests.StringToTime("2023-09-03T03:27:06.882206Z"),
					IsLive:             false,
					ID:                 "prm_3eTlttAEF84045",
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
					Req(gomock.Any(), "GET", PATH_PROMO, nil, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.FetchPromos(parseArgs.ctx)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchPromos() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchPromos() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
