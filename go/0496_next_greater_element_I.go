package main

// time: O(n)
// space: O(n)

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	stack := []int{}
	nextGreater := make(map[int]int)

	for _, num := range nums2 {
		// While the stack is not empty and the current number is greater than the top of the stack
		for len(stack) > 0 && num > stack[len(stack)-1] {
			last := stack[len(stack)-1]
			nextGreater[last] = num
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, num)
	}

	// Set missing nextGreater to -1
	for len(stack) > 0 {
		nextGreater[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	// map nextGreater to nums1
	for i, v := range nums1 {
		nums1[i] = nextGreater[v]
	}

	return nums1
}
