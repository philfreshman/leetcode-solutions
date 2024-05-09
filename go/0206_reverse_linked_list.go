package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursive solution
// time: O(n)
// space: O(n)

func reverseList0(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversedRest := reverseList0(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversedRest
}

// in-place solution
// time: O(n)
// space: O(1)

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

// new-copy solution
// time: O(n)
// space: O(n)

func reverseList2(head *ListNode) *ListNode {
	var newHead *ListNode

	for head != nil {
		newHead = &ListNode{
			Val:  head.Val,
			Next: newHead,
		}
		head = head.Next
	}
	return newHead
}
