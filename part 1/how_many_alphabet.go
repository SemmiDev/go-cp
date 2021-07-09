package main

import (
	"fmt"
	"strings"
)

var (
	msg = "LAMBORGHINI"
	key = "A"
)

func main() {
	fmt.Println(len(howManyAlphabet(7)) == (len(msg) - 1) + 7)
	fmt.Println(len(howManyAlphabet(2)) == (len(msg) - 1) + 2)
	fmt.Println(len(howManyAlphabet(10)) == (len(msg) - 1) + 10)
}

func howManyAlphabet(count int) string {
	msg = strings.ToLower(msg)
	key = strings.ToLower(key)

	splitted := strings.Split(msg, "")
	for i := 0; i < len(msg); i++ {
		if splitted[i] ==  key{
			temp := ""
			for j := 0; j < count ; j++ {
				temp += key
			}
			temp2 := splitted[i+1:]
			splitted = append(splitted[:i], temp)
			splitted = append(splitted, temp2...)

			return strings.Join(splitted, "")
		}
	}
	return strings.Join(splitted, "")
}
