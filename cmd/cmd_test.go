package main

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	for x := 0; x < 100; x++ {
		A(x)
	}
}

func A(x int) {
	result = ""
	quotient = x
	for quotient >= 0 {
		remainder = quotient % 26
		result = string(remainder+97) + result
		quotient = int(quotient/26) - 1
	}

	fmt.Println(result)
}
