package main

import "fmt"

func main() {

	fmt.Println(consecutiveCombo([]int{1, 4, 5, 7}, []int{2, 3, 6}))
	fmt.Println(consecutiveCombo([]int{33, 34, 40}, []int{39, 38, 37, 36, 35, 32, 31, 30}))
	fmt.Println(consecutiveCombo([]int{4, 3, 1}, []int{2, 5, 7, 6}))
	fmt.Println(consecutiveCombo([]int{4, 3, 1}, []int{7, 6, 5}))
}

func consecutiveCombo(args1, args2 []int) bool {
	// join
	var joined []int
	joined = append(joined, args1...)
	joined = append(joined, args2...)

	var total int
	for _, v := range joined {
		total += v
	}

	min := joined[0]
	max := joined[0]

	for i := 0; i < len(joined); i++ {
		if joined[i] < min {
			min = joined[i]
		}
		if joined[i] > max {
			max = joined[i]
		}
	}

	var totalTemp int
	for min <= max {
		totalTemp += min	
		min++
	}

	return total == totalTemp
}
