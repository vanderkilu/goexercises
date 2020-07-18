package main

import (
	"fmt"
	"strings"
)

func atBash(str string) string {
	signature := "abcdefghijklmnopqrstuvwxyz"
	newStr := ""
	for _, char := range str {
		index := strings.Index(str, string(char))
		newStr += string(signature[index%26])
	}
	return newStr
}

func main() {
	test := atBash("gvhg")
	fmt.Println(test)
}
