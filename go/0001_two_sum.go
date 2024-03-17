package main

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

Assumptions:
- Each input would have exactly one solution.
- The same element cannot be used twice.
- The answer can be returned in any order.

Examples:
1. Input: nums = [2,7,11,15], target = 9
   Output: [0,1]
   Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

2. Input: nums = [3,2,4], target = 6
   Output: [1,2]

3. Input: nums = [3,3], target = 6
   Output: [0,1]

Constraints:
- 2 <= nums.length <= 104
- -109 <= nums[i] <= 109
- -109 <= target <= 109
- Only one valid answer exists.
*/

func twoSum(numbers []int, target int) []int {

	indicesMap := make(map[int]int)

	for index, currentNumber := range numbers {

		diff := target - currentNumber

		if targetIdx, found := indicesMap[diff]; found {
			return []int{targetIdx, index}
		}

		indicesMap[currentNumber] = index
	}
	return nil
}
