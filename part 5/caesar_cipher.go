package main

import (
	"fmt"
	"strings"
)

// ups many bugs :(

func main() {
	fmt.Println(caesarCipher("A Fool and His Money Are Soon Parted", 27))
}
func caesarCipher(s string, n int) string {
	s = strings.ToLower(s)
	splitted := strings.Split(s, "")
	var base int32 = 97
	var temp []int32
	for _, v := range splitted {
		runes := []rune(v)
		for _, run := range runes {
			t := run + int32(n)
			if t > 122 {
				l := t - 122 - 1
				b := base + l
				temp = append(temp, b)
			} else {
				temp = append(temp, t)
			}
		}
	}
	var t strings.Builder
	for _, v := range temp {
		t.Write([]byte(string(v)))
	}
	return t.String()
}
