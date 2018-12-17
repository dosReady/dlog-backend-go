package user

import (
	"time"

	common "github.com/dosReady/dlog/backend/models/common"
	dao "github.com/dosReady/dlog/backend/modules/dao"
	jwt "github.com/dosReady/dlog/backend/modules/jwt"
	"github.com/gin-gonic/gin"
)

// DB MODEL
type DlogUser struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserCall     string `json:"user_call"`
	RefreshToken string `json:"refresh_token"`
	common.Base
}

// API 파라미터 데이터
type BodyData struct {
	Email string
	Pwd   string
}

func UserList() *[]DlogUser {
	users := make([]DlogUser, 0)
	conn := dao.GetConnection()
	conn.Find(&users)
	conn.Close()
	return &users
}

func SignedUser(c *gin.Context) string {
	var user DlogUser
	var bodyData BodyData
	_ = c.ShouldBindJSON(&bodyData)

	conn := dao.GetConnection()
	conn.Select("user_email, user_call").Find(&user)
	accessToken := jwt.CreateAccessToken(&user)
	refreshToken := jwt.CreateRefreshToken(&user)
	conn.Model(&user).Where(DlogUser{
		UserEmail: user.UserEmail,
	}).Update(DlogUser{
		RefreshToken: refreshToken,
		Base: common.Base{
			UpdateDate: time.Now(),
		},
	})
	conn.Close()
	return accessToken
}

func AuthenticationUser(c *gin.Context) string {
	var param struct {
		Token string
		Email string
	}
	_ = c.ShouldBindJSON(&param)

	var user DlogUser
	decode := jwt.VaildAccessToken(param.Token)
	if decode == nil {
		conn := dao.GetConnection()
		conn.Select("refresh_token").Where(DlogUser{
			UserEmail: param.Email,
		}).Find(&user)

		jwt.CreateAccessToken()
	}
	return ""
}
