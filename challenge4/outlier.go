package main

import (
	"fmt"
)

func FindOutlier(integers []int) int {
	var even_nums []int
	var odd_nums []int
	for _, v := range integers {
		fmt.Println(v)
		if v%2 == 0 {
			even_nums = append(even_nums, v)
		} else {
			odd_nums = append(odd_nums, v)
		}
	}
	if len(even_nums) == 1 {
		return even_nums[0]
	} else {
		return odd_nums[0]
	}
}

func main() {
	FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781})
}
