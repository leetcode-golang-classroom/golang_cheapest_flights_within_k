# golang_cheapest_flights_within_k

There are `n` cities connected by some number of flights. You are given an array `flights` where `flights[i] = [fromi, toi, pricei]` indicates that there is a flight from city `fromi` to city `toi` with cost `pricei`.

You are also given three integers `src`, `dst`, and `k`, return ***the cheapest price** from* `src` *to* `dst` *with at most* `k` *stops.* If there is no such route, return **`-1`.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-3drawio.png](https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-3drawio.png)

```
Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
Output: 700
Explanation:
The graph is shown above.
The optimal path with at most 1 stop from city 0 to 3 is marked in red and has cost 100 + 600 = 700.
Note that the path through cities [0,1,2,3] is cheaper but is invalid because it uses 2 stops.

```

**Example 2:**

![https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-1drawio.png](https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-1drawio.png)

```
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
Output: 200
Explanation:
The graph is shown above.
The optimal path with at most 1 stop from city 0 to 2 is marked in red and has cost 100 + 100 = 200.

```

**Example 3:**

![https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-2drawio.png](https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-2drawio.png)

```
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
Output: 500
Explanation:
The graph is shown above.
The optimal path with no stops from city 0 to 2 is marked in red and has cost 500.

```

**Constraints:**

- `1 <= n <= 100`
- `0 <= flights.length <= (n * (n - 1) / 2)`
- `flights[i].length == 3`
- `0 <= fromi, toi < n`
- `fromi != toi`
- `1 <= pricei <= 104`
- There will not be any multiple flights between two cities.
- `0 <= src, dst, k < n`
- `src != dst`

## 解析

題目給定一個整數 n 代表有 0 到 n-1 的飛機站

給定一個 flights 矩陣 , 矩陣中每個 entry , flight[i] = [$source_i, target_i, price_i$] 

代表從 $source_i$ 出發到 $target_i$ 需要花費 $price_i$

給定一個整數 src 代表出發站

給定一個整數 dst 代表到達站

給定一個整數 k 代表最多只經過 k 個中繼站

要求在以上條件下實作一個演算法找出最便宜的從 src 到 dst 的票價

這題如果要使用 [Dijkstra Algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)

會需要注意有最多 k 中繼站的限制

也就是可能要跑 1 到 k 次 [Dijkstra Algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm) 來處理

而這題剛好適合用 [Bellman ford Algorithm](https://zh.wikipedia.org/zh-tw/%E8%B4%9D%E5%B0%94%E6%9B%BC-%E7%A6%8F%E7%89%B9%E7%AE%97%E6%B3%95)

可以用每次從 src 出發透過可行的邊去更新 在level L 的限制之下的 prices

這樣只要更新 k+1 次(除了 source 之外有k+1個點)最後 prices[dst] 就是結果

![](https://i.imgur.com/E7AbnwK.png)

這樣時間複雜度會是 O(E*k)

其中 E 代表邊的個數, k 是 level 

## 程式碼
```go
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

```
## 困難點

1. 比一般的 shortest path 多了一個 最多 k 個中繼站的限制
2. 理解如何逐步更新每個 level 的 prices

## Solve Point

- [x]  需要理解[Bellman ford Algorithm](https://zh.wikipedia.org/zh-tw/%E8%B4%9D%E5%B0%94%E6%9B%BC-%E7%A6%8F%E7%89%B9%E7%AE%97%E6%B3%95)
- [x]  建立一個 prices 陣列來儲存當下 level 的 prices 更新結果
- [x]  每次更新都需要用上一個 level 的更新結果來當做基準