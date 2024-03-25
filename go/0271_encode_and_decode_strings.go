package main

import (
	"fmt"
	"strconv"
)

// time: 0(n)
// space: O(n)

func encode(input []string) string {
	output := ""

	for _, word := range input {
		output += fmt.Sprintf("%d#%s", len(word), word)
	}

	return output
}

func decode(input string) []string {
	var output []string
	i := 0

	for i < len(input) {
		j := i
		for input[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(input[i:j])
		i = j + 1
		j = i + length
		output = append(output, input[i:j])
		i = j
	}

	return output
}
