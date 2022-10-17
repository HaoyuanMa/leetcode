package main

import "fmt"

//188.买卖股票的最佳时机 IV.best-time-to-buy-and-sell-stock-iv

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(k int, prices []int) int {
	sz := len(prices)
	ans := make([][]int, sz)
	for i := 0; i < sz; i++ {
		ans[i] = make([]int, k+1)
	}
	curMin := prices[0]
	ans[0][1] = 0
	curMax := ans[0][1]
	for i := 1; i < sz; i++ {
		if prices[i] > curMin {
			if prices[i]-curMin > curMax {
				curMax = prices[i] - curMin
			}
		} else {
			curMin = prices[i]
		}
		ans[i][1] = curMax
	}
	// ans[i][k]
	// ans[i][1] O(n)

	for i := 2; i < sz; i++ {
		for t := 2; t <= k; t++ {
			max := ans[i][t-1]
			if ans[i-1][t] > max {
				max = ans[i-1][t]
			}
			for j := 0; j < i; j++ {
				if prices[i]-prices[j]+ans[j][t-1] > max {
					max = prices[i] - prices[j] + ans[j][t-1]
				}
			}
			ans[i][t] = max
		}
	}
	// A[i,k] = max{(max{ s[i]-s[j]+A[j,k-1] } j = i-1 down to 0),A[i,k-1],A[i-1,k]}
	// t = A[j,k] - s[j]

	return ans[sz-1][k]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() { ///////////////////////0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11,12
	result := maxProfit(2, []int{3, 2, 6, 5, 0, 3})
	fmt.Println(result)
}
