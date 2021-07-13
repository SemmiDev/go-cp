package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(oddishOrEvenish(43))
	fmt.Println(oddishOrEvenish(55551))
	fmt.Println(oddishOrEvenish(11))
	fmt.Println(oddishOrEvenish(2111122212312312312))
	fmt.Println("elapsed: ", time.Since(now))
}

func oddishOrEvenish(i int) (r string) {
	// divide % 2 -> sisa 0 ? evenish : oddish
	r = "Evenish"
	if i%2 != 0 {
		r = "Oddish"
	}
	return
}
