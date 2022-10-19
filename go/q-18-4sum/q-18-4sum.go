package main

import (
	"fmt"
	"sort"
)

//18.四数之和.4sum

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, target, 4)
}
func nSum(nums []int, target, n int) (ans [][]int) {
	if n == 2 {
		return twoSum(nums, target)
	}
	for i := 0; i <= len(nums)-n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := nSum(nums[i+1:], target-nums[i], n-1)
		for _, t := range tmp {
			ta := []int{nums[i]}
			ta = append(ta, t...)
			ans = append(ans, ta)
		}
	}
	return ans
}
func twoSum(nums []int, target int) (ans [][]int) {
	i, j := 0, len(nums)-1
	for i < j {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		if j < len(nums)-1 && nums[j] == nums[j+1] {
			j--
			continue
		}
		sum := nums[i] + nums[j]
		if sum == target {
			ans = append(ans, []int{nums[i], nums[j]})
			i++
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := nSum([]int{0, 0, 0}, 0, 3)
	fmt.Println(result)
}
