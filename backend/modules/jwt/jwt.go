package modules

import (
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
	paylod := PayLoad{
		Data: obj,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "dlog",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.GetAlg()), &paylod)
	tokenstr, _ := token.SignedString([]byte(cfg.GetJwtAccessSecret()))
	return tokenstr
}
func CreateRefreshToken(obj interface{}) string {
	cfg := config.New()
	paylod := PayLoad{
		Data: obj,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 30, 0).Unix(),
			Issuer:    "dlog",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.GetAlg()), &paylod)
	tokenstr, _ := token.SignedString([]byte(cfg.GetJwtRefreshSecret()))
	return tokenstr
}
