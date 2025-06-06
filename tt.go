package main

import (
	"fmt"
)

func main() {

	mp := map[string]int{
		"a": 2,
		"b": 3,
	}
	fmt.Println(mp["a"])

	for key, value := range mp {
		mp[key] = value + 1
		fmt.Println(key, mp[key])
	}

	s := "abc"
	runes := []rune(s)
	length := len(runes)
	mid := (length + 1) / 2
	for i := 0; i < mid; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}
	result := string(runes)
	fmt.Println(result)
}
