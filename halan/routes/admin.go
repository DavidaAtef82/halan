package routes

import (
	"github.com/gin-gonic/gin"
	"halan/app/challenge"
)

func Admin(r *gin.RouterGroup) *gin.RouterGroup {
	challenge.Routes(r)
	return r
}
