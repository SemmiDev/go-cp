package main

import (
	"fmt"
	"strings"
)

func doubleChar(s string) string {
	var temp []string
	for _, v := range strings.Split(s, "") {
		temp = append(temp, v, v)
	}
	return strings.Join(temp, "")
}

func main() {
	fmt.Println(doubleChar("String") == "SSttrriinngg")
	fmt.Println(doubleChar("Hello World!") == "HHeelllloo  WWoorrlldd!!")
	fmt.Println(doubleChar("1234!_ ") == "11223344!!__  ")
	fmt.Println(doubleChar("_______") == "______________")
}
