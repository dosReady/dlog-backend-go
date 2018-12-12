package routers

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "dos"})
	})

	dlogApi := r.Group("/dlog")
	{
		dlogApi.GET("/me", controllers.DlogMe)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"name": "NotFound"})
	})
}
