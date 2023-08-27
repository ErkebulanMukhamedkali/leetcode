package main

import (
	"log"
	"unicode"
)

func main() {
	log.Println(isPalindrome("A man, a plan, a canal: Panama"))
	log.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	arr := []rune(s)

	check := func(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) }

	for l < r {
		left, right := unicode.ToLower(arr[l]), unicode.ToLower(arr[r])
		if !check(left) {
			l++
			continue
		}

		if !check(right) {
			r--
			continue
		}

		if left != right {
			return false
		}

		r--
		l++
	}
	return true
}
