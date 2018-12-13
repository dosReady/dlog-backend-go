package router

import (
	"net/http"

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

	api := r.Group("/api")
	{
		api.POST("/dlog", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"name": "dos"})
		})
	}
}
