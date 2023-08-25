/*
 * File Created: Thursday, 24th August 2023 6:36:25 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package order

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
	path_response_order = "../internal/tests/response/order/"
	path_response       = "../internal/tests/response/"
	path_payload_order  = "../internal/tests/payload/order/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_Create(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.OrderPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *OrderCreate
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				payload: durianpay.OrderPayload{
					Amount:        "10000.67",
					PaymentOption: "full_payment",
					Currency:      "IDR",
					OrderRefID:    "order_ref_001",
					Customer: durianpay.OrderCustomer{
						CustomerRefID: "cust_001",
						GivenName:     "Jane Doe",
						Email:         "jane_doe@nomail.com",
						Mobile:        "85722173217",
						Address: durianpay.OrderCustomerAddress{
							ReceiverName:  "Jude Casper",
							ReceiverPhone: "8987654321",
							Label:         "Home Address",
							AddressLine1:  "Jl. HR. Rasuna Said",
							AddressLine2:  "Apartment #786",
							City:          "Jakarta Selatan",
							Region:        "Jakarta",
							Country:       "Indonesia",
							PostalCode:    "560008",
							Landmark:      "Kota Jakarta Selatan",
						},
					},
					Items: []durianpay.OrderItem{
						{
							Name:  "LED Television",
							Qty:   1,
							Price: "10001.00",
							Logo:  "https://merchant.com/tv_image.jpg",
						},
					},
					Metadata: durianpay.OrderMetadata{
						MyMetaKey:       "my-meta-value",
						SettlementGroup: "BranchName",
					},
					ExpiryDate: tests.StringToTime("2022-03-29T10:00:00.000Z"),
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", durianpay.DURIANPAY_URL+PATH_ORDER, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_order+"create_order_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &OrderCreate{
				ID:            "ord_0dIWbuDJQ84078",
				CustomerID:    "cus_ViPeX4iBYp2233",
				OrderRefID:    "order_ref_001",
				Amount:        "10001.00",
				PaymentOption: "full_payment",
				Currency:      "IDR",
				Status:        "started",
				CreatedAt:     tests.StringToTime("2021-04-01T14:39:37.860426Z"),
				UpdatedAt:     tests.StringToTime("2021-04-01T14:39:37.860426Z"),
				ExpiryDate:    tests.StringToTime("2022-03-29T10:00:00Z"),
				MetaData: durianpay.OrderMetadata{
					MyMetaKey:       "my-meta-value",
					SettlementGroup: "BranchName",
				},
				Items: []durianpay.OrderItem{
					{
						Name:  "LED Television",
						Qty:   1,
						Price: "10001.00",
						Logo:  "https://merchant.com/tv_image.jpg",
					},
				},
				AccessToken:    "adsyoi12sdASd123ASX@qqsda231",
				ExpireTime:     tests.StringToTime("2022-09-15T13:40:17.064478772Z"),
				AddressID:      7526,
				AdminFeeMethod: "included",
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", durianpay.DURIANPAY_URL+PATH_ORDER, nil, args.payload, nil, gomock.Any()).
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
				payload := durianpay.OrderPayload{}

				if !featureWrap.DeepEqualPayload(path_payload_order+"create.json", &payload, &tt.args.payload) {
					t.Errorf("Client.Create() gotPayload = %v, wantPayload %v", tt.args.payload, payload)
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

func TestClient_FetchOrders(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		opt durianpay.OrderFetchOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *FetchOrders
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{ctx: context.Background(), opt: durianpay.OrderFetchOption{Skip: 1}},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", durianpay.DURIANPAY_URL+PATH_ORDER, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_order+"fetch_orders_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchOrders{
				Orders: []Orders{
					{
						ID:                    "ord_jcI3YWlYbD5367",
						CustomerID:            "cus_IwDIb0MDY20938",
						Amount:                "10000.00",
						Currency:              "IDR",
						Status:                "completed",
						IsLive:                false,
						CreatedAt:             tests.StringToTime("2023-07-27T04:16:33.380918Z"),
						UpdatedAt:             tests.StringToTime("2023-07-27T04:24:29.382684Z"),
						ExpiryDate:            tests.StringToTime("2023-07-28T11:16:17.37Z"),
						GivenName:             "Abdul",
						Email:                 "abdul.surel@gmail.com",
						Mobile:                "811111111",
						PaymentOption:         "full_payment",
						PaymentID:             "pay_y2yKEEWBYe1299",
						PaymentDetailsType:    "va_details",
						PaymentStatus:         "completed",
						PaymentDate:           tests.StringToTime("2023-07-27T04:17:42.997975Z"),
						Description:           "Test Description",
						PaymentLinkUrl:        "VGXbuJ",
						IsNotificationEnabled: false,
						PaymentMethodID:       "BCA",
					},
				},
				Count: 1,
			},
		},
		{
			name: "Invalid Options",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", durianpay.DURIANPAY_URL+PATH_ORDER, args.opt, nil, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response_order+"fetch_orders_400.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response_order+"fetch_orders_400.json")),
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

			gotRes, gotErr := c.FetchOrders(tt.args.ctx, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchOrders() got = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchOrders() got1 = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
