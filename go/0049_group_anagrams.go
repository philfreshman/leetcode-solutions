package main

type Key [26]byte

func groupAnagrams(strs []string) [][]string {
	var pairs = make(map[Key][]string)
	for _, word := range strs {
		var hashKey = Key{}
		for _, char := range word {
			hashKey[char-'a']++
		}

		pairs[hashKey] = append(pairs[hashKey], word)
	}

	var result [][]string
	for _, pair := range pairs {
		result = append(result, pair)
	}

	return result
}
