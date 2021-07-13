package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(numInStr([]string{"this IS", "10xYZ", "xy2K77", "Z1K2W0", "xYz"}))
	fmt.Println(numInStr([]string{"abc", "ab10c", "a10bc", "bcd"}))
	fmt.Println(numInStr([]string{"adsd", "aiJASJasda"}))
}

func numInStr(s []string) (results []string) {
	var passed []int
	errTemp := 0
	for i, v := range s {
		tempSplit := strings.Split(v, "")
		for _, v2 := range tempSplit {
			_, err := strconv.Atoi(v2)
			if err != nil {
				errTemp += 1
			}
		}
		if errTemp < len(v) {
			passed = append(passed, i)
		}
		errTemp = 0
	}
	if len(passed) != 0 {
		for _, v := range passed {
			results = append(results, s[v])
		}
	}
	return
}
