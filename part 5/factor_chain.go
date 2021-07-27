package main

import (
	"fmt"
)

func factorChain(ns ...int) (b bool) {
	sum := 0
	for _, v := range ns {
		sum += v
	}
	if len(ns)*ns[0] == sum {
		return true
	}

	for i := 0; i < len(ns)-1; i++ {
		if ns[i]*2 != ns[i+1] {
			return
		}
	}
	b = true
	return
}

func main() {
	fmt.Println(factorChain(1, 2, 4, 8, 16, 32))
	fmt.Println(factorChain(2, 4, 6, 7, 12))
	fmt.Println(factorChain(1, 1, 1, 1, 1, 1))
	fmt.Println(factorChain(10, 1))
	fmt.Println(factorChain(10, 20, 30, 40))
	fmt.Println(factorChain(10, 20, 40))
}
