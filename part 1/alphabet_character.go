package main

import (
	"fmt"
	"sort"
	"strings"
)

// wkwkwkk
func alphabetCharacter(msg string) (splitted []string) {
	msg = strings.Replace(msg, " ", "", -1)
	splitted = strings.Split(msg, "")
	sort.Strings(splitted)
	return
}

func main() {
	fmt.Println(alphabetCharacter("aaaagkkmnsuuy"))
}
