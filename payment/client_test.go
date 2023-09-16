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
	"strings"
	"testing"
	"time"

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
					Metadata:     map[string]string{},
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

func TestClient_ChargeRetailStore(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type retailStorePayload struct {
		Type    string                                    `json:"type"`
		Request durianpay.PaymentChargeRetailStorePayload `json:"request"`
	}

	type args struct {
		ctx     context.Context
		payload durianpay.PaymentChargeRetailStorePayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *ChargeRetailStore
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PaymentChargeRetailStorePayload{
					OrderID:      "ord_mJH2hKOSYb3514",
					BankCode:     "ALFAMART",
					Name:         "Name Appear in ATM",
					Amount:       "20000.00",
					PaymentRefID: "pay_ref_123",
				},
			},
			prepare: func(m mocks, args args) {
				payload := chargePayload{
					Type:          "RETAILSTORE",
					Request:       args.payload,
					SandboxOption: args.payload.SandboxOption,
				}

				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"charge_retailstore_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &ChargeRetailStore{
				Type: "RETAILSTORE",
				Response: chargeResponseRetailStore{
					PaymentID:      "pay_Ln1PZECuqf3748",
					OrderID:        "ord_VN5nVJpSW27112",
					AccountNumber:  "1111111111",
					PaymentRefID:   "pay_ref_123",
					ExpirationTime: tests.StringToTime("2023-09-05T10:31:39.672938538Z"),
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
				validPayload := retailStorePayload{}
				argsPayload := retailStorePayload{
					Type:    "RETAILSTORE",
					Request: parseArgs.payload,
				}

				if !featureWrap.DeepEqualPayload(path_payload_payment+"charge_retailstore.json", &validPayload, &argsPayload) {
					t.Errorf("Client.ChargeRetailStore() validPayload = %v, argsPayload %v", validPayload, argsPayload)
				}
			}

			gotRes, gotErr := c.ChargeRetailStore(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ChargeRetailStore() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ChargeRetailStore() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_ChargeOnlineBank(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type onlineBankingPayload struct {
		Type    string                                      `json:"type"`
		Request durianpay.PaymentChargeOnlineBankingPayload `json:"request"`
	}

	type args struct {
		ctx     context.Context
		payload durianpay.PaymentChargeOnlineBankingPayload
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *ChargeOnlineBank
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PaymentChargeOnlineBankingPayload{
					OrderID: "ord_mJH2hKOSYb3514",
					Type:    "JENIUSPAY",
					Name:    "Name Appear in ATM",
					Amount:  "20000.00",
					CustomerInfo: durianpay.PaymentCustomerInfo{
						Email:     "jude_kasper@koss.in",
						GivenName: "Jude Kasper",
						ID:        "cus_aGn5UD0m7F0994",
					},
					Mobile: "+6285722173217",
				},
			},
			prepare: func(m mocks, args args) {
				payload := chargePayload{
					Type:    "ONLINE_BANKING",
					Request: args.payload,
				}

				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"charge_onlinebanking_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &ChargeOnlineBank{
				Type: "ONLINE_BANKING",
				Response: chargeResponseOnlineBank{
					PaymentID:      "pay_RGEkDpZZWR9662",
					OrderID:        "ord_VN5nVJpSW27112",
					Mobile:         "+6285722173217",
					Status:         "processing",
					ExpirationTime: tests.StringToTime("2023-09-05T10:32:27.273180959Z"),
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
				validPayload := onlineBankingPayload{}
				argsPayload := onlineBankingPayload{
					Type:    "ONLINE_BANKING",
					Request: parseArgs.payload,
				}

				if !featureWrap.DeepEqualPayload(path_payload_payment+"charge_onlinebanking.json", &validPayload, &argsPayload) {
					t.Errorf("Client.ChargeOnlineBank() validPayload = %v, argsPayload %v", validPayload, argsPayload)
				}
			}

			gotRes, gotErr := c.ChargeOnlineBank(parseArgs.ctx, parseArgs.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ChargeOnlineBank() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ChargeOnlineBank() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_ChargeQRIS(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type qrisPayload struct {
		Type    string                             `json:"type"`
		Request durianpay.PaymentChargeQRISPayload `json:"request"`
	}

	type args struct {
		ctx     context.Context
		payload durianpay.PaymentChargeQRISPayload
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *ChargeQRIS
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PaymentChargeQRISPayload{
					OrderID: "ord_ZSHipeBgUd4740",
					Type:    "DANA",
					Amount:  "80001.00",
					Name:    "Name Appear in ATM",
				},
			},
			prepare: func(m mocks, args args) {
				payload := chargePayload{
					Type:    "QRIS",
					Request: args.payload,
				}

				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"charge_qris_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &ChargeQRIS{
				Type: "QRIS",
				Response: chargeResponseQRIS{
					PaymentID:      "pay_s2sSBlDSWv4167",
					OrderID:        "ord_QETgbs2UGL3100",
					Status:         "processing",
					ExpirationTime: tests.StringToTime("2021-09-15T15:44:37Z"),
					CreationTime:   tests.StringToTime("2021-09-12T15:44:37Z"),
					QRString:       "data:image/png;base64, long_qr_string",
					UniqueID:       "QRIS",
					Metadata: map[string]string{
						"merchant_name": "Durianpay",
						"merchant_id":   "sample_national_merchant_id",
					},
					Amount: "80001.00",
					QRCode: "00020101021226590013ID.CO.BNI.WWW011893600009150002286002092107061320303UME51470015ID.OR.GPNQR.WWW0217ID2107271315771960303UME520454995303360540880001.005802ID5905Ajesh6013JAKARTA PUSAT6105101406214011038291492856304E1F",
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
				validPayload := qrisPayload{}
				argsPayload := qrisPayload{
					Type:    "QRIS",
					Request: parseArgs.payload,
				}

				if !featureWrap.DeepEqualPayload(path_payload_payment+"charge_qris.json", &validPayload, &argsPayload) {
					t.Errorf("Client.ChargeQRIS() validPayload = %v, argsPayload %v", validPayload, argsPayload)
				}
			}

			gotRes, gotErr := c.ChargeQRIS(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ChargeQRIS() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ChargeQRIS() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_ChargeCard(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type cardPayload struct {
		Type    string                             `json:"type"`
		Request durianpay.PaymentChargeCardPayload `json:"request"`
	}

	type args struct {
		ctx     context.Context
		payload durianpay.PaymentChargeCardPayload
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *ChargeCard
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PaymentChargeCardPayload{
					OrderID:      "ord_1EcWGI2xSs7216",
					Amount:       "10000.00",
					PaymentRefID: "pay_ref_123",
					CustomerInfo: durianpay.PaymentCustomerInfo{
						ID:        "cus_aGn5UD0m7F0994",
						Email:     "jude_kasper@koss.in",
						GivenName: "Jude Kasper",
					},
				},
			},
			prepare: func(m mocks, args args) {
				payload := chargePayload{
					Type:    "CARD",
					Request: args.payload,
				}

				m.api.EXPECT().
					Req(gomock.Any(), "POST", PATH_PAYMENT_CHARGE, nil, payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"charge_card_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &ChargeCard{
				Type: "CARD",
				Response: chargeResponseCard{
					PaymentID:    "pay_TMlTVT3wvr3598",
					OrderID:      "ord_Gf7LimyjMk7270",
					PaymentRefID: "pay_ref_123",
					Status:       "completed",
					PaidAmount:   "10001.00",
					CheckoutURL:  "https://link.to/card-checkout-url",
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
				validPayload := cardPayload{}
				argsPayload := cardPayload{
					Type:    "CARD",
					Request: parseArgs.payload,
				}

				if !featureWrap.DeepEqualPayload(path_payload_payment+"charge_card.json", &validPayload, &argsPayload) {
					t.Errorf("Client.ChargeCard() validPayload = %v, argsPayload %v", validPayload, argsPayload)
				}
			}

			gotRes, gotErr := c.ChargeCard(parseArgs.ctx, parseArgs.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ChargeCard() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ChargeCard() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchPayments(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		opt durianpay.PaymentFetchOption
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *FetchPayments
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				opt: durianpay.PaymentFetchOption{
					From: time.Now().Format(time.DateOnly),
					To:   time.Now().Format(time.DateOnly),
				},
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", PATH_PAYMENT_CHARGE, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"fetch_payments_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &FetchPayments{
				Payments: []Payments{
					{
						ID:                 "pay_80pgxEcUbO8054",
						OrderID:            "ord_NDmLvwTTh95152",
						PaymentRefID:       "pay_ref_123",
						Amount:             "10001.00",
						Status:             "processing",
						IsLive:             false,
						ExpirationDate:     tests.StringToTime("0001-01-01T00:00:00Z"),
						PaymentDetailsType: "bnpl_details",
						MethodID:           "AKULAKU",
						CreatedAt:          tests.StringToTime("2023-09-04T10:38:12.122665Z"),
						UpdatedAt:          tests.StringToTime("2023-09-04T10:38:12.122665Z"),
						RetryCount:         0,
						CustomerID:         "cus_L6WNLJrTKT6000",
						GivenName:          "Jane Doe",
						Email:              "jane_doe@nomail.com",
						OrderRefID:         "order_ref_002",
						Currency:           "IDR",
						FailureReason:      make(map[string]string),
						Metadata:           make(map[string]string),
					},
				},
				Total: 1,
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				m.api.EXPECT().
					Req(gomock.Any(), "GET", PATH_PAYMENT_CHARGE, args.opt, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.FetchPayments(tt.args.ctx, tt.args.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchPayments() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchPayments() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_FetchPaymentByID(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		ID  string
		opt durianpay.PaymentFetchByIDOption
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *Payment
		wantErr *durianpay.Error
	}{
		{
			name: "Success without params",
			args: args{
				ctx: context.Background(),
				ID:  "pay_Ln1PZECuqf3748",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_FETCH_BY_ID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"fetch_payment_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Payment{
				ID:                 "pay_Ln1PZECuqf3748",
				OrderID:            "ord_VN5nVJpSW27112",
				MerchantID:         "mer_pHXgBZ2Qx95625",
				PaymentRefID:       "pay_ref_123",
				Amount:             "10001.00",
				Status:             "processing",
				IsLive:             false,
				ExpirationDate:     tests.StringToTime("2023-09-05T10:31:39.672939Z"),
				CreatedAt:          tests.StringToTime("2023-09-04T10:31:39.673491Z"),
				UpdatedAt:          tests.StringToTime("2023-09-04T10:31:39.673491Z"),
				PaymentDetailsType: "RETAILSTORE",
				MethodID:           "ALFAMART",
				FailureReason:      make(map[string]string),
			},
		},
		{
			name: "Success with expand customer",
			args: args{
				ctx: context.Background(),
				ID:  "pay_Ln1PZECuqf3748",
				opt: durianpay.PaymentFetchByIDOption{
					Expand: "customer",
				},
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_FETCH_BY_ID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"fetch_payment_customer_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Payment{
				ID:                 "pay_Ln1PZECuqf3748",
				OrderID:            "ord_VN5nVJpSW27112",
				MerchantID:         "mer_pHXgBZ2Qx95625",
				PaymentRefID:       "pay_ref_123",
				Amount:             "10001.00",
				Status:             "cancelled",
				IsLive:             false,
				ExpirationDate:     tests.StringToTime("2023-09-05T10:31:39.672939Z"),
				CreatedAt:          tests.StringToTime("2023-09-04T10:31:39.673491Z"),
				UpdatedAt:          tests.StringToTime("2023-09-05T17:10:28.802621Z"),
				PaymentDetailsType: "RETAILSTORE",
				MethodID:           "ALFAMART",
				FailureReason:      make(map[string]string),
				Customer: PaymentCustomer{
					ID:            "cus_L6WNLJrTKT6000",
					CustomerRefID: "cust_001",
					Email:         "jane_doe@nomail.com",
					GivenName:     "Jane Doe",
					CreatedAt:     tests.StringToTime("0001-01-01T00:00:00Z"),
				},
			},
		},
		{
			name: "Success with expand order",
			args: args{
				ctx: context.Background(),
				ID:  "pay_Ln1PZECuqf3748",
				opt: durianpay.PaymentFetchByIDOption{
					Expand: "order",
				},
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_FETCH_BY_ID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"fetch_payment_order_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Payment{
				ID:                 "pay_Ln1PZECuqf3748",
				OrderID:            "ord_VN5nVJpSW27112",
				MerchantID:         "mer_pHXgBZ2Qx95625",
				PaymentRefID:       "pay_ref_123",
				Amount:             "10001.00",
				Status:             "cancelled",
				IsLive:             false,
				ExpirationDate:     tests.StringToTime("2023-09-05T10:31:39.672939Z"),
				CreatedAt:          tests.StringToTime("2023-09-04T10:31:39.673491Z"),
				UpdatedAt:          tests.StringToTime("2023-09-05T17:10:28.802621Z"),
				PaymentDetailsType: "RETAILSTORE",
				MethodID:           "ALFAMART",
				FailureReason:      make(map[string]string),
				Order: PaymentOrder{
					ID:         "ord_VN5nVJpSW27112",
					MerchantID: "mer_pHXgBZ2Qx95625",
					CustomerID: "cus_L6WNLJrTKT6000",
					OrderRefID: "order_ref_002",
					Amount:     "10001.00",
					Currency:   "IDR",
					Status:     "completed",
					IsLive:     false,
					CreatedAt:  tests.StringToTime("2023-09-04T10:26:33.723962Z"),
				},
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_FETCH_BY_ID, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, args.opt, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.FetchPaymentByID(parseArgs.ctx, parseArgs.ID, parseArgs.opt)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.FetchPaymentByID() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.FetchPaymentByID() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_CheckPaymentStatus(t *testing.T) {
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
		wantRes *CheckPaymentStatus
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				ID:  "pay_wA2X2Mvm2d4965",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_CHECK_STATUS, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"check_payment_status_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &CheckPaymentStatus{
				Status:      "processing",
				IsCompleted: false,
				Signature:   "9206137de45a992ca3416e4571d14dce6493b104fac00849da41dfb04a913ef9",
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_CHECK_STATUS, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, nil, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.CheckPaymentStatus(parseArgs.ctx, parseArgs.ID)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.CheckPaymentStatus() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.CheckPaymentStatus() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_Verify(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
		ID  string
		opt durianpay.PaymentVerifyOption
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes bool
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				opt: durianpay.PaymentVerifyOption{
					VerificationSignature: "adf9a1a37af514c91225f6680e2df723fefebb7638519bcc7e7c9de02f2a3ab2",
				},
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_VERIFY, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, args.opt, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"verify_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: true,
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_VERIFY, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "GET", url, args.opt, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.Verify(parseArgs.ctx, parseArgs.ID, parseArgs.opt)
			if gotRes != tt.wantRes {
				t.Errorf("Client.Verify() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Verify() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_Capture(t *testing.T) {
	featureWrap := tests.FeatureWrap(t)
	defer featureWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		ID      string
		payload durianpay.PaymentCapturePayload
	}

	tests := []struct {
		name    string
		args    args
		prepare func(m mocks, args args)
		wantRes *Capture
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: durianpay.PaymentCapturePayload{
					Amount: "1000.00",
				},
				ID: "pay_wA2X2Mvm2d4965",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_CAPTURE, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "POST", url, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"capture_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Capture{
				PaymentID:           "pay_123",
				OrderID:             "ord_123",
				PreauthorizedAmount: "1000.00",
				PaidAmount:          "1000.00",
				Status:              "processing",
				CreatedAt:           tests.StringToTime("2022-12-15T10:51:47.829636Z"),
				UpdatedAt:           tests.StringToTime("2022-12-15T16:26:25.076181Z"),
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_CAPTURE, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "POST", url, nil, args.payload, nil, gomock.Any()).
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

			gotRes, gotErr := c.Capture(tt.args.ctx, tt.args.ID, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Capture() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Capture() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_Cancel(t *testing.T) {
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
		wantRes *Cancel
		wantErr *durianpay.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				ID:  "pay_wA2X2Mvm2d4965",
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_CANCEL, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "PUT", url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *durianpay.Error {
						err := json.Unmarshal(featureWrap.ResJSONByte(path_response_payment+"cancel_200.json"), response)
						if err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Cancel{
				ID:             "pay_Ln1PZECuqf3748",
				OrderID:        "ord_VN5nVJpSW27112",
				Status:         "cancelled",
				IsLive:         false,
				ExpirationDate: tests.StringToTime("0001-01-01T00:00:00Z"),
				CreatedAt:      tests.StringToTime("2023-09-04T10:31:39.673491Z"),
				UpdatedAt:      tests.StringToTime("2023-09-05T17:10:28.802621Z"),
				RetryCount:     0,
				FailureReason:  make(map[string]string),
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m mocks, args args) {
				url := strings.ReplaceAll(PATH_PAYMENT_CANCEL, ":id", args.ID)

				m.api.EXPECT().
					Req(gomock.Any(), "PUT", url, nil, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.Cancel(parseArgs.ctx, parseArgs.ID)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Cancel() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Cancel() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
