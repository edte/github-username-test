原因：多次尝试更改 github 用户名，但都被占用，故怒开此 rep
原理：通过判断 https://github.com/username 的 statuscode 是否为 404 来判断 username 是否占用
问题：尚存在一些问题，如一段时间后会返回 504, 大概是 github 的保护措施，还存在误判的问题。