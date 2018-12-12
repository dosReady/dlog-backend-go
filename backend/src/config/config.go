package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config : struct
type Config struct {
	Db        Db
	Gmail     Gmail
	Jwt       Jwt
	SecretKey string
}

type Db struct {
	Host     string
	Port     string
	User     string
	Database string
	Dbtype   string
	Url      string
}
type Gmail struct {
	User         string
	Password     string
	ClientID     string
	ClientSecret string
	RefeshToekn  string
	AccessToken  string
	Expires      int
}
type Jwt struct {
	AccessSecret   string
	RefreshSecret  string
	AccessAlg      string
	AccessExpires  int
	AefreshExpires int
}

func _initConfig() *Config {
	c := Config{}
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file := fmt.Sprintf("%v/src/config/config.json", pwd)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(b, &c); err != nil {
		panic(err)
	}
	return &c
}

func GetDbConfig() *Db {
	c := _initConfig()
	return &c.Db
}

func GetJwtConfig() *Jwt {
	c := _initConfig()
	return &c.Jwt
}

func GetGmailConfig() *Gmail {
	c := _initConfig()
	return &c.Gmail
}

func GetSecretKey() string {
	c := _initConfig()
	return c.SecretKey
}
