package main

import (
	"github.com/subosito/gotenv"
	_ "halan/docs"
	"halan/providers"
)

// @title Halan problem-solving challenge
// @version 1.0
// @description The APIs of Halan problem-solving challenge
// @termsOfService http://swagger.io/terms/
// @BasePath /halan/test/
func main() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
	/**
	* Run gin framework
	* run routing
	* serve app
	 */
	providers.Run()
}
