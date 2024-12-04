package main

import (
	"fmt"
	"strings"
)

func Solution(word string) string {
	var answer strings.Builder
	for i := len(word) - 1; i > -1; i-- {
		fmt.Println(i)
		letter := word[i]
		answer.WriteString(string(letter))

	}
	return answer.String()
}

func Solution2(word string) string {
	var result string
	for _, c := range word {
		result = string(c) + result
	}
	return result
}

func main() {
	fmt.Println(Solution2("Missouri"))
}
