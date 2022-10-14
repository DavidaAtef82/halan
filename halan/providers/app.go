package providers

import (
	"github.com/gin-gonic/gin"
	"halan/common/helper"
	"os"
)

func Run() {
	/**
	* init gin frame work
	* run default middleware like CROS
	 */
	r := gin.Default()
	/**
	* setup routes with three default groups
	* admin / auth / visitors
	 */
	Routing(r)
	/**
	* start app on port to change port just edit APP_PORT in
	* env file like :9090
	 */
	if os.Getenv("APP_ENV") == "local" {
		err := r.Run(":" + os.Getenv("APP_PORT_LOCAL"))
		helper.CheckError(err)
	} else {
		err := r.Run(":" + os.Getenv("APP_PORT_PRODUCTION"))
		helper.CheckError(err)
	}
}
