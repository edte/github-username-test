//原理：
//* 通过判断 https://github.com/username 的 statuscode 是否为 404 来判断 username 是否占用
//* 通过判断 https://api.github.com/users/username 的 statuscode 是否为 404 来判断 username 是否被占用
//* 通过判断 https://github.com/signup_check/username?suggest_usernames=true 的 statuscode 是否为 404 来判断 username

//TODO: 增加其他网站的 feat
//TODO: 解决误判问题
//TODO: 解决 504 问题
//TODO: 加入数字
//TODO: 将结果按 含韵母 筛选一遍
//TODO: 找出有意义的结果, 如单词，特定缩写，字母表如 orz 等
//TODO: 建立 Ip 池解决 ip 限制问题

package main

import (
	"github-username-test/app"
	"github-username-test/database"
)

func main() {
	database.Start()
	app.Start()
}
