package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitException() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("test")
		c.Next()
	}
}
