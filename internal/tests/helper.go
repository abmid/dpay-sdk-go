/*
 * File Created: Sunday, 30th July 2023 3:57:52 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package tests

import (
	"errors"
	"fmt"
	"net/http"
	"os"
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

// BoolPtr return value pointer for boolean
func BoolPtr[V bool](value V) *V {
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
