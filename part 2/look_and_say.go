package main

import (
	"fmt"
	"math"
)

func profit(biaya, jual, modal float32) float64 {
	sub := jual - biaya
	return math.Ceil(float64(sub * modal))
}

func main() {
	fmt.Println(profit(32.67, 45.00, 1200))
	fmt.Println(profit(378.11, 990.00, 99))
	fmt.Println(profit(4.67, 5.00, 78000))
}
