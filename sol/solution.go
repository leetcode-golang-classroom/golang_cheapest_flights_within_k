package sol

import (
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// Bellman algorithm
	prices := make([]int, n)
	for idx := range prices {
		prices[idx] = math.MaxInt16
	}
	prices[src] = 0
	for level := 1; level <= k+1; level++ {
		levelPrices := make([]int, n)
		copy(levelPrices, prices)
		for _, flight := range flights {
			source := flight[0]
			target := flight[1]
			price := flight[2]
			if prices[source] == math.MaxInt16 {
				continue
			}
			if levelPrices[target] > prices[source]+price {
				levelPrices[target] = prices[source] + price
			}
		}
		copy(prices, levelPrices)
	}
	if prices[dst] == math.MaxInt16 {
		return -1
	}
	return prices[dst]
}
