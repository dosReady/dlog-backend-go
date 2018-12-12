package middleware

import (
	"config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitConnection() gin.HandlerFunc {
	db := config.GetDbConfig()
	conn, err := gorm.Open("mysql", db.Url)
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	m["DB"] = conn
	return func(c *gin.Context) {
		c.Keys = m
	}
}
