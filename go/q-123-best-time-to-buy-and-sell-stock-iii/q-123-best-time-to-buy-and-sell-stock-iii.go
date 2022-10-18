package main

import (
	"fmt"
	"math"
)

//123.买卖股票的最佳时机 III.best-time-to-buy-and-sell-stock-iii

// leetcode submit region begin(Prohibit modification and deletion)
const MIN = math.MinInt

func maxProfit(prices []int) int {
	sz := len(prices)
	if sz <= 1 {
		return 0
	}
	a10, a11 := 0, -prices[0]
	a20, a21 := 0, -prices[0]

	a01 := -prices[0]
	for i := 1; i < sz; i++ {
		// buy count k
		//t10 := max(a10,a11+prices[i])
		//t11 := max(a11, -prices[i])
		//t20 := max(a20,a21+prices[i])
		//t21 := max(a21,a10-prices[i])

		// sell count k
		t01 := max(a01, -prices[i])
		t10 := max(a10, a01+prices[i])
		t11 := max(a11, a10-prices[i])
		t20 := max(a20, a11+prices[i])
		t21 := max(a21, a20-prices[i])
		a01, a10, a11, a20, a21 = t01, t10, t11, t20, t21
	}
	return a20
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := maxProfit([]int{1, 3, 3, 2})
	fmt.Println(result)
}
