package config

import (
	"github.com/ginzhu/gorm"
	_ "github.com/ginzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
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
