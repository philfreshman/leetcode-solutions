package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("twoSum(%v, %v) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
		}
	}
}
