package main

import "sort"

// time: O(n log n)
// space: O(n)

func carFleet(target int, position []int, speed []int) int {
	pairs := make([][2]int, len(position))
	for i := range position {
		pairs[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	stack := []float64{}
	for _, pair := range pairs {
		time := float64(target-pair[0]) / float64(pair[1])
		if len(stack) == 0 || time > stack[len(stack)-1] {
			stack = append(stack, time)
		}
	}

	return len(stack)
}
