package main

import "log"

func main() {
	log.Println(twoSum2([]int{2, 3, 4}, 6))
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

func twoSum2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for numbers[i]+numbers[j] != target {
		if numbers[j]+numbers[i] > target {
			j -= 1
			continue
		}
		i -= -1
	}
	return []int{i + 1, j + 1}
}
