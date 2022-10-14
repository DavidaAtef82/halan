package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	_ "testing"
)

type Case struct {
	ID                      int                    `json:"id"`
	Body                    interface{}            `json:"body"`
	TypeRequest             string                 `json:"type_request"`
	Url                     string                 `json:"url"`
	ExpectedOutputInPayload map[string]interface{} `json:"expected_output_in_payload"`
	ExpectedCode            int                    `json:"expected_code"`
}

func ReadJson(filename string) []Case {
	var cases map[string][]Case
	filepath := path.Join(path.Dir(filename), filename)
	jsonFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &cases)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return cases["test_cases"]
}
