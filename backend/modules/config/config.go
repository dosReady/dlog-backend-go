package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type Config struct {
	DB struct {
		Host     string
		Port     uint `default:"3306"`
		Name     string
		User     string
		Password string `required:"true"`
	}
}

func _new() *Config {
	c := Config{}
	_ = configor.Load(&c, "config.yml")
	return &c
}

func GetDbURL() string {
	cfg := _new()
	parms := []interface{}{cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name}
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", parms...)
	return url
}
