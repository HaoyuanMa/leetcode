package main

import (
	"fmt"
	"math"
)

//213.打家劫舍 II.house-robber-ii

// leetcode submit region begin(Prohibit modification and deletion)
const MIN = math.MinInt

func rob(nums []int) int {
	sz := len(nums)
	if sz == 0 {
		return 0
	}
	a00, a01 := 0, MIN
	a10, a11 := MIN, nums[0]
	for i := 1; i < sz; i++ {
		t00 := max(a00, a01)
		t01 := a00 + nums[i]
		t10 := max(a10, a11)
		t11 := a10 + nums[i]
		if i == sz-1 {
			// t00 t01 t10 no conflict
			t11 = MIN
		}
		a00, a01, a10, a11 = t00, t01, t10, t11
	}
	ans0 := max(a00, a01)
	ans1 := max(a11, a10)
	return max(ans0, ans1)
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
	result := rob([]int{1, 2, 3})
	fmt.Println(result)
}
