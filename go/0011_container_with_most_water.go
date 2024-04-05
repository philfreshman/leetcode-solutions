package main

// time: O(n)
// space: O(1)

func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1

	for l < r {
		area := (r - l) * min(height[l], height[r])
		if area > res {
			res = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
