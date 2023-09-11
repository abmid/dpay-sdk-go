/*
 * File Created: Tuesday, 5th September 2023 11:13:26 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package payment

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
	path_response_payment = "../internal/tests/response/payment/"
	path_response         = "../internal/tests/response/"
	path_payload_payment  = "../internal/tests/payload/payment/"
)

type mocks struct {
	api *mock_common.MockApi
}

func TestClient_ChargeVA(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type vaPayload struct {
		Type          string                           `json:"type"`
		Request       durianpay.PaymentChargeVAPayload `json:"request"`
		SandboxOption *durianpay.PaymentSandboxOption  `json:"sandbox_options"`
	}

	type args struct {
		ctx     context.Context
		payload durianpay.PaymentChargeVAPayload
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *ChargeVA
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PaymentChargeVAPayload{
					OrderID:      "ord_WkJWY1ysZ57194",
					BankCode:     "MANDIRI",
					Name:         "Name Appear in ATM",
					Amount:       "20000",
					PaymentRefID: "pay_ref_123",
				},
			},
			prepare: func(m mocks, args args) {
				payload := chargePayload{
					Type:          "VA",
					Request:       args.payload,
					SandboxOption: args.payload.SandboxOption,
				}

				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"charge_va_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &ChargeVA{
				Type: "VA",
				Response: chargeResponseVA{
					PaymentID:      "pay_pYQ319c4qo5956",
					OrderID:        "ord_VN5nVJpSW27112",
					AccountNumber:  "7893572945724867",
					PaymentRefID:   "pay_ref_123",
					ExpirationTime: tests.StringToTime("2023-09-05T10:00:00Z"),
					PaymentInstruction: struct {
						EN paymentInstruction "json:\"en\""
						ID paymentInstruction "json:\"ID\""
					}{
						EN: paymentInstruction{
							Atm: struct {
								Heading         string "json:\"heading\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "ATM Mandiri",
								InstructionText: "<ol><li>Insert your ATM card and select \"ENGLISH\"</li></ol>",
							},
							MobileApp: struct {
								Heading         string "json:\"heading\""
								AppStoreURL     string "json:\"appstore_url\""
								PlayStoreURL    string "json:\"playstore_url\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "Mandiri Online",
								AppStoreURL:     "<<http://onelink.to/dvs8pn",
								PlayStoreURL:    "http://onelink.to/dvs8pn",
								InstructionText: "<ol><li>Insert your ATM card and select \"ENGLISH\"</li></ol>",
							},
							InternetBanking: struct {
								Heading         string "json:\"heading\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "Internet Banking",
								InstructionText: "<ol><li>Insert your ATM card and select \"ENGLISH\"</li></ol>",
							},
						},
						ID: paymentInstruction{
							Atm: struct {
								Heading         string "json:\"heading\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "ATM Mandiri",
								InstructionText: "<ol><li>Masukkan kartu ATM dan pilih \"Bahasa Indonesia\"</li></ol>",
							},
							MobileApp: struct {
								Heading         string "json:\"heading\""
								AppStoreURL     string "json:\"appstore_url\""
								PlayStoreURL    string "json:\"playstore_url\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "Mandiri Online",
								AppStoreURL:     "<<http://onelink.to/dvs8pn",
								PlayStoreURL:    "http://onelink.to/dvs8pn",
								InstructionText: "<ol><li>Masukkan kartu ATM dan pilih \"Bahasa Indonesia\"</li></ol>",
							},
							InternetBanking: struct {
								Heading         string "json:\"heading\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "Internet Banking",
								InstructionText: "<ol><li>Masukkan kartu ATM dan pilih \"Bahasa Indonesia\"</li></ol>",
							},
						},
					},
				},
			},
		},
		{
			name: "Success with Sandbox Mode",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PaymentChargeVAPayload{
					OrderID:      "ord_WkJWY1ysZ57194",
					BankCode:     "MANDIRI",
					Name:         "Name Appear in ATM",
					Amount:       "20000",
					PaymentRefID: "pay_ref_123",
					SandboxOption: &durianpay.PaymentSandboxOption{
						ForceFail: true,
						DelayMS:   10000,
					},
				},
			},
			prepare: func(m mocks, args args) {
				payload := chargePayload{
					Type:          "VA",
					Request:       args.payload,
					SandboxOption: args.payload.SandboxOption,
				}

				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"charge_va_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &ChargeVA{
				Type: "VA",
				Response: chargeResponseVA{
					PaymentID:      "pay_pYQ319c4qo5956",
					OrderID:        "ord_VN5nVJpSW27112",
					AccountNumber:  "7893572945724867",
					PaymentRefID:   "pay_ref_123",
					ExpirationTime: tests.StringToTime("2023-09-05T10:00:00Z"),
					PaymentInstruction: struct {
						EN paymentInstruction "json:\"en\""
						ID paymentInstruction "json:\"ID\""
					}{
						EN: paymentInstruction{
							Atm: struct {
								Heading         string "json:\"heading\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "ATM Mandiri",
								InstructionText: "<ol><li>Insert your ATM card and select \"ENGLISH\"</li></ol>",
							},
							MobileApp: struct {
								Heading         string "json:\"heading\""
								AppStoreURL     string "json:\"appstore_url\""
								PlayStoreURL    string "json:\"playstore_url\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "Mandiri Online",
								AppStoreURL:     "<<http://onelink.to/dvs8pn",
								PlayStoreURL:    "http://onelink.to/dvs8pn",
								InstructionText: "<ol><li>Insert your ATM card and select \"ENGLISH\"</li></ol>",
							},
							InternetBanking: struct {
								Heading         string "json:\"heading\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "Internet Banking",
								InstructionText: "<ol><li>Insert your ATM card and select \"ENGLISH\"</li></ol>",
							},
						},
						ID: paymentInstruction{
							Atm: struct {
								Heading         string "json:\"heading\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "ATM Mandiri",
								InstructionText: "<ol><li>Masukkan kartu ATM dan pilih \"Bahasa Indonesia\"</li></ol>",
							},
							MobileApp: struct {
								Heading         string "json:\"heading\""
								AppStoreURL     string "json:\"appstore_url\""
								PlayStoreURL    string "json:\"playstore_url\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "Mandiri Online",
								AppStoreURL:     "<<http://onelink.to/dvs8pn",
								PlayStoreURL:    "http://onelink.to/dvs8pn",
								InstructionText: "<ol><li>Masukkan kartu ATM dan pilih \"Bahasa Indonesia\"</li></ol>",
							},
							InternetBanking: struct {
								Heading         string "json:\"heading\""
								InstructionText string "json:\"instruction_text\""
							}{
								Heading:         "Internet Banking",
								InstructionText: "<ol><li>Masukkan kartu ATM dan pilih \"Bahasa Indonesia\"</li></ol>",
							},
						},
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
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, gomock.Any(), nil, gomock.Any()).
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
				var filename string
				isSandbox := parseArgs.payload.SandboxOption != nil
				validPayload := vaPayload{}
				argsPayload := vaPayload{
					Type:          "VA",
					Request:       parseArgs.payload,
					SandboxOption: parseArgs.payload.SandboxOption,
				}

				if isSandbox {
					filename = "charge_sandbox.json"
				} else {
					filename = "charge_va.json"
				}

				err := json.Unmarshal(featureWrap.ResJSONByte(path_payload_payment+filename), &validPayload)
				if err != nil {
					panic(err)
				}

				if isSandbox {
					validPayload.Request.SandboxOption = validPayload.SandboxOption
				}

				if !reflect.DeepEqual(validPayload, argsPayload) {
					t.Errorf("Client.ChargeVA() validPayload = %v, argsPayload %v", validPayload, argsPayload)
				}
			}

			gotRes, gotErr := c.ChargeVA(tt.args.ctx, tt.args.payload)

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ChargeVA() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ChargeVA() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_ChargeBNPL(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type bnplPayload struct {
		Type          string                             `json:"type"`
		Request       durianpay.PaymentChargeBNPLPayload `json:"request"`
		SandboxOption *durianpay.PaymentSandboxOption    `json:"sandbox_options"`
	}

	type args struct {
		ctx     context.Context
		payload durianpay.PaymentChargeBNPLPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *ChargeBNPL
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PaymentChargeBNPLPayload{
					OrderID:               "ord_1EcWGI2xSs7216",
					Amount:                "10000.00",
					PaymentRefID:          "pay_ref_123",
					PaymentMethodUniqueID: "AKULAKU",
					CustomerInfo: durianpay.PaymentCustomerInfo{
						ID:        "cus_aGn5UD0m7F0994",
						Email:     "jude_kasper@koss.in",
						GivenName: "Jude Kasper",
					},
				},
			},
			prepare: func(m mocks, args args) {
				payload := chargePayload{
					Type:          "BNPL",
					Request:       args.payload,
					SandboxOption: args.payload.SandboxOption,
				}

				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"charge_bnpl_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &ChargeBNPL{
				Type: "BNPL",
				Response: chargeResponseBNPL{
					PaymentID:    "pay_80pgxEcUbO8054",
					OrderID:      "ord_NDmLvwTTh95152",
					PaymentRefID: "pay_ref_123",
					RedirectURL:  "https://redirect-url.com/",
					PaidAmount:   "80001.00",
					Metadata:     Metadata{},
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
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, gomock.Any(), nil, gomock.Any()).
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
				validPayload := bnplPayload{}
				argsPayload := bnplPayload{
					Type:          "BNPL",
					Request:       parseArgs.payload,
					SandboxOption: parseArgs.payload.SandboxOption,
				}

				if !featureWrap.DeepEqualPayload(path_payload_payment+"charge_bnpl.json", &validPayload, &argsPayload) {
					t.Errorf("Client.ChargeBNPL() validPayload = %v, argsPayload %v", validPayload, argsPayload)
				}
			}

			gotRes, gotErr := c.ChargeBNPL(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ChargeBNPL() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ChargeBNPL() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_ChargeEwallet(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type ewalletPayload struct {
		Type          string                                `json:"type"`
		Request       durianpay.PaymentChargeEwalletPayload `json:"request"`
		SandboxOption *durianpay.PaymentSandboxOption       `json:"sandbox_options"`
	}

	type args struct {
		ctx     context.Context
		payload durianpay.PaymentChargeEwalletPayload
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *ChargeEwallet
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PaymentChargeEwalletPayload{
					OrderID:    "ord_mJH2hKOSYb3514",
					Amount:     "20000.00",
					Mobile:     "08123456789",
					WalletType: "DANA",
				},
			},
			prepare: func(m mocks, args args) {
				payload := chargePayload{
					Type:          "EWALLET",
					Request:       args.payload,
					SandboxOption: args.payload.SandboxOption,
				}

				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"charge_ewallet_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &ChargeEwallet{
				Type: "EWALLET",
				Response: chargeResponseEwallet{
					PaymentID:      "pay_PoVnlDmGts4956",
					OrderID:        "ord_VN5nVJpSW27112",
					Mobile:         "08123456789",
					Status:         "processing",
					ExpirationTime: tests.StringToTime("0001-01-01T00:00:00Z"),
					CheckoutURL:    "https://checkout.durianpay.id/callback",
					PaidAmount:     "10001.00",
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
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, gomock.Any(), nil, gomock.Any()).
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
				validPayload := ewalletPayload{}
				argsPayload := ewalletPayload{
					Type:          "EWALLET",
					Request:       parseArgs.payload,
					SandboxOption: parseArgs.payload.SandboxOption,
				}

				if !featureWrap.DeepEqualPayload(path_payload_payment+"charge_ewallet.json", &validPayload, &argsPayload) {
					t.Errorf("Client.ChargeEwallet() validPayload = %v, argsPayload %v", validPayload, argsPayload)
				}
			}

			gotRes, gotErr := c.ChargeEwallet(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ChargeEwallet() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ChargeEwallet() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
