package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range testCases {
		result := containsDuplicate(tc.arr)
		if !result == tc.expected {
			t.Errorf("containsDuplicate(%v) = %v; want %v", tc.arr, result, tc.expected)
		}
	}
}
