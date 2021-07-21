package main

import "fmt"

func main() {
	fmt.Println(factorGroup(33))
	fmt.Println(factorGroup(36))
	fmt.Println(factorGroup(7))
}

func factorGroup(i int) (result string){
	result = "even"
	if i % 2 !=  0{
		result = "odd"
	}
	return
}