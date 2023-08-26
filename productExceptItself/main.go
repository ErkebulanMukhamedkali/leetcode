package main

import "log"

func main() {
	samples := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
	}

	for _, sample := range samples {
		log.Println(productExceptSelf(sample))
	}
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	first := make([]int, len(nums))
	second := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			first[i], second[len(nums)-1] = 1, 1
			continue
		}

		first[i] = nums[i-1] * first[i-1]
		second[len(nums)-i-1] = second[len(nums)-i] * nums[len(nums)-i]
	}

	for i := range result {
		result[i] = first[i] * second[i]
	}

	return result
}
