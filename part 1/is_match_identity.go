package main

import (
	"fmt"
	"strings"
)

func isMatchIdentity(msg string) bool {
	vowels := []string{"a","i","u","e","o"}
	splitted := strings.Split(msg, "")

	eachSplittedContainsVowels := 0
	for i := 0; i < len(splitted); i++ {
		for j := 0; j < len(vowels); j++ {
			if splitted[i] == vowels[j] {
				eachSplittedContainsVowels++
			}
		}
	}

	if eachSplittedContainsVowels != len(splitted) {
		return false
	}

	key := splitted[0]
	for i := 1; i < len(splitted); i++ {
		if splitted[i] != key{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isMatchIdentity("aaaa"))
	fmt.Println(isMatchIdentity("aiaaa"))
	fmt.Println(isMatchIdentity("ccccccccccc"))
	fmt.Println(isMatchIdentity("bbaa"))
	fmt.Println(isMatchIdentity("oooooooo"))
}
