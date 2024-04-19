package main

// time: O(log n)
// space: O(1)

func findMin(nums []int) int {

	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
