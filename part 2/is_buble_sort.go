package main

import "fmt"

func BubbleSort(array[] int)[]int {
	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				fmt.Println(array)
			}
		}
	}
	return array
}
func main() {
	array := []int{303, 6, 90, 21, 33, 80, 212}
	array2 := []int{101, 1002, 0, 212, 12, 4, 10}

	fmt.Println(BubbleSort(array))
	fmt.Println(BubbleSort(array2))
}