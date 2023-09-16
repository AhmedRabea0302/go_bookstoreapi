package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:0123456/simplerest?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
