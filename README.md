原因：多次尝试更改 github 用户名，但都被占用，故怒开此 rep

原理：

* 通过判断 https://github.com/username 的 statuscode 是否为 404 来判断 username 是否占用
* 通过判断 https://api.github.com/users/username 的 statuscode 是否为 404 来判断 username 是否被占用
* 通过判断 https://github.com/signup_check/username?suggest_usernames=true 的 statuscode 是否为 404 来判断 username 是否被占用