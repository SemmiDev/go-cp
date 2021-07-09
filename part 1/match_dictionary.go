package main

import "fmt"

func main() {
	fmt.Println(matchDictonary("ku", []string{"kuda", "kurus", "kece", "kucel"}))
	fmt.Println(matchDictonary("di", []string{"diman", "dih", "debora", "deh", "die"}))
}

func matchDictonary(s string, strings []string) (result []string) {
	for _, v := range strings {
		if (string(v[0]) + string(v[1])) == s{
			result = append(result, v)
		}
	}
	return
}
