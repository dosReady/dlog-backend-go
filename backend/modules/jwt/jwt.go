package modules

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	config "github.com/dosReady/dlog/backend/modules/config"
)

type JwtException struct {
	Code uint32
}

const (
	INVAILD uint32 = 20
	EXPIRED uint32 = 16
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
	Data interface{}
	jwt.StandardClaims
}

func CreateAccessToken(obj interface{}) string {
	cfg := config.New()

	if payload, ok := obj.(PayLoad); ok {
		obj = payload.Data
	}

	payload := PayLoad{
		Data: obj,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "dlog",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.GetAlg()), &payload)
	tokenstr, _ := token.SignedString([]byte(cfg.GetJwtAccessSecret()))
	return tokenstr
}
func CreateRefreshToken(obj interface{}) string {
	cfg := config.New()
	payload := PayLoad{
		Data: obj,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 30, 0).Unix(),
			Issuer:    "dlog",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.GetAlg()), &payload)
	tokenstr, _ := token.SignedString([]byte(cfg.GetJwtRefreshSecret()))
	return tokenstr
}
func VaildAccessToken(tokenString string) (*PayLoad, *JwtException) {
	cfg := config.New()
	decodeAccess, err := _decodeToken(tokenString, cfg.GetJwtAccessSecret())
	if err != nil {
		return nil, err
	} else {
		return decodeAccess, nil
	}
}
func VaildRefreshToken(tokenString string) (*PayLoad, *JwtException) {
	cfg := config.New()
	decdoeRefresh, err := _decodeToken(tokenString, cfg.GetJwtRefreshSecret())
	if err != nil {
		return nil, err
	} else {
		return decdoeRefresh, nil
	}
}

func _decodeToken(tokenString string, secret string) (*PayLoad, *JwtException) {
	token, err := jwt.ParseWithClaims(tokenString, &PayLoad{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &JwtException{Code: INVAILD}
		}
		return []byte(secret), nil
	})
	if err != nil {
		parseE, _ := err.(*jwt.ValidationError)
		if parseE.Errors == EXPIRED {
			return nil, &JwtException{Code: EXPIRED}
		} else {
			return nil, &JwtException{Code: INVAILD}
		}
	}

	if claims, ok := token.Claims.(*PayLoad); ok && token.Valid {
		return claims, nil
	} else {
		return nil, &JwtException{Code: INVAILD}
	}
}
