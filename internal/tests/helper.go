/*
 * File Created: Sunday, 30th July 2023 3:57:52 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package tests

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
)

type Feature struct {
	Ctrl      *gomock.Controller
	ServerKey string
}

func HttpMockResJSON(statusCode int, filePath string) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") == "" {
			return nil, errors.New("access_key not present!")
		}

		if r.Header.Get("Content-Type") != "application/json" {
			return nil, errors.New("Invalid content-type != application/json")
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
