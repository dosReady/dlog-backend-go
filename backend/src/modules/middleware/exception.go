package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitException() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(*c)
		c.Next()
	}
}
