package main

import (
	"fmt"
	"strings"
	"regexp"
)

func main() {
	s1 := "sAmmi aldhi yantO"
	//s2 := "rahmatUl izzAh annisa"
	fmt.Println(getCount(s1))
	fmt.Println(GetCount2(s1))
	fmt.Println(GetCount3(s1))
}

func getCount(s string) (count int) {
	for _, v := range strings.ToLower(s) {
		switch v {
		case 'a','i','u','e','o':
			count++
		}
	}
	return
}

func GetCount2(s string) int {
	s = strings.ToLower(s)
	count := 0
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		count += strings.Count(s, vowel)
	}
	return count
}

func GetCount3(s string) int {
	s = strings.ToLower(s)
	r := regexp.MustCompile("[aeiou]")
	vowels := r.FindAllString(s, -1)
	return len(vowels)
}
