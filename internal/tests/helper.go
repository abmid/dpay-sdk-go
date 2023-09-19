/*
 * File Created: Sunday, 30th July 2023 3:57:52 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
)

type Feature struct {
	Ctrl      *gomock.Controller
	ServerKey string
}

func HttpMockResJSON(statusCode int, filePath string, headers map[string]string) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") == "" {
			return nil, errors.New("access_key not present!")
		}

		if r.Header.Get("Content-Type") != "application/json" {
			return nil, errors.New("Invalid content-type != application/json")
		}

		if headers != nil {
			for key, value := range headers {
				if r.Header.Get(key) != value {
					return nil, errors.New(fmt.Sprintf("Headers: Invalid %s != %s", key, value))
				}
			}
		}

		resp, _ := httpmock.NewJsonResponderOrPanic(statusCode, httpmock.File(filePath))(r)

		return resp, nil
	}
}

func FeatureWrap(t *testing.T) *Feature {
	ctrl := gomock.NewController(t)

	return &Feature{
		Ctrl:      ctrl,
		ServerKey: "dpay_test_xxx",
	}
}

func (f *Feature) ResJSONByte(jsonFile string) []byte {
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}

	return file
}

// DeepEqualPayload checks whether the payload for the request matches the example json or not.
// Value payload and argPayload must be assign as pointer
func (f *Feature) DeepEqualPayload(fileJson string, payload any, argPayload any) bool {
	json.Unmarshal(f.ResJSONByte(fileJson), payload)

	if reflect.DeepEqual(payload, argPayload) {
		return true
	}

	var castPayload, castArgsPayload any

	bytesCastPayload, err := json.Marshal(payload)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	bytesCastArgsPayload, err := json.Marshal(argPayload)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	err = json.Unmarshal(bytesCastPayload, &castPayload)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	err = json.Unmarshal(bytesCastArgsPayload, &castArgsPayload)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	if reflect.DeepEqual(castPayload, castArgsPayload) {
		return true
	}

	return false
}

// DeepEqualResponse is used to check responses that have interface values
func (f *Feature) DeepEqualResponse(gotRes any, wantRes any) bool {
	if reflect.DeepEqual(wantRes, gotRes) {
		return true
	}

	var castWantRes, castGotRes any

	bytesCastWantRes, err := json.Marshal(wantRes)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	bytesCastGotRes, err := json.Marshal(gotRes)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	err = json.Unmarshal(bytesCastWantRes, &castWantRes)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	err = json.Unmarshal(bytesCastGotRes, &castGotRes)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	if reflect.DeepEqual(castWantRes, castGotRes) {
		return true
	}

	return false
}

// ToPtr return value pointer for anything data types.
func ToPtr[V any](value V) *V {
	return &value
}

// StringToTime return string to time without return error.
// If when parsing encounters an error, it will return the default value
func StringToTime(timeString string) time.Time {
	parse, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return time.Time{}
	}

	return parse
}
