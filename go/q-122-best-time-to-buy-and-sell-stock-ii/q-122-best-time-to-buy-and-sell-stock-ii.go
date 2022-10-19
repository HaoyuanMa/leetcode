package main

import (
	"fmt"
)

//122.买卖股票的最佳时机 II.best-time-to-buy-and-sell-stock-ii

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	sz := len(prices)
	if sz <= 1 {
		return 0
	}
	//ans := make([][2]int,sz)
	//ans[0][0] = 0
	//ans[0][1] = -prices[0]
	//for i := 1; i < sz; i++ {
	//    ans[i][0] = max(ans[i-1][0],ans[i-1][1]+prices[i])
	//    ans[i][1] = max(ans[i-1][1],ans[i-1][0]-prices[i])
	//}
	//return ans[sz-1][0]

	ans0, ans1 := 0, -prices[0]
	for i := 0; i < sz; i++ {
		ans0 = max(ans0, ans1+prices[i])
		ans1 = max(ans1, ans0-prices[i])
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
	result := maxProfit([]int{3, 2, 6})
	fmt.Println(result)
}
