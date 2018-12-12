package middleware

import (
	"github.com/gin-gonic/gin"
)

func HanddleException() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errorToPrint := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if errorToPrint != nil {
			c.JSON(500, gin.H{
				"status":  500,
				"message": errorToPrint.Error(),
			})
		}
	}
}
