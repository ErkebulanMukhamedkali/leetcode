package main

import (
	"fmt"
	"log"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//method used just for printing
func (l *ListNode) loop(source []int) []int {
	if l.Next == nil {
		return append(source, l.Val)
	}
	return l.Next.loop(append(source, l.Val))
}

func (l *ListNode) String() string {
	return strings.Join(strings.Fields(fmt.Sprint(l.loop(nil))), ",")
}

func main() {
	cases := [][]*ListNode{
		{&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}},
		{&ListNode{0, nil}, &ListNode{0, nil}},
		{
			&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
		},
	}

	for _, c := range cases {
		log.Println(addTwoNumbers(c[0], c[1]))
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return helper(l1, l2, 0)
}

func helper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil {
		l1 = &ListNode{0, nil}
	}
	if l2 == nil {
		l2 = &ListNode{0, nil}
	}

	if l1.Next != nil || l2.Next != nil {
		sum := l1.Val + l2.Val + carry
		if sum > 9 {
			return &ListNode{sum % 10, helper(l1.Next, l2.Next, sum/10)}
		}
		return &ListNode{sum, helper(l1.Next, l2.Next, 0)}
	}

	sum := l1.Val + l2.Val + carry
	if sum > 9 {
		return &ListNode{sum % 10, &ListNode{sum / 10, nil}}
	}
	return &ListNode{sum, nil}
}
