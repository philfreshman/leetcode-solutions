package main

// time O(1)
// space O(n)

type MinStack struct {
	stack []pair
}

type pair struct {
	val int
	min int
}

// func Constructor() MinStack {
// 	return MinStack{}
// }

func (s *MinStack) Push(val int) {
	minVal := val
	if len(s.stack) > 0 && s.GetMin() < val {
		minVal = s.GetMin()
	}
	s.stack = append(s.stack, pair{val: val, min: minVal})
}

func (s *MinStack) Pop() {
	s.stack = s.stack[0 : len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].val
}

func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].min
}
