package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

type Config struct {
	DB struct {
		Host     string
		Port     uint `default:"3306"`
		Database string
		User     string
		Password string `required:"true"`
	}
	Jwt struct {
		Accesssecret  string
		Refreshsecret string
		Alg           string
	}
}

func New() *Config {
	c := Config{}
	dir, _ := os.Getwd()
	_ = configor.Load(&c, fmt.Sprintf("%s/modules/config/%s", dir, "config.yml"))
	return &c
}

func (c Config) GetDbURL() string {
	parms := []interface{}{c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database}
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", parms...)
	return url
}

func (c Config) GetJwtAccessSecret() string {
	return c.Jwt.Accesssecret
}
func (c Config) GetJwtRefreshSecret() string {
	return c.Jwt.Refreshsecret
}
func (c Config) GetAlg() string {
	return c.Jwt.Alg
}
