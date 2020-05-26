package str

// NewStr 生成一个新字符串，输入的 x 决定一个字符串，连续的 x 则输出连续的字符串
// 原理是把数字转换为字母，26进制，再利用 Ascii 码之间的关系
func NewStr(x int) string {
	result := ""
	quotient := x
	for quotient >= 0 {
		remainder := quotient % 26
		result = string(remainder+97) + result
		quotient = int(quotient/26) - 1
	}
	return result
}
