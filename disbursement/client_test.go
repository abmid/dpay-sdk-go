/*
 * File Created: Sunday, 30th July 2023 12:43:19 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package disbursement

import (
	"context"
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
	path_response_disbursement = "../internal/tests/response/disbursement/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_ValidateDisbursement(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload durianpay.ValidateDisbursementPayload
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *durianpay.ValidateDisbursement
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.ValidateDisbursementPayload{
					XIdempotencyKey: "1",
					AccountNumber:   "123",
					BankCode:        "bca",
				},
			},
			prepare: func(m mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")
				m.api.EXPECT().Req(gomock.Any(), "POST", durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_VALIDATE, nil, args.payload, headers).
					Return(featureWrap.ResJSONByte(path_response_disbursement+"validate_disbursement_200.json"), nil)
			},
			wantRes: &durianpay.ValidateDisbursement{
				Data: durianpay.ValidateDisbursementData{
					AccountNumber: "123737383830",
					BankCode:      "bca",
					Status:        "processing",
				},
			},
		},
		{
			name: "Invalid requests",
			args: args{
				ctx: context.Background(),
				payload: durianpay.ValidateDisbursementPayload{
					XIdempotencyKey: "1",
					AccountNumber:   "123",
					BankCode:        "bca",
				},
			},
			prepare: func(m mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")
				m.api.EXPECT().Req(gomock.Any(), "POST", durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_VALIDATE, nil, args.payload, headers).
					Return([]byte{}, &durianpay.Error{
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

			gotRes, gotErr := c.ValidateDisbursement(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ValidateDisbursement() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ValidateDisbursement() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_SubmitDisbursement(t *testing.T) {
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
		wantRes *durianpay.Disbursement
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
					ForceDisburse:  tests.BoolPtr(true),
					SkipValidation: tests.BoolPtr(false),
				},
			},
			prepare: func(mock mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, args.payload.IdempotencyKey)

				mock.api.EXPECT().Req(args.ctx, "POST", durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_SUBMIT, args.opt, args.payload, headers).
					Return(featureWrap.ResJSONByte(path_response_disbursement+"disbursement_200.json"), nil)
			},
			wantRes: &durianpay.Disbursement{
				Message: "request already processed",
				Data: durianpay.DisbursementData{
					ID:                 "dis_LjxhDKq8Am3427",
					IdempotencyKey:     "0d5cb9a6-2488-4c86-1000-1502",
					Name:               "test disb",
					TotalAmount:        "20000",
					TotalDisbursements: 2,
					Description:        "description",
				},
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

				mock.api.EXPECT().Req(args.ctx, "POST", durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_SUBMIT, args.opt, args.payload, headers).
					Return(nil, durianpay.FromAPI(400, featureWrap.ResJSONByte(path_response_disbursement+"disbursement_400.json")))
			},
			wantErr: durianpay.FromAPI(400, featureWrap.ResJSONByte(path_response_disbursement+"disbursement_400.json")),
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

				mock.api.EXPECT().Req(args.ctx, "POST", durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_SUBMIT, args.opt, args.payload, headers).
					Return(nil, durianpay.FromAPI(403, featureWrap.ResJSONByte(path_response_disbursement+"disbursement_403.json")))
			},
			wantErr: durianpay.FromAPI(403, featureWrap.ResJSONByte(path_response_disbursement+"disbursement_403.json")),
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(mock mocks, args args) {
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, args.payload.IdempotencyKey)

				mock.api.EXPECT().Req(args.ctx, "POST", durianpay.DURIANPAY_URL+PATH_DISBURSEMENT_SUBMIT, args.opt, args.payload, headers).
					Return(nil, durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response_disbursement+"disbursement_500.json")))
			},
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response_disbursement+"disbursement_500.json")),
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

			gotRes, gotErr := c.SubmitDisbursement(tt.args.ctx, tt.args.payload, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.SubmitDisbursement() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.SubmitDisbursement() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_ApproveDisbursement(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type fields struct {
		Api common.Api
	}
	type args struct {
		ctx     context.Context
		payload durianpay.ApproveDisbursementPayload
		opt     *durianpay.ApproveDisbursementOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes *durianpay.Disbursement
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.ApproveDisbursementPayload{
					XIdempotencyKey: "x-123",
					ID:              "dis_XXXX",
				},
				opt: &durianpay.ApproveDisbursementOption{
					IgnoreInvalid: tests.BoolPtr(true),
				},
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_APPROVE
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")
				url = strings.ReplaceAll(url, ":id", args.payload.ID)

				mock.api.EXPECT().Req(args.ctx, "POST", url, args.opt, args.payload, headers).
					Return(featureWrap.ResJSONByte(path_response_disbursement+"approve_disbursement_200.json"), nil)
			},
			wantRes: &durianpay.Disbursement{
				Message: "request already submitted",
				Data: durianpay.DisbursementData{
					ID:                 "dis_XXXX",
					Name:               "sample disbursement",
					Type:               "batch",
					Status:             "approved",
					TotalAmount:        "10000.00",
					TotalDisbursements: 1,
					Description:        "this is a sample disbursement",
				},
			},
			wantErr: nil,
		},
		{
			name: "Invalid Request (Payload)",
			args: args{
				ctx: context.Background(),
				payload: durianpay.ApproveDisbursementPayload{
					ID: "dis_xxx",
				},
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_APPROVE
				headers := common.HeaderIdempotencyKey("", "")
				url = strings.ReplaceAll(url, ":id", args.payload.ID)

				mock.api.EXPECT().Req(args.ctx, "POST", url, args.opt, args.payload, headers).
					Return(nil, durianpay.FromAPI(400, featureWrap.ResJSONByte(path_response_disbursement+"approve_disbursement_400.json")))
			},
			wantRes: nil,
			wantErr: durianpay.FromAPI(400, featureWrap.ResJSONByte(path_response_disbursement+"approve_disbursement_400.json")),
		},
		{
			name: "Invalid Request (Already Submit)",
			args: args{
				ctx: context.Background(),
				payload: durianpay.ApproveDisbursementPayload{
					XIdempotencyKey: "x-123",
					ID:              "dis_xxx",
				},
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_APPROVE
				headers := common.HeaderIdempotencyKey(args.payload.XIdempotencyKey, "")
				url = strings.ReplaceAll(url, ":id", args.payload.ID)

				mock.api.EXPECT().Req(args.ctx, "POST", url, args.opt, args.payload, headers).
					Return(nil, durianpay.FromAPI(409, featureWrap.ResJSONByte(path_response_disbursement+"approve_disbursement_409.json")))
			},
			wantRes: nil,
			wantErr: durianpay.FromAPI(409, featureWrap.ResJSONByte(path_response_disbursement+"approve_disbursement_409.json")),
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

			gotRes, gotErr := c.ApproveDisbursement(tt.args.ctx, tt.args.payload, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ApproveDisbursement() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ApproveDisbursement() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchDisbursementItemsByID(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		ID  string
		opt *durianpay.FetchDisbursementItemsOption
	}
	tests := []struct {
		name    string
		args    args
		prepare func(mock mocks, args args)
		wantRes *durianpay.DisbursementItem
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				ID:  "dis_xxx",
				opt: &durianpay.FetchDisbursementItemsOption{
					Skip:  10,
					Limit: 10,
				},
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_FETCH_ITEMS_BY_ID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().Req(args.ctx, "GET", url, args.opt, nil, nil).
					Return(featureWrap.ResJSONByte(path_response_disbursement+"fetch_disbursement_items_200.json"), nil)
			},
			wantRes: &durianpay.DisbursementItem{
				SubmissionStatus: "completed",
				Count:            1,
				DisbursementBatchItems: []durianpay.DisbursementBatchItem{
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
						InvalidFields: []durianpay.DisbursementBatchItemInvalidField{
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
				url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_FETCH_ITEMS_BY_ID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().Req(args.ctx, "GET", url, args.opt, nil, nil).
					Return(nil, durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response_disbursement+"fetch_disbursement_items_500.json")))
			},
			wantRes: nil,
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response_disbursement+"fetch_disbursement_items_500.json")),
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

			gotRes, gotErr := c.FetchDisbursementItemsByID(tt.args.ctx, tt.args.ID, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchDisbursementItemsByID() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchDisbursementItemsByID() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchDisbursementByID(t *testing.T) {
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
		wantRes *durianpay.DisbursementData
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				ID:  "dis_xxx",
			},
			prepare: func(mock mocks, args args) {
				url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_FETCH_BY_ID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().Req(args.ctx, "GET", url, nil, nil, nil).
					Return(featureWrap.ResJSONByte(path_response_disbursement+"fetch_disbursement_200.json"), nil)
			},
			wantRes: &durianpay.DisbursementData{
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
				url := durianpay.DURIANPAY_URL + PATH_DISBURSEMENT_FETCH_BY_ID
				url = strings.ReplaceAll(url, ":id", args.ID)

				mock.api.EXPECT().Req(args.ctx, "GET", url, nil, nil, nil).
					Return(nil, durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response_disbursement+"fetch_disbursement_500.json")))
			},
			wantRes: nil,
			wantErr: durianpay.FromAPI(500, featureWrap.ResJSONByte(path_response_disbursement+"fetch_disbursement_500.json")),
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

			gotRes, gotErr := c.FetchDisbursementByID(parseArgs.ctx, parseArgs.ID)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchDisbursementByID() gotRes = %v, want %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchDisbursementByID() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
