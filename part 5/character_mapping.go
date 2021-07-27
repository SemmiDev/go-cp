package main

import (
	"fmt"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}

func characterMapping(s string) (r []int) {
	var temp []KeyValue
	splitted := strings.Split(s, "")
	temp = append(temp, KeyValue{Key: splitted[0], Value: 0})
	count := 1
	for i := 1; i < len(splitted); i++ {
		prev := len(temp)
		for j := 0; j < len(temp); j++ {
			if temp[j].Key == splitted[i] {
				temp = append(temp, KeyValue{Key: splitted[i], Value: temp[j].Value})
				break
			}
		}
		if len(temp) != prev {
			continue
		}
		temp = append(temp, KeyValue{Key: splitted[i], Value: count})
		count++
	}
	for _, v := range temp {
		r = append(r, v.Value)
	}
	return
}

func main() {
	fmt.Println(characterMapping("babbcb"))
	fmt.Println(characterMapping("aaabb"))
	fmt.Println(characterMapping("fluffy"))
	fmt.Println(characterMapping("buttery"))
	fmt.Println(characterMapping("canine"))
}
