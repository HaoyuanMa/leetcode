package main

import "fmt"

//121.买卖股票的最佳时机.best-time-to-buy-and-sell-stock

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	sz := len(prices)
	if sz <= 1 {
		return 0
	}
	curMin := prices[0]
	curMax := 0
	for i := 1; i < sz; i++ {
		curMax = max(curMax, prices[i]-curMin)
		curMin = min(curMin, prices[i])
	}
	return curMax
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := maxProfit([]int{1, 2, 3, 4})
	fmt.Println(result)
}
