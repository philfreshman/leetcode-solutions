package main

// time: O(n)
// space: O(1)

func twoSum2(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]

		if sum == target {
			break
		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return []int{l + 1, r + 1}
}
