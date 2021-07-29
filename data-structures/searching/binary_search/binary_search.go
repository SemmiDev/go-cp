package main

import (
	"fmt"
	"math"
)

func search(arr []int, x int) int {
	low, high := 0, len(arr) - 1
	count := 0

	for low <= high {
		mid := (low + high) / 2 // 3
		fmt.Println("--------------------")
		fmt.Println("LANGKAH KE -> ", count)
		fmt.Println("LOW -> ", low)
		fmt.Println("HIGH -> ", high)
		fmt.Println("MID -> ", mid)
		fmt.Println("ARR[MID] -> ", arr[mid])
		fmt.Println("WANT -> ", x)
		fmt.Println("--------------------")
		if arr[mid] == x {
			fmt.Println("TOTAL LANGKAH: ", count)
			fmt.Println("LOG BASIS 2: ", math.Log2(float64(x)))
			return mid
		}else if arr[mid] < x  {
			low = mid + 1
		}else {
			high = mid - 1
		}
		count++
	}

	return -1
}

func main() {
	var arr []int
	for i := 1; i <= 256; i++ {
		arr = append(arr, i)
	}

	fmt.Println(search(arr, 256))
}
