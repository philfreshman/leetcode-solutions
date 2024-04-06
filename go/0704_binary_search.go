package main

func search(arr []int, n int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		middle := (low + high) >> 1 // bitwise right shift to divide by 2

		if arr[middle] == n {
			return middle
		}
		if arr[middle] < n {
			low = middle + 1
		}
		if arr[middle] > n {
			high = middle - 1
		}
	}

	return -1
}
