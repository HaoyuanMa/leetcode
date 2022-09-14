package main

import "fmt"

//283.移动零.move-zeroes

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	l, p, q := len(nums), 0, 0
	for q < l {
		if nums[q] != 0 {
			nums[p] = nums[q]
			p++
		}
		q++
	}
	for p < l {
		nums[p] = 0
		p++
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []int{3, 0, 1, 3, 0, 3, 3, 0, 4, 0, 6, 0, 7}
	moveZeroes(a)
	fmt.Println(a)
}
