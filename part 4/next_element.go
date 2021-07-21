package main

import (
	"fmt"
)

func main() {
	caseA := []int{3, 5, 7, 9}
	caseB := []int{-5, -6, -7}
	caseD := []int{2, 2, 2, 2, 2}

	fmt.Println(nextElement(caseA))
	fmt.Println(nextElement(caseB))
	fmt.Println(nextElement(caseD))
}

func nextElement(args []int) int {
	sub := args[1] - args[0]
	nextElem := args[len(args)-1] + sub
	return nextElem
}
