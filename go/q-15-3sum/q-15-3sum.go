package main

import (
	"fmt"
	"sort"
)

//15.三数之和.3sum

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		two := twoSum(nums[i+1:], -nums[i])
		for _, p := range two {
			res = append(res, []int{nums[i], p[0], p[1]})
		}
	}
	return res
}

func twoSum(nums []int, target int) (ans [][2]int) {
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
			ans = append(ans, [2]int{nums[i], nums[j]})
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
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result)
}
