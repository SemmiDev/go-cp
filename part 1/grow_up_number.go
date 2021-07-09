package main

import "fmt"

func growUpNumber(n int) (arr []int){
	if n == 0 {
		return
	}
	for i := n; i >= 0; i-- {
		arr = append(arr, i)
	}
	return
}

func main() {
	fmt.Println(growUpNumber(5))
	fmt.Println(growUpNumber(3))
	fmt.Println(growUpNumber(0))
}