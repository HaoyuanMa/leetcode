package main

import "fmt"

//198.打家劫舍.house-robber

// leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	sz := len(nums)
	if sz == 0 {
		return 0
	}
	a0, a1 := 0, nums[0]
	for i := 1; i < sz; i++ {
		t0 := max(a0, a1)
		t1 := a0 + nums[i]
		a0, a1 = t0, t1
	}
	return max(a0, a1)
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
	result := rob([]int{1, 2, 5, 4, 3})
	fmt.Println(result)
}
