package main

import "fmt"

func main() {
	fmt.Println(additiveInverse([]int{12, -7, -1, 3, 5}))
	fmt.Println(additiveInverse([]int{-5, -25, 35}))
}

func additiveInverse(ints []int) (result []int) {
	for _, v := range ints {
		result = append(result, -v)
	}
	return
}
