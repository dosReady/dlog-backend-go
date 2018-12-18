package dlog

import (
	"net/http"

	userModel "github.com/dosReady/dlog/backend/models/user"
	"github.com/gin-gonic/gin"
)

func UserSelect(c *gin.Context) {
	/*
		users := make([]DlogUser, 0)
		dao.List(&users, `SELECT * FROM dlog_user`)
		c.JSON(http.StatusOK, gin.H{"name": users})
	*/
	result := userModel.UserList()
	c.JSON(http.StatusOK, gin.H{"name": result})
}

func UserLogin(c *gin.Context) {
	if accesstoken, status := userModel.SignedUser(c); status > 0 {
		c.JSON(http.StatusOK, gin.H{"accessToken": "", "status": status})
	} else {
		c.JSON(http.StatusOK, gin.H{"accessToken": accesstoken, "status": 0})
	}
}
