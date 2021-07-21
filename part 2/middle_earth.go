package main

import "fmt"

func main() {
	fmt.Println(middleEarth([]string{"Ari", "Danang", "Jamal"}))
	fmt.Println(middleEarth([]string{"usi", "Ari", "Danang", "Jono"}))
	fmt.Println(middleEarth([]string{"Ruben", "Ari", "Restu", "Danang", "Arif"}))
	fmt.Println(middleEarth([]string{"Ruben", "Ari", "Restu", "Danang", "Arif"}))
}

func middleEarth(strings []string) (status bool) {
	key1, key2 := "Ari", "Danang"
	var tempKeyFound1, tempKeyFound2 int

	for i := 0; i < len(strings); i++ {
		if strings[i] == key1 {
			tempKeyFound1 = i
		}
		if strings[i] == key2 {
			tempKeyFound2 = i
		}
	}

	sub := tempKeyFound2 - tempKeyFound1
	if sub == 1 {
		status = true
	}
	return
}
