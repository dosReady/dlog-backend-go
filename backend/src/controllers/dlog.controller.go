package controllers

import (
	"github.com/gin-gonic/gin"
)

func DlogMe(c *gin.Context) {
	c.JSON(200, gin.H{"name": "me"})
}
