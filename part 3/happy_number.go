package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(happyNumber(203)) // expect -> true
	fmt.Println(happyNumber(11))  // expect -> true
	fmt.Println(happyNumber(107)) // expect -> true
}

func happyNumber(num int) (status bool) {
	tempBefore, temp := num, num
	for temp != 1 {
		strToInt := splitted(temp)
		temp = 0
		for _, v := range strToInt {
			temp += v * v
			if temp > tempBefore {
				return
			}
		}
	}
	return true
}

func splitted(num int) (strToInt []int) {
	numSpl := strings.Split(strconv.Itoa(num), "")
	for _, v := range numSpl {
		n, _ := strconv.Atoi(v)
		strToInt = append(strToInt, n)
	}
	return
}
