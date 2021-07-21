package main

import (
	"fmt"
)

func main() {

	fmt.Println(canConcatenate([]int{2, 1, 3}, []int{5, 4, 7, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(canConcatenate([]int{1, 4}, []int{3}, []int{1, 3, 4}))
	fmt.Println(canConcatenate([]int{1, 4}, []int{3}, []int{4, 3, 1}))
}

func canConcatenate(args1, args2, target []int) bool {
	var joined []int
	joined = append(joined, args1...)
	joined = append(joined, args2...)

	joined = BubbleSort(joined)
	target = BubbleSort(target)

	if len(joined) != len(target) {
		return false
	}

	return comparing(joined, target)
}

func comparing(args1, args2 []int) bool {
	for i := 0; i < len(args1); i++ {
		if args1[i] != args2[i] {
			return false
		}
	}
	return true
}

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
