package main

import (
	"log"
)

func main() {
	var strings = []string{"abcabcbb", "bbbbb", "pwwkew", "", "aab", "dvdf"}
	for _, s := range strings {
		log.Println(lengthOfLongestSubstring(s))
	}
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	for i := range s {
		curResult := 0
		exist := make(map[string]bool)
		for j := i; j >= 0; j-- {
			if _, ok := exist[string(s[j])]; !ok {
				exist[string(s[j])] = true
				curResult++
			} else {
				break
			}
		}
		if curResult > result {
			result = curResult
		}
	}

	return result
}
