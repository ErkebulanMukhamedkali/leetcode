package main

import "log"

func main() {
	var nums = [][]int{{2, 7, 11, 15}, {3, 2, 4}, {3, 3}, {3, 2, 3}}
	var targets = []int{9, 6, 6, 6}

	for i := range targets {
		log.Println(twoSum(nums[i], targets[i]))
	}
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	buffer := make(map[int]int)

	for i, num := range nums {
		if index, ok := buffer[target-num]; ok {
			return []int{index, i}
		}
		buffer[num] = i

	}
	return nil
}
