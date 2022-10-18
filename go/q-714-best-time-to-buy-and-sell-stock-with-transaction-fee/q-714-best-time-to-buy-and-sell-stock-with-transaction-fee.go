package main

import "fmt"

//714.买卖股票的最佳时机含手续费.best-time-to-buy-and-sell-stock-with-transaction-fee

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int, fee int) int {
	sz := len(prices)
	if sz <= 1 {
		return 0
	}

	ans0, ans1 := 0, -prices[0]
	for i := 1; i < sz; i++ {
		t0 := max(ans0, ans1+prices[i]-fee)
		t1 := max(ans1, ans0-prices[i])
		ans0, ans1 = t0, t1
	}
	return ans0
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
	result := maxProfit([]int{1, 2, 3}, 2)
	fmt.Println(result)
}
