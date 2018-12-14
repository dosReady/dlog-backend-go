package dlog

import (
	"net/http"
	"time"

	"github.com/dosReady/dlog/backend/modules/dao"
	"github.com/gin-gonic/gin"
)

type DlogUser struct {
	UserEmail    string
	UserPassword string
	UserCall     string
	CreateDate   time.Time
	UpdateDate   time.Time
}

func UserSelect(c *gin.Context) {
	/*
		conn := dao.GetConnection()
		rows, _ := conn.Raw("SELECT * FROM dlog_user").Rows()
		defer rows.Close()
		for rows.Next() {
			var user DlogUser
			if err := conn.ScanRows(rows, &user); err != nil {
				fmt.Println(err)
			}
			dlogusers = append(dlogusers, user)
		}
	*/
	var user DlogUser
	dao.List(user, "SELECT * FROM dlog_user")
	c.JSON(http.StatusOK, gin.H{"name": "qwe"})
}
