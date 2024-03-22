package main

// time: O(n)
// space: O(n)

func isValid(s string) bool {
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]byte, 0)

	for _, char := range []byte(s) {
		pair, isClosingBracket := pairs[char]

		// if it's not a closing bracket, add it to the stack
		if !isClosingBracket {
			stack = append(stack, char)
			continue
		}

		// if opening pair exist, but not in the stack => exit
		if len(stack) == 0 {
			return false
		}

		// if last element is different that the pair
		if stack[len(stack)-1] != pair {
			return false
		}

		// remove last element from stack
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
