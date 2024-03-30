package main

// time O(n)
// space O(n)

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([][2]int, 0) // pair: [temp, index]

	for i, t := range temperatures {

		for len(stack) > 0 && t > stack[len(stack)-1][0] {
			stackInd := stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			res[stackInd] = i - stackInd
		}

		stack = append(stack, [2]int{t, i})
	}

	return res
}
