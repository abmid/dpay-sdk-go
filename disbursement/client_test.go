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
