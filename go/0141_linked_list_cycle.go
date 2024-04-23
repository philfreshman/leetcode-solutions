package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// time: O(n)
// space O(1)

func hasCycle(head *ListNode) bool {
	s := head
	f := head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}
