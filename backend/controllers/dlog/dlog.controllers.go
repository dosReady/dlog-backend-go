package dlog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	dao "github.com/dosReady/dlog/backend/modules/dao"
)

type DlogUser struct {
	UserEmail    string
	UserPassword string
	UserCall     string
	CreateDate   time.Time
	UpdateDate   time.Time
}

func UserSelect(c *gin.Context) {
	var dlogusers []DlogUser
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

	c.JSON(http.StatusOK, gin.H{"name": dlogusers})
}
