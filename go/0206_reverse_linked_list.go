package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// in-place solution
// time: O(n)
// space: O(1)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	var next *ListNode

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// new-copy solution
// time: O(n)
// space: O(n)

// func reverseList(head *ListNode) *ListNode {
// 	var newHead *ListNode
//
// 	for head != nil {
// 		newHead = &ListNode{
// 			Val:  head.Val,
// 			Next: newHead,
// 		}
// 		head = head.Next
// 	}
// 	return newHead
// }
