package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.LoadHTMLGlob("views/*")
	r.GET("/*", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{})
	})
	/*
		dlogApi := r.Group("/dlog")
		{
			dlogApi.GET("/me", controllers.DlogMe)
		}
	*/
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"name": "NotFound"})
	})
}
