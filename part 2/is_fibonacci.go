package main

import (
	"fmt"
)

func isFibbonaci(n int) {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 1
	f[1] = 2
	for i := 2; i <= n-2; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	fmt.Println(f[:len(f)-2])
}

func main() {
	isFibbonaci(3)
	isFibbonaci(8)
}