package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(balanceWords("forgetful") == false)
	fmt.Println(balanceWords("vegetation") == true)
	fmt.Println(balanceWords("disillusioned") == false)
	fmt.Println(balanceWords("abstract") == true)
	fmt.Println(balanceWords("conditionalities") == true)
}

func balanceWords(s string) bool {
	splitted := strings.Split(strings.ToLower(s), "")
	kv := make(map[string]int)
	for i := 'a'; i <= 'z'; i++ {
		toint, _ := strconv.Atoi(fmt.Sprint(i - 96))
		kv[string(i)] = toint
	}
	left, right := 0, 0
	for i := 0; i < len(splitted)/2; i++ {
		left += kv[splitted[i]]
	}
	base := len(splitted) / 2
	if len(splitted)%2 != 0 {
		base = len(splitted)/2 + 1
	}
	for i := base; i < len(splitted); i++ {
		right += kv[splitted[i]]
	}
	return left == right
}
