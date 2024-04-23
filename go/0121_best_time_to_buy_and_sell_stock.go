package main

// time: O(n)
// space: O(1)

func maxProfit(prices []int) int {
	res := 0
	lowest := prices[0]
	for i := 0; i < len(prices); i++ {
		lowest = min(lowest, prices[i])
		res = max(prices[i]-lowest, res)
	}
	return res
}
