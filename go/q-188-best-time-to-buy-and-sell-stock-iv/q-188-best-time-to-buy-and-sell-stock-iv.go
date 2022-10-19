package main

import (
	"fmt"
	"math"
)

//188.买卖股票的最佳时机 IV.best-time-to-buy-and-sell-stock-iv

// leetcode submit region begin(Prohibit modification and deletion)
const MIN = math.MinInt

func maxProfit(k int, prices []int) int {
	sz := len(prices)
	ans := make([][][2]int, sz)
	for i := 0; i < sz; i++ {
		ans[i] = make([][2]int, k+1)
		ans[i][0][0] = 0
		ans[i][0][1] = MIN
	}
	for i := 0; i <= k; i++ {
		ans[0][i][0] = 0
		ans[0][i][1] = -prices[0]
	}
	for i := 1; i < sz; i++ {
		for j := 1; j <= k; j++ {
			ans[i][j][0] = max(ans[i-1][j][0], ans[i-1][j][1]+prices[i])
			ans[i][j][1] = max(ans[i-1][j][1], ans[i-1][j-1][0]-prices[i])
		}
	}

	return ans[sz-1][k][0]
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

func main() { ///////////////////////0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11,12
	result := maxProfit(2, []int{3, 2, 6, 5, 0, 3})
	fmt.Println(result)
}

func maxProfit_mine(k int, prices []int) int {
	sz := len(prices)
	ans := make([][]int, sz)
	for i := 0; i < sz; i++ {
		ans[i] = make([]int, k+1)
	}
	curMin := prices[0]
	ans[0][1] = 0
	curMax := ans[0][1]
	for i := 1; i < sz; i++ {
		curMax = max(curMax, prices[i]-curMin)
		curMin = min(curMin, prices[i])
		ans[i][1] = curMax
	}
	// ans[i][k]
	// ans[i][1] O(n)

	for i := 2; i < sz; i++ {
		for t := 2; t <= k; t++ {
			r := max(ans[i][t-1], ans[i-1][t])
			for j := 0; j < i; j++ {
				r = max(r, prices[i]-prices[j]+ans[j][t-1])
			}
			ans[i][t] = r
		}
	}
	// A[i,k] = max{(max{ s[i]-s[j]+A[j,k-1] } j = i-1 down to 0),A[i,k-1],A[i-1,k]}
	// t = A[j,k] - s[j]

	return ans[sz-1][k]
}
