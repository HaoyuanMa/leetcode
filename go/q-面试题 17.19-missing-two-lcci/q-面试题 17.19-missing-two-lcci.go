package main

import (
	"fmt"
	"math"
)

//面试题 17.19.消失的两个数字.missing-two-lcci

// leetcode submit region begin(Prohibit modification and deletion)
func missingTwo(nums []int) []int {
	n := len(nums)
	a, b := 0, 0
	for i := 0; i < n; i++ {
		a += i - nums[i]
		b += i*i - nums[i]*nums[i]
	}
	a += n + n + 1 + n + 2
	b += 3*n*n + 6*n + 5
	sum := a
	product := (a*a - b) / 2
	dt := sum*sum - 4*product
	dt_sqrt := int(math.Sqrt(float64(dt)))
	x := (sum + dt_sqrt) / 2
	y := sum - x
	return []int{x, y}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := missingTwo([]int{})
	fmt.Println(result)
}
