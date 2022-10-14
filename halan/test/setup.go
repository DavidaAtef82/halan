package test

import (
	"github.com/gin-gonic/gin"
	helperTest "halan/common/test/helper"
	"halan/providers"
)

var TestRouter = false

func Init() {
	if !TestRouter {
		providers.Routing(helperTest.R)
	}
}

/**
* init gin and return gin engine
 */
func setupRouter() *gin.Engine {
	Init()
	return helperTest.R
}
