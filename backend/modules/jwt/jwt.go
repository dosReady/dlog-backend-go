package modules

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"

	jwt "github.com/dgrijalva/jwt-go"
	config "github.com/dosReady/dlog/backend/modules/config"
)

type PayLoad struct {
	Data interface{}
	jwt.StandardClaims
}

func CreateAccessToken(obj interface{}) string {
	cfg := config.New()
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

func VaildAccessToken(tokenString string) {
	cfg := config.New()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("JWT ERROR")
		}
		return []byte(cfg.GetJwtAccessSecret()), nil
	})
	if err != nil {
		panic(err)
	}

	var payload struct {
		UserEmail string
		UserCall  string
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_ = mapstructure.Decode(claims, &payload)
	}
	fmt.Println(payload)
}
