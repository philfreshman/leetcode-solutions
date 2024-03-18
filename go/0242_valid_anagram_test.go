package main

import "testing"

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		a        string
		b        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, tc := range testCases {
		result := isAnagram(tc.a, tc.b)
		if !result == tc.expected {
			t.Errorf("isAnagram(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.expected)
		}
	}
}
