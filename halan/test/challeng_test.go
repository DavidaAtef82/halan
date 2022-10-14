package test

import (
	helperTest "halan/common/test/helper"
	"testing"
)

// all direction has all json and run all this cases
// make factory by struct to call function

func TestUser(t *testing.T) {
	setupRouter()
	cases := helperTest.ReadJson("challenge.json")
	helperTest.TestGenerator(t, cases)
	return
}
