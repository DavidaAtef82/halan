package providers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"halan/routes"
	"os"
)

func Routing(r *gin.Engine) *gin.Engine {
	admin := r.Group(os.Getenv("MS_HALAN_URI"))
	routes.Admin(admin)
	swaggerRouting(r)
	return r
}

func swaggerRouting(r *gin.Engine) {
	if os.Getenv("ENABLE_SWAGGER") == "1" || os.Getenv("ENABLE_SWAGGER") == "true" {
		swagger := r.Group(os.Getenv("MS_HALAN_URI") + "/test")
		{
			swagger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
			routes.Admin(swagger)
		}
	}
}
