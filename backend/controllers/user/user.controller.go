package user

import (
	"net/http"

	userModel "github.com/dosReady/dlog/backend/models/user"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	userModel.Create(c)
	c.JSON(http.StatusOK, gin.H{})
}

func UserDelete(c *gin.Context) {
	email := c.Param("email")
	userModel.Delete(email)
	c.JSON(http.StatusOK, gin.H{})
}
