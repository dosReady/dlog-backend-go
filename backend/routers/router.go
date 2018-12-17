package router

import (
	"net/http"

	dlogCtrl "github.com/dosReady/dlog/backend/controllers/dlog"
	userModel "github.com/dosReady/dlog/backend/models/user"
	"github.com/gin-gonic/gin"
)

type BodyJSON struct {
	Val string
}

// STATIC 파일, API 호출 URL 아닌 경우에만 html 파일 호출
// html 파일 호출이유는 Vue Js, React Js 같은것을 사용할때 새로고침 할 경우
// 로드 파일을 불러올수 있도록 하기위함
func _urlvalidator(c *gin.Context) bool {
	result := false
	if c.Request.Method == http.MethodGet {
		result = true
	}
	return result
}

func SettingRouters(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		if _urlvalidator(c) {
			c.HTML(http.StatusOK, "app.html", "")
		} else {
			c.Next()
		}
	})
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		api.POST("/dlog", dlogCtrl.UserSelect)
		api.POST("/dlog/login", dlogCtrl.UserLogin)
	}
	apitest := r.Group("/test")
	{
		apitest.Use(func(c *gin.Context) {
			var param struct {
				Token string
			}
			_ = c.ShouldBindJSON(&param)
			userModel.AuthenticationUser(param.Token)
			c.Next()
		})
		apitest.POST("/echo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"name": "123"})
		})
	}
}
