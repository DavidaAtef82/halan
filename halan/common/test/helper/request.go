package helper

import "net/http/httptest"

func MakeRequest(testCase Case) *httptest.ResponseRecorder {
	switch testCase.TypeRequest {
	case "post":
		return post(testCase.Body, testCase.Url, nil)
	}
	return nil
}
