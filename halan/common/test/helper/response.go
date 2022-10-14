package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
)

type GlobalResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
	Code    int               `json:"code"`
}

type GlobalResponseObject struct {
	Payload map[string]interface{} `json:"payload"`
	GlobalResponse
}

func ReturnResponseObject(k *httptest.ResponseRecorder, result interface{}) {
	var globalObject GlobalResponseObject
	body, _ := ioutil.ReadAll(k.Result().Body)
	err := json.Unmarshal(body, &globalObject)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonString, err := json.Marshal(globalObject.Payload)
	err = json.Unmarshal(jsonString, result)
	if err != nil {
		fmt.Println(err.Error())
	}
}
