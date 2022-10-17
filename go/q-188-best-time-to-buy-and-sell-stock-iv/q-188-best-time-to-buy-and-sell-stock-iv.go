package main

import "fmt"

//188.买卖股票的最佳时机 IV.best-time-to-buy-and-sell-stock-iv

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(k int, prices []int) int {
	sz := len(prices)
	ans := make([][]int, sz)
	for i := 0; i < sz; i++ {
		ans[i] = make([]int, k)
	}
	curMin := prices[0]
	ans[0][1] = 0
	for i := 1; i < sz; i++ {
		if prices[i] > curMin {
			ans[i][1] = prices[i] - curMin
		} else {
			ans[i][1] = 0
			curMin = prices[i]
		}
	}
	// ans[i][k]
	// ans[i][1] O(n)
	// A[i,k] = max{ s[i]-s[j]+A[j,k-1] } j = i-1 down to 0
	// t = A[j,k] - s[j]

	return ans[sz-1][k-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := maxProfit(2, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(result)
}
