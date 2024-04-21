package main

// time: O(log n)
// space: O(1)

func searchB(nums []int, target int) int {

	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			// left sorted portion
			if target < nums[l] || target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// right sorted portion
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

	}
	return -1
}
