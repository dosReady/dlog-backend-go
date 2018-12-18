package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	userModel "github.com/dosReady/dlog/backend/models/user"
	"github.com/gin-gonic/gin"
)

func BodyParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]interface{}
		decoder := json.NewDecoder(c.Request.Body)
		defer c.Request.Body.Close()
		_ = decoder.Decode(&m)
		fmt.Println(m)
		c.Next()
	}
}

// https://www.curioustore.com/#!/
// 변수명 작명 사이트
func CertifiedMdlw() gin.HandlerFunc {
	return func(c *gin.Context) {
		if result := userModel.AuthenticationUser(c); result != "" {
			c.Set("auth", result)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}
	}
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

func VerificationURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _urlvalidator(c) {
			c.HTML(http.StatusOK, "app.html", "")
		} else {
			c.Next()
		}
	}
}
