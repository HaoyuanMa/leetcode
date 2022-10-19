package main

import (
	"fmt"
	"math"
)

//309.最佳买卖股票时机含冷冻期.best-time-to-buy-and-sell-stock-with-cooldown

// leetcode submit region begin(Prohibit modification and deletion)
const MIN = math.MinInt

func maxProfit(prices []int) int {
	sz := len(prices)
	if sz <= 1 {
		return 0
	}
	//ans := make([][2][2]int, sz)
	//ans[0][0][0] = 0
	//ans[0][0][1] = MIN
	//ans[0][1][0] = -prices[0]
	//ans[0][1][1] = MIN
	//for i := 1; i < sz; i++ {
	//	ans[i][1][1] = MIN
	//	ans[i][0][1] = ans[i-1][1][0] + prices[i]
	//	ans[i][0][0] = max(ans[i-1][0][0],ans[i-1][0][1])
	//	ans[i][1][0] = max(ans[i-1][1][0],ans[i-1][0][0]-prices[i])
	//}
	//return max(ans[sz-1][0][0],ans[sz-1][0][1])

	ans00 := 0
	ans01 := MIN
	ans10 := -prices[0]
	for i := 1; i < sz; i++ {
		t00 := max(ans00, ans01)
		t01 := ans10 + prices[i]
		t10 := max(ans10, ans00-prices[i])
		ans00, ans01, ans10 = t00, t01, t10
	}
	return max(ans00, ans01)
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
	result := maxProfit([]int{1, 2, 3})
	fmt.Println(result)
}
