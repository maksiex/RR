package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	db   *gorm.DB
	pass = os.Getenv("MYSQL_PASS")
)

func Connect() {
	dsn := pass + "root:rootpass1234@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
