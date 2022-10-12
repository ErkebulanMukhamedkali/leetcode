package main

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
