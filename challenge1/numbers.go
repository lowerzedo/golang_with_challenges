package main

import "fmt"

func createRange(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func sumAndDivide(num1, num2 int) float64 {
	range_of_nums := createRange(num1, num2)

	sum := 0
	for _, v := range range_of_nums {
		sum = v + sum
	}

	final_answer := float64(sum) / (float64(num1) + float64(num2))
	fmt.Println(final_answer)

	return final_answer
}

func main() {
	sumAndDivide(23, 123)
}
