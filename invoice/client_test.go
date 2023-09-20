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
	"strings"
	"testing"
	"time"

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

func TestClient_GenerateCheckoutURL(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx        context.Context
		customerID string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *GenerateCheckoutURL
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx:        context.Background(),
				customerID: "cus_ViPeX4iBYp2233",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(urlGenerateCheckoutURL, ":customer_id", args.customerID)

				m.api.EXPECT().
					Req(gomock.Any(), "POST", url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(dirResponseInvoice+"generate_url_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &GenerateCheckoutURL{
				URL:    "/Twkx37",
				Expiry: tests.StringToTime("2023-12-16T12:06:41+07:00"),
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(urlGenerateCheckoutURL, ":customer_id", args.customerID)

				m.api.EXPECT().
					Req(gomock.Any(), "POST", url, nil, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.GenerateCheckoutURL(tt.args.ctx, tt.args.customerID)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.GenerateCheckoutURL() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.GenerateCheckoutURL() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchInvoiceByID(t *testing.T) {
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
		wantRes *FetchInvoiceByID
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				ID:  "inv_2J17tdoUed3468",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(urlFetchByID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(dirResponseInvoice+"fetch_invoice_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchInvoiceByID{
				ID:                          "inv_2J17tdoUed3468",
				InvoiceRefID:                "inv_ref_001",
				CustomerID:                  "cus_xWI6twzZbr7065",
				IsLive:                      false,
				Title:                       "sample",
				Status:                      "outstanding",
				Amount:                      "20001",
				RemainingAmount:             "5001",
				StartDate:                   tests.StringToTime("2023-09-18T10:00:00Z"),
				DueDate:                     tests.StringToTime("2023-09-19T10:00:00Z"),
				CreatedAt:                   tests.StringToTime("2023-09-17T05:06:41.816313Z"),
				IsPartialTransactionEnabled: true,
				PartialTransactionConfig: map[string]any{
					"min_acceptable_amount": 10000,
				},
				Transactions: []Transaction{
					{
						ID:     "inv_txn_sAMwRDEqcE0554",
						Amount: "15000",
						Status: "paid_manually",
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
				url := strings.ReplaceAll(urlFetchByID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, nil, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.FetchInvoiceByID(tt.args.ctx, tt.args.ID)
			if !featureWrap.DeepEqualResponse(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchInvoiceByID() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchInvoiceByID() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchInvoices(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		opt durianpay.InvoiceFetchOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *FetchInvoices
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				opt: durianpay.InvoiceFetchOption{
					From:   time.Now().Format("2006-01-02"),
					To:     time.Now().Format("2006-01-02"),
					Skip:   1,
					Limit:  10,
					Status: "paid",
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", urlInvoice, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(dirResponseInvoice+"fetch_invoices_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchInvoices{
				Invoices: []Invoices{
					{
						ID:                          "inv_2J17tdoUed3468",
						InvoiceRefID:                "inv_ref_001",
						CustomerID:                  "cus_xWI6twzZbr7065",
						Status:                      "outstanding",
						Title:                       "sample",
						StartDate:                   tests.StringToTime("2023-09-18T10:00:00Z"),
						DueDate:                     tests.StringToTime("2023-09-19T10:00:00Z"),
						CreatedAt:                   tests.StringToTime("2023-09-17T05:06:41.816313Z"),
						IsPartialTransactionEnabled: true,
						PartialTransactionConfig: map[string]any{
							"min_acceptable_amount": 10000,
						},
						Amount:          "20001",
						RemainingAmount: "5001",
						IsLive:          false,
						IsBlocked:       false,
					},
				},
				TotalCount: 1,
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", urlInvoice, args.opt, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.FetchInvoices(tt.args.ctx, tt.args.opt)
			if !featureWrap.DeepEqualResponse(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchInvoices() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchInvoices() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_Update(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		ID      string
		payload durianpay.InvoiceUpdatePayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *Update
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.InvoiceUpdatePayload{
					InvoiceRefID:             "inv_ref_001",
					RemainingAmount:          "5000.67",
					Title:                    "sample invoice",
					EnablePartialTransaction: true,
					PartialTransactionConfig: map[string]any{
						"min_acceptable_amount": 10000,
					},
					StartDate: tests.StringToTime("2023-03-28T00:00:00.000Z"),
					DueDate:   tests.StringToTime("2023-03-29T00:00:00.000Z"),
					Metadata: map[string]any{
						"invoice_type": "internal",
					},
				},
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(urlUpdateByID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "PUT", url, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(dirResponseInvoice+"update_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Update{
				ID:                       "inv_2J17tdoUed3468",
				InvoiceRefID:             "inv_ref_001",
				CustomerID:               "cus_xWI6twzZbr7065",
				Title:                    "sample invoice",
				Status:                   "outstanding",
				Amount:                   "20001",
				RemainingAmount:          "5001",
				StartDate:                tests.StringToTime("2023-09-19T00:00:00Z"),
				DueDate:                  tests.StringToTime("2023-09-20T00:00:00Z"),
				CreatedAt:                tests.StringToTime("2023-09-17T05:06:41.816313Z"),
				UpdatedAt:                tests.StringToTime("2023-09-17T05:14:10.107641Z"),
				EnablePartialTransaction: true,
				PartialTransactionConfig: map[string]any{
					"min_acceptable_amount": 10000,
				},
				Metadata: map[string]any{
					"invoice_type": "internal",
				},
				IsBlocked: false,
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(urlUpdateByID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "PUT", url, nil, args.payload, nil, gomock.Any()).
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
				payload := durianpay.InvoiceUpdatePayload{}

				if !featureWrap.DeepEqualPayload(dirPayloadInvoice+"update.json", &payload, &tt.args.payload) {
					t.Errorf("Client.update() gotPayload = %v, wantPayload %v", tt.args.payload, payload)
				}
			}

			gotRes, gotErr := c.Update(tt.args.ctx, tt.args.ID, tt.args.payload)
			if !featureWrap.DeepEqualResponse(gotRes, tt.wantRes) {
				t.Errorf("Client.Update() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Update() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
