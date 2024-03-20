package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		strs     []string
		expected [][]string
	}{
		{
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"eat", "tea", "ate"}, {"tan", "nat"}},
		},
		{
			strs:     []string{"bdddddddddd", "bbbbbbbbbbc"},
			expected: [][]string{{"bbbbbbbbbbc"}, {"bdddddddddd"}},
		},
		{
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, tc := range testCases {
		result := groupAnagrams(tc.strs)
		sort.Slice(result, func(i, j int) bool {
			return result[i][0] < result[j][0]
		})

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("groupAnagrams(%v) = %v; expected %v", tc.strs, result, tc.expected)
		}
	}
}
