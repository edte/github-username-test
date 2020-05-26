package app

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	fmt.Println(GetDay(1))
}

func GetDay(mouth int) int {
	if mouth == 1 {
		return 31
	} else if mouth == 2 {
		return 23
	} else if mouth == 3 {
		return 30
	} else {
		return 0
	}
}
