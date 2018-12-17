package modules

import (
	"fmt"
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
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
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

func DecodeAccessToken(tokenString string) *PayLoad {
	cfg := config.New()
	token, err := jwt.ParseWithClaims(tokenString, &PayLoad{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("JWT ERROR")
		}
		return []byte(cfg.GetJwtAccessSecret()), nil
	})
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(*PayLoad); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}
func DecodeRereshToken(tokenString string) *PayLoad {
	cfg := config.New()
	token, err := jwt.ParseWithClaims(tokenString, &PayLoad{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("JWT ERROR")
		}
		return []byte(cfg.GetJwtRefreshSecret()), nil
	})
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(*PayLoad); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}
