package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config : struct
type Config struct {
	Db struct {
		Host     string
		Port     string
		User     string
		Database string
		Dbtype   string
	}
	Gmail struct {
		User         string
		Password     string
		ClientID     string
		ClientSecret string
		RefeshToekn  string
		AccessToken  string
		Expires      int
	}
	Jwt struct {
		AccessSecret   string
		RefreshSecret  string
		AccessAlg      string
		AccessExpires  int
		AefreshExpires int
	}
	SecretKey string
}

func (c Config) GetConfig() *Config {
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
	fmt.Println(c.Db.Host)
	return &c
}
