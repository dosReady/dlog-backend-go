package dao

import (
	"fmt"

	config "github.com/dosReady/dlog/backend/modules/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {
	c := config.New()
	db, err := gorm.Open("mysql", c.GetDbURL())
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	} else {
		db.LogMode(true)
	}
	return db
}

/*
func List(ref interface{}, sqlstr string, values ...interface{}) {
	conn := GetConnection()
	rows, err := conn.Raw(sqlstr, values...).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var objects []map[string]interface{}

	for rows.Next() {
		columns, err := rows.ColumnTypes()
		if err != nil {
			panic(err)
		}

		results := make([]interface{}, len(columns))
		object := make(map[string]interface{})

		for i, column := range columns {
			object[column.Name()] = new(*string)
			results[i] = object[column.Name()]
		}
		if err := rows.Scan(results...); err != nil {
			panic(err)
		}
		objects = append(objects, object)
	}
	jsonByte, _ := json.Marshal(objects)
	if err := json.Unmarshal(jsonByte, &ref); err != nil {
		panic(err)
	}
}
*/
