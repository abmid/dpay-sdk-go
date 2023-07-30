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
	"github.com/abmid/dpay-sdk-go/internal/tests"
	mock_common "github.com/abmid/dpay-sdk-go/internal/tests/mock"
	"github.com/golang/mock/gomock"
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
					IdempotenKey:  "1",
					AccountNumber: "123",
					BankCode:      "bca",
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().Req(gomock.Any(), "POST", durianpay.DURIANPAY_URL+PATH, args.payload, args.payload.IdempotenKey).
					Return(featureWrap.ResJSONByte("../internal/tests/response/validate_disbursement_200.json"), nil)
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
					IdempotenKey:  "1",
					AccountNumber: "123",
					BankCode:      "bca",
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().Req(gomock.Any(), "POST", durianpay.DURIANPAY_URL+PATH, args.payload, args.payload.IdempotenKey).
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
