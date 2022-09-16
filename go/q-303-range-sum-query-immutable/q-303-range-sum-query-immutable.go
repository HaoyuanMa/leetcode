package main

import "fmt"

//303.区域和检索 - 数组不可变.range-sum-query-immutable

// leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	array []int
}

func Constructor(nums []int) NumArray {
	res := NumArray{}
	for i := 0; i < len(nums); i++ {
		res.array = append(res.array, nums[i])
	}
	return res
}

func (this *NumArray) SumRange(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		ans += this.array[i]
	}
	return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := 0
	fmt.Println(result)
}
