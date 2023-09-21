/*
 * File Created: Sunday, 30th July 2023 12:43:19 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package disbursement

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"testing"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/common"
	"github.com/abmid/dpay-sdk-go/internal/tests"
	mock_common "github.com/abmid/dpay-sdk-go/internal/tests/mock"
	"github.com/golang/mock/gomock"
)

const (
	pathResponse             = "../internal/tests/response/"
	pathResponseDisbursement = pathResponse + "disbursement/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_DisbursementValidate(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.DisbursementValidatePayload
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *DisbursementValidate
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.DisbursementValidatePayload{
					XIdempotencyKey: "1",
					AccountNumber:   "123",
					BankCode:        "bca",
				},
			},
			prepare: func(m mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")

				m.api.EXPECT().
					Req(gomock.Any(), "POST", durianpay.DurianpayURL+pathValidate, nil, args.payload, headers, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseDisbursement+"validate_disbursement_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &DisbursementValidate{
				AccountNumber: "123737383830",
				BankCode:      "bca",
				Status:        "processing",
			},
		},
		{
			name: "Invalid requests",
			args: args{
				ctx: context.Background(),
				payload: durianpay.DisbursementValidatePayload{
					XIdempotencyKey: "1",
					AccountNumber:   "123",
					BankCode:        "bca",
				},
			},
			prepare: func(m mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")

				m.api.EXPECT().
					Req(gomock.Any(), "POST", durianpay.DurianpayURL+pathValidate, nil, args.payload, headers, gomock.Any()).
					Return(&durianpay.Error{
						Error:     "error reading request body",
						ErrorCode: "DPAY_INTERNAL_ERROR",
					})
			},
			wantErr: &durianpay.Error{
				Error:     "error reading request body",
				ErrorCode: "DPAY_INTERNAL_ERROR",
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

			gotRes, gotErr := c.Validate(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Validate() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Validate() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_Submit(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.DisbursementPayload
		opt     *durianpay.DisbursementOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes *Disbursement
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.DisbursementPayload{
					XIdempotencyKey: "x-123",
					IdempotencyKey:  "0d5cb9a6-2488-4c86-1000-1502",
					Name:            "Test Submit",
					Description:     "Desc test submit",
					Items: []durianpay.DisbursementItemPayload{
						{
							AccountOwnerName: "Abdul Hamid",
							BankCode:         "bca",
							Amount:           "10000",
							AccountNumber:    "8422647",
							EmailRecipient:   "abdul.surel@gmail.com",
							PhoneNumber:      "081234567890",
							Notes:            "test notes",
						},
					},
				},
				opt: &durianpay.DisbursementOption{
					ForceDisburse:  tests.ToPtr(true),
					SkipValidation: tests.ToPtr(false),
				},
			},
			prepare: func(mock mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, args.payload.IdempotencyKey)

				mock.api.EXPECT().
					Req(args.ctx, "POST", durianpay.DurianpayURL+pathSubmit, args.opt, args.payload, headers, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseDisbursement+"disbursement_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Disbursement{
				ID:                 "dis_LjxhDKq8Am3427",
				IdempotencyKey:     "0d5cb9a6-2488-4c86-1000-1502",
				Name:               "test disb",
				TotalAmount:        "20000",
				TotalDisbursements: 2,
				Description:        "description",
			},
			wantErr: nil,
		},
		{
			name: "Invalid Requests (Validation)",
			args: args{
				ctx: context.Background(),
				payload: durianpay.DisbursementPayload{
					XIdempotencyKey: "x-123",
					IdempotencyKey:  "0d5cb9a6-2488-4c86-1000-1502",
					Name:            "Test Submit",
					Description:     "Desc test submit",
					Items: []durianpay.DisbursementItemPayload{
						{
							AccountOwnerName: "Abdul Hamid",
							BankCode:         "bca",
							AccountNumber:    "8422647",
							EmailRecipient:   "abdul.surel@gmail.com",
							PhoneNumber:      "081234567890",
							Notes:            "test notes",
						},
					},
				},
			},
			prepare: func(mock mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, args.payload.IdempotencyKey)

				mock.api.EXPECT().
					Req(args.ctx, "POST", durianpay.DurianpayURL+pathSubmit, args.opt, args.payload, headers, gomock.Any()).
					Return(durianpay.FromAPI(400, featureWrap.ResJSONByte(pathResponseDisbursement+"disbursement_400.json")))
			},
			wantErr: durianpay.FromAPI(400, featureWrap.ResJSONByte(pathResponseDisbursement+"disbursement_400.json")),
		},
		{
			name: "Invalid Requests (Idempotency Key Exists)",
			args: args{
				ctx: context.Background(),
				payload: durianpay.DisbursementPayload{
					XIdempotencyKey: "x-123",
					IdempotencyKey:  "0d5cb9a6-2488-4c86-1000-1502",
					Name:            "Test Submit",
					Description:     "Desc test submit",
				},
			},
			prepare: func(mock mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, args.payload.IdempotencyKey)

				mock.api.EXPECT().
					Req(args.ctx, "POST", durianpay.DurianpayURL+pathSubmit, args.opt, args.payload, headers, gomock.Any()).
					Return(durianpay.FromAPI(403, featureWrap.ResJSONByte(pathResponseDisbursement+"disbursement_403.json")))
			},
			wantErr: durianpay.FromAPI(403, featureWrap.ResJSONByte(pathResponseDisbursement+"disbursement_403.json")),
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(mock mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, args.payload.IdempotencyKey)

				mock.api.EXPECT().
					Req(args.ctx, "POST", durianpay.DurianpayURL+pathSubmit, args.opt, args.payload, headers, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponseDisbursement+"disbursement_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponseDisbursement+"disbursement_500.json")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := mock_common.NewMockApi(featureWrap.Ctrl)
			parseArgs := tt.args

			tt.prepare(mocks{api: apiMock}, parseArgs)

			c := &Client{
				ServerKey: featureWrap.ServerKey,
				Api:       apiMock,
			}

			gotRes, gotErr := c.Submit(tt.args.ctx, tt.args.payload, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Submit() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Submit() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_Approve(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type fields struct {
		Api common.Api
	}
	type args struct {
		ctx     context.Context
		payload durianpay.DisbursementApprovePayload
		opt     *durianpay.DisbursementApproveOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes *Disbursement
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.DisbursementApprovePayload{
					XIdempotencyKey: "x-123",
					ID:              "dis_XXXX",
				},
				opt: &durianpay.DisbursementApproveOption{
					IgnoreInvalid: tests.ToPtr(true),
				},
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathApprove
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")
				url = strings.ReplaceAll(url, ":id", args.payload.ID)

				mock.api.EXPECT().
					Req(args.ctx, "POST", url, args.opt, args.payload, headers, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseDisbursement+"approve_disbursement_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Disbursement{
				ID:                 "dis_XXXX",
				Name:               "sample disbursement",
				Type:               "batch",
				Status:             "approved",
				TotalAmount:        "10000.00",
				TotalDisbursements: 1,
				Description:        "this is a sample disbursement",
			},
			wantErr: nil,
		},
		{
			name: "Invalid Request (Payload)",
			args: args{
				ctx: context.Background(),
				payload: durianpay.DisbursementApprovePayload{
					ID: "dis_xxx",
				},
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathApprove
				headers := common.HeaderIdempotencyKey("", "")
				url = strings.ReplaceAll(url, ":id", args.payload.ID)

				mock.api.EXPECT().Req(args.ctx, "POST", url, args.opt, args.payload, headers, gomock.Any()).
					Return(durianpay.FromAPI(400, featureWrap.ResJSONByte(pathResponseDisbursement+"approve_disbursement_400.json")))
			},
			wantErr: durianpay.FromAPI(400, featureWrap.ResJSONByte(pathResponseDisbursement+"approve_disbursement_400.json")),
		},
		{
			name: "Invalid Request (Already Submit)",
			args: args{
				ctx: context.Background(),
				payload: durianpay.DisbursementApprovePayload{
					XIdempotencyKey: "x-123",
					ID:              "dis_xxx",
				},
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathApprove
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")
				url = strings.ReplaceAll(url, ":id", args.payload.ID)

				mock.api.EXPECT().Req(args.ctx, "POST", url, args.opt, args.payload, headers, gomock.Any()).
					Return(durianpay.FromAPI(409, featureWrap.ResJSONByte(pathResponseDisbursement+"approve_disbursement_409.json")))
			},
			wantErr: durianpay.FromAPI(409, featureWrap.ResJSONByte(pathResponseDisbursement+"approve_disbursement_409.json")),
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

			gotRes, gotErr := c.Approve(tt.args.ctx, tt.args.payload, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Approve() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Approve() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchItemsByID(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		ID  string
		opt *durianpay.DisbursementFetchItemsOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes *DisbursementItem
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				ID:  "dis_xxx",
				opt: &durianpay.DisbursementFetchItemsOption{
					Skip:  10,
					Limit: 10,
				},
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathFetchItemsByID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().
					Req(args.ctx, "GET", url, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseDisbursement+"fetch_disbursement_items_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &DisbursementItem{
				SubmissionStatus: "completed",
				Count:            1,
				DisbursementBatchItems: []DisbursementBatchItem{
					{
						ID:                  "dis_item_XXXXX",
						DisbursementBatchID: "dis_XXXXX",
						AccountOwnerName:    "John Doe",
						RealName:            "Dummy Name",
						BankCode:            "bca",
						Amount:              "10000",
						AccountNumber:       "8422647",
						EmailRecipient:      "john@nomail.com",
						PhoneNumber:         "85609873209",
						InvalidFields: []DisbursementBatchItemInvalidField{
							{
								Key:     "bank_code",
								Message: "Invalid BankCode/AccountNumber",
							},
						},
						Status:    "invalid",
						Notes:     "salary",
						IsDeleted: false,
						CreatedAt: tests.StringToTime("2021-05-03T13:54:28.842634Z"),
						UpdatedAt: tests.StringToTime("2021-05-03T13:54:28.842635Z"),
						SplitID:   "",
						Receipt:   "",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
				ID:  "dis_xxx",
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathFetchItemsByID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().Req(args.ctx, "GET", url, args.opt, nil, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponseDisbursement+"fetch_disbursement_items_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponseDisbursement+"fetch_disbursement_items_500.json")),
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

			gotRes, gotErr := c.FetchItemsByID(tt.args.ctx, tt.args.ID, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchItemsByID() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchItemsByID() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchByID(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		ID  string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes *Disbursement
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				ID:  "dis_xxx",
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathFetchByID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().
					Req(args.ctx, "GET", url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseDisbursement+"fetch_disbursement_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Disbursement{
				ID:                 "dis_XXXXXXX",
				Name:               "sample disbursement",
				Type:               "batch",
				Status:             "approved",
				TotalAmount:        "10000.00",
				TotalDisbursements: 1,
				Description:        "this is a sample description",
				Fees:               4000,
				CreatedAt:          tests.StringToTime("2021-05-03T12:57:07.296575Z"),
			},
			wantErr: nil,
		},
		{
			name: "Server Internal Error",
			args: args{
				ctx: context.TODO(),
				ID:  "dis_xxx",
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathFetchByID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().Req(args.ctx, "GET", url, nil, nil, nil, gomock.Any()).
					Return(durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponseDisbursement+"fetch_disbursement_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(pathResponseDisbursement+"fetch_disbursement_500.json")),
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

			gotRes, gotErr := c.FetchByID(parseArgs.ctx, parseArgs.ID)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchByID() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchByID() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_Delete(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		ID  string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes string
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				ID:  "dis_xxx",
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathFetchByID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().
					Req(args.ctx, http.MethodDelete, url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseDisbursement+"delete_disbursement_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: "Deleted Successfully",
			wantErr: nil,
		},
		{
			name: "Invalid Requests",
			args: args{
				ctx: context.TODO(),
				ID:  "dis_xxx",
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathFetchByID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().Req(args.ctx, http.MethodDelete, url, nil, nil, nil, gomock.Any()).
					Return(durianpay.FromAPI(403, featureWrap.ResJSONByte(pathResponseDisbursement+"delete_disbursement_403.json")))
			},
			wantRes: "",
			wantErr: durianpay.FromAPI(403, featureWrap.ResJSONByte(pathResponseDisbursement+"delete_disbursement_403.json")),
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

			gotRes, gotErr := c.Delete(parseArgs.ctx, parseArgs.ID)
			if gotRes != tt.wantRes {
				t.Errorf("Client.Delete() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Delete() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchBanks(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes []DisbursementBank
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{ctx: context.Background()},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathFetchBanks

				mock.api.EXPECT().
					Req(args.ctx, http.MethodGet, url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseDisbursement+"fetch_bank_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: []DisbursementBank{
				{
					ID:        1,
					Name:      "BCA",
					Code:      "BCA",
					Type:      "disbursement",
					CreatedAt: tests.StringToTime("2021-03-18T09:51:13.13246Z"),
					UpdatedAt: tests.StringToTime("2021-03-18T09:51:13.13246Z"),
				},
				{
					ID:        2,
					Name:      "MANDIRI",
					Code:      "MANDIRI",
					Type:      "disbursement",
					CreatedAt: tests.StringToTime("2021-03-18T09:51:13.13246Z"),
					UpdatedAt: tests.StringToTime("2021-03-18T09:51:13.13246Z"),
				},
			},
		},
		{
			name: "Internal Server Error",
			args: args{ctx: context.Background()},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathFetchBanks

				mock.api.EXPECT().Req(args.ctx, http.MethodGet, url, nil, nil, nil, gomock.Any()).
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

			gotRes, gotError := c.FetchBanks(tt.args.ctx)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchBanks() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotError, tt.wantErr) {
				t.Errorf("Client.FetchBanks() gotError = %v, want %v", gotError, tt.wantErr)
			}
		})
	}
}

func TestClient_TopupAmount(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.DisbursementTopupPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes *DisbursementTopup
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{ctx: context.Background()},
			prepare: func(mock mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")
				url := durianpay.DurianpayURL + pathTopupAmount

				mock.api.EXPECT().
					Req(args.ctx, http.MethodPost, url, nil, args.payload, headers, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseDisbursement+"topup_amount_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &DisbursementTopup{
				SenderBank:  "bni",
				TotalAmount: "10000",
				Status:      "processing",
				ExpiryDate:  tests.StringToTime("2021-03-21T09:58:53Z"),
				TransferTo: DisbursementTopupTransferTo{
					BankCode:          "bni",
					BankName:          "BNI / BNI Syariah",
					AtmBersamaCode:    "009",
					BankAccountNumber: "0437051936",
					AccountHolderName: "PT Fliptech Lentera Inspirasi Pertiwi",
					UniqueCode:        10,
				},
			},
		},
		{
			name: "Invalid Request",
			args: args{ctx: context.Background()},
			prepare: func(mock mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")
				url := durianpay.DurianpayURL + pathTopupAmount

				mock.api.EXPECT().Req(args.ctx, http.MethodPost, url, nil, args.payload, headers, gomock.Any()).
					Return(durianpay.FromAPI(400, featureWrap.ResJSONByte(pathResponseDisbursement+"topup_amount_400.json")))
			},
			wantErr: durianpay.FromAPI(400, featureWrap.ResJSONByte(pathResponseDisbursement+"topup_amount_400.json")),
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

			gotRes, gotErr := c.TopupAmount(parseArgs.ctx, parseArgs.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.TopupAmount() got = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.TopupAmount() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchBalance(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes *int
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{ctx: context.Background()},
			prepare: func(mock mocks, args args) {
				url := durianpay.DurianpayURL + pathFetchBalance

				mock.api.EXPECT().
					Req(args.ctx, http.MethodGet, url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(pathResponseDisbursement+"fetch_balance_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: tests.ToPtr(949859471313),
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

			gotRes, gotErr := c.FetchBalance(tt.args.ctx)
			if *gotRes != *tt.wantRes {
				t.Errorf("Client.FetchBalance() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchBalance() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
