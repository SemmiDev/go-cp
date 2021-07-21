package main

import "fmt"

func main() {
	matrix := [][]int {
		{1,2},
		{3,4},
	}

	matrix2 := [][]int {
		{5,3},
		{3,1},
	}

	matrix3 := [][]int {
		{1,1},
		{1,1},
	}

	fmt.Println(calcDeterminant(matrix))
	fmt.Println(calcDeterminant(matrix2))
	fmt.Println(calcDeterminant(matrix3))
}

func calcDeterminant(matrix [][]int) int {
	mainDiagonal := matrix[0][0] * matrix[1][1]
	mainDiagona2 := matrix[0][1] * matrix[1][0]
	return mainDiagonal - mainDiagona2
}
