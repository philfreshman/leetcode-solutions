package main

// time: O(n)
// space: O(n)

func twoSum(numbers []int, target int) []int {
	hash := map[int]int{}

	for i, n := range numbers {
		diff := target - n
		idx, contains := hash[diff]
		if contains {
			return []int{idx, i}
		}
		hash[n] = i
	}
	return nil
}
