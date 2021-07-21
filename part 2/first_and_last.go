package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(firstAndLast("another"))
	fmt.Println(firstAndLast("welcome"))
	fmt.Println(firstAndLast("anonmymous"))
}

func firstAndLast(msg string) (result []string) {
	sorted := strings.Split(msg, "")
	sort.Strings(sorted)
	result = append(result, strings.Join(sorted, ""))
	// reverse
	j := len(msg)-1
	var sortedReverse []string

	for i := j; i >= 0; i-- {
		sortedReverse = append(sortedReverse , sorted[i])
	}
	result = append(result, strings.Join(sortedReverse, ""))
	return
}