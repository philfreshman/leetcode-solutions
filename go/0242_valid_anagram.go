package main

// IsAnagram checks if the two input strings are anagrams of each other.
// An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Examples:
// 1. Input: s = "anagram", t = "nagaram"
//    Output: true
// 2. Input: s = "rat", t = "car"
//    Output: false

func isAnagram(a string, b string) bool {

	aMap := countLetters(a)
	bMap := countLetters(b)

	if len(aMap) != len(bMap) {
		return false
	}

	for key, value1 := range aMap {
		if value2, ok := bMap[key]; !ok || value2 != value1 {
			return false
		}
	}

	return true
}

func countLetters(s string) map[int32]int32 {
	letterMap := make(map[int32]int32)

	for _, r := range s {
		_, found := letterMap[r]
		if found {
			letterMap[r]++
		} else {
			letterMap[r] = 1
		}
	}
	return letterMap
}
