package main

import (
	"github.com/zenazn/goji/web"
	
	log "github.com/Sirupsen/logrus"

	"net/http"
	"html/template"
	"models"
	"strconv"
)

var tpl *template.Template

type FormData struct {
	User models.User
	Mess string
}

func UserIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	Users := [] models.User{}
	db.Find(&Users)
	tpl = template.Must(template.ParseFiles("view/user/index.html"))
	log.WithFields(log.Fields {
		"users": Users,
	}).Info("")
	tpl.Execute(w, Users)
}

func UserNew(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("view/user/new.html"))
	tpl.Execute(w, FormData{models.User{}, ""})
}

// TODO: validation
func UserCreate(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{Name: r.FormValue("Name")}
	db.Create(&User)
	http.Redirect(w, r, "/user/index", 301)
}

// TODO: validation
func UserEdit(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Find(&User)
	tpl = template.Must(template.ParseFiles("view/user/edit.html"))
	tpl.Execute(w, FormData{User, ""})
}

// TODO: validation
func UserUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Find(&User)
	User.Name = r.FormValue("Name")
	db.Save(&User)
	http.Redirect(w, r, "/user/index", 301)
}

func UserDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Delete(&User)
	http.Redirect(w, r, "/user/index", 301)
}
