package dao

import (
	"fmt"
	"reflect"

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
func List(ref interface{}, sql string, values ...interface{}) {
	ch := make(chan []interface{})
	fch := make(chan int)
	// 참조 필드 가져오기
	go func() {
		var fields []interface{}
		t := reflect.TypeOf(ref)
		fmt.Println(t)
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fields = append(fields, reflect.New(f.Type))
		}
		ch <- fields
		fch <- 123
	}()
	// 참조 필드 값 셋팅하기
	// https://forum.golangbridge.org/t/database-rows-scan-unknown-number-of-columns-json/7378
	go func(ch chan []interface{}) {
		i := <-fch
		fch <- i + 123
		fields := <-ch
		conn := GetConnection()
		rows, _ := conn.Raw(sql, values).Rows()
		for rows.Next() {
			if err := rows.Scan(fields...); err != nil {
				panic(err)
			}
		}
		ch <- fields
	}(ch)

	// 결과 값 만들기
	go func(ch chan []interface{}) {
		fields := <-ch
		i := <-fch
		fmt.Println(fields)
		fmt.Println(i)
	}(ch)
	<-ch
	close(ch)
}
