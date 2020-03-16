//原理：
//* 通过判断 https://github.com/username 的 statuscode 是否为 404 来判断 username 是否占用
//* 通过判断 https://api.github.com/users/username 的 statuscode 是否为 404 来判断 username 是否被占用
//* 通过判断 https://github.com/signup_check/username?suggest_usernames=true 的 statuscode 是否为 404 来判断 username

package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"sync"
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

//TODO: 增加其他网站的 feat
//TODO: 解决误判问题
//TODO: 解决 504 问题
//TODO: 加入数字
//TODO: 将结果按 含韵母 筛选一遍
//TODO: 找出有意义的结果, 如单词，特定缩写，字母表如 orz 等
//TODO: 建立 Ip 池解决 ip 限制问题

func main() {
	DatabasePre()
	for i := 0; i < 100; i++ {
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

		//fmt.Println(result)
		code := GetStatusCode(GenderateUrl(result))
		fmt.Println(code)
		//if code == 200 {
		//	//fmt.Println(result + " is registe!")
		//} else if code == 404 {
		//	fmt.Println(result + " is not registe!")
		//	Inser(result)
		//} else {
		//	//todo：爬一会 会返回 504 ，猜测是 github 的防护措施
		//	fmt.Println(code)
		//	time.Sleep(time.Second * 5)
		c <- 0

	}
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
	return "https://api.github.com/users/" + str
}

func Inser(name string) {
	G_db.Create(&Username{Name: name})
}
