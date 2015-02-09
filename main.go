package main

import (
	"github.com/zenazn/goji"
	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db gorm.DB

func main() {

	log.Printf(`main`)

	user := web.New()
	goji.Handle("/user/*", user)
	goji.Get("/css/*", http.FileServer(http.Dir(".")))	
	goji.Get("/js/*", http.FileServer(http.Dir(".")))

	user.Use(middleware.SubRouter)
	user.Get("/", UserIndex)
	user.Get("/index", UserIndex)
	user.Get("/new", UserNew)
	user.Post("/new", UserCreate)
	user.Get("/edit/:id", UserEdit)
	user.Post("/update/:id", UserUpdate)
	user.Get("/delete/:id", UserDelete)

	goji.Serve()
}

func init() {
	db, _  = gorm.Open("mysql", "root:pass@/go?charset=utf8&parseTime=True")
}
