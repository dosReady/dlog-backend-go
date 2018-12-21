package user

import (
	"time"

	common "github.com/dosReady/dlog/backend/models/common"
	dao "github.com/dosReady/dlog/backend/modules/dao"
	jwt "github.com/dosReady/dlog/backend/modules/jwt"
	"github.com/dosReady/dlog/backend/modules/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DB MODEL
type DlogUser struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password,omitempty"`
	UserCall     string `json:"user_call,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
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
func _generateToken(user DlogUser, conn *gorm.DB) string {
	var obj = struct {
		UserEmail string
		UserCall  string
	}{
		UserEmail: user.UserEmail,
		UserCall:  user.UserCall,
	}
	accessToken, xidval := jwt.CreateAccessToken(obj)
	refreshToken := jwt.CreateRefreshToken(xidval)

	conn.Model(&user).Where(DlogUser{
		UserEmail: user.UserEmail,
	}).Update(DlogUser{
		RefreshToken: refreshToken,
		Base: common.Base{
			UpdateDate: time.Now(),
		},
	})
	return accessToken
}

func SignedUser(c *gin.Context) (string, int) {
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
				accessToken := _generateToken(user, conn)
				return accessToken, 0
			} else {
				return "", UserNotMatch
			}
		}
	}

	return "", UserNotFound
}

// 1. accesstoken 검증
// 2. refreshtoken 검증
// 3. 검증: accesstoken 재 발급, 유효하지않은 검증: 빈 문자열
func AuthenticationUser(c *gin.Context) (string, uint32) {
	var accessToken string
	var status uint32
	token, tokenErr := c.Cookie("token")
	if tokenErr != nil {
		status = jwt.INVAILD
	} else {
		var user DlogUser
		// 만료될때만 재 발급 로직 수행
		if decodeAccess, accessErr := jwt.VaildAccessToken(token); accessErr == jwt.EXPIRED {
			var accessPayLoad struct {
				UserEmail string
			}
			utils.DecodingJson(decodeAccess.Data, &accessPayLoad)
			conn := dao.GetConnection()
			conn.Where(DlogUser{
				UserEmail: accessPayLoad.UserEmail,
			}).Find(&user)
			if user != (DlogUser{}) {
				decodeRefresh, refreshErr := jwt.VaildRefreshToken(user.RefreshToken)
				if refreshErr == 0 && decodeAccess.Xid == decodeRefresh.Xid {
					accessToken = _generateToken(user, conn)
					status = jwt.EXPIRED
				} else {
					status = jwt.INVAILD
				}
			} else {
				status = jwt.INVAILD
			}
		} else {
			status = jwt.INVAILD
		}
	}

	return accessToken, status
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
