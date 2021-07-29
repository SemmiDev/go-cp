package main

import "fmt"

func search(arr []int , x int) (i int) {
	for i, v := range arr {
		if x == v {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1,4,6,8}, 10))
	fmt.Println(search([]int{1,4,6,8}, 4))
}
