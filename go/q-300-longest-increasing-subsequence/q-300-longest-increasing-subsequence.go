package main

import (
	"fmt"
)

//300.最长递增子序列.longest-increasing-subsequence

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	size := len(nums)
	ans := make([]int, size)
	for i := 0; i < size; i++ {
		ans[i] = 1
	}
	for i := 1; i < size; i++ {
		tmp := 0
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && ans[j] > tmp {
				tmp = ans[j]
			}
		}
		ans[i] = tmp + 1
	}
	res := ans[0]
	for i := 0; i < size; i++ {
		if ans[i] > res {
			res = ans[i]
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
	fmt.Println(result)
}
