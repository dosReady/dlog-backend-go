package middleware

import (
	"config"
	"fmt"
	"modules/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func _initConnection() *gorm.DB {
	db := config.GetDbConfig()
	conn, err := gorm.Open("mysql", db.Url)
	if err != nil {
		panic(err)
	}
	return conn
}

func InitTables(isStart bool) {
	if isStart {
		fmt.Println("======================== DB 테이블 초기화 ========================")
		conn := _initConnection()
		var tables []interface{}
		tables = append(tables, models.DlogUser{})
		for _, table := range tables {
			conn.DropTableIfExists(table)
			conn.CreateTable(table)
		}
		fmt.Println("======================== DB 테이블 설정완료 ========================")
	}
}

func InitConnection() gin.HandlerFunc {
	conn := _initConnection()
	m := make(map[string]interface{})
	m["DB"] = conn
	return func(c *gin.Context) {
		c.Keys = m
	}
}
