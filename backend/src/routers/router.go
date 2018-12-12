package routers

import (
	"config"
	"controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		cfg := new(config.Config)
		cfg.GetConfig()
		c.JSON(200, gin.H{"name": "dos"})
	})

	dlogApi := r.Group("/dlog")
	{
		dlogApi.GET("/me", controllers.DlogMe)
	}
}
