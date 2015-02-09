package main

import(
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"models"
)

func main() {
	db, _ := gorm.Open("mysql", "root:pass@/go?charset=utf8&parseTime=True")
	db.CreateTable(&models.User{})
}
