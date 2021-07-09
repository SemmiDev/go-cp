package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseBool(true))
	fmt.Println(reverseBool(false))
}

func reverseBool(b bool) (result bool) {
	result = !b
	return
}
