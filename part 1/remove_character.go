package main

import "fmt"

const Char = "C"

func main() {
	fmt.Println(removeCharacter([]string{"Kambing", "Chapung",  "Kalong"}))
	fmt.Println(removeCharacter([]string{"Cacing", "Bebek", "Cicak", "Beruang"}))
}

func removeCharacter(slc []string) (result []string) {
	for _, v := range slc {
		if string(v[0]) != Char {
			result = append(result, v)
		}
	}
	return
}
