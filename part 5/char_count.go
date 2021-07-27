package main

import (
	"fmt"
	"strings"
)

func charCount(s string, msg string) (n int) {
	splitted := strings.Split(msg, "")
	for _, v := range splitted {
		if v == s {
			n++
		}
	}
	return
}

func main() {
	fmt.Println(charCount("b", "big fat bubble") == 4)
	fmt.Println(charCount("c", "Chamber of secrets") == 1)
	fmt.Println(charCount("f", "frank and his friends have offered five foxes for sale") == 7)
	fmt.Println(charCount("x", "edabit") == 0)
	fmt.Println(charCount("a", "Adam and Eve bit the apple and found a snake") == 6)
	fmt.Println(charCount("7", "10795426697") == 2)
}
