package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"sync"
	"time"
)

var (
	G_db      *gorm.DB
	quotient  int
	remainder int
	wg        sync.WaitGroup
	result    string
	c         chan int
	j         int
)

type Username struct {
	gorm.Model
	Name string
}

func main() {
	DatabasePre()
	for i := 18000; i > 1200; i-- {
		j = i * 10
		go Start(j, j+9999)
	}
	<-c
}

//穷举字符串的思路是 数字转换为字母，26进制，再利用 Ascii 码之间的关系
func Start(begin int, end int) {
	for x := begin; x < end; x++ {
		result = ""
		quotient = x
		for quotient >= 0 {
			remainder = quotient % 26
			result = string(remainder+97) + result
			quotient = int(quotient/26) - 1
		}
		fmt.Println(result)
		code := GetStatusCode(GenderateUrl(result))
		if code == 200 {
			//fmt.Println(result + " is registe!")
		} else if code == 404 {
			fmt.Println(result + " is not registe!")
			Inser(result)
		} else {
			//todo：爬一会 会返回 504 ，猜测是 github 的防护措施
			fmt.Println(code)
			time.Sleep(time.Second * 5)
		}
	}
	c <- 0
}

func DatabasePre() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/github_username?charset=utf8&parseTime=true")
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

func GetStatusCode(url string) int {
	res, err := http.Get(url)
	if err != nil {
		errors.New("errors!")
	}
	return res.StatusCode
}

func GenderateUrl(str string) string {
	return "https://github.com/" + str
}

func Inser(name string) {
	G_db.Create(&Username{Name: name})
}
