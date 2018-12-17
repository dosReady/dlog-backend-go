package modules

import (
	"errors"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	config "github.com/dosReady/dlog/backend/modules/config"
)

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
func VaildAccessToken(tokenString string) *PayLoad {
	cfg := config.New()
	decodeAccess, err := DecodeToken(tokenString, cfg.GetJwtAccessSecret())
	if err != nil {
		log.Println(err)
		return nil
	} else {
		return decodeAccess
	}
}
func VaildRefreshToken(tokenString string) *PayLoad {
	cfg := config.New()
	decdoeRefresh, err := DecodeToken(tokenString, cfg.GetJwtRefreshSecret())
	if err != nil {
		log.Println(err)
		return nil
	} else {
		return decdoeRefresh
	}
}

func DecodeToken(tokenString string, secret string) (*PayLoad, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PayLoad{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token error")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*PayLoad); ok && token.Valid {
		return claims, nil
	} else {
		err := errors.New("tokenIsInVaild")
		return nil, err
	}
}
