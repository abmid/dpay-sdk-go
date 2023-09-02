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
	"testing"
	"time"

	durianpay "github.com/abmid/dpay-sdk-go"
	"github.com/abmid/dpay-sdk-go/internal/tests"
	mock_common "github.com/abmid/dpay-sdk-go/internal/tests/mock"
	"github.com/golang/mock/gomock"
)

const (
	path_response_settlement = "../internal/tests/response/settlement/"
	path_response            = "../internal/tests/response/"
	path_payload_settlement  = "../internal/tests/payload/settlement/"
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
					From: time.Now().Unix(),
					To:   time.Now().Unix(),
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", PATH_SETTLEMENT, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_settlement+"fetch_settlements_200.json"), response)
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
					Req(gomock.Any(), "GET", PATH_SETTLEMENT, args.opt, nil, nil, gomock.Any()).
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