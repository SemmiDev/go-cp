package main

import "fmt"

func isPalindrome(msg string) bool {
	for i := 0; i < len(msg)/2; i++ {
		if msg[i] != msg[len(msg)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("aaaaa"))
	fmt.Println(isPalindrome("asa"))
	fmt.Println(isPalindrome("sammas"))
}
