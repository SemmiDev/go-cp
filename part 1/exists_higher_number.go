package main

import "fmt"

func main() {
	fmt.Println(existsHigherNumber([]int{8, 4, 20, 32, 1}, 10))
	fmt.Println(existsHigherNumber([]int{1, 3, 7, 4, 6}, 8))
	fmt.Println(existsHigherNumber([]int{}, 1))
}

func existsHigherNumber(num []int, i int) bool {
	if len(num) == 0 {
		return false
	}

	for _, v := range num {
		if v > i {
			return true
		}
	}
	return false
}
