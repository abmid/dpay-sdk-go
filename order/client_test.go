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
	"strings"
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

func TestClient_FetchOrderByID(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		ID  string
		opt durianpay.OrderFetchByIDOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *FetchOrder
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{ctx: context.Background(), ID: "ord_wNSShKTAsL1204"},
			prepare: func(m mocks, args args) {
				url := durianpay.DURIANPAY_URL + PATCH_FETCH_ORDER_BY_ID
				url = strings.ReplaceAll(url, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_order+"fetch_order_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchOrder{
				ID:                    "ord_wNSShKTAsL1204",
				CustomerID:            "cus_dkRHbkDXrn2354",
				OrderRefID:            "order_ref_8",
				Amount:                "80000.00",
				Currency:              "IDR",
				Status:                "completed",
				IsLive:                false,
				CreatedAt:             tests.StringToTime("2023-07-27T04:12:14.084026Z"),
				UpdatedAt:             tests.StringToTime("2023-07-27T04:12:14.112331Z"),
				ExpiryDate:            tests.StringToTime("2023-08-26T04:12:14.079634Z"),
				IsNotificationEnabled: false,
				Payments: []Payment{
					{
						ID:                 "pay_sample_5b2tSNVDXn5148",
						Amount:             "80000.00",
						Status:             "completed",
						IsLive:             false,
						ExpirationDate:     tests.StringToTime("2023-07-27T04:12:30.887217Z"),
						PaymentDetailsType: "card_details",
						CreatedAt:          tests.StringToTime("2023-07-27T04:12:30.887927Z"),
						UpdatedAt:          tests.StringToTime("2023-07-27T04:12:30.887927Z"),
						RetryCount:         0,
					},
				},
			},
		},
		{
			name: "Not Found",
			args: args{
				ctx: context.Background(),
				ID:  "Wrong",
			},
			prepare: func(m mocks, args args) {
				url := durianpay.DURIANPAY_URL + PATCH_FETCH_ORDER_BY_ID
				url = strings.ReplaceAll(url, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, args.opt, nil, nil, gomock.Any()).
					Return(durianpay.FromAPI(404, featureWrap.ResJSONByte(path_response_order+"fetch_order_404.json")))
			},
			wantErr: durianpay.FromAPI(404, featureWrap.ResJSONByte(path_response_order+"fetch_order_404.json")),
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

			gotRes, gotErr := c.FetchOrderByID(tt.args.ctx, tt.args.ID, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchOrderByID() got = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchOrderByID() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_CreatePaymentLink(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.OrderPaymentLinkPayload
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
				ctx: context.Background(),
				payload: durianpay.OrderPaymentLinkPayload{
					Amount:        "20000",
					Currency:      "IDR",
					OrderRefID:    "order2314",
					IsPaymentLink: true,
					Customer: durianpay.OrderPaymentLinkCustomer{
						Email: "jude.casper@durianpay.id",
					},
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", durianpay.DURIANPAY_URL+PATH_ORDER, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_order+"create_payment_link_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &OrderCreate{
				ID:             "ord_n7WUecCLkz5074",
				CustomerID:     "cus_0l6ZMxd9cW6365",
				OrderRefID:     "order_ref_001",
				Amount:         "10001.00",
				Currency:       "IDR",
				Status:         "started",
				IsLive:         false,
				CreatedAt:      tests.StringToTime("2023-08-25T16:48:31.956761Z"),
				UpdatedAt:      tests.StringToTime("2023-08-25T16:48:31.956761Z"),
				ExpireTime:     tests.StringToTime("2023-09-24T16:48:33.010938526Z"),
				ExpiryDate:     tests.StringToTime("2023-09-24T16:48:31.953648Z"),
				AccessToken:    "5f79c29d9ba7715490d31566e1a30417da96499f5088fc090ef777ebfab53ac1",
				PaymentLinkUrl: "yUyERa",
				AdminFeeMethod: "included",
				PaymentOption:  "full_payment",
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
				payload := durianpay.OrderPaymentLinkPayload{}

				if !featureWrap.DeepEqualPayload(path_payload_order+"create_payment_link.json", &payload, &tt.args.payload) {
					t.Errorf("Client.CreatePaymentLink() gotPayload = %v, wantPayload %v", tt.args.payload, payload)
				}
			}

			gotRes, gotErr := c.CreatePaymentLink(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.CreatePaymentLink() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.CreatePaymentLink() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

// gotRes =
// &{ord_n7WUecCLkz5074 cus_0l6ZMxd9cW6365 order_ref_001  10001.00 full_payment  IDR started false 2023-08-25 16:48:31.956761 +0000 UTC 2023-08-25 16:48:31.956761 +0000 UTC { } [] 5f79c29d9ba7715490d31566e1a30417da96499f5088fc090ef777ebfab53ac1 2023-09-24 16:48:33.010938526 +0000 UTC 2023-09-24 16:48:31.953648 +0000 UTC yUyERa 0   included}, wantRes
// &{ord_n7WUecCLkz5074 cus_0l6ZMxd9cW6365 order_ref_001  10001.00 full_payment  IDR started false 2023-08-25 16:48:31.956761 +0000 UTC 2023-08-25 16:48:31.956761 +0000 UTC { } [] 5f79c29d9ba7715490d31566e1a30417da96499f5088fc090ef777ebfab53ac1 2023-09-24 16:48:33.010938526 +0000 UTC 0001-01-01 00:00:00 +0000 UTC yUyERa 0   included}
