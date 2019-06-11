package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {
	params := "root:dev.spacej.tech@(www.spacej.vip:33306)/gyTest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", params)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
