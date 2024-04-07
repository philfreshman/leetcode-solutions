package main

func maxProfit(prices []int) int {
	lowest := prices[0]
	highest := prices[0]
	result := 0

	for _, price := range prices {
		if price < lowest {
			lowest = price
			highest = price
			continue
		}
		if price > highest {
			highest = price
		}
		if highest-lowest > result {
			result = highest - lowest
		}
	}

	return result
}
