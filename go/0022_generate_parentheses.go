package main

/*

Here's a breakdown of how this complexity is derived:

Structure of the Function:

It uses a recursive backtracking approach to explore possible combinations of balanced parentheses.
The backtrack function takes two parameters: openN (number of open parentheses) and closedN (number of closed parentheses).
It builds candidate strings by adding parentheses and recursively calls itself for further exploration.
Factors Contributing to Time Complexity:

Decision Tree Size:

At each level of recursion, the function has two choices: add an open parenthesis or a closed parenthesis.
This creates a binary decision tree with up to 2^n levels (2 choices for each of the n parentheses).
However, some branches are pruned due to the constraints on valid combinations.
Catalan Numbers:

The number of valid parentheses combinations for n pairs is given by the nth Catalan number, C_n.
The asymptotic growth of Catalan numbers is approximately 4^n / sqrt(n * π).
Combining Factors:

The function explores a decision tree with up to 4^n / sqrt(n) valid paths (considering the pruning).
Each path involves a constant amount of work for string operations and function calls.
Therefore, the overall time complexity is approximately O(4^n / sqrt(n)).
Additional Considerations:

The space complexity of the function is also O(4^n / sqrt(n)), primarily due to the recursion stack and the resulting set of valid combinations.
This exponential time complexity means that the function's execution time grows rapidly as the input size n increases.

*/

func generateParenthesis(n int) []string {
	var stack string
	var res []string

	// valid if open == closed == n
	// only add openings if open < n
	// only add closing if open > closed

	var backtrack = func(openN int, closedN int) {}
	backtrack = func(openN int, closedN int) {
		if openN == closedN && closedN == n {
			res = append(res, stack)
		}
		if openN < n {
			stack += "("
			backtrack(openN+1, closedN)
			stack = stack[:len(stack)-1]
		}
		if closedN < openN {
			stack += ")"
			backtrack(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return res
}
