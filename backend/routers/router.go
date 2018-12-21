package router

import (
	"net/http"

	commonCtrl "github.com/dosReady/dlog/backend/controllers/common"
	userCtrl "github.com/dosReady/dlog/backend/controllers/user"
	middleware "github.com/dosReady/dlog/backend/modules/middleware"
	"github.com/gin-gonic/gin"
)

func SettingRouters(r *gin.Engine) {
	r.Use(middleware.VerificationURL())
	r.Use(middleware.BodyParser())
	r.Use(gin.Recovery())

	apir1 := r.Group("/api/dlog")
	{
		apir1.POST("/login", commonCtrl.UserLogin)
		apir1.POST("/logout", commonCtrl.UserLogin)
	}

	apir2 := r.Group("/api/user", middleware.CertifiedMdlw())
	{
		apir2.POST("/create", userCtrl.UserCreate)
		apir2.POST("/delete/:email", userCtrl.UserDelete)
	}

	apitest := r.Group("/api/test")
	apitest.Use(middleware.CertifiedMdlw())
	{
		apitest.POST("/echo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"name": "Hi"})
		})
	}
}
