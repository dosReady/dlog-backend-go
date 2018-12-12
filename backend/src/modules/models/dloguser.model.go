package models

type DlogUser struct {
	UserEmail    string `gorm:"primary_key"`
	UserName     string
	UserPassword string
	Base
}
