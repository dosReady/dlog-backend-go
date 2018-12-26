package modules

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	config "github.com/dosReady/dlog/backend/modules/config"
	utils "github.com/dosReady/dlog/backend/modules/utils"
	"github.com/rs/xid"
)

type JwtException struct {
	Code uint32
}

const (
	INVAILD uint32 = 20
	EXPIRED uint32 = 16
	PASS    uint32 = 0
)

func (je JwtException) Error() string {
	switch {
	case je.Code == INVAILD:
		return fmt.Sprintln("[JWT] 유효하지않은 토큰입니다.")
	case je.Code == EXPIRED:
		return fmt.Sprintln("[JWT] 만료된 토큰입니다.")
	default:
		return fmt.Sprintln("[JWT] 알수없는 오류입니다.")
	}
}

type PayLoad struct {
	Data []byte
	Xid  string
	jwt.StandardClaims
}

func CreateAccessToken(obj interface{}) (string, string) {
	cfg := config.New()
	xidstr := xid.New().String()
	jsonobj := utils.EncodingJson(obj)

	payload := PayLoad{
		Data: jsonobj,
		Xid:  xidstr,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "dlog",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.GetAlg()), &payload)
	tokenstr, _ := token.SignedString([]byte(cfg.GetJwtAccessSecret()))
	return tokenstr, xidstr
}
func CreateRefreshToken(xidval string) string {
	cfg := config.New()
	payload := struct {
		Xid string
		jwt.StandardClaims
	}{
		Xid: xidval,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 30, 0).Unix(),
			Issuer:    "dlog",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.GetAlg()), payload)
	tokenstr, _ := token.SignedString([]byte(cfg.GetJwtRefreshSecret()))
	return tokenstr
}
func VaildAccessToken(tokenString string) (*PayLoad, uint32) {
	cfg := config.New()
	decodeAccess, err := _decodeToken(tokenString, cfg.GetJwtAccessSecret())
	return decodeAccess, err
}
func VaildRefreshToken(tokenString string) (*PayLoad, uint32) {
	cfg := config.New()
	decdoeRefresh, err := _decodeToken(tokenString, cfg.GetJwtRefreshSecret())
	return decdoeRefresh, err
}

func _decodeToken(tokenString string, secret string) (*PayLoad, uint32) {
	var payload PayLoad
	token, err := jwt.ParseWithClaims(tokenString, &payload, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &JwtException{Code: INVAILD}
		}
		return []byte(secret), nil
	})

	var exception JwtException
	if err != nil {
		parseE, _ := err.(*jwt.ValidationError)
		if parseE.Errors == EXPIRED {
			exception = JwtException{Code: EXPIRED}
		} else {
			exception = JwtException{Code: INVAILD}
		}
	} else if !token.Valid {
		exception = JwtException{Code: INVAILD}
	}

	return &payload, exception.Code
}
