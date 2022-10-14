package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

type RequestData struct {
	RequestType string
	Url         string
	Data        interface{}
	Header      map[string]string
}

var R = gin.Default()

func post(data interface{}, url string, headers map[string]string) *httptest.ResponseRecorder {
	return request(RequestData{
		Url:         url,
		RequestType: "POST",
		Data:        data,
		Header:      headers,
	})
}

/**
* Make new request
 */
func request(request RequestData) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	sendData, _ := json.Marshal(&request.Data)
	req, _ := http.NewRequest(request.RequestType, request.Url, bytes.NewReader(sendData))
	if len(request.Header) > 0 {
		for headerName, headerValue := range request.Header {
			req.Header.Set(headerName, headerValue)
		}
	}
	fmt.Println(req)
	R.ServeHTTP(w, req)
	fmt.Println(w)
	return w
}
