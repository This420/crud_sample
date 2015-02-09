package models

import(
	"time"
	"regexp"
)

type User struct {
	Id int64
	Name string `sql:"size:255"`
	CreatedAt time.Time
	updatedAt time.Time
	DeletedAt time.Time
}
