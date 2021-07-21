package main

import "fmt"

func main() {
	a := []int{1,2,3}
	b := []int{4,5}
	c := []int{6,7}
	
	fmt.Println(mergeArray(a,b,c))
}

func mergeArray(data ...[]int) []int {
	merged := []int{}
	for _, v := range data {
		for _, v2 := range v {
			merged = append(merged, v2)
		}
	}
	return merged
}