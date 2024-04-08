package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
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

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	newHead := reverseList(head.Next)
//	head.Next.Next = head
//
//	head.Next = nil
//	return newHead
//}
