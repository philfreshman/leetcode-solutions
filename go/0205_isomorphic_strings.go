package main

// time O(n)
// space O(n)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashST := make(map[byte]byte)
	hashTS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar, tChar := s[i], t[i]
		if sMappedChar, exists := hashST[sChar]; exists && sMappedChar != tChar {
			return false
		}
		if tMappedChar, exists := hashTS[tChar]; exists && tMappedChar != sChar {
			return false
		}
		hashST[sChar] = tChar
		hashTS[tChar] = sChar
	}

	return true
}
