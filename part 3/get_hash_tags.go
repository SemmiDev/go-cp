package main

import (
	"fmt"
	"strings"
)

func getHashTags(s string) (result []string) {
	if s != "" {
		splitted := strings.Split(s, " ")
		wordMap := make(map[int]int)
		for i, v := range splitted {
			wordMap[i] = len(v)
		}
		if len(splitted) == 1 {
			return addHashToPrefix(splitted)
		} else if len(splitted) == 2 {
			if len(splitted[0]) < len(splitted[1]) {
				result = append(result, splitted[1])
				result = append(result, splitted[0])
			} else {
				result = append(result, splitted[0])
				result = append(result, splitted[1])
			}
			return addHashToPrefix(result)
		}
		i := reduce(splitted, wordMap)
		result = append(result, splitted[i])
		delete(wordMap, i)
		result = append(result, splitted[i])
		delete(wordMap, i)
		result = append(result, splitted[i])
		delete(wordMap, i)
		return addHashToPrefix(result)
	}
	return
}

func reduce(splitted []string, wordMap map[int]int) (indexTempLv int) {
	for i := 1; i < len(splitted); i++ {
		if wordMap[i] > wordMap[indexTempLv] {
			indexTempLv = i
		}
	}
	return
}

func addHashToPrefix(s []string) (result []string) {
	for _, v := range s {
		temp := strings.Split(v, "")
		resTemp := []string{"#"}
		resTemp = append(resTemp, temp...)
		result = append(result, strings.Join(resTemp, ""))
	}
	return
}

func main() {
	fmt.Println(getHashTags("How the Avocado Became the Fruit of the Global Trade"))
	fmt.Println(getHashTags("Why You Will Probably Pay More for Your Christmas Tree This Year"))
	fmt.Println(getHashTags("Minecraft Stars on Youtube Share Secrets to Their Fame"))
	fmt.Println(getHashTags("Visualizing Science"))
	fmt.Println(getHashTags("xixixi"))
	fmt.Println(getHashTags(""))
}
