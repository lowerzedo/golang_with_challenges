package main

import (
	"fmt"
	"strings"
)

func createRange(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func generatePattern(n int) string {
	if n > 10 {
		fmt.Println("The number should be less than 10")
	}

	var numbers []int
	var temp string
	// divider basically
	var delim string = ""

	numbers = createRange(1, n)

	for len(numbers) > 0 {
		temp = temp + strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", delim, -1), "[]") + "\n"
		numbers = numbers[:len(numbers)-1]
	}

	fmt.Println(temp)

	return temp
}

func main() {
	generatePattern(10)
}
