package utils

import "github.com/gin-gonic/gin"

func SetCookieWithHttpOnly(key string, value string, c *gin.Context) {
	c.SetCookie(key, value, 0, "", "", false, true)
}

func SetCookie(key string, value string, c *gin.Context) {
	c.SetCookie(key, value, 0, "", "", false, false)
}

func DeleteCookie(key string, c *gin.Context) {
	c.SetCookie(key, "", -1, "", "", false, false)
}

func GetCookie(key string, c *gin.Context) string {
	cookie, err := c.Cookie(key)
	if err != nil {
		panic(err)
	}
	return cookie
}
