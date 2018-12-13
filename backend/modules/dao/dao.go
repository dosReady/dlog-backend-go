package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "dlog:dlog@tcp(106.10.33.118:3306)/dlog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	} else {
		db.LogMode(true)
	}
	return db
}

func Select(sql string, values ...interface{}) *interface{} {
	var result interface{}
	conn := GetConnection()
	conn.Raw(sql, values).Scan(&result)
	conn.Close()
	return &result
}

func List(sql string, values ...interface{}) *[]interface{} {
	var returnValue []interface{}
	conn := GetConnection()
	rows, err := conn.Raw(sql, values).Rows()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	} else {
		for rows.Next() {
			var result interface{}
			if err := rows.Scan(&result); err != nil {
				fmt.Println("ERROR")
				fmt.Println(err)
			}
			returnValue = append(returnValue, &result)
		}
	}
	conn.Close()
	return &returnValue
}
