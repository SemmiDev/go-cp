// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	caseA, caseB, caseC := "aaaaaaa", "abcdabcd", "shopper"
// 	fmt.Println(distanceToNearestVowel(caseA))	
// 	fmt.Println(distanceToNearestVowel(caseB))	
// 	fmt.Println(distanceToNearestVowel(caseC))	
// }

// func distanceToNearestVowel(s string) (result []int) {
// 	vowels := []string{"a","i","u","e","o"}
// 	sSplitted := strings.Split(s, "")	
// 	var tempValues []int 
// 	for j := 0; j < len(sSplitted); j++ {
// 		tempLen := len(tempValues)
// 		for i := 0; i < len(vowels); i++ {
// 			if sSplitted[j] == vowels[i] {
// 				tempValues = append(tempValues, 0)}
// 		}
// 		if tempLen == len(tempValues) {
// 			tempValues = append(tempValues, 2)}
// 	}
// 	allIsSame := 1
// 	i := 0
// 	for j := 1; j < len(tempValues); j++ {
// 		if tempValues[i] == tempValues[j] {
// 		allIsSame++}}
// 	if allIsSame == len(tempValues) {return tempValues}
// 	var zeroIndexes []int
// 	for i, v := range tempValues {
// 		if v == 0 {
// 			zeroIndexes = append(zeroIndexes, i)
// 		}
// 	}
// 	for i, _ := range zeroIndexes {
// 		if i != len(zeroIndexes)-1 {
// 			zeroIndexes[i] = zeroIndexes[i]+1 
// 		}
// 	}

// 	subDistance := 
// 	fmt.Println(zeroIndexes)

// 	return tempValues
// }