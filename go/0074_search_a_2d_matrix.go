package main

func searchMatrix(matrix [][]int, target int) bool {
	y := len(matrix) // y-axis
	if y == 0 {
		return false
	}
	x := len(matrix[0]) // x-axis

	// binary search
	left := 0
	right := x*y - 1

	var midIdx, midElement int
	for left <= right {
		midIdx = (left + right) >> 1
		midElement = matrix[midIdx/x][midIdx%x]

		if target == midElement {
			return true
		} else if target < midElement {
			right = midIdx - 1
		} else {
			left = midIdx + 1
		}
	}

	return false
}
