package main

import "fmt"

//303.区域和检索 - 数组不可变.range-sum-query-immutable

// leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	array  []int
	preSum []int
}

func Constructor(nums []int) NumArray {
	res := NumArray{}
	for i := 0; i < len(nums); i++ {
		res.array = append(res.array, nums[i])
		if i == 0 {
			res.preSum = append(res.preSum, nums[0])
		} else {
			res.preSum = append(res.preSum, res.preSum[i-1]+nums[i])
		}
	}
	return res
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right] - this.preSum[left] + this.array[left]
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
