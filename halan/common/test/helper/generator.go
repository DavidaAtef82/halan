package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
)

func TestGenerator(t *testing.T, cases []Case) {
	var seqMutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(cases))
	for _, testCase := range cases {
		seqMutex.Lock()
		k := MakeRequest(testCase)
		fmt.Println(k.Body)
		if !assert.Equal(t, testCase.ExpectedCode, k.Code) {
			fmt.Println("test case number failed ", testCase.ID)
			fmt.Println("url ", testCase.Url)
			return
		}
		if CheckExpectedInPayload(testCase, k, t) {
			fmt.Println("test case number failed ", testCase.ID)
			fmt.Println("url ", testCase.Url)
			return
		}
		fmt.Println("test case number pass", testCase.ID)
		seqMutex.Unlock()
		wg.Done()
	}
	wg.Wait()
	return
}

func CheckExpectedInPayload(testCase Case, k *httptest.ResponseRecorder, t *testing.T) bool {
	if len(testCase.ExpectedOutputInPayload) != 0 {
		var result map[string]interface{}
		ReturnResponseObject(k, &result)
		for key, expected := range testCase.ExpectedOutputInPayload {
			if CheckType(expected) {
				if !assert.Equal(t, expected, result[key]) {
					fmt.Println("failed in assert", testCase.ID)
					return true
				}
			}
		}
	}
	return false
}

func CheckType(value interface{}) bool {
	typeOFValue := reflect.TypeOf(value).Kind()
	switch typeOFValue {
	case reflect.Float64:
		return !(value.(float64) == 0)
	case reflect.String:
		return !(value.(string) == "")
	case reflect.Slice:
		return !(reflect.ValueOf(value).Len() == 0)
	}
	return true
}
