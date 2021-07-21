package main

import "fmt"

func main() {
	caseA := []int{1, 1, 1, 1, -3, -4}
	caseB := []int{5, 4, 3, 0, 0, -1, -1, -2, -2}
	caseC := []int{52, 52, 52, -3, 2, 2, 2, -4}
	caseD := []int{4, -2, -1, -4}

	fmt.Println(isPositiveDominant(caseA) == true)
	fmt.Println(isPositiveDominant(caseB) == true)
	fmt.Println(isPositiveDominant(caseC) == true)
	fmt.Println(isPositiveDominant(caseD) == false)
}

func isPositiveDominant(args []int) bool {
	pCount, nCount := 0, 0
	for _, v := range args {
		if v < 0 {
			nCount++
		} else {
			pCount++
		}
	}
	return pCount > nCount
}
