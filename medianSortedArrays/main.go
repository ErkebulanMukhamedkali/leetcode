package main

import (
	"log"
	"sort"
)

func main() {
	nums := [][]int{
		{1, 3}, {2},
		{1, 2}, {3, 4},
	}

	for i := 0; i < len(nums); i += 2 {
		log.Println(findMedianSortedArrays(nums[i], nums[i+1]))
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l == 0 {
		return 0
	}
	nums := make([]int, l)
	copy(nums, append(nums2, nums1...))
	sort.Ints(nums)
	if l%2 == 0 {
		return float64(nums[l/2-1]+nums[l/2]) / 2
	}
	return float64(nums[l/2])
}
