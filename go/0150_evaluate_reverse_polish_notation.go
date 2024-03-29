package main

// time: O(n)
// space: O(n)

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {

		if token != "+" && token != "-" && token != "*" && token != "/" {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
			continue
		}

		num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]

		switch token {
		case "+":
			stack = append(stack, num1+num2)
			break
		case "-":
			stack = append(stack, num1-num2)
			break
		case "*":
			stack = append(stack, num1*num2)
			break
		case "/":
			stack = append(stack, num1/num2)
			break
		}
	}

	return stack[0]
}
