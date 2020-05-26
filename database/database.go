package database

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	G_db *gorm.DB
)

type Username struct {
	gorm.Model
	Name string
}

func Start() {
	db, err := gorm.Open("mysql", "root:mima@(127.0.0.1:3306)/github_username?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err)
		errors.New("open database error!")
	}
	G_db = db

	if db.HasTable(Username{}) {
		db.AutoMigrate(Username{})
	} else {
		if err = db.CreateTable(&Username{}).Error; err != nil {
			fmt.Println(err)
			errors.New("create table named usernames failed!")
		}
	}
}

func Insert(name string) {
	G_db.Create(&Username{Name: name})
}
