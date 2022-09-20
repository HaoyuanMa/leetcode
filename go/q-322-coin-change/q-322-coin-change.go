package main

import "fmt"

//322.零钱兑换.coin-change

// leetcode submit region begin(Prohibit modification and deletion)
var memo [10001]int

func coinChange(coins []int, amount int) int {
	for i := 0; i < 10001; i++ {
		memo[i] = -5
	}
	return dp(coins, amount)
}

func dp(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	memo[0] = 0
	for i := 1; i <= amount; i++ {
		memo[i] = 10001
		for _, coin := range coins {
			if i-coin >= 0 && memo[i-coin] >= 0 && memo[i-coin]+1 < memo[i] {
				memo[i] = memo[i-coin] + 1
			}
		}
		if memo[i] == 10001 {
			memo[i] = -1
		}
	}
	return memo[amount]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	c := []int{2}
	result := coinChange(c, 3)
	fmt.Println(result)
}

func dp_recursion(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if memo[amount] >= -1 {
		return memo[amount]
	}
	tmpMin := 10001
	for _, coin := range coins {
		t := dp(coins, amount-coin) + 1
		if t > 0 && t < tmpMin {
			tmpMin = t
		}
	}
	if tmpMin < 10001 {
		memo[amount] = tmpMin
		return tmpMin
	} else {
		memo[amount] = -1
		return -1
	}

}

// TLE

//func coinChange(coins []int, amount int) int {
//	if amount == 0 {
//		return 0
//	}
//	var memo [10001]int
//	for i := 0; i <= amount; i++ {
//		memo[i] = 10001
//	}
//	for i := 0; i < len(coins); i++ {
//		if coins[i] <= amount {
//			memo[coins[i]] = 1
//		}
//	}
//	for i := 1; i <= amount; i++ {
//		for j := 1; j <= i/2; j++ {
//			if memo[j] < 10001 && memo[i-j] < 10001 && memo[j]+memo[i-j] < memo[i] {
//				memo[i] = memo[j] + memo[i-j]
//			}
//		}
//	}
//	if memo[amount] < 10001 {
//		return memo[amount]
//	}
//	return -1
//}
