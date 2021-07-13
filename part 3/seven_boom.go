package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{8, 6, 33, 100}
	c := []int{2, 55, 60, 97, 86}
	d := []int{2, 55, 60, 97, 86, 10000007}

	fmt.Println(sevenBoom(a))
	fmt.Println(sevenBoom(b))
	fmt.Println(sevenBoom(c))
	fmt.Println(sevenBoom(d))
}

func sevenBoom(a []int) string {
	var aToString []string
	for _, v := range a {
		if v > 10 {
			temp := strconv.Itoa(v)
			splitted := strings.Split(temp, "")
			aToString = append(aToString, splitted...)
		}
		aToString = append(aToString, strconv.Itoa(v))
	}

	for _, v := range aToString {
		if v == "7" {
			return "Boom Exploded"
		}
	}
	return "Boom Defused"
}
