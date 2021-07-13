package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncensoredString("Wh*r* d*d my v*w*ls g*?", "eeioeo"))
	fmt.Println(uncensoredString("abcd", ""))
	fmt.Println(uncensoredString("*PP*RC*S*", "UEAE"))
	fmt.Println(uncensoredString("*l*ph*nt", "Eea"))
}

func uncensoredString(msg string, lose string) string {
	if len(lose) == 0 {
		return msg
	}

	loseChars := strings.Split(lose, "")
	msgChars := strings.Split(msg, "")
	i := 0

	for index, v := range msgChars {
		if v == "*" {
			msgChars[index] = loseChars[i]
			i++
		}
	}
	return strings.Join(msgChars, "")
}
