package app

import (
	"errors"
	"fmt"
	"github-username-test/database"
	str "github-username-test/string"
	"net/http"
)

var (
	strs string
	ch   chan int
)

var codeToresult = map[int]func(){
	0: func() {
		fmt.Println("some error")
	},
	200: func() {
		fmt.Println(strs + "had been registered")
	},
	403: func() {
		fmt.Println("403 error")
	},
	404: func() {
		database.Insert(strs)
		fmt.Println(strs + "ok")
	},
}

func Start() {
	for i := 0; i < 100; i++ {
		go Judge(30000+i*5000, 9999999)
	}
	ch <- 1
}

func Judge(begin int, end int) {
	for x := begin; x < end; x++ {
		strs = str.NewStr(x)
		fmt.Println(strs)
		code := GetStatusCode(GenderateUrl(strs))
		f := codeToresult[code]
		f()
	}
}

//GetStatusCode 获取状态码
func GetStatusCode(url string) int {
	res, err := http.Get(url)
	if err != nil {
		errors.New("errors!")
	}
	if res == nil {
		return 0
	} else {
		return res.StatusCode
	}
}

//GenderateUrl 生成 url
func GenderateUrl(str string) string {
	return "https://api.github.com/users/" + str
}
