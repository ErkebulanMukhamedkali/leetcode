package main

import "log"

func main() {
	var nums = [][]int{
		{2, 7, 11, 15}, {3, 2, 4}, {3, 3},
		{3, 2, 3},
	}
	var targets = []int{
		9, 6, 6,
		6}

	for i := range targets {
		log.Println(twoSum(nums[i], targets[i]))
	}
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
