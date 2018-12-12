package controllers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"

	"modules/models"

	"github.com/gin-gonic/gin"
)

func DlogMe(c *gin.Context) {
	m := c.Keys
	db := gorm.DB{}
	if err := mapstructure.Decode(m["DB"], &db); err != nil {
		panic(err)
	}
	user := models.DlogUser{}
	db.Find(&user)
	fmt.Println(user.UserEmail)
	c.JSON(200, gin.H{"name": "me"})
}
