package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var slow, fast = head, head

	if head == nil || head.Next == nil {
		return false
	}

	for {
		fast = fast.Next
		if fast == nil {
			return false
		}

		fast = fast.Next
		if fast == nil {
			return false
		}

		slow = slow.Next

		if fast == slow {
			return true
		}
	}
}
