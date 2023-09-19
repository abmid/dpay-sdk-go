/*
 * File Created: Monday, 18th September 2023 11:34:03 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package invoice

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
	dirResponse        = "../internal/tests/response/"
	dirResponseInvoice = dirResponse + "invoice/"
	dirPayloadInvoice  = "../internal/tests/payload/invoice/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_Create(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.InvoiceCreatePayload
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
				payload: durianpay.InvoiceCreatePayload{
					Amount:          "20000.67",
					RemainingAmount: "5000.67",
					Title:           "sample",
					InvoiceRefID:    "inv_ref_001",
					Customer: durianpay.Customer{
						CustomerRefID: "cust_001",
						GivenName:     "Jane Doe",
						Email:         "jane_doe@nomail.com",
						Mobile:        "85722173217",
						Address: durianpay.CustomerAddress{
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
					EnablePartialTransaction: true,
					PartialTransactionConfig: map[string]any{
						"min_acceptable_amount": 10000,
					},
					StartDate: tests.StringToTime("2022-11-22T10:00:00.000Z"),
					DueDate:   tests.StringToTime("2022-11-22T10:00:00.000Z"),
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", urlInvoice, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(dirResponseInvoice+"create_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Create{
				ID:                       "inv_2J17tdoUed3468",
				InvoiceRefID:             "inv_ref_001",
				Title:                    "sample",
				Status:                   "outstanding",
				Amount:                   "20001",
				RemainingAmount:          "5000.67",
				DueDate:                  tests.StringToTime("2023-09-19T10:00:00Z"),
				StartDate:                tests.StringToTime("2023-09-18T10:00:00Z"),
				CreatedAt:                tests.StringToTime("2023-09-17T05:06:41.816313Z"),
				CustomerID:               "cus_xWI6twzZbr7065",
				EnablePartialTransaction: true,
				PartialTransactionConfig: map[string]any{
					"min_acceptable_amount": 10000,
				},
				CheckoutURL:         "/Twkx37",
				CheckoutURLExpiryAt: tests.StringToTime("2023-12-16T12:06:41.900592201+07:00"),
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "POST", urlInvoice, nil, args.payload, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(dirResponseInvoice+"internal_server_error_500.json")))
			},
			wantErr: &durianpay.Error{
				StatusCode:   500,
				Error:        "error creating invoice",
				ResponseCode: "0005",
			},
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
				payload := durianpay.InvoiceCreatePayload{}

				if !featureWrap.DeepEqualPayload(dirPayloadInvoice+"create.json", &payload, &tt.args.payload) {
					t.Errorf("Client.Create() gotPayload = %v, wantPayload %v", tt.args.payload, payload)
				}
			}

			gotRes, gotErr := c.Create(parseArgs.ctx, parseArgs.payload)
			if !featureWrap.DeepEqualResponse(gotRes, tt.wantRes) {
				t.Errorf("Client.Create() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Create() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}