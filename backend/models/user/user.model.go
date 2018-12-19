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
	UserPassword string `json:"user_password,omitempty"`
	UserCall     string `json:"user_call,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	common.Base
}

type AuthData struct {
	AccessToken string
	Id          string
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

func SignedUser(c *gin.Context) (*AuthData, int) {
	if body, exists := c.Get("body"); exists {
		body, _ := body.(map[string]interface{})
		email := body["email"].(string)
		pwd := body["pwd"].(string)

		var user DlogUser
		conn := dao.GetConnection()
		conn.Select("user_email, user_password, user_call").Where(DlogUser{
			UserEmail: email,
		}).Find(&user)

		if user != (DlogUser{}) {
			if user.UserPassword == pwd {
				accessToken, xidval := jwt.CreateAccessToken(DlogUser{
					UserEmail: user.UserEmail,
					UserCall:  user.UserCall,
				})
				refreshToken := jwt.CreateRefreshToken(xidval)
				conn.Model(&user).Where(DlogUser{
					UserEmail: user.UserEmail,
				}).Update(DlogUser{
					RefreshToken: refreshToken,
					Base: common.Base{
						UpdateDate: time.Now(),
					},
				})

				authData := AuthData{
					AccessToken: accessToken,
					Id:          xidval,
				}

				return &authData, 0
			} else {
				return nil, UserNotMatch
			}
		}
	}

	return nil, UserNotFound
}

// 1. accesstoken 검증
// 2. refreshtoken 검증
// 3. 검증: accesstoken 재 발급, 유효하지않은 검증: 빈 문자열
func AuthenticationUser(c *gin.Context) (string, uint32) {
	var resultTkn string
	var status uint32

	if body, exists := c.Get("body"); exists {
		body := body.(map[string]interface{})
		token := body["token"].(string)
		email := " "
		if body["email"] != nil {
			email = body["email"].(string)
		}

		var user DlogUser
		// 만료될때만 재 발급 로직 수행
		if decode, err := jwt.VaildAccessToken(token); err.Code == jwt.EXPIRED {
			conn := dao.GetConnection()
			conn.Where(DlogUser{
				UserEmail: email,
			}).Find(&user)
			if user != (DlogUser{}) {
				if _, err := jwt.VaildRefreshToken(user.RefreshToken); err == nil {
					var data = struct {
						UserEmail string
						UserCall  string
					}{
						UserEmail: user.UserEmail,
						UserCall:  user.UserCall,
					}
					resultTkn, _ = jwt.CreateAccessToken(&data)
					status = jwt.EXPIRED
				}
			} else {
				status = jwt.INVAILD
			}
		} else if err.Code == jwt.INVAILD {
			status = jwt.INVAILD
		}
	}

	return resultTkn, status
}

func Create(c *gin.Context) {
	if body, exists := c.Get("body"); exists {
		body := body.(map[string]interface{})
		var param = struct {
			Email string
			Pwd   string
		}{
			Email: body["email"].(string),
			Pwd:   body["pwd"].(string),
		}
		conn := dao.GetConnection()
		if err := conn.Create(DlogUser{
			UserEmail:    param.Email,
			UserPassword: param.Pwd,
			Base: common.Base{
				CreateDate: time.Now(),
				UpdateDate: time.Now(),
			},
		}).Error; err != nil {
			panic(err)
		}
	}
}
func Delete(email string) {
	conn := dao.GetConnection()
	if err := conn.Where(DlogUser{UserEmail: email}).Delete(DlogUser{}).Error; err != nil {
		panic(err)
	}
}
