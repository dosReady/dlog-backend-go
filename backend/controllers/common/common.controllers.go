package dlog

import (
	"net/http"

	userModel "github.com/dosReady/dlog/backend/models/user"
	"github.com/dosReady/dlog/backend/modules/utils"
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
	err := userModel.SignedUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

func UserLogout(c *gin.Context) {
	utils.DeleteCookie("user", c)
	utils.DeleteCookie("token", c)
	c.JSON(http.StatusOK, gin.H{})
}
