package dlog

import (
	"net/http"

	userModel "github.com/dosReady/dlog/backend/models/user"
	utils "github.com/dosReady/dlog/backend/modules/utils"
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
	authData, status := userModel.SignedUser(c)
	utils.SetCookie("token", authData, c)
	c.JSON(http.StatusOK, gin.H{"token": authData, "status": status})
}
