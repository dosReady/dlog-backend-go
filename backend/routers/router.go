package router

import (
	"net/http"

	dlogCtrl "github.com/dosReady/dlog/backend/controllers/dlog"
	middleware "github.com/dosReady/dlog/backend/modules/middleware"
	"github.com/gin-gonic/gin"
)

func SettingRouters(r *gin.Engine) {
	r.Use(middleware.VerificationURL())
	r.Use(gin.Recovery())

	api := r.Group("/api/dlog")
	{
		api.POST("/login", dlogCtrl.UserLogin)
	}
	apitest := r.Group("/test")
	apitest.Use(middleware.CertifiedMdlw())
	{
		apitest.POST("/echo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"name": "123"})
		})
	}
}
