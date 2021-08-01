package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceVowel("karachi"))
	fmt.Println(replaceVowel("dang"))
	fmt.Println(replaceVowel("aen"))
	fmt.Println(replaceVowel("chembur"))
	fmt.Println(replaceVowel("khandbari"))
	fmt.Println(replaceVowel("thamel"))
}

func replaceVowel(s string) string {
	split := strings.Split(s, "")

	for i := 0; i < len(split); i++ {
		switch split[i] {
		case "a":
			split[i] = "1"
		case "e":
			split[i] = "2"
		case "i":
			split[i] = "3"
		case "o":
			split[i] = "4"
		case "u":
			split[i] = "5"
		}
	}

	return strings.Join(split, "")
}
