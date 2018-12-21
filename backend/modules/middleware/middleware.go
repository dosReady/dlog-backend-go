package middleware

import (
	"encoding/json"
	"net/http"

	userModel "github.com/dosReady/dlog/backend/models/user"
	jwt "github.com/dosReady/dlog/backend/modules/jwt"
	"github.com/dosReady/dlog/backend/modules/utils"
	"github.com/gin-gonic/gin"
)

func BodyParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]interface{}
		decoder := json.NewDecoder(c.Request.Body)
		defer c.Request.Body.Close()
		_ = decoder.Decode(&m)
		c.Set("body", m)
		c.Next()
	}
}

// https://www.curioustore.com/#!/
// 변수명 작명 사이트
func CertifiedMdlw() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, status := userModel.AuthenticationUser(c)
		if status == jwt.INVAILD {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "토큰이 유효하지않습니다."})
		} else if status == jwt.EXPIRED {
			utils.SetCookie("token", result, true, c)
			c.Next()
		} else {
			c.Next()
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
