package main

func lengthOfLongestSubstring(s string) int {
	windowStart := 0
	maxLength := 0
	charIndexMap := make(map[string]int)

	for i, char := range s {
		if _, ok := charIndexMap[string(char)]; ok {
			windowStart = max(windowStart, charIndexMap[string(char)]+1)
		}
		charIndexMap[string(char)] = i
		maxLength = max(maxLength, i-windowStart+1)
	}

	return maxLength
}
