package models

import (
	config "command-line-arguments/home/fadloovich/Public/D/Tuts/go_tuts/goprojects/bookmanagerapi/pkg/config/app.go"

	"github.com/ginzhu/gorm"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}
