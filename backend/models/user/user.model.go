package user

import (
	"encoding/json"
	"fmt"
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

const (
	UserNotFound int = 1
	UserNotMatch int = 2
)

func UserList() *[]DlogUser {
	users := make([]DlogUser, 0)
	conn := dao.GetConnection()
	conn.Find(&users)
	conn.Close()
	return &users
}

func SignedUser(c *gin.Context) (string, int) {
	var user DlogUser
	var bodyData struct {
		Email, Pwd string
	}
	_ = c.ShouldBindJSON(&bodyData)
	conn := dao.GetConnection()
	defer conn.Close()
	conn.Select("user_email, user_password, user_call").Where(DlogUser{
		UserEmail: bodyData.Email,
	}).Find(&user)

	if user != (DlogUser{}) {
		if user.UserPassword == bodyData.Pwd {
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
			return accessToken, 0
		} else {
			return "", UserNotMatch
		}
	}

	return "", UserNotFound
}

// 1. accesstoken 검증
// 2. refreshtoken 검증
// 3. 검증: accesstoken 재 발급, 유효하지않은 검증: 빈 문자열
func AuthenticationUser(c *gin.Context) string {
	var param struct {
		Token string
		Email string
	}
	_ = c.ShouldBindJSON(&param)

	var user DlogUser
	// 만료될때만 재 발급 로직 수행
	if _, err := jwt.VaildAccessToken(param.Token); err.Code == jwt.EXPIRED {
		conn := dao.GetConnection()
		conn.Where(DlogUser{
			UserEmail: param.Email,
		}).Find(&user)

		if _, err := jwt.VaildRefreshToken(user.RefreshToken); err != nil {
			var data = struct {
				UserEmail string
				UserCall  string
			}{
				UserEmail: user.UserEmail,
				UserCall:  user.UserCall,
			}
			return jwt.CreateAccessToken(&data)
		}
	}

	return ""
}

func Create(c *gin.Context) {
	var param struct {
		Email string
		Pwd   string
	}{
		Email: 
	}
	_ = c.ShouldBindJSON(&param)
	conn := dao.GetConnection()
	if err := conn.Create(DlogUser{UserEmail: param.Email, UserPassword: param.Pwd}).Error; err != nil {
		panic(err)
	}
}
func Delete(email string) {
	conn := dao.GetConnection()
	if err := conn.Where(DlogUser{UserEmail: email}).Delete(DlogUser{}).Error; err != nil {
		panic(err)
	}
}
