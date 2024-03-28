package main

// time: O(n)
// space: O(n)

func isPalindrome(s string) bool {
	lp := 0
	rp := len(s) - 1

	for lp < rp {

		l, r := s[lp], s[rp]

		// IF UPPERCASE
		if l >= 65 && l <= 90 {
			l += 32
		}

		if r >= 65 && r <= 90 {
			r += 32
		}

		// Skip non-alphanumeric characters
		if !((l >= 97 && l <= 122) || (l >= 48 && l <= 57)) {
			lp++
			continue
		}

		if !((r >= 97 && r <= 122) || (r >= 48 && r <= 57)) {
			rp--
			continue
		}

		if l != r {
			return false
		}

		lp++
		rp--
	}

	return true
}
