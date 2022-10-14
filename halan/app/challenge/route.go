package challenge

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) *gin.RouterGroup {

	routesGroup := r.Group("/challenge")
	{
		routesGroup.POST("/problem-1", Problem1)
		routesGroup.POST("/problem-2", Problem2)
		routesGroup.POST("/problem-3/rle", Problem3RLE)
		routesGroup.POST("/problem-3/rld", Problem3RLD)
		routesGroup.POST("/problem-4", Problem4)
		routesGroup.POST("/problem-5", Problem5)
		routesGroup.POST("/problem-6", Problem6)
		routesGroup.POST("/problem-7", Problem7)
		routesGroup.POST("/problem-8", Problem8)
		routesGroup.POST("/problem-9", Problem9)
	}
	return r
}
